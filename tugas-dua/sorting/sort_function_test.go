package sorting

import (
	"testing"
)

func TestGetMapAverage(t *testing.T) {
	type args struct {
		values map[string]float32
	}
	tests := []struct {
		name string
		args args
		want float32
	}{
		// TODO: Add test cases.
		{
			name: "Get Average Value", args: args{},
			want: 81.1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetMapAverage(tt.args.values); got != tt.want {
				t.Errorf("GetMapAverage() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetMapMinValue(t *testing.T) {
	type args struct {
		values map[string]float32
	}
	tests := []struct {
		name string
		args args
		want float32
	}{
		// TODO: Add test cases.
		{name: "Get Minimal Value",
			args: args{}, want: 69},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetMapMinValue(tt.args.values); got != tt.want {
				t.Errorf("GetMapMinValue() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetMapMaxValue(t *testing.T) {
	type args struct {
		values map[string]float32
	}
	tests := []struct {
		name string
		args args
		want float32
	}{
		// TODO: Add test cases.
		{name: "Get Maximal Value",
			args: args{}, want: 90},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetMapMaxValue(tt.args.values); got != tt.want {
				t.Errorf("GetMapMaxValue() = %v, want %v", got, tt.want)
			}
		})
	}
}
