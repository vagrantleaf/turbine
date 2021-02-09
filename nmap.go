package main

import (
	"os/exec"
)

func RegisterActivePortScanAction() {
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

func ActivePortScanCommand() {
	Log("Running active port scan")
	ActivePortScanCompleted()
}

func ActivePortScanCompleted() {

}

