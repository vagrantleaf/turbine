package main

import (
	"fmt"
	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
	"os"
	"os/user"
	"path/filepath"
)

var (
	workspace       string
	app             *tview.Application
	pages           *tview.Pages
	actionsModal    *ActionsModal
	ribbonView      *tview.TextView
	logView         *tview.TextView
	ipsView         *tview.List
	portsView       *tview.Table
	hotkeysView     *tview.TextView
	inputField      *tview.InputField
	inputFieldAdded bool
	flexH           *tview.Flex
	flexV           *tview.Flex
)

type OnInputFieldClosedCallback func(tcell.Key)

func CreateLogView() {
	logView = tview.NewTextView().
		SetDynamicColors(true).
		SetRegions(true).
		SetChangedFunc(func() {
			app.Draw()
		})
	logView.SetBorder(true)
	logView.SetTitle("Log")
}

func CreateInputField() {
	inputFieldAdded = false
	inputField = tview.NewInputField().
		SetLabel("input: ").
		SetFieldWidth(15)
}

func ShowInputField(label string, fieldWidth int, fn OnInputFieldClosedCallback) {
	if inputFieldAdded == false {
		flexV.RemoveItem(hotkeysView)
		flexV.AddItem(inputField, 1, 1, false)
		inputField.SetLabel(label)
		inputField.SetDoneFunc(fn)
		app.SetFocus(inputField)
		inputFieldAdded = true
	}
}

func CloseInputField(key tcell.Key) {
	if inputFieldAdded == true {
		flexV.RemoveItem(inputField)
		flexV.AddItem(hotkeysView, 1, 1, false)
		inputFieldAdded = false
	}
}

func Log(text string) {
	fmt.Fprintf(logView, "%s\n", text)
}

func CreateLayout() *tview.Flex {
	CreateLogView()
	CreateHotkeysView()
	CreateRibbonView()
	CreateIPsView()
	CreatePortsView()
	CreateInputField()
	UpdateRibbon()

	flexV = tview.NewFlex().SetDirection(tview.FlexRow)
	flexH = tview.NewFlex()

	flexH.AddItem(ipsView, 21, 1, false)
	flexH.AddItem(portsView, 0, 2, false)

	flexV.AddItem(ribbonView, 1, 1, false)
	flexV.AddItem(flexH, 0, 1, false)
	flexV.AddItem(logView, 10, 1, false)

	flexV.AddItem(hotkeysView, 1, 1, false)

	return flexV
}

func LoadProject() {
	DeserialiseIPs()
}

func SaveProject() {
	SerialiseIPs()
}

func main() {
	if len(os.Args) == 1 {
		fmt.Printf("Usage: turbine <workspace_name>")
		return
	}

	user, userErr := user.Current()
	if userErr != nil {
		return
	}

	// TODO: Check that the workspace name is valid for the OS.
	workspacePath := filepath.Join(user.HomeDir, ".turbine")
	os.Mkdir(workspacePath, 0700)

	workspace = filepath.Join(workspacePath, os.Args[1])
	os.Mkdir(workspace, 0700)

	app = tview.NewApplication()

	flex := CreateLayout()
	actionsModal = CreateActionsModal()

	pages = tview.NewPages().
		AddPage("background", flex, true, true).
		AddPage("actionsModal", actionsModal.flex, true, false) // Actions modal starts hidden.

	LoadProject()
	RegisterActions()

	if err := app.SetRoot(pages, true).SetFocus(ipsView).Run(); err != nil {
		fmt.Printf("Error running application: %s\n", err)
	}

	SaveProject()
}
