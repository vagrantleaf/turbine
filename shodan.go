package main

func RegisterPassivePortScan() {
	action := Action{
		"Passive port scan",
		"Use Shodan to retrieve port scan information.",
		"IP",
		nil,
	}
	RegisterAction(action)
}
