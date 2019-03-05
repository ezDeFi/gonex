// Copyright 2017 The go-ethereum Authors
// This file is part of go-ethereum.
//
// go-ethereum is free software: you can redistribute it and/or modify
// it under the terms of the GNU General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// go-ethereum is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
// GNU General Public License for more details.
//
// You should have received a copy of the GNU General Public License
// along with go-ethereum. If not, see <http://www.gnu.org/licenses/>.

package main

import (
	"fmt"
	"time"

	"github.com/ethereum/go-ethereum/log"
)

// deployExplorer creates a new block explorer based on some user input.
func (w *wizard) deployExplorer() {
	// Select the server to interact with
	server := w.selectServer()
	if server == "" {
		return
	}
	client := w.servers[server]

	// Retrieve any active node configurations from the server
	infos, err := checkExplorer(client, w.network)
	if err != nil {
		infos = &explorerInfos{
			httpRPCURL: "http://localhost:8545",
			wsRPCURL:   "ws://localhost:8546",
			webPort:    80,
			webHost:    client.server,
		}
	}
	existed := err == nil

	// Figure out which port to listen on
	fmt.Println()
	fmt.Printf("Which port should the explorer listen on? (default = %d)\n", infos.webPort)
	infos.webPort = w.readDefaultInt(infos.webPort)

	// Figure which virtual-host to deploy ethstats on
	if infos.webHost, err = w.ensureVirtualHost(client, infos.webPort, infos.webHost); err != nil {
		log.Error("Failed to decide on explorer host", "err", err)
		return
	}

	// Figure out which JSONRPC_HTTP_URL for explorer to connect
	fmt.Println()
	fmt.Printf("Which JSON RPC http url? (default = %s)\n", infos.httpRPCURL)
	infos.httpRPCURL = w.readDefaultString(infos.httpRPCURL)

	// Figure out which JSONRPC_HTTP_URL for explorer to connect
	fmt.Println()
	fmt.Printf("Which JSON RPC websocket url? (default = %s)\n", infos.wsRPCURL)
	infos.wsRPCURL = w.readDefaultString(infos.wsRPCURL)

	// Figure out where the user wants to store the persistent data for backend database
	fmt.Println()
	if infos.dbdir == "" {
		fmt.Printf("Where should postgres data be stored on the remote machine?\n")
		infos.dbdir = w.readString()
	} else {
		fmt.Printf("Where should postgres data be stored on the remote machine? (default = %s)\n", infos.dbdir)
		infos.dbdir = w.readDefaultString(infos.dbdir)
	}

	// Try to deploy the explorer on the host
	nocache := false
	if existed {
		fmt.Println()
		fmt.Printf("Should the explorer be built from scratch (y/n)? (default = no)\n")
		nocache = w.readDefaultYesNo(false)
	}
	if out, err := deployExplorer(client, w.network, infos, nocache); err != nil {
		log.Error("Failed to deploy explorer container", "err", err)
		if len(out) > 0 {
			fmt.Printf("%s\n", out)
		}
		return
	}
	// All ok, run a network scan to pick any changes up
	log.Info("Waiting for node to finish booting")
	time.Sleep(3 * time.Second)

	w.networkStats()
}
