package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"google.golang.org/api/calendar/v3"
	"google.golang.org/api/option"
	"gopkg.in/yaml.v2"
)

var apiKey = os.Getenv("GOOGLE_CALENDAR_API_KEY")

// excludedSummaries は Google カレンダーの祝日イベントに混入する、祝日ではない記念日・行事の名称。
var excludedSummaries = map[string]bool{
	"銀行休業日": true,
	"節分":    true,
	"雛祭り":   true,
	"母の日":   true,
	"七夕":    true,
	"七五三":   true,
	"クリスマス": true,
	"大晦日":   true,
}

func main() {
	if err := handler(); err != nil {
		log.Fatal(err)
	}
}

func handler() error {
	// Google Calendar API で祝日判定する
	ctx := context.Background()
	client, err := calendar.NewService(ctx, option.WithAPIKey(apiKey))
	if err != nil {
		return fmt.Errorf("unable to retrieve Calendar client: %v", err)
	}

	// 実行日時を含む祝日を1件のみ取得する
	events, err := client.Events.List("ja.japanese#holiday@group.v.calendar.google.com").
		SingleEvents(true).OrderBy("startTime").Do()
	if err != nil {
		return err
	}

	for _, item := range events.Items {
		fmt.Println(item.Start.Date, item.Summary)
	}

	data := make(map[string]string, len(events.Items))
	for _, item := range events.Items {
		if excludedSummaries[item.Summary] {
			continue
		}
		data[item.Start.Date] = item.Summary
	}

	buf, err := yaml.Marshal(data)
	if err != nil {
		return err
	}

	// gosec:ignore:G306
	if err := os.WriteFile("shukujitsu.yml", buf, 0600); err != nil {
		return err
	}

	return nil
}
