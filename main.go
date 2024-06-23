package main

import (
	"fmt"
	"os"
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"github.com/faiface/beep/mp3"
	"github.com/faiface/beep/speaker"
)

type config struct {
	WindowWidth  float32
	WindowHeight float32
	WindowTitle  string
	TimerLength  int
	StartText    string
	StopText     string
}

var cfg = config{
	WindowWidth:  100,
	WindowHeight: 20,
	WindowTitle:  "🍅",
	TimerLength:  25 * 60,
	StartText:    "▶",
	StopText:     "⏹️",
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
	fmt.Printf("Waiting for %d seconds...\n", p.Second)

	go func() {
		for {
			select {
			case <-p.Timer.C:
				setText(
					p.Label,
					fmt.Sprintf("%d seconds have passed!", p.Second),
				)
				playSound()
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
