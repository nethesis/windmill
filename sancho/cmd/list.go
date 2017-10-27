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

package cmd

import (
	"fmt"
	"net/http"
	"io/ioutil"
	"encoding/json"
	"bytes"

	"github.com/spf13/cobra"

	"sancho/model"
	"sancho/config"
	"sancho/helper"
)

var (
	jsonFlag = false
)

func printJSON(body []byte) {
	var prettyJSON bytes.Buffer
	err := json.Indent(&prettyJSON, body, "", "  ")
	if err != nil {
		helper.RedPanic(err.Error())
	}
	fmt.Println(string(prettyJSON.Bytes()))
}

func printSession(session model.Session) {
	if (jsonFlag) {
		jsonPrint := []byte(`{
			"session":"` +session.SessionId + `",
			"lk":"` + session.Lk + `",
			"vpn":"` + session.VpnIp + `",
			"started":"` + session.Started + `"
		}`)
		printJSON(jsonPrint);
	} else {
		fmt.Printf("session: %s\n", helper.GreenString(session.SessionId))
		fmt.Printf("  lk:\t\t%s\n", session.Lk)
		fmt.Printf("  vpn:\t\t%s\n", session.VpnIp)
		fmt.Printf("  started:\t%s\n\n", session.Started)
	}
}

func listSessions() {
	resp, err := http.Get(config.API + "sessions")

	if err != nil {
		helper.RedPanic(err.Error())
	}
	defer resp.Body.Close()

	if (resp.StatusCode < 300) {
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			helper.RedPanic(err.Error())
		}

		var sessions[] model.Session
		err = json.Unmarshal(body, &sessions)
		if err != nil {
			helper.RedPanic(err.Error())
		}

		for i := 0; i < len(sessions); i++ {
			printSession(sessions[i])
		}
	} else {
		helper.ErrorLog("No sessions found\n")
	}
}

func listSession(sessionId string) {
	resp, err := http.Get(config.API + "sessions/" + sessionId)

	if err != nil {
		helper.RedPanic(err.Error())
	}
	defer resp.Body.Close()

	if (resp.StatusCode < 300) {
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			helper.RedPanic(err.Error())
		}

		var session model.Session
		err = json.Unmarshal(body, &session)
		if err != nil {
			helper.RedPanic(err.Error())
		}

		printSession(session)
	} else {
		helper.ErrorLog("No session %s found\n", sessionId)
	}
}

var listCmd = &cobra.Command{
	Use: "list [session-id]",
	Short: "Show all VPNs of connected servers",
	Args: cobra.MaximumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 1 {
            listSession(args[0])
        } else {
			listSessions()
        }
	},
}

func init() {
	RootCmd.AddCommand(listCmd)

	listCmd.Flags().BoolVarP(&jsonFlag, "json", "j", false, "Print output in JSON format")
}