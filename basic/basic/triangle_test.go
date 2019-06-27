package main

import (
	"testing"
)

func TestTriangle(t *testing.T) {

	tests := []struct{
		a, b, c int
	}{
		{3,4,5},
		{6,8,10},
		{1,2,3},
		{30000,40000,50000},
	}

	for _, tt := range tests{
		actual := calcTriangle(tt.a, tt.b)
		if actual != tt.c{
			/**
				用t.Errorf()输出错误信息
			 */
			t.Errorf("%d^2 + %d^2 != %d^2", tt.a, tt.b, tt.c)
		}
	}

}