// Copyright 2024 The localFS Authors.
// Use of this source code is governed by a GPL
// license that can be found in the LICENSE file.

package main

const (
	appBuild string = "0.1.0"
)

const (
	defaultHost    string = "0.0.0.0"
	defaultPort    string = "5000"
	defaultStorage string = ".localfs"
)

const ckey_storage = "storage"
const ckey_port = "port"

var appCache = map[string]string{}
