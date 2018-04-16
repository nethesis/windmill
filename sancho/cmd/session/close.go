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

package session

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"os/exec"

	"github.com/spf13/cobra"

	"github.com/nethesis/windmill/sancho/configuration"
	"github.com/nethesis/windmill/sancho/helper"
	"github.com/nethesis/windmill/sancho/model"
)

func closeConnections() {
	resp, err := http.Get(configuration.Config.APIEndpoint + "sessions")

	if err != nil {
		helper.RedPanic(err.Error())
	}
	defer resp.Body.Close()

	if resp.StatusCode < 300 {
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			helper.RedPanic(err.Error())
		}

		var sessions []model.Session
		err = json.Unmarshal(body, &sessions)
		if err != nil {
			helper.RedPanic(err.Error())
		}

		for i := 0; i < len(sessions); i++ {
			closeConnection(sessions[i].SessionId)
		}
	} else {
		helper.ErrorLog("No sessions found\n")
	}
}

func closeConnection(sessionId string) {
	vpnIp := helper.GetSessionIp(sessionId)
	port := configuration.Config.SSHPort

	if len(vpnIp) > 0 {
		helper.StartLoader()
		fmt.Printf("Try to close %s session...\n", helper.GreenString(sessionId))

		vpnCmd := exec.Command("/opt/windmill/helpers/windmill-stop-ssh", vpnIp, port)

		if err := vpnCmd.Start(); err != nil {
			helper.RedPanic(err.Error())
		}

		if err := vpnCmd.Wait(); err != nil {
			if exiterr, ok := err.(*exec.ExitError); ok {
				helper.RedPanic(exiterr.Error())
			} else {
				helper.RedPanic(err.Error())
			}
		}

		helper.StopLoader()
		fmt.Printf("Session %s closed!\n", helper.GreenString(sessionId))
	} else {
		helper.ErrorLog("Error: session %s not found\n", sessionId)
	}
}

var CloseCmd = &cobra.Command{
	Use:   "close <session-id>",
	Short: "Close Session ID and remove VPN connection",
	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) < 1 {
			return errors.New(helper.RedString("requires session-id"))
		}
		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {
		closeConnection(args[0])
	},
}

func init() {
}
