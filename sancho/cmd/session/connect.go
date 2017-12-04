/*
 * Copyright (C) 2017 Nethesis S.r.l.
 * http://www.nethesis.it - info@nethesis.it
 *
 * This file is part of Windmill project.
 *
 * NethServer is free software: you can redistribute it and/or modify
 * it under the terms of the GNU General Public License as published by
 * the Free Software Foundation, either version 3 of the License,
 * or any later version.
 *
 * NethServer is distributed in the hope that it will be useful,
 * but WITHOUT ANY WARRANTY; without even the implied warranty of
 * MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
 * GNU General Public License for more details.
 *
 * You should have received a copy of the GNU General Public License
 * along with NethServer.  If not, see COPYING.
 *
 * author: Edoardo Spadoni <edoardo.spadoni@nethesis.it>
 */

package session

import (
	"fmt"
	"errors"
	"os"
	"os/exec"

	"github.com/spf13/cobra"

	"sancho/helper"
	"sancho/configuration"
)

func connectSession(sessionId string) {
	vpnIp := helper.GetSessionIp(sessionId)
	port := configuration.Config.SSHPort

	if (len(vpnIp) > 0) {
		fmt.Printf("Try connection on %s session...\n", helper.GreenString(sessionId))

		vpnCmd := exec.Command("/opt/windmill/helpers/windmill-start-ssh", vpnIp, port, sessionId)
		vpnCmd.Stdin = os.Stdin
		vpnCmd.Stdout = os.Stdout
		vpnCmd.Stderr = os.Stderr

		vpnCmd.Run()
	} else {
		helper.ErrorLog("Error: session %s not found\n", sessionId)
	}
}

var ConnectCmd = &cobra.Command{
	Use: "ssh <session-id>",
	Short: "Connect to server by specify Session ID",
	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) < 1 {
			return errors.New(helper.RedString("requires session-id"))
		}
		return nil;
	},
	Run: func(cmd *cobra.Command, args []string) {
		sessionId := args[0]

		connectSession(sessionId)
	},
}

func init() {
}
