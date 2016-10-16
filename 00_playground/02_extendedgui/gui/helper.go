package gui

import "github.com/andlabs/ui"

func handleClose(w *ui.Window) {
	w.OnClosing(func(*ui.Window) bool {
		ui.Quit()
		return true
	})
}
