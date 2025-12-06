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
			want:  243,
		},
		{
			name:  "case_3",
			input: "998-1012",
			want:  2009,
		},
		{
			name:  "case_4",
			input: "1188511880-1188511890",
			want:  1188511885,
		},
		{
			name:  "case_5",
			input: "824824821-824824827",
			want:  824824824,
		},
		{
			name:  "case_6",
			input: "1212121212-1212121212",
			want:  1212121212,
		},
		{
			name:  "case_7",
			input: "8336728382-8336728384",
			want:  0,
		},
		{
			name:  "case_8",
			input: "2121212118-2121212124",
			want:  2121212121,
		},
		{
			name:  "case_9",
			input: "824824821-824824827",
			want:  824824824,
		},
		{
			name:  "case_10",
			input: "1-22",
			want:  33,
		},
		{
			name:  "case_11",
			input: "11111-11112",
			want:  11111,
		},
		{
			name:  "case_12",
			input: "12-13",
			want:  0,
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
