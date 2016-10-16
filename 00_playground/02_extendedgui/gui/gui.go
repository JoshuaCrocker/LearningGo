package gui

import (
	"log"

	"github.com/andlabs/ui"
)

func Init() {
	err := ui.Main(buildGUI)

	if err != nil {
		log.Fatal(err)
	}
}

func buildGUI() {
	// name := ui.NewEntry()
	// button := ui.NewButton("Greet")
	// greeting := ui.NewLabel("")
	// box := ui.NewVerticalBox()
	//
	// box.Append(ui.NewLabel("Enter your name: "), false)
	// box.Append(name, false)
	// box.Append(button, false)
	// box.Append(greeting, false)
	//
	// window := ui.NewWindow("Hello", 200, 100, false)
	// window.SetChild(box)
	//
	// button.OnClicked(func(*ui.Button) {
	// 	greeting.SetText("Hello, " + name.Text() + "!")
	// })
	//
	// window.OnClosing(func(*ui.Window) bool {
	// 	ui.Quit()
	// 	return true
	// })
	//
	// window.Show()

	buildLoginScreen()
}
