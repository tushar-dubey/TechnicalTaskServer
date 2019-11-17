package main

import (
	"TechnicalTaskServer/calculator/proto"
	"context"
	"testing"
)

func TestServer_SumIntegerTest(t *testing.T) {
	s := server{}

	// set up test cases
	tests := []struct{
		int1 int64
		int2 int64
		result int64
	} {
		{
			int1: 1,
			int2: 2,
			result: 3,
		},
		{
			int1: 3,
			int2: 4,
			result: 7,
		},
	}

	for _, tt := range tests {
		req := &proto.SumRequest{Int1: tt.int1, Int2:tt.int2,}
		resp, err := s.SumIntegers(context.Background(), req)
		if err != nil {
			t.Errorf("SumTest(%v, %v) got unexpected error", tt.int1, tt.int2)
		}
		if resp.Result != tt.result {
			t.Errorf("SumTest(%v,%v)=%v, wanted %v", tt.int1, tt.int2, resp.Result, tt.result)
		}
	}
}


func TestServer_SubtractIntegers(t *testing.T) {
	s := server{}

	// set up test cases
	tests := []struct{
		int1 int64
		int2 int64
		result int64
	} {
		{
			int1: 1,
			int2: 2,
			result: -1,
		},
		{
			int1: 3,
			int2: 4,
			result: -1,
		},
	}

	for _, tt := range tests {
		req := &proto.SubtractRequest{ Int1:tt.int1, Int2:tt.int2,}
		resp, err := s.SubtractIntegers(context.Background(), req)
		if err != nil {
			t.Errorf("SumTest(%v, %v) got unexpected error", tt.int1, tt.int2)
		}
		if resp.Result != tt.result {
			t.Errorf("SumTest(%v,%v)=%v, wanted %v", tt.int1, tt.int2, resp.Result, tt.result)
		}
	}
}