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

package helper

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/briandowns/spinner"
	"github.com/fatih/color"

	"sancho/configuration"
	"sancho/model"
)

var (
	SuccessLog  = color.New(color.FgHiGreen).PrintfFunc()
	ErrorLog    = color.New(color.FgHiRed).PrintfFunc()
	GreenString = color.HiGreenString
	RedString   = color.HiRedString
	CyanString  = color.HiCyanString
	Loader      = spinner.New(spinner.CharSets[41], 100*time.Millisecond)
)

func RedPanic(err string) {
	panic(color.HiRedString(err))
}

func StartLoader() {
	Loader.Start()
}

func StopLoader() {
	Loader.Stop()
}

func GetSessionIp(sessionId string) string {
	resp, err := http.Get(configuration.Config.APIEndpoint + "sessions/" + sessionId)

	if err != nil {
		RedPanic(err.Error())
	}
	defer resp.Body.Close()

	if resp.StatusCode < 300 {
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			RedPanic(err.Error())
		}

		var session model.Session
		err = json.Unmarshal(body, &session)
		if err != nil {
			RedPanic(err.Error())
		}

		return session.VpnIp
	} else {
		return ""
	}
}
