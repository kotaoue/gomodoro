package main

import (
	"fmt"
	"time"

	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func main() {
	runTimer(25)

	// createWindow()
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
	myApp := app.New()
	myWindow := myApp.NewWindow("Hello")

	myWindow.SetContent(container.NewVBox(
		widget.NewLabel("Hello Fyne!"),
		widget.NewButton("Quit", func() {
			myApp.Quit()
		}),
	))

	myWindow.ShowAndRun()
}
