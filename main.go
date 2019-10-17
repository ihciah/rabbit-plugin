package main

import (
	"github.com/ihciah/rabbit-tcp/client"
	"github.com/ihciah/rabbit-tcp/logger"
	"github.com/ihciah/rabbit-tcp/tunnel"
	"log"
	"os"
	"strconv"
)

const (
	DefaultTunnelCount = 6
)

func main() {
	opts, err := parseEnv()
	if err != nil {
		log.Println("Cannot parse ENV")
		os.Exit(1)
	}
	if mode, ok := opts.Get("mode"); ok && mode == "s" {
		log.Println("Start in server mode")
		startServer(opts)
	} else {
		log.Println("Start in client mode")
		startClient(opts)
	}
}

func startServer(opts Args) {
	log.Println("Server mode is not implemented yet.")
}

func startClient(opts Args) {
	remoteHost, _ := opts.Get("remoteAddr")
	remotePort, _ := opts.Get("remotePort")
	localHost, _ := opts.Get("localAddr")
	localPort, _ := opts.Get("localPort")

	serviceAddr, ok := opts.Get("serviceAddr") // ip:port
	if !ok {
		log.Println("Cannot get serviceAddr")
		os.Exit(1)
	}
	password, ok := opts.Get("password")
	if !ok {
		log.Println("Cannot get password")
		os.Exit(1)
	}

	tunnelCount := DefaultTunnelCount
	tunnelN, ok := opts.Get("tunnelN")
	if ok {
		tunnelInt, err := strconv.Atoi(tunnelN)
		if err != nil {
			log.Println("Cannot parse tunnelN")
			os.Exit(1)
		}
		tunnelCount = tunnelInt
	}

	remoteAddr := remoteHost + ":" + remotePort
	localAddr := localHost + ":" + localPort

	cipher, _ := tunnel.NewAEADCipher("CHACHA20-IETF-POLY1305", nil, password)
	logger.LEVEL = 2
	c := client.NewClient(tunnelCount, remoteAddr, cipher)
	err := c.ServeForward(localAddr, serviceAddr)
	log.Println(err)
}
