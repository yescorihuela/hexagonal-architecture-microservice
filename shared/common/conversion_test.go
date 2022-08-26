package common

import (
	"reflect"
	"testing"
)

func TestGetSlicedUrls(t *testing.T) {
	tests := []struct {
		description string
		urls        string
		expected    []string
	}{
		{
			"Sliced urls",
			"https://go.dev/play/,https://go.dev/solutions/,https://pkg.go.dev/",
			[]string{"https://go.dev/play/", "https://go.dev/solutions/", "https://pkg.go.dev/"},
		},
		{
			"Sliced Url",
			"https://www.solotodo.cl/,",
			[]string{"https://www.solotodo.cl/"},
		},
		{
			"Empty urls",
			"",
			[]string{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.description, func(t *testing.T) {
			got := GetSlicedUrls(tt.urls)
			if !reflect.DeepEqual(got, tt.expected) {
				t.Errorf("expected %s, but got %s", tt.expected, got)
			}
		})
	}
}

func TestGetStringFromSlicedUrls(t *testing.T) {
	tests := []struct {
		description string
		slicedUrls  []string
		expected    string
	}{
		{
			"Sliced urls",
			[]string{"https://go.dev/play/", "https://go.dev/solutions/", "https://pkg.go.dev/"},
			"https://go.dev/play/,https://go.dev/solutions/,https://pkg.go.dev/",
		},
		{
			"Sliced Url",
			[]string{"https://www.solotodo.cl/"},
			"https://www.solotodo.cl/",
		},
		{
			"Empty urls",
			[]string{},
			"",
		},
	}
	for _, tt := range tests {
		t.Run(tt.description, func(t *testing.T) {
			got := GetStringFromSlicedUrls(tt.slicedUrls)
			if !reflect.DeepEqual(got, tt.expected) {
				t.Errorf("expected %s, but got %s", tt.expected, got)
			}
		})
	}
}
