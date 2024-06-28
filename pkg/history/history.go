package history

import (
	"os"
)

func Append(s string) error {
	f, err := os.OpenFile("pomodoro.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return err
	}
	defer f.Close()

	// テキストを書き込む
	_, err = f.WriteString(s)
	return err
}
