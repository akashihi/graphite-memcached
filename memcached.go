/*
   conntrack-logger
   Copyright (C) 2015 Denis V Chapligin <akashihi@gmail.com>
   This program is free software: you can redistribute it and/or modify
   it under the terms of the GNU General Public License as published by
   the Free Software Foundation, either version 3 of the License, or
   (at your option) any later version.
   This program is distributed in the hope that it will be useful,
   but WITHOUT ANY WARRANTY; without even the implied warranty of
   MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
   GNU General Public License for more details.
   You should have received a copy of the GNU General Public License
   along with this program.  If not, see <http://www.gnu.org/licenses/>.
*/

package main

import (
	"bufio"
	"net"
)

func getStatusData(address string) ([]string, error) {
	conn, err := net.Dial("tcp", address)
	if err != nil {
		log.Error("Can't connect to memcache: %v", err)
		return []string{}, err
	}
	defer conn.Close()

	var body []string
	for {
		reader := bufio.NewReader(conn)
		line, err := reader.ReadString('\n')
		if err != nil {
			log.Error("Can't read statistics from memcache: %v", err)
			return []string{}, err
		}
		if line == "END" {
			return body, nil
		}
		body = append(body, line)
	}
	return body, nil
}
