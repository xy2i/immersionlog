#!/bin/sh
set -a
. ./secrets.env

if [ "$1" = "run" ] || [ "$1" = "r" ] ; then
	go run -ldflags "-X main.botToken=$BOT_TOKEN" .
else 
	go build -ldflags "-X main.botToken=$BOT_TOKEN" -o ./bin/immersionlog
fi

