package shukujitsu

import (
	_ "embed"
	"testing"
	"time"
)

var (
	notHoliday, _ = time.Parse("2006-01-02", "2021-07-19")
	holiday, _    = time.Parse("2006-01-02", "2021-07-22")
)

func TestIsHoliday(t *testing.T) {
	type args struct {
		t time.Time
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "there is a holiday",
			args: args{
				holiday,
			},
			want: true,
		},
		{
			name: "there is not a holiday",
			args: args{
				notHoliday,
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsHoliday(tt.args.t); got != tt.want {
				t.Errorf("IsHoliday() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestHolidayName(t *testing.T) {
	type args struct {
		t time.Time
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{
			name: "there is a holiday",
			args: args{
				holiday,
			},
			want:    "海の日",
			wantErr: false,
		},
		{
			name: "there is not a holiday",
			args: args{
				notHoliday,
			},
			want:    "",
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := HolidayName(tt.args.t)
			if (err != nil) != tt.wantErr {
				t.Errorf("HolidayName() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("HolidayName() = %v, want %v", got, tt.want)
			}
		})
	}
}
