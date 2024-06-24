package main

import (
	"github.com/kotaoue/gomodoro/pkg/pomodoro"
)

func main() {
	p := pomodoro.NewPomodoro(
		pomodoro.Config{
			WindowWidth:  100,
			WindowHeight: 20,
			WindowTitle:  "🍅",
			TimerLength:  25 * 60,
			StartText:    "▶",
			StopText:     "⏹️",
			StopSound:    "assets/expiry.mp3",
		},
	)

	w := p.CreateWindow()
	w.ShowAndRun()
}
