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
	"regexp"
)

type Status struct {
	Items            string
	TotalItems       string
	Connections      string
	TotalConnections string
	Get              string
	Set              string
	Hit              string
	Miss             string
	Evict            string
	BytesIn          string
	BytesOut         string
}

var ItemsRX, _ = regexp.Compile(`curr_items:\s+(\d+)\s+`)
var TotalItemsRX, _ = regexp.Compile(`total_items:\s+(\d+)\s+`)
var ConnectionsRX, _ = regexp.Compile(`curr_connections:\s+(\d+)\s+`)
var TotalConnectionsRX, _ = regexp.Compile(`total_connections:\s+(\d+)\s+`)
var GetRX, _ = regexp.Compile(`cmd_get:\s+(\d+)\s+`)
var SetRX, _ = regexp.Compile(`cmd_set:\s+(\d+)\s+`)
var HitRX, _ = regexp.Compile(`get_hits:\s+(\d+)\s+`)
var MissRX, _ = regexp.Compile(`get_misses:\s+(\d+)\s+`)
var EvictRX, _ = regexp.Compile(`evictions:\s+(\d+)\s+`)
var BytesInRx, _ = regexp.Compile(`bytes_read:\s+(\d+)\s+`)
var BytesOutRx, _ = regexp.Compile(`bytes_written:\s+(\d+)\s+`)

func parse(statusData []string) Status {
	var result = Status{}

	for _, element := range statusData {
		if ItemsRX.MatchString(element) {
			result.Items = ItemsRX.FindStringSubmatch(element)[1]
		}

		if TotalItemsRX.MatchString(element) {
			result.TotalItems = TotalItemsRX.FindStringSubmatch(element)[1]
		}

		if ConnectionsRX.MatchString(element) {
			result.Connections = ConnectionsRX.FindStringSubmatch(element)[1]
		}

		if TotalConnectionsRX.MatchString(element) {
			result.TotalConnections = TotalConnectionsRX.FindStringSubmatch(element)[1]
		}

		if GetRX.MatchString(element) {
			result.Get = GetRX.FindStringSubmatch(element)[1]
		}

		if SetRX.MatchString(element) {
			result.Set = SetRX.FindStringSubmatch(element)[1]
		}

		if HitRX.MatchString(element) {
			result.Hit = HitRX.FindStringSubmatch(element)[1]
		}

		if MissRX.MatchString(element) {
			result.Miss = MissRX.FindStringSubmatch(element)[1]
		}

		if EvictRX.MatchString(element) {
			result.Evict = EvictRX.FindStringSubmatch(element)[1]
		}

		if BytesInRx.MatchString(element) {
			result.BytesIn = BytesInRx.FindStringSubmatch(element)[1]
		}

		if BytesOutRx.MatchString(element) {
			result.BytesOut = BytesOutRx.FindStringSubmatch(element)[1]
		}
	}

	return result
}
