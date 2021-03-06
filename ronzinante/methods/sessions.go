/*
 * Copyright (C) 2017 Nethesis S.r.l.
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
 *
 * author: Edoardo Spadoni <edoardo.spadoni@nethesis.it>
 */

package methods

import (
	"net/http"
	"os/exec"
	"time"

	"github.com/gin-gonic/gin"
	_ "github.com/jinzhu/gorm/dialects/mysql"

	"github.com/nethesis/windmill/ronzinante/database"
	"github.com/nethesis/windmill/ronzinante/models"
)

func CreateSession(c *gin.Context) {
	var session models.Session
	sessionId := c.PostForm("session_id")
	serverId := c.PostForm("server_id")
	started := time.Now().String()

	db := database.Instance()
	db.Where("session_id = ?", sessionId).First(&session)

	if session.Id > 0 {
		c.JSON(http.StatusConflict, gin.H{"message": "Session already exists!"})
		return
	}

	session = models.Session{
		ServerId:  serverId,
		SessionId: sessionId,
		VpnIp:     "",
		Started:   started,
	}

	db.Save(&session)


	c.JSON(http.StatusCreated, gin.H{"id": session.Id})
}

func UpdateSession(c *gin.Context) {
	var session models.Session
	serverId := c.Param("server_id")
	vpnIp := c.PostForm("vpn_ip")

	db := database.Instance()
	db.Where("server_id = ?", serverId).First(&session)

	if session.Id == 0 {
		c.JSON(http.StatusNotFound, gin.H{"message": "No session found!"})
		return
	}

	session.VpnIp = vpnIp
	db.Save(&session)

	c.JSON(http.StatusOK, gin.H{"message": "Session updated successfully!"})
}

func GetSessions(c *gin.Context) {
	var sessions []models.Session

	db := database.Instance()
	db.Find(&sessions)

	if len(sessions) <= 0 {
		c.JSON(http.StatusNotFound, gin.H{"message": "No sessions found!"})
		return
	}

	// check vpn ips
	for _, s := range sessions {
		if s.VpnIp == "" {
			// call helper to get ip
			out, _ := exec.Command("/opt/windmill/helpers/windmill-get-ip", s.ServerId).Output()
			s.VpnIp = string(out)

			// save vpn ip
			db.Save(&s)
		}
	}

	c.JSON(http.StatusOK, sessions)
}

func GetSession(c *gin.Context) {
	var session models.Session
	sessionId := c.Param("session_id")

	db := database.Instance()
	db.Where("session_id = ?", sessionId).First(&session)

	if session.Id == 0 {
		c.JSON(http.StatusNotFound, gin.H{"message": "No session found!"})
		return
	}

	if session.VpnIp == "" {
		// call helper to get ip
		out, _ := exec.Command("/opt/windmill/helpers/windmill-get-ip", session.ServerId).Output()
		session.VpnIp = string(out)

		// save vpn ip
		db.Save(&session)
	}

	c.JSON(http.StatusOK, session)
}

func DeleteSession(c *gin.Context) {
	var session models.Session
	var history models.History
	serverId := c.Param("server_id")

	db := database.Instance()
	db.Where("server_id = ?", serverId).First(&session)

	if session.Id == 0 {
		c.JSON(http.StatusNotFound, gin.H{"message": "No session found!"})
		return
	}

	// add to history this session
	history.SessionId = session.SessionId
	history.ServerId = session.ServerId
	history.Started = time.Now().String()
	db.Save(&history)

	db.Delete(&session)

	c.JSON(http.StatusOK, gin.H{"message": "Session deleted successfully!"})
}
