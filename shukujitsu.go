package shukujitsu

import (
	_ "embed" // Register a standard stuff
	"errors"
	"log"
	"time"

	"gopkg.in/yaml.v2"
)

//go:embed shukujitsu.yml
var b []byte

var holidays map[string]string

func init() {
	if err := yaml.Unmarshal(b, &holidays); err != nil {
		log.Fatal(err)
	}
}

// IsHoliday : 祝日判定
func IsHoliday(t time.Time) bool {
	_, ok := holidays[dateStr(t)]
	return ok
}

// HolidayName : 祝日の名前
func HolidayName(t time.Time) (string, error) {
	name, ok := holidays[dateStr(t)]
	if !ok {
		return name, errors.New("not a holiday")
	}
	return name, nil
}

// csv の日付が 2021/7/22 形式である為、フォーマットを合わせる
func dateStr(t time.Time) string {
	return t.Format("2006/1/2")
}
