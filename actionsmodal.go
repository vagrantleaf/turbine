package main

import (
	"github.com/rivo/tview"
)

type ActionsModal struct {
	flex *tview.Flex
	list *tview.List
}

func CreateActionsModal() *ActionsModal {
	modal := new(ActionsModal)
	modal.list = tview.NewList()
	modal.list.SetBorder(true).SetTitle("Actions")
	width := 40
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
		modal.list.AddItem(action.Name, action.Command, '.', nil)
	}

	pages.ShowPage("actionsModal")	
}

