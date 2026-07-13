package main

import (
	"bytes"
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strings"
	"time"

	"golang.org/x/text/encoding/japanese"
	"golang.org/x/text/transform"
	"gopkg.in/yaml.v2"
)

// csvURL : 内閣府が公表している国民の祝日一覧の CSV
const csvURL = "https://www8.cao.go.jp/chosei/shukujitsu/syukujitsu.csv"

func main() {
	if err := handler(); err != nil {
		log.Fatal(err)
	}
}

func handler() error {
	rows, err := fetchHolidays()
	if err != nil {
		return err
	}

	data := make(map[string]string, len(rows))
	for _, row := range rows {
		date, parseErr := time.Parse("2006/1/2", strings.TrimSpace(row[0]))
		if parseErr != nil {
			return fmt.Errorf("failed to parse date %q: %w", row[0], parseErr)
		}
		data[date.Format("2006-01-02")] = row[1]
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

// fetchHolidays : 内閣府 CSV (Shift_JIS) を取得し、ヘッダー行を除いたレコードを返す
func fetchHolidays() ([][]string, error) {
	resp, err := http.Get(csvURL)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}

	body, err := io.ReadAll(transform.NewReader(resp.Body, japanese.ShiftJIS.NewDecoder()))
	if err != nil {
		return nil, err
	}

	r := csv.NewReader(bytes.NewReader(body))
	rows, err := r.ReadAll()
	if err != nil {
		return nil, err
	}
	if len(rows) < 2 {
		return nil, fmt.Errorf("unexpected csv format: %d rows", len(rows))
	}

	return rows[1:], nil
}
