package tcpserver

import (
	"MWrap/wrapper"
	"fmt"
	"net"
)

func ServerTCPStart(port string, w *wrapper.Wrapper) {
	PORT := ":" + port
	fmt.Printf("GoChat listening on port: %s\n", PORT)

	l, err := net.Listen("tcp", PORT)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer l.Close()

	for {
		c, err := l.Accept()
		if err != nil {
			fmt.Println(err)
			return
		}
		go HandleNewMsg(c, w)

	}
}
