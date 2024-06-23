package main

import (
	"fmt"
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"github.com/kotaoue/gomodoro/pkg/sound"
)

type config struct {
	WindowWidth  float32
	WindowHeight float32
	WindowTitle  string
	TimerLength  int
	StartText    string
	StopText     string
	StopSound    string
}

var cfg = config{
	WindowWidth:  100,
	WindowHeight: 20,
	WindowTitle:  "🍅",
	TimerLength:  25 * 60,
	StartText:    "▶",
	StopText:     "⏹️",
	StopSound:    "assets/expiry.mp3",
}

type pomodoro struct {
	Label   *widget.Label
	Button  *widget.Button
	Second  int
	Ticker  *time.Ticker
	Timer   *time.Timer
	Stopper chan struct{}
}

func (p *pomodoro) start() {
	p.Button.SetText(cfg.StopText)
	p.Button.OnTapped = p.stop

	p.Ticker = time.NewTicker(1 * time.Second)
	p.Timer = time.NewTimer(time.Duration(p.Second) * time.Second)
	p.Stopper = make(chan struct{})

	start := time.Now()
	fmt.Println("START")

	go func() {
		for {
			select {
			case <-p.Timer.C:
				setText(
					p.Label,
					"FINISH",
				)
				sound.Play(cfg.StopSound)
				p.stop()
				return
			case <-p.Ticker.C:
				setText(
					p.Label,
					secToMD(p.Second-int(time.Since(start).Seconds())),
				)
			case <-p.Stopper:
				p.Timer.Stop()
				p.Ticker.Stop()
				return
			}
		}
	}()
}

func (p *pomodoro) stop() {
	p.Second = cfg.TimerLength
	setText(p.Label, secToMD(p.Second))
	p.Label.SetText(secToMD(p.Second))

	p.Button.SetText(cfg.StartText)
	p.Button.OnTapped = p.start
	if p.Stopper != nil {
		close(p.Stopper)
	}
}

func main() {
	p := &pomodoro{}
	p.Second = cfg.TimerLength
	p.Second = 10
	p.Label = widget.NewLabel(secToMD(p.Second))
	p.Button = widget.NewButton(cfg.StartText, p.start)

	w := createWindow(p)
	w.ShowAndRun()
}

func createWindow(p *pomodoro) fyne.Window {
	a := app.New()
	w := a.NewWindow(cfg.WindowTitle)
	w.Resize(fyne.NewSize(cfg.WindowWidth, cfg.WindowHeight))

	w.SetContent(
		container.NewVBox(
			container.NewCenter(p.Label),
			p.Button,
		))

	return w
}

func setText(l *widget.Label, s string) {
	fmt.Println(s)
	l.SetText(s)
}

func secToMD(sec int) string {
	return fmt.Sprintf("%02d:%02d", sec/60, sec%60)
}
