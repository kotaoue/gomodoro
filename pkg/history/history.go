package history

import (
	"encoding/csv"
	"os"
)

func Append(start, end string) error {
	name := "history.csv"

	records := [][]string{{start, end}}
	if newFile(name) {
		records = append([][]string{{"start", "end"}}, records...)
	}

	f, err := os.OpenFile(name, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return err
	}
	defer f.Close()

	w := csv.NewWriter(f)
	defer w.Flush()

	return w.WriteAll(records)
}

func newFile(name string) bool {
	_, err := os.Stat(name)
	return os.IsNotExist(err)
}
