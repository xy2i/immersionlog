#!/bin/sh
set -a
. ./secrets.env

if [ "$1" = "run" ] || [ "$1" = "r" ] ; then
	go run -ldflags "-X main.botToken=$BOT_TOKEN -X main.guildID=$GUILD_ID" .
else 
	go build -ldflags "-X main.botToken=$BOT_TOKEN -X main.guildID=$GUILD_ID" -o ./bin/immersionlog
fi

