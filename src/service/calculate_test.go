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
		c    *service.Calculate
		args args
		want int
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &service.Calculate{}
			if got := c.Service(tt.args.loopNumber); got != tt.want {
				t.Errorf("Calculate.Service() = %v, want %v", got, tt.want)
			}
		})
	}
}
