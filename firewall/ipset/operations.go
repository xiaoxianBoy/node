/*
 * Copyright (C) 2020 The "MysteriumNetwork/node" Authors.
 *
 * This program is free software: you can redistribute it and/or modify
 * it under the terms of the GNU General Public License as published by
 * the Free Software Foundation, either version 3 of the License, or
 * (at your option) any later version.
 *
 * This program is distributed in the hope that it will be useful,
 * but WITHOUT ANY WARRANTY; without even the implied warranty of
 * MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
 * GNU General Public License for more details.
 *
 * You should have received a copy of the GNU General Public License
 * along with this program.  If not, see <http://www.gnu.org/licenses/>.
 */

package ipset

import (
	"net"
	"strconv"
)

// SetType defines type of IP set.
type SetType string

var (
	// SetTypeHashIP set type uses a hash to store IP addresses where clashing is resolved by storing the clashing elements in an array and, as a last resort, by dynamically growing the hash.
	SetTypeHashIP = SetType("hash:ip")
)

// OpVersion is an operation which prints version information.
func OpVersion() []string {
	return []string{"version"}
}

// OpSetCreate is an operation which creates a new set.
func OpSetCreate(setName string, setType SetType, netMask net.IPMask, hashSize int) []string {
	args := []string{"create", setName, string(setType)}
	if netMask != nil {
		ones, _ := netMask.Size()
		args = append(args, "--netmask", strconv.Itoa(ones))
	}
	if hashSize != 0 {
		args = append(args, "--hashsize", strconv.Itoa(hashSize))
	}
	return args
}

// OpSetDelete is an operation which destroys a named set.
func OpSetDelete(setName string) []string {
	return []string{"destroy", setName}
}

// OpSetIPAdd is an operation which adds IP entry to the named set.
func OpSetIPAdd(setName string, ip net.IP) []string {
	return []string{"add", setName, ip.String()}
}

// OpSetIPRemove is an operation which deletes IP entry from the named set.
func OpSetIPRemove(setName string, ip net.IP) []string {
	return []string{"del", setName, ip.String()}
}
