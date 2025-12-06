package main

import "testing"

func TestSumInvalidIDs(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  int64
	}{
		{
			name:  "case_1",
			input: "11-22",
			want:  33,
		},
		{
			name:  "case_2",
			input: "11-22,95-115",
			want:  132,
		},
		{
			name:  "case_3",
			input: "998-1012",
			want:  1010,
		},
		{
			name:  "case_4",
			input: "1188511880-1188511890",
			want:  1188511885,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := sumInvalidIDs(tt.input)
			if err != nil {
				t.Fatalf("got err %v, want err == nil", err)
			}
			if got != tt.want {
				t.Errorf("got %v, want %v", got, tt.want)
			}
		})
	}
}
