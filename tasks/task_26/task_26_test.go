package task_26

import (
	"testing"
)

func Test_checkUniqueSymbols(t *testing.T) {
	tests := []struct {
		name     string
		inputStr string
		want     bool
	}{
		{
			name:     `checkUniqueSymbols(): "abcd"`,
			inputStr: "abcd",
			want:     true,
		},
		{
			name:     `checkUniqueSymbols(): "abCdefAaf"`,
			inputStr: "abCdefAaf",
			want:     false,
		},
		{
			name:     `checkUniqueSymbols(): "aabcd"`,
			inputStr: "aabcd",
			want:     false,
		},
		{
			name:     `checkUniqueSymbols(): "aAbcd"`,
			inputStr: "aAbcd",
			want:     false,
		},
		{
			name:     `checkUniqueSymbols(): "a"`,
			inputStr: "a",
			want:     true,
		},
		{
			name:     `checkUniqueSymbols(): ""`,
			inputStr: "",
			want:     true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			res := checkUniqueSymbols(tt.inputStr)

			if res != tt.want {
				t.Errorf("task_26 checkUniqueSymbols() wanted - %t, got - %t", tt.want, res)
				return
			}
		})
	}
}
