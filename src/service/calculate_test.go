package service_test

import (
	"golang-devcontainer/src/service"
	"testing"
)

func TestCalculate_Service(t *testing.T) {
	type args struct {
		loopNumber int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"正常系_1を渡した場合、1が返却される",
			args{1},
			1,
		},
		{"正常系_10を渡した場合、55が返却される",
			args{10},
			55,
		},
		{"正常系_100を渡した場合、5050が返却される",
			args{100},
			5050,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := service.NewCalculate()
			if got := c.Service(tt.args.loopNumber); got != tt.want {
				t.Errorf("Calculate.Service() = %v, want %v", got, tt.want)
			}
		})
	}
}
