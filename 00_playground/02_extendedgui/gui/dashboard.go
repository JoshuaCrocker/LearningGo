package gui

import "github.com/andlabs/ui"

var windowDash *ui.Window

func buildDashboardScreen() {
	lblTest := ui.NewLabel("Test")

	btnLogout := ui.NewButton("Logout")

	boxContainer := ui.NewVerticalBox()

	boxTest := ui.NewVerticalBox()
	boxControls := ui.NewVerticalBox()

	boxTest.Append(lblTest, false)

	boxControls.Append(btnLogout, false)
	boxControls.SetPadded(true)

	boxContainer.Append(boxTest, false)
	boxContainer.Append(boxControls, false)
	boxContainer.SetPadded(true)

	windowDash = ui.NewWindow("Test", 600, 800, false)
	windowDash.SetChild(boxContainer)

	btnLogout.OnClicked(func(*ui.Button) {
		destroyDashboardScreen()
		buildLoginScreen()
	})

	handleClose(windowDash)

	windowDash.Show()
}

func destroyDashboardScreen() {
	windowDash.Destroy()
}
