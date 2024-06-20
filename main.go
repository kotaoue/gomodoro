package main

import (
	"fmt"
	"strconv"
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
	t.Button.SetText("Running...")

	counter := time.After(time.Duration(t.Second) * time.Second)
	ticker := time.Tick(1 * time.Second)

	start := time.Now()
	fmt.Printf("Waiting for %d seconds...\n", t.Second)
	for {
		select {
		case <-counter:
			s := fmt.Sprintf("%d seconds have passed!", t.Second)
			fmt.Println(s)
			t.Label.SetText(s)
			return
		case <-ticker:
			s := t.Second - int(time.Since(start).Seconds())
			fmt.Println(s)
			t.Label.SetText(strconv.Itoa(s))
		}
	}
}

func main() {
	t := &timer{}
	t.Second = 25 * 60
	t.Label = widget.NewLabel(strconv.Itoa(t.Second))
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
