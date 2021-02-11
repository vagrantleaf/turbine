package main

import (
	"github.com/rivo/tview"
)

func CreatePortsView() {
	//portsView = tview.NewBox()
	//portsView.SetBorder(true)
	//portsView.SetTitle("Ports")
	//portsView.SetInputCapture(RibbonInputHandler)
	//AddRibbonEntry(portsView, "Ports", nil, nil)

	portsView = tview.NewTable().SetBorders(true)
	portsView.SetBorder(true).SetTitle("Ports")

	portsView.SetCell(0, 0, tview.NewTableCell("Port"))
	portsView.SetCell(0, 1, tview.NewTableCell("Protocol"))
	portsView.SetCell(0, 2, tview.NewTableCell("State"))
	portsView.SetCell(0, 3, tview.NewTableCell("Service"))
	portsView.SetCell(0, 4, tview.NewTableCell("Version"))

	portsView.SetCell(1, 0, tview.NewTableCell("80"))
	portsView.SetCell(1, 1, tview.NewTableCell("TCP"))
	portsView.SetCell(1, 2, tview.NewTableCell("Open"))
	portsView.SetCell(1, 3, tview.NewTableCell("HTTP"))
	portsView.SetCell(1, 4, tview.NewTableCell("Apache"))

	AddRibbonEntry(portsView.Box, "Ports", nil, nil)
}
