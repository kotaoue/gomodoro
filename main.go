package main

import (
	"fmt"
	"os"
	"strconv"
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"github.com/faiface/beep/mp3"
	"github.com/faiface/beep/speaker"
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
			playSound()
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

func playSound() error {
	f, err := os.Open("assets/expiry.mp3")
	if err != nil {
		return err
	}
	defer f.Close()

	streamer, format, err := mp3.Decode(f)
	if err != nil {
		return err
	}
	defer streamer.Close()

	speaker.Init(format.SampleRate, format.SampleRate.N(time.Second/10))

	speaker.Play(streamer)

	time.Sleep(format.SampleRate.D(streamer.Len()))
	return nil
}
