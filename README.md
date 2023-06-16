# immersionlog
Discord immersion logger


# Build Instructions

To build and run this, you need to create a secrets.env file, which will be sourced for two variables at BUILD time.

secrets.env should contain BOT_TOKEN which is the discord api token for your bot as well as GUILD_ID which is the ID for your guild.

```
BOT_TOKEN=
GUILD_ID=
```

To run the program

```
./build.sh r
```
OR 
```
./build.sh run
```

and to build the executable(THIS EXECUTABLE IS FOR DEPLOYMENT, IT WILL CONTAIN YOUR BOT TOKEN AND GUILD ID)

```
./build.sh
```
