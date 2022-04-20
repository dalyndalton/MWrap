package main

import (
	"os"
	"os/signal"

	"github.com/dalyndalton/MWrap/wrapper"
	"github.com/dalyndalton/MWrap/ws"
)

const PATH string = "../Minecraft1.18.2"

var PORT = "7777"

func main() {

	// Spin up Minecraft server
	cmd := wrapper.JavaExecCmd(PATH, 1024, 1024)
	console := wrapper.NewConsole(cmd)
	w := wrapper.NewWrapper(console)

	// Start Web server
	go ws.SetupWebsocket(w)

	// Listen for ctrl + c
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)

	<-quit
	w.Stop()
}
