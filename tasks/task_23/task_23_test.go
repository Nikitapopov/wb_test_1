package task_23

import (
	"errors"
	"reflect"
	"testing"
)

func Test_deleteElem(t *testing.T) {
	type args struct {
		slc *[]int
		num int
	}

	errNumberExceedsSliceRange := errors.New("number exceeds slice range")

	tests := []struct {
		name string
		args args
		want []int
		err  error
	}{
		{
			name: `deleteElem(): slice = {3, 4, 5, 6, 7, 8}, num = 4`,
			args: args{
				slc: &[]int{3, 4, 5, 6, 7, 8},
				num: 4,
			},
			want: []int{3, 4, 5, 7, 8},
		},
		{
			name: `deleteElem(): slice = {3, 4, 5, 6, 7, 8}, num = 1`,
			args: args{
				slc: &[]int{3, 4, 5, 6, 7, 8},
				num: 1,
			},
			want: []int{4, 5, 6, 7, 8},
		},
		{
			name: `deleteElem(): slice = {3, 4, 5, 6, 7, 8}, num = 6`,
			args: args{
				slc: &[]int{3, 4, 5, 6, 7, 8},
				num: 6,
			},
			want: []int{3, 4, 5, 6, 7},
		},
		{
			name: `deleteElem(): slice = {}, num = 1`,
			args: args{
				slc: &[]int{},
				num: 1,
			},
			err: errNumberExceedsSliceRange,
		},
		{
			name: `deleteElem(): slice = {3, 4, 5, 6, 7, 8}, num = 0`,
			args: args{
				slc: &[]int{3, 4, 5, 6, 7, 8},
				num: 0,
			},
			err: errNumberExceedsSliceRange,
		},
		{
			name: `deleteElem(): slice = {3, 4, 5, 6, 7, 8}, num = 7`,
			args: args{
				slc: &[]int{3, 4, 5, 6, 7, 8},
				num: 7,
			},
			err: errNumberExceedsSliceRange,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := deleteElem(tt.args.slc, tt.args.num)

			if err != nil {
				if err.Error() != tt.err.Error() {
					t.Errorf("task_23 deleteElem() wanted error - %v, got - %v", tt.err, err)
					return
				}
			} else {
				if !reflect.DeepEqual(*tt.args.slc, tt.want) {
					t.Errorf("task_23 deleteElem() wanted - %v, got - %v", tt.want, *tt.args.slc)
					return
				}
			}
		})
	}
}
