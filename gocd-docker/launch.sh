#!/bin/bash
docker run --rm -d --name gocd-server -p 8153:8153 gocd/gocd-server

docker run --rm -d --link gocd-server:go-server  gocd/gocd-agent
docker run --rm -d --link gocd-server:go-server  gocd/gocd-agent
