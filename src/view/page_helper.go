// Copyright 2024 The localFS Authors.
// Use of this source code is governed by a GPL
// license that can be found in the LICENSE file.

package view

var ListingIndex = func(idx int) int {
	return idx + 1
}

var ListingZebraCss = func(idx int) bool {
	return idx%2 != 0
}

type NavItem struct {
	Name string
	Link string
}

type NavBar struct {
	NavItem    []NavItem
	ActiveItem string
}
