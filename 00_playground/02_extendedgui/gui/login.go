package gui

import "github.com/andlabs/ui"

var windowLogin *ui.Window

func buildLoginScreen() {
	// TODO Title Label

	lblUsername := ui.NewLabel("Username")
	txtUsername := ui.NewEntry()

	lblPassword := ui.NewLabel("Password")
	txtPassword := ui.NewEntry() // TODO passwords

	btnLogin := ui.NewButton("Login")

	boxContainer := ui.NewVerticalBox()
	boxUsername := ui.NewVerticalBox()
	boxPassword := ui.NewVerticalBox()
	boxControls := ui.NewVerticalBox()

	boxUsername.Append(lblUsername, false)
	boxUsername.Append(txtUsername, false)
	boxUsername.SetPadded(true)

	boxPassword.Append(lblPassword, false)
	boxPassword.Append(txtPassword, false)
	boxPassword.SetPadded(true)

	boxControls.Append(btnLogin, false)
	boxControls.SetPadded(true)

	boxContainer.Append(boxUsername, false)
	boxContainer.Append(boxPassword, false)
	boxContainer.Append(boxControls, false)
	boxContainer.SetPadded(true)

	windowLogin = ui.NewWindow("Login", 300, 400, false)
	windowLogin.SetChild(boxContainer)

	btnLogin.OnClicked(func(*ui.Button) {
		destroyLoginScreen()
		buildDashboardScreen()
	})

	handleClose(windowLogin)

	windowLogin.Show()
}

func destroyLoginScreen() {
	windowLogin.Destroy()
}
