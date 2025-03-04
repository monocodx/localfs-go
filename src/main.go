// Copyright 2024 The localFS Authors.
// Use of this source code is governed by a GPL
// license that can be found in the LICENSE file.

package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"runtime"
)

func commandLineFlag() (port *string) {
	//override flag usage
	flag.Usage = func() {
		fmt.Printf("LocalFS %s, a portable web-based local file server.\n", appBuild)
		fmt.Printf("Usage: %s [options]\n", os.Args[0])
		fmt.Printf("options\n")
		fmt.Printf("  %-20s server port to use (default %s).\n", "-p, --port", defaultPort)
		fmt.Printf("  %-20s use '/var/tmp' for the temporary directory instead of\n"+
			"%-23s'tmpfs' to handle large file uploads (linux systems only).\n", "    --no-tmpfs", "")
		fmt.Printf("  %-20s print this list and exit.\n", "-h, --help")
		fmt.Printf("  %-20s print the version and exit.\n", "-v, --version")
		fmt.Printf("\n")
	}

	// server port
	port = flag.String("port", defaultPort, "server port to use")
	flag.StringVar(port, "p", defaultPort, "server port to use")
	// build version
	version := flag.Bool("version", false, "print the version and exit")
	flag.BoolVar(version, "v", false, "print the version and exit")
	// temporary directory
	tmpfs := flag.Bool("no-tmpfs", false, "use /var/tmp for the temporary directory"+
		" instead of tmpfs to handle large file uploads (linux systems only)")
	flag.Parse()

	// handle flag -v
	if *version {
		fmt.Printf("build#%s\n", appBuild)
		os.Exit(0)
	}

	// handle flag --no-tmpfs
	if *tmpfs {
		if runtime.GOOS == "linux" {
			log.Printf("INFO set '/var/tmp' as temporary directory.\n")
			os.Setenv("TMPDIR", "/var/tmp")
		} else {
			log.Printf("ERROR '--no-tmpfs' flag is for linux systems only.\n")
			flag.Usage()
			os.Exit(0)
		}
	}

	return
}

func main() {
	port := commandLineFlag()

	s := httpServer(*port)
	s.initStorage()
	s.initRoutes()
	s.run()
}
