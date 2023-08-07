package task_20

import "testing"

func Test_revertWords(t *testing.T) {
	tests := []struct {
		name     string
		inputStr string
		want     string
	}{
		{
			name:     `revertWords(): "snow dog sun"`,
			inputStr: "snow dog sun",
			want:     "sun dog snow",
		},
		{
			name:     `revertWords(): "snow  dog sun"`,
			inputStr: "snow  dog sun",
			want:     "sun dog snow",
		},
		{
			name:     `revertWords(): " snow dog sun"`,
			inputStr: " snow dog sun",
			want:     "sun dog snow",
		},
		{
			name:     `revertWords(): "snow dog sun "`,
			inputStr: "snow dog sun ",
			want:     "sun dog snow",
		},
		{
			name:     `revertWords(): " snow dog sun "`,
			inputStr: " snow dog sun ",
			want:     "sun dog snow",
		},
		{
			name:     `revertWords(): "x y"`,
			inputStr: "x y",
			want:     "y x",
		},
		{
			name:     `revertWords(): "x"`,
			inputStr: "x",
			want:     "x",
		},
		{
			name:     `revertWords(): ""`,
			inputStr: "",
			want:     "",
		},
		{
			name:     `revertWords(): "111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111 2"`,
			inputStr: "",
			want:     "",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			res := revertWords(&tt.inputStr)

			if res != tt.want {
				t.Errorf("task_20 revertWords() wanted - %s, got - %s", tt.want, res)
				return
			}
		})
	}
}
