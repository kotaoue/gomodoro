package main

import (
	"fmt"
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

type timer struct {
	Label  *widget.Label
	Button *widget.Button
	Second int
}

func (t *timer) run() {
	counter := time.After(time.Duration(10) * time.Second)
	ticker := time.Tick(1 * time.Second)

	fmt.Printf("Waiting for %d seconds...\n", t.Second)
	for {
		select {
		case <-counter:
			s := fmt.Sprintf("%d seconds have passed!", t.Second)
			fmt.Println(s)
			t.Label.SetText(s)
			return
		case tic := <-ticker:
			s := tic.Format("15:04:05")
			fmt.Println(s)
			t.Label.SetText(s)
		}
	}
}

func main() {
	// runTimer(25)
	t := &timer{}
	t.Second = 25
	t.Label = widget.NewLabel("")
	t.Button = widget.NewButton("Start!", t.run)

	w := createWindow(t)
	w.ShowAndRun()
}

func createWindow(t *timer) fyne.Window {
	a := app.New()
	w := a.NewWindow("Timer")
	w.Resize(fyne.NewSize(300, 20))

	w.SetContent(
		container.NewVBox(
			t.Label,
			t.Button,
		))

	return w
}
