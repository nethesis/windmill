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

package main

import (
	"github.com/gin-gonic/gin"

	"github.com/nethesis/windmill/ronzinante/configuration"
	"github.com/nethesis/windmill/ronzinante/methods"
)

func main() {
	// read and init configuration
	configuration.Init()

	// init routers
	router := gin.Default()

	sessions := router.Group("/api/sessions")
	{
		sessions.GET("/", methods.GetSessions)
		sessions.GET("/:session_id", methods.GetSession)
		sessions.POST("/", methods.CreateSession)
		sessions.PUT("/:server_id", methods.UpdateSession)
		sessions.DELETE("/:server_id", methods.DeleteSession)
	}

	histories := router.Group("/api/histories")
	{
		histories.GET("/", methods.GetHistories)
		histories.GET("/:server_id", methods.GetHistory)
		histories.PUT("/:server_id", methods.UpdateHistory)
	}

	logs := router.Group("/api/logs")
	{
		logs.GET("/", methods.GetLogs)
		logs.GET("/:session_id", methods.GetLog)
		logs.POST("/", methods.CreateLog)
		logs.PUT("/:log_id", methods.UpdateLog)
	}

	router.Run()

}
