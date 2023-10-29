package main

import (
	"testing"
	"time"

	"github.com/7adidaz/internal/assert"
)

func TestHumanDate(t *testing.T) {
	tests := []struct {
		name string
		tm   time.Time
		want string
	}{
		{
			name: "UTC",
			tm:   time.Date(2022, 3, 17, 10, 15, 0, 0, time.UTC),
			want: "17 Mar 2022 at 10:15",
		},
		// {
		// 	name: "Empty",
		// 	tm:   time.Time{},
		// 	want: "",
		// },
		{
			name: "CET",
			tm:   time.Date(2022, 3, 17, 10, 15, 0, 0, time.FixedZone("CET", 1*60*60)),
			want: "17 Mar 2022 at 10:15",
		},
	}
	// Loop over the test cases.
	for _, tt := range tests {
		// t.Run() runs a sub-test for each test case.

		// first params is the name of the test.
		// second is anony func to containing the test.
		t.Run(tt.name, func(t *testing.T) {
			hd := humanDate(tt.tm)

			assert.Equal(t, hd, tt.want)
		})
	}

}
