package task_24

import (
	"testing"
)

func Test_calculateDistance(t *testing.T) {
	type args struct {
		x1 int
		y1 int
		x2 int
		y2 int
	}

	tests := []struct {
		name string
		args args
		want float64
		err  error
	}{
		{
			name: "calculateDistance(): x1 = 5, y1 = 13, x2 = 17, y2: 10",
			args: args{
				x1: 5,
				y1: 13,
				x2: 17,
				y2: 10,
			},
			want: 12.369,
		},
		{
			name: "calculateDistance(): x1 = 5, y1 = 15, x2 = 20, y2: 3",
			args: args{
				x1: 5,
				y1: 15,
				x2: 20,
				y2: 3,
			},
			want: 19.209,
		},
		{
			name: "calculateDistance(): x1 = -1, y1 = -6, x2 = -2, y2: -5",
			args: args{
				x1: -1,
				y1: -6,
				x2: -2,
				y2: -5,
			},
			want: 1.414,
		},
		{
			name: "calculateDistance(): x1 = 0, y1 = 4, x2 = 3, y2: 0",
			args: args{
				x1: 0,
				y1: 4,
				x2: 3,
				y2: 0,
			},
			want: 5,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			res := calculateDistance(tt.args.x1, tt.args.y1, tt.args.x2, tt.args.y2)

			if res != tt.want {
				t.Errorf("task_24 calculateDistance() wanted - %f, got - %f", tt.want, res)
				return
			}
		})
	}
}
