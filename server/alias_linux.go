// +build linux

package server

import "os/exec"

func installNetworkAlias() ([]byte, error) {
	return exec.Command("ip", "addr", "add", "169.254.169.254/24", "dev", "lo", "label", "lo:0").CombinedOutput()
}
