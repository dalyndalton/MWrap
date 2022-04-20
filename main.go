package main

import (
	"MWrap/tcpserver"
	"MWrap/wrapper"
	"fmt"
	"os"
	"os/signal"
)

const PATH string = "../Minecraft1.18.2"

var PORT = "7777"

func main() {

	// Spin up Minecraft server
	cmd := wrapper.JavaExecCmd(PATH, 1024, 1024)
	console := wrapper.NewConsole(cmd)
	w := wrapper.NewWrapper(console)
	w.Start()
	fmt.Println("Go server up and running ! ✅")

	// Start Web server
	go tcpserver.ServerTCPStart(PORT, w)

	// Listen for ctrl + c
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)

	<-quit
	w.Stop()
}
