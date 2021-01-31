package main

import (
	"fmt"
	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

type HotkeyEntry struct {
	view     *tview.Box
	name     string
	callback hotkeyFn
	hotkey   rune
}

type hotkeyFn func()

var hotkeys []HotkeyEntry

func CreateHotkeysView() {
	hotkeysView = tview.NewTextView().SetDynamicColors(true)

	AddHotkeyEntry("Quit", 'x', func(){ app.Stop() })
}

func AddHotkeyEntry(name string, hotkey rune, callback hotkeyFn) {
	var _, existingEntry = FindHotkeyEntry(name)
	if existingEntry != nil {
		return
	}

	entry := HotkeyEntry{hotkeysView.Box, name, callback, hotkey}
	hotkeys = append(hotkeys, entry)

	UpdateHotkeys()
}

func RemoveHotkeyEntry(name string) {
	var index, existingEntry = FindHotkeyEntry(name)
	if existingEntry == nil {
		return
	}

	hotkeys = append(hotkeys[:index], hotkeys[index+1:]...)

	UpdateHotkeys()
}

func FindHotkeyEntry(name string) (int, *HotkeyEntry) {
	for index, entry := range hotkeys {
		if entry.name == name {
			return index, &entry
		}
	}

	return -1, nil
}

func HotkeysInputHandler(event *tcell.EventKey) *tcell.EventKey {
	for _, entry := range hotkeys {
		if event.Rune() == entry.hotkey {
			entry.callback()
			return nil
		}
	}

	return event
}

func UpdateHotkeys() {
	hotkeysView.Clear()
	l := len(hotkeys)
	for i, entry := range hotkeys {
		fmt.Fprintf(hotkeysView, "([::u]%c[::-]) %s", entry.hotkey, entry.name)

		if i < l-1 {
			fmt.Fprintf(hotkeysView, " | ")
		}
	}
}
