package log_test

import (
	"testing"
	"github.com/go-ozzo/ozzo-log"
	"strings"
)

func TestFilterAllow(t *testing.T) {
	tests := []struct {
		cats     []string
		cat      string
		expected bool
	}{
		{[]string{}, "", true},
		{[]string{}, "system", true},
		{[]string{"system"}, "", false},
		{[]string{"system"}, "system", true},
		{[]string{"system"}, "system.db", false},
		{[]string{"system.*"}, "", false},
		{[]string{"system.*"}, "system", false},
		{[]string{"system.*"}, "system.", true},
		{[]string{"system.*"}, "system.db", true},
	}
	for _, test := range tests {
		filter := log.Filter{MaxLevel: log.LevelDebug, Categories: test.cats}
		filter.Init()
		e := &log.Entry{Category: test.cat}
		if filter.Allow(e) != test.expected {
			t.Errorf("filter(%q).Allow(%q) = %v, expected %v", strings.Join(test.cats, ","), test.cat, filter.Allow(e), test.expected)
		}
	}
}