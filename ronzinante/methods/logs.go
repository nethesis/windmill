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

 package methods

  import (
	  "net/http"
	  "time"

	  "github.com/gin-gonic/gin"
	  _ "github.com/jinzhu/gorm/dialects/mysql"

	  "ronzinante/database"
	  "ronzinante/models"
  )

  func CreateLog(c *gin.Context) {
	  sessionId := c.PostForm("session_id")
	  operatorId := c.PostForm("operator_id")

	  var session models.Session
	  db := database.Database()
	  db.Where("session_id = ?", sessionId).First(&session)

	  if session.Id == 0 {
		  c.JSON(http.StatusNotFound, gin.H{"message": "No session found!"})
		  return
	  }

	  sessionCreated := session.Started
	  sessionConnected := time.Now().String()

	  log := models.Log{
		  SessionId: sessionId,
		  OperatorId: operatorId,
		  SessionCreated: sessionCreated,
		  SessionConnected: sessionConnected,
	  }

	  db.Save(&log)

	  c.JSON(http.StatusCreated, gin.H{"id": log.Id})
  }

  func UpdateLog(c *gin.Context) {
	var log models.Log
	logId := c.Param("log_id")

	db := database.Database()
	db.Where("id = ?", logId).First(&log)

	if log.Id == 0 {
		c.JSON(http.StatusNotFound, gin.H{"message": "No log found!"})
		return
	}

	log.SessionDisconnected = time.Now().String()
	db.Save(&log)
}

  func GetLogs(c *gin.Context) {
	  var logs []models.Log

	  db := database.Database()
	  db.Find(&logs)

	  if len(logs) <= 0 {
		  c.JSON(http.StatusNotFound, gin.H{"message": "No logs found!"})
		  return
	  }

	  c.JSON(http.StatusOK, logs)
  }

  func GetLog(c *gin.Context) {
	  var log models.Log
	  sessionId := c.Param("session_id")

	  db := database.Database()
	  db.Where("session_id = ?", sessionId).Find(&log)

	  if log.Id == 0 {
		  c.JSON(http.StatusNotFound, gin.H{"message": "No log found!"})
		  return
	  }

	  c.JSON(http.StatusOK, log)
  }