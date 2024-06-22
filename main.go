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

type pomodoro struct {
	Label  *widget.Label
	Button *widget.Button
	Second int
	Timer  <-chan time.Time
	Ticker <-chan time.Time
	// TODO
	// counter と ticker をここに持つ
}

/*
func (t *timer) stop() {
	t.Button.SetText("Start!")
	// counter と ticker を止める
	if !t.counter.Stop() {
			<-t.counter.C // タイマーが既に満了している場合はチャネルを読み捨てる
	}
	ボタンのテキストと関数を元に戻す
}
*/

func (p *pomodoro) run() {
	p.Button.SetText("⏹️")

	p.Timer = time.After(time.Duration(p.Second) * time.Second)
	p.Ticker = time.Tick(1 * time.Second)

	start := time.Now()
	fmt.Printf("Waiting for %d seconds...\n", p.Second)
	for {
		select {
		case <-p.Timer:
			s := fmt.Sprintf("%d seconds have passed!", p.Second)
			fmt.Println(s)
			p.Label.SetText(s)
			playSound()
			return
		case <-p.Ticker:
			s := p.Second - int(time.Since(start).Seconds())
			fmt.Println(s)
			p.Label.SetText(strconv.Itoa(s))
		}
	}
}

func main() {
	p := &pomodoro{}
	p.Second = 25 * 60
	p.Label = widget.NewLabel(strconv.Itoa(p.Second))
	p.Button = widget.NewButton("▶ ", p.run)

	w := createWindow(p)
	w.ShowAndRun()
}

func createWindow(p *pomodoro) fyne.Window {
	a := app.New()
	w := a.NewWindow("🍅")
	w.Resize(fyne.NewSize(100, 20))

	w.SetContent(
		container.NewVBox(
			p.Label,
			p.Button,
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
