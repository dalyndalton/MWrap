package wrapper

import (
	"fmt"
	"os/exec"
)

func JavaExecCmd(serverPath string, initialHeapSize, maxHeapSize int) *exec.Cmd {
	initialHeapFlag := fmt.Sprintf("-Xms%dM", initialHeapSize)
	maxHeapFlag := fmt.Sprintf("-Xmx%dM", maxHeapSize)
	cmd := exec.Command("java", initialHeapFlag, maxHeapFlag, "-jar", "server.jar", "nogui")
	cmd.Dir = serverPath
	return cmd
}
