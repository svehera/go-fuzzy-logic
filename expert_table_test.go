package fuzzy

import (
	"reflect"
	"testing"
)

func Test_membershipFunc(t *testing.T) {
	type args struct {
		table [][]float64
	}
	tests := []struct {
		name string
		args args
		want []float64
	}{
		{"Test", args{[][]float64{{0, 1, 1, 1, 0}, {0, 0, 1, 1, 0}}}, []float64{0, 0.5, 1, 1, 0}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := membershipFunc(tt.args.table); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("membershipFunc() = %v, want %v", got, tt.want)
			}
		})
	}
}
