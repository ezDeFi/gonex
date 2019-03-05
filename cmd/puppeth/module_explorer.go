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
	"bytes"
	"fmt"
	"html/template"
	"math/rand"
	"path/filepath"
	"strconv"
)

// explorerDockerfile is the Dockerfile required to run a block explorer.
var explorerDockerfile = `
FROM nextyio/blockscout:latest

ENV \
  ETHEREUM_JSONRPC_HTTP_URL={{.HttpRPCURL}} \
	ETHEREUM_JSONRPC_WS_URL={{.WsRPCURL}}

RUN \
  echo '/usr/local/bin/docker-entrypoint.sh postgres &' >> explorer.sh && \
  echo 'sleep 5' >> explorer.sh && \
	echo 'mix do ecto.drop --force, ecto.create, ecto.migrate' >> explorer.sh && \
  echo 'mix phx.server' >> explorer.sh

ENTRYPOINT ["/bin/sh", "explorer.sh"]
`

// explorerComposefile is the docker-compose.yml file required to deploy and
// maintain a block explorer.
var explorerComposefile = `
version: '2'
services:
  explorer:
    build: .
    image: {{.Network}}/explorer
    container_name: {{.Network}}_explorer_1
    ports:
      - {{if not .VHost}}"{{.WebPort}}:4000"{{end}}
    environment:
      - NETWORK={{.NetworkName}}{{if .VHost}}
      - VIRTUAL_HOST={{.VHost}}
      - VIRTUAL_PORT=4000{{end}}
    volumes:
      - {{.DBDir}}:/var/lib/postgresql/data
    logging:
      driver: "json-file"
      options:
        max-size: "1m"
        max-file: "10"
    restart: always
`

// deployExplorer deploys a new block explorer container to a remote machine via
// SSH, docker and docker-compose. If an instance with the specified network name
// already exists there, it will be overwritten!
func deployExplorer(client *sshClient, network string, config *explorerInfos, nocache bool) ([]byte, error) {
	// Generate the content to upload to the server
	workdir := fmt.Sprintf("%d", rand.Int63())
	files := make(map[string][]byte)

	dockerfile := new(bytes.Buffer)
	template.Must(template.New("").Parse(explorerDockerfile)).Execute(dockerfile, map[string]interface{}{
		"HttpRPCURL": config.httpRPCURL,
		"WsRPCURL":   config.wsRPCURL,
	})
	files[filepath.Join(workdir, "Dockerfile")] = dockerfile.Bytes()

	composefile := new(bytes.Buffer)
	template.Must(template.New("").Parse(explorerComposefile)).Execute(composefile, map[string]interface{}{
		"NetworkName": network,
		"VHost":       config.webHost,
		"DBDir":       config.dbdir,
		"Network":     network,
		"WebPort":     config.webPort,
	})
	files[filepath.Join(workdir, "docker-compose.yaml")] = composefile.Bytes()

	// Upload the deployment files to the remote server (and clean up afterwards)
	if out, err := client.Upload(files); err != nil {
		return out, err
	}
	defer client.Run("rm -rf " + workdir)

	// Build and deploy the boot or seal node service
	if nocache {
		return nil, client.Stream(fmt.Sprintf("cd %s && docker-compose -p %s build --pull --no-cache && docker-compose -p %s up -d --force-recreate --timeout 60", workdir, network, network))
	}
	return nil, client.Stream(fmt.Sprintf("cd %s && docker-compose -p %s up -d --build --force-recreate --timeout 60", workdir, network))
}

// explorerInfos is returned from a block explorer status check to allow reporting
// various configuration parameters.
type explorerInfos struct {
	httpRPCURL string
	wsRPCURL   string
	dbdir      string
	webHost    string
	webPort    int
}

// Report converts the typed struct into a plain string->string map, containing
// most - but not all - fields for reporting to the user.
func (info *explorerInfos) Report() map[string]string {
	report := map[string]string{
		"Website address ":       info.webHost,
		"Website listener port ": strconv.Itoa(info.webPort),
	}
	return report
}

// checkExplorer does a health-check against a block explorer server to verify
// whether it's running, and if yes, whether it's responsive.
func checkExplorer(client *sshClient, network string) (*explorerInfos, error) {
	// Inspect a possible explorer container on the host
	infos, err := inspectContainer(client, fmt.Sprintf("%s_explorer_1", network))
	if err != nil {
		return nil, err
	}
	if !infos.running {
		return nil, ErrServiceOffline
	}
	// Resolve the port from the host, or the reverse proxy
	webPort := infos.portmap["4000/tcp"]
	if webPort == 0 {
		if proxy, _ := checkNginx(client, network); proxy != nil {
			webPort = proxy.port
		}
	}
	if webPort == 0 {
		return nil, ErrNotExposed
	}
	// Resolve the host from the reverse-proxy and the config values
	host := infos.envvars["VIRTUAL_HOST"]
	if host == "" {
		host = client.server
	}

	// Assemble and return the useful infos
	stats := &explorerInfos{
		dbdir:   infos.volumes["/var/lib/postgresql/data"],
		webHost: host,
		webPort: webPort,
	}
	return stats, nil
}
