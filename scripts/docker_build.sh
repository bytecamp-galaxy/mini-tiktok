#!/usr/bin/env bash
docker build -f docker/api-server/Dockerfile -t vgalaxy/api-server .
docker build -f docker/user-server/Dockerfile -t vgalaxy/user-server .