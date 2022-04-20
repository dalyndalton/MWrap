package wrapper

import (
	"fmt"
	"io"
)

type Wrapper struct {
	console *Console
	state   *Event
}

func NewWrapper(c *Console) *Wrapper {
	return &Wrapper{
		console: c,
		state:   nil,
	}
}

func (w *Wrapper) Start() {
	w.console.cmd.Start()
	w.DisplayLogs()
}

func (w *Wrapper) Stop() {
	w.console.cmd.Process.Kill()
}

func (w *Wrapper) DisplayLogs() {
	go func() {
		for {
			line, err := w.console.ReadLine()
			if err == io.EOF {
				continue
			}

			log, state := LogParser(line)

			if state != nil {
				fmt.Println("Server state: ", *state)
				w.state = state
				go onStateChange()

			}

			fmt.Printf("%s : %s\n", log.timestamp, log.output)
		}
	}()
}

func (w *Wrapper) SendMessage(msg string) {
	w.console.WriteCmd(msg)
}

func onStateChange() {

}
