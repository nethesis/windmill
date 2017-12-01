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

	 "github.com/gin-gonic/gin"
	 _ "github.com/jinzhu/gorm/dialects/mysql"

	 "ronzinante/database"
	 "ronzinante/models"
 )

 func UpdateHistory(c *gin.Context) {
	 var history models.History
	 serverId := c.Param("server_id")

	 bytes_received := c.PostForm("bytes_received")
	 bytes_sent := c.PostForm("bytes_sent")
	 ifconfig_pool_local_ip := c.PostForm("ifconfig_pool_local_ip")
	 ifconfig_pool_netmask := c.PostForm("ifconfig_pool_netmask")
	 ifconfig_pool_remote_ip := c.PostForm("ifconfig_pool_remote_ip")
	 time_duration := c.PostForm("time_duration")
	 trusted_ip := c.PostForm("trusted_ip")
	 trusted_port := c.PostForm("trusted_port")

	 db := database.Database()
	 db.Where("server_id = ?", serverId).First(&history)

	 if history.Id == 0 {
		 c.JSON(http.StatusNotFound, gin.H{"message": "No history found!"})
		 return
	 }

	 history.BytesReceived = bytes_received
	 history.BytesSent = bytes_sent
	 history.IfconfigPoolLocalIp = ifconfig_pool_local_ip
	 history.IfconfigPoolNetmask = ifconfig_pool_netmask
	 history.IfconfigPoolRemoteIp = ifconfig_pool_remote_ip
	 history.TimeDuration = time_duration
	 history.TrustedIp = trusted_ip
	 history.TrustedPort = trusted_port
	 db.Save(&history)
 }

 func GetHistories(c *gin.Context) {
	 var histories []models.History

	 db := database.Database()
	 db.Find(&histories)

	 if len(histories) <= 0 {
		 c.JSON(http.StatusNotFound, gin.H{"message": "No histories found!"})
		 return
	 }

	 c.JSON(http.StatusOK, histories)
 }

 func GetHistory(c *gin.Context) {
	 var history models.History
	 serverId := c.Param("server_id")

	 db := database.Database()
	 db.Where("server_id = ?", serverId).Find(&history)

	 if history.Id == 0 {
		 c.JSON(http.StatusNotFound, gin.H{"message": "No history found!"})
		 return
	 }

	 c.JSON(http.StatusOK, history)
 }