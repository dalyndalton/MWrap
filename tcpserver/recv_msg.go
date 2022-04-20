package tcpserver

import (
	"MWrap/wrapper"
	"bufio"
	"fmt"
	"net"
	"strings"
)

func HandleNewMsg(c net.Conn, w *wrapper.Wrapper) {

	netData, err := bufio.NewReader(c).ReadString('\n')
	if err != nil {
		fmt.Println(err)
		return
	}
	msg := strings.TrimSpace(string(netData))

	w.SendMessage(msg)
	c.Close()
}
