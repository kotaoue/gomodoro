package main

import (
	"fmt"
	"time"

	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func main() {
	i := 25
	timer := time.After(time.Duration(10) * time.Second)
	ticker := time.Tick(1 * time.Second)

	fmt.Printf("Waiting for %d seconds...\n", i)
	for {
		select {
		case <-timer:
			fmt.Printf("%d seconds have passed!\n", i)
			return
		case t := <-ticker:
			fmt.Println("Current time:", t.Format("2006-01-02 15:04:05"))
		}
	}

	// createWindow()
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
