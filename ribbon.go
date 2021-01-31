package main

import (
	"fmt"
	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

type RibbonEntry struct {
	view  *tview.Box
	name  string
	index int
	onFocusedCallback OnFocusedCallback
	onUnfocusedCallback OnUnfocusedCallback
}

type OnFocusedCallback func()
type OnUnfocusedCallback func()

var ribbon []RibbonEntry
var ribbonInitialised bool
var selectedRibbonEntry RibbonEntry

func CreateRibbonView() {
	ribbonInitialised = false
	ribbonView = tview.NewTextView().SetDynamicColors(true)
}

func AddRibbonEntry(view *tview.Box, name string, onFocused OnFocusedCallback, onUnfocused OnUnfocusedCallback) {
	entry := RibbonEntry{view, name, len(ribbon), onFocused, onUnfocused}
	ribbon = append(ribbon, entry)

	if ribbonInitialised == false {
		selectedRibbonEntry = entry
		ribbonInitialised = true

		if onFocused != nil {
			onFocused()
		}
	}

	UpdateRibbon()
}

func RibbonInputHandler(event *tcell.EventKey) *tcell.EventKey {
	switch event.Rune() {
	case 'q':
		PreviousRibbonEntry()
		return nil
	case 'e':
		NextRibbonEntry()
		return nil
	}
	return HotkeysInputHandler(event)
}

func NextRibbonEntry() {
	if selectedRibbonEntry.index < len(ribbon)-1 {
		if selectedRibbonEntry.onUnfocusedCallback != nil {
			selectedRibbonEntry.onUnfocusedCallback()
		}

		selectedRibbonEntry = ribbon[selectedRibbonEntry.index+1]

		if selectedRibbonEntry.onFocusedCallback != nil {
			selectedRibbonEntry.onFocusedCallback()
		}

		UpdateRibbon()
		app.SetFocus(selectedRibbonEntry.view)
	}
}

func PreviousRibbonEntry() {
	if selectedRibbonEntry.index > 0 {
		if selectedRibbonEntry.onUnfocusedCallback != nil {
			selectedRibbonEntry.onUnfocusedCallback()
		}

		selectedRibbonEntry = ribbon[selectedRibbonEntry.index-1]

		if selectedRibbonEntry.onFocusedCallback != nil {
			selectedRibbonEntry.onFocusedCallback()
		}

		UpdateRibbon()
		app.SetFocus(selectedRibbonEntry.view)
	}
}

func UpdateRibbon() {
	ribbonView.Clear()
	l := len(ribbon)
	for i, entry := range ribbon {
		if entry.index == selectedRibbonEntry.index {
			fmt.Fprintf(ribbonView, "[::u]%s[::-]", entry.name)
		} else {
			fmt.Fprintf(ribbonView, "%s", entry.name)
		}

		if i < l-1 {
			fmt.Fprintf(ribbonView, " >> ")
		}
	}
}
