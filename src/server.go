// Copyright 2024 The localFS Authors.
// Use of this source code is governed by a GPL
// license that can be found in the LICENSE file.

package main

import (
	"localfs/util/fsutil"
	"log"
	"net"
	"net/http"
	"os"
	"path/filepath"
)

type server struct {
	port string
}

func httpServer(port string) *server {
	return &server{
		port: port,
	}
}

func (s *server) initRoutes() {
	// handle index page and all invalid routes
	http.HandleFunc("/", routesHandler)
	// handle upload routes
	http.HandleFunc("/upload", uploadPageHandler)
	http.HandleFunc("/upload/file", uploadFileHandler)
	http.HandleFunc("/upload/status", uploadStatusPageHandler)
	// handle files download
	http.Handle("/download/", fileHandler("/download/"))
}

func (s *server) initStorage() {
	userHome, err := os.UserHomeDir()
	if err != nil {
		log.Fatal("FATAL", err)
		return
	}

	path := filepath.Join(userHome, defaultStorage)
	err = fsutil.Mkdir(path)
	if err != nil {
		log.Fatal("FATAL", err)
		return
	}

	appCache[ckey_storage] = path
	log.Printf("INFO storage '%s'.\n", path)
}

func (s *server) run() {
	appCache[ckey_port] = s.port

	addr := net.JoinHostPort(defaultHost, s.port)
	log.Printf("INFO server is listening on %s...\n", addr)
	log.Fatal(http.ListenAndServe(addr, nil))
}
