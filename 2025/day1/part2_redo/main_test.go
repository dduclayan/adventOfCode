package main

import (
	"testing"
)

func TestTimesAtZero(t *testing.T) {
	tests := []struct {
		name string
		path string
		want float64
	}{
		{
			name: "right_1000",
			path: "test_data/r1000.txt",
			want: 10,
		},
		{
			name: "left_1000",
			path: "test_data/l1000.txt",
			want: 10,
		},
		{
			name: "aoc_test_input",
			path: "test_data/test_input.txt",
			want: 6,
		},
		{
			name: "return1",
			path: "test_data/return1.txt",
			want: 1,
		},
		{
			name: "return2",
			path: "test_data/return2.txt",
			want: 2,
		},
		{
			name: "return3",
			path: "test_data/return3.txt",
			want: 3,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := timesAtZero(tt.path)
			if err != nil {
				t.Fatalf("timesAtZero() got = %v, want no error", err)
			}
			if got != tt.want {
				t.Errorf("timesAtZero() got = %v, want %v", got, tt.want)
			}

		})
	}

}
