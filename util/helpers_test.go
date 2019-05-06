package util

import (
	"reflect"
	"testing"
)

func TestSetHeights(t *testing.T) {
	type args struct {
		minHeight uint8
		maxHeight uint8
		increment uint8
	}
	tests := []struct {
		name        string
		args        args
		wantHeights []float64
		wantErr     bool
	}{
		{"Positive heights", args{minHeight: 150, maxHeight: 200, increment: 5}, []float64{150, 155, 160, 165, 170, 175, 180, 185, 190, 195, 200}, false},
		{"Min greater than Max", args{minHeight: 200, maxHeight: 150, increment: 5}, nil, true},
		{"Min equal zero", args{minHeight: 0, maxHeight: 150, increment: 5}, nil, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotHeights, err := SetHeights(tt.args.minHeight, tt.args.maxHeight, tt.args.increment)
			if (err != nil) != tt.wantErr {
				t.Errorf("SetHeights() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotHeights, tt.wantHeights) {
				t.Errorf("SetHeights() = %v, want %v", gotHeights, tt.wantHeights)
			}
		})
	}
}
