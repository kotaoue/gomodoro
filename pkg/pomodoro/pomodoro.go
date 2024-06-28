package pomodoro

import (
	"fmt"
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"github.com/kotaoue/gomodoro/pkg/history"
	"github.com/kotaoue/gomodoro/pkg/sound"
)

type Config struct {
	WindowWidth  float32
	WindowHeight float32
	WindowTitle  string
	TimerLength  int
	StartText    string
	StopText     string
	StopSound    string
}

type Pomodoro struct {
	Label   *widget.Label
	Button  *widget.Button
	Second  int
	Ticker  *time.Ticker
	Timer   *time.Timer
	Stopper chan struct{}
	Config  Config
}

func NewPomodoro(cfg Config) *Pomodoro {
	p := &Pomodoro{}
	p.Second = cfg.TimerLength
	p.Label = widget.NewLabel(p.secToMD(p.Second))
	p.Button = widget.NewButton(cfg.StartText, p.start)
	p.Config = cfg

	return p
}

func (p *Pomodoro) CreateWindow() fyne.Window {
	a := app.New()
	w := a.NewWindow(p.Config.WindowTitle)
	w.Resize(fyne.NewSize(p.Config.WindowWidth, p.Config.WindowHeight))

	w.SetContent(
		container.NewVBox(
			container.NewCenter(p.Label),
			p.Button,
		))

	return w
}

func (p *Pomodoro) start() {
	p.Button.SetText(p.Config.StopText)
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
				history.Append(start.Format("2006-01-02 15:04:05"), time.Now().Format("2006-01-02 15:04:05"))
				p.setText("FINISH")
				sound.Play(p.Config.StopSound)
				p.stop()
				return
			case <-p.Ticker.C:
				p.setText(p.secToMD(p.Second - int(time.Since(start).Seconds())))
			case <-p.Stopper:
				p.Timer.Stop()
				p.Ticker.Stop()
				return
			}
		}
	}()
}

func (p *Pomodoro) stop() {
	p.Second = p.Config.TimerLength
	p.setText(p.secToMD(p.Second))

	p.Button.SetText(p.Config.StartText)
	p.Button.OnTapped = p.start
	if p.Stopper != nil {
		close(p.Stopper)
	}
}

func (p *Pomodoro) setText(s string) {
	fmt.Println(s)
	p.Label.SetText(s)
}

func (p *Pomodoro) secToMD(sec int) string {
	return fmt.Sprintf("%02d:%02d", sec/60, sec%60)
}
