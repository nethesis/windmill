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

	 "github.com/spf13/cobra"

	 "sancho/cmd/session"
 )

 var sessionCmd = &cobra.Command{
	 Use: "session <command>",
	 Short: "Perform action to sessions",
	 Args: cobra.MinimumNArgs(1),
	 Run: func(cmd *cobra.Command, args []string) {
		 fmt.Println(RootCmd.Use)
	 },
 }

 func init() {
	 RootCmd.AddCommand(sessionCmd)
	 sessionCmd.AddCommand(session.ListCmd);
	 sessionCmd.AddCommand(session.CloseCmd);
	 sessionCmd.AddCommand(session.ConnectCmd);
 }
