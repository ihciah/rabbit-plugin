# Rabbit TCP SIP003 Plugin

![BuildBadge](https://github.com/ihciah/rabbit-plugin/workflows/Build/badge.svg)

A shadowsocks plugin with [rabbit-tcp](https://github.com/ihciah/rabbit-tcp) support.

Currently only client mode is supported. And currently this project is very unstable.

## Usage

1. Put the binary file to the same directory of shadowsocks
2. Fill in `Plugin Options` with `serviceAddr=your.ssserver:ssport;password=mypassword;tunnelN=6`

## How to run a rabbit tcp server

#### Choice 1: Run binary
1. Download executable from [rabbit-tcp release page](https://github.com/ihciah/rabbit-tcp/releases)
2. Keep it running using `systemd` or `supervisor`, etc. 

(You may have to write configure files by your own)

#### Choice 2: Run with docker (Recommended)
1. Install docker and docker-compose if you don\'t have them
2. Download [docker-compose file](https://github.com/ihciah/rabbit-tcp/blob/master/docker-compose-server.yml)
3. Modify it and `docker-compose -f docker-compose-server.yml up -d`
