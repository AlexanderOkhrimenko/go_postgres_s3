#!/usr/bin/env bash

docker-compose -f ./docker-compose.yml  up -d --scale worker=1
