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
