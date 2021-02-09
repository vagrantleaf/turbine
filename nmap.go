package main

import (
	"fmt"
	"os/exec"
)

func RegisterActivePortScan() {
	if IsActivePortScanAvailable() {
		action := Action{
			"Active port scan",
			"Comprehensive TCP port scan with service discovery.",
			"IP",
			ActivePortScanCommand,
		}
		RegisterAction(action)
	}
}

func IsActivePortScanAvailable() bool {
	_, err := exec.LookPath("nmap")
	if err != nil {
		Log("WARNING: Active Port Scan action not available: couldn't find 'nmap'.")
		return false
	}
	return true
}

func ActivePortScanCommand(node *Node) {
	Log(fmt.Sprintf("Running active port scan on %s.", node.Name))

	go func() {
		cmd := exec.Command("nmap", "-T4", "-p80,443", "-oX", "-", "--stats-every", "1s", "45.33.32.156")
		var out outstream
		cmd.Stdout = out
		if err := cmd.Start(); err != nil {
			Log("Error while running nmap")
		}
		cmd.Wait()
	}()

	ActivePortScanCompleted()
}

func ActivePortScanCompleted() {

}

type outstream struct{}

func (out outstream) Write(p []byte) (int, error) {
	Log(string(p))
	return len(p), nil
}
