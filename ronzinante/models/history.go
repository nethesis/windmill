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

package models

type History struct {
	Id						uint		`db:"id" json:"id"`
	SessionId				string		`db:"session_id" json:"session"`
	ServerId				string		`db:"server_id" json:"server"`
	Started					string		`db:"started" json:"started"`
	BytesReceived			string		`db:"bytes_received" json:"bytes_received"`
	BytesSent				string		`db:"bytes_sent" json:"bytes_sent"`
	IfconfigPoolLocalIp		string		`db:"ifconfig_pool_local_ip" json:"ifconfig_pool_local_ip"`
	IfconfigPoolNetmask		string		`db:"ifconfig_pool_netmask" json:"ifconfig_pool_netmask"`
	IfconfigPoolRemoteIp	string		`db:"ifconfig_pool_remote_ip" json:"ifconfig_pool_remote_ip"`
	TimeDuration			string		`db:"time_duration" json:"time_duration"`
	TrustedIp				string		`db:"trusted_ip" json:"trusted_ip"`
	TrustedPort				string		`db:"trusted_port" json:"trusted_port"`
}