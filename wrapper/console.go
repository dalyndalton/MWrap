package wrapper

import (
	"bufio"
	"fmt"
	"os/exec"
)

type Console struct {
	cmd    *exec.Cmd
	stdout *bufio.Reader
	stdin  *bufio.Writer
}

/*
Returns a pointer to a console object bound to
the specified command
*/

func NewConsole(cmd *exec.Cmd) (c *Console) {
	c = &Console{
		cmd: cmd,
	}

	stdout, _ := cmd.StdoutPipe()
	c.stdout = bufio.NewReader(stdout)

	stdin, _ := c.cmd.StdinPipe()
	c.stdin = bufio.NewWriter(stdin)

	return

}

/* 	We dont need classes, we can just specify the type that
each function can run on
*/
func (c *Console) WriteCmd(cmd string) error {
	wrappedCmd := fmt.Sprintf("%s\r\n", cmd)
	_, err := c.stdin.WriteString(wrappedCmd)
	if err != nil {
		return err
	}
	return c.stdin.Flush()
}

func (c *Console) ReadLine() (string, error) {
	return c.stdout.ReadString('\n')
}

func (c *Console) Start() error {
	return c.cmd.Start()
}

func (c *Console) Kill() {
	c.cmd.Process.Kill()
}
