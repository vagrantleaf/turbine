package main

import (
	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

type ActionsModal struct {
	flex          *tview.Flex
	list          *tview.List
}

func CreateActionsModal() *ActionsModal {
	modal := new(ActionsModal)
	modal.list = tview.NewList()
	modal.list.SetBorder(true).SetTitle("Actions")
	modal.list.SetInputCapture(ActionsModalInputHandler)
	width := 60
	height := 10
	modal.flex = tview.NewFlex().
		AddItem(nil, 0, 1, false).
		AddItem(tview.NewFlex().SetDirection(tview.FlexRow).
			AddItem(nil, 0, 1, false).
			AddItem(modal.list, height, 1, false).
			AddItem(nil, 0, 1, false), width, 1, false).
		AddItem(nil, 0, 1, false)
	return modal
}

func (modal *ActionsModal) Show(node Node) {
	modal.list.Clear()
	for _, action := range(actions) {
		modal.list.AddItem(action.Name, action.Description, '.', nil)
	}

	pages.ShowPage("actionsModal")	
	app.SetFocus(modal.list)
}

func (modal *ActionsModal) Hide() {
	pages.HidePage("actionsModal")
	SelectRibbonEntry()
}

func ActionsModalInputHandler(event *tcell.EventKey) *tcell.EventKey {
	if event.Key() == tcell.KeyEnter {
		actionsModal.Hide()
		return nil
	} else if event.Key() == tcell.KeyEscape {
		actionsModal.Hide()
		return nil
	}

	return event
}

