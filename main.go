package main

import (
	"fmt"
	"time"

	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func main() {
	// runTimer(25)
	createWindow()
}

func runTimer(i int) {
	timer := time.After(time.Duration(10) * time.Second)
	ticker := time.Tick(1 * time.Second)

	fmt.Printf("Waiting for %d seconds...\n", i)
	for {
		select {
		case <-timer:
			fmt.Printf("%d seconds have passed!\n", i)
			return
		case t := <-ticker:
			fmt.Println(t.Format("15:04:05"))
		}
	}
}

func createWindow() {
	app := app.New()
	window := app.NewWindow("Hello")

	window.SetContent(container.NewVBox(
		widget.NewLabel("Hello Fyne!"),
		widget.NewButton("Run", func() {
			runTimer(10)
		}),
		widget.NewButton("Quit", func() {
			app.Quit()
		}),
	))

	window.ShowAndRun()
}
