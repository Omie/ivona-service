#!/bin/bash

export IVONA_ACCESSKEY=
export IVONA_SECRETKEY=
export IVONA_SERVICE_HOST=localhost
export IVONA_SERVICE_PORT=9575

rm ivona-service
go build
./ivona-service

