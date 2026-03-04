/*
 * Copyright (C) 2026 Nethesis S.r.l.
 * http://www.nethesis.it - info@nethesis.it
 *
 * This file is part of Windmill project.
 *
 * WindMill is free software: you can redistribute it and/or modify
 * it under the terms of the GNU General Public License as published by
 * the Free Software Foundation, either version 3 of the License,
 * or any later version.
 *
 * WindMill is distributed in the hope that it will be useful,
 * but WITHOUT ANY WARRANTY; without even the implied warranty of
 * MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
 * GNU General Public License for more details.
 *
 * You should have received a copy of the GNU General Public License
 * along with WindMill.  If not, see COPYING.
 */

package tasks

import (
	"fmt"
	"net"
	"strings"
	"time"

	"github.com/nethesis/windmill/ronzinante/configuration"
	"github.com/nethesis/windmill/ronzinante/database"
	"github.com/nethesis/windmill/ronzinante/models"
)

func StartSessionCleaner() {
	go func() {
		// run immediately at startup
		cleanExpiredSessions()

		ticker := time.NewTicker(1 * time.Hour)
		for range ticker.C {
			cleanExpiredSessions()
		}
	}()
}

func cleanExpiredSessions() {
	maxAge := time.Duration(configuration.Config.SessionMaxAge) * time.Hour

	db := database.Instance()
	var sessions []models.Session
	db.Find(&sessions)

	for _, session := range sessions {
		started, err := parseStartedTime(session.Started)
		if err != nil {
			fmt.Printf("[session-cleaner] error parsing started time for session %s: %v\n", session.ServerId, err)
			continue
		}

		if time.Since(started) < maxAge {
			continue
		}

		fmt.Printf("[session-cleaner] session %s expired (started: %s)\n", session.ServerId, session.Started)

		// try to kill the client on all OpenVPN sockets
		killed := false
		for _, socketPath := range configuration.Config.OpenVPNSockets {
			if killVPNClient(socketPath, session.ServerId) {
				fmt.Printf("[session-cleaner] killed %s via %s\n", session.ServerId, socketPath)
				killed = true
				break
			}
		}

		if killed {
			// disconnect hook (windmill-disconnect) will handle history + cleanup
			continue
		}

		// orphan session: client not found on any socket, cleanup directly
		fmt.Printf("[session-cleaner] session %s is orphan, cleaning up directly\n", session.ServerId)
		var history models.History
		db.Where("session_id = ?", session.SessionId).First(&history)
		if history.Id == 0 {
			history.SessionId = session.SessionId
			history.ServerId = session.ServerId
			history.Started = time.Now().String()
			db.Save(&history)
		}
		db.Delete(&session)
	}
}

func killVPNClient(socketPath string, serverID string) bool {
	conn, err := net.DialTimeout("unix", socketPath, 5*time.Second)
	if err != nil {
		fmt.Printf("[session-cleaner] cannot connect to %s: %v\n", socketPath, err)
		return false
	}
	defer conn.Close()

	conn.SetDeadline(time.Now().Add(10 * time.Second))

	// read the initial banner
	buf := make([]byte, 4096)
	conn.Read(buf)

	// send kill command
	_, err = fmt.Fprintf(conn, "kill %s\r\n", serverID)
	if err != nil {
		fmt.Printf("[session-cleaner] error sending kill to %s: %v\n", socketPath, err)
		return false
	}

	// read response
	n, err := conn.Read(buf)
	if err != nil {
		fmt.Printf("[session-cleaner] error reading response from %s: %v\n", socketPath, err)
		return false
	}

	response := string(buf[:n])
	return strings.Contains(response, "SUCCESS")
}

func parseStartedTime(started string) (time.Time, error) {
	// time.Now().String() produces: "2006-01-02 15:04:05.999999999 -0700 MST m=+0.000000001"
	// strip the monotonic clock suffix if present
	if idx := strings.Index(started, " m="); idx != -1 {
		started = started[:idx]
	}
	return time.Parse("2006-01-02 15:04:05.999999999 -0700 MST", started)
}
