package main

import (
	"encoding/xml"
	"fmt"
	"os/exec"
	"strings"
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

func ActivePortScanCommand(instance *ActionInstance) {
	Log(fmt.Sprintf("Running active port scan on %s.", instance.Node.Name))

	go func() {
		cmd := exec.Command("nmap", "-T4", "-p-", "-oX", "-", "--stats-every", "1s", "45.33.32.156")
		var out outstream
		cmd.Stdout = out
		if err := cmd.Start(); err != nil {
			Log("Error while running nmap")
		}
		cmd.Wait()

		outputBytes, _ := cmd.Output()
		instance.Output = string(outputBytes)
	}()

	ActivePortScanCompleted()
}

func ActivePortScanCompleted() {

}

type outstream struct{}

type TaskProgress struct {
	XMLName xml.Name `xml:"taskprogress"`
	Percent string   `xml:"percent,attr"`
}

func (out outstream) Write(p []byte) (int, error) {
	output := string(p)
	if strings.Contains(output, "<taskprogress") {
		var taskProgress TaskProgress
		err := xml.Unmarshal(p, &taskProgress)
		if err != nil {
			Log(err.Error())
		}
		Log(fmt.Sprintf("Nmap progress: %s", taskProgress.Percent))
	}
	return len(p), nil
}
