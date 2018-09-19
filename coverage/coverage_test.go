package coverage

import "testing"

func TestConvert(t *testing.T) {
	tests := []struct {
		input  int
		output bool
	}{
		{
			0,
			false,
		},
		{
			1,
			true,
		},
		{
			2,
			true,
		},
		{
			3,
			false,
		},
	}

	for _, tt := range tests {
		b := convert(tt.input)
		if b != tt.output {
			t.Errorf("invalid result. want=%t, got=%t", tt.output, b)
		}
	}
}
