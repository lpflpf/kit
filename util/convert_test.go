package util

import (
	"fmt"
	"math"
	"testing"
)

func TestString(t *testing.T) {
	str := "test"
	units := []struct {
		input    interface{}
		expected string
	}{
		{input: int64(10), expected: "10"},
		{input: uint8(10), expected: "10"},
		{input: 10.1, expected: "10.1"},
		{input: float32(10.1), expected: "10.1"},
		{input: float64(10.1), expected: "10.1"},
		{input: []byte{'1', '0', '.', '9'}, expected: "10.9"},
		{input: "10", expected: "10"},
		{input: true, expected: "true"},
		{input: false, expected: "false"},
		{input: 1 + 1i, expected: "(1+1i)"},
		{input: &str, expected: "test"},
	}

	for _, unit := range units {
		output := String(unit.input)
		if output != unit.expected {
			fmt.Printf("input: %v, type:%T, expected: %v, output: %v\n", unit.input, unit.input, unit.expected, output)
			t.Fail()
		}
	}
}

func TestFloat64(t *testing.T) {
	units := []struct {
		input    interface{}
		expected float64
	}{
		{input: int64(10), expected: 10},
		{input: uint8(10), expected: 10},
		{input: 10.1, expected: 10.1},
		{input: float32(10.1), expected: 10.1},
		{input: float64(10.1), expected: 10.1},
		{input: []byte{'1', '0', '.', '9'}, expected: 10.9},
		{input: "10", expected: 10},
		{input: true, expected: 1},
		{input: false, expected: 0},
	}

	for _, unit := range units {
		output := Float64(unit.input, 2)
		if math.Abs(output-unit.expected) > 0.000001 {
			fmt.Printf("input: %v, type:%T, expected: %v, output: %v\n", unit.input, unit.input, unit.expected, output)
			t.Fail()
		}
	}

	units = []struct {
		input    interface{}
		expected float64
	}{
		{input: 1 + 1i, expected: 0},
		{input: struct{}{}, expected: 0},
	}
	for _, unit := range units {
		func(input interface{}, expected float64) {
			defer func() {
				if err := recover(); err == nil {
					fmt.Printf("input: %v, type:%T, expected: %v, need expected, but not\n", input, input, expected)
					t.Fail()
				}
			}()
			output := Float64(input, 2)
			fmt.Println(output)
			if math.Abs(output-expected) > 0.000001 {
				fmt.Printf("input: %v, type:%T, expected: %v, output: %v\n", input, input, expected, output)
				t.Fail()
			}
		}(unit.input, unit.expected)
	}

}

func TestInt64(t *testing.T) {
	units := []struct {
		input    interface{}
		expected int64
	}{
		{input: int64(10), expected: 10},
		{input: uint8(10), expected: 10},
		{input: 10.1, expected: 10},
		{input: float32(10.9), expected: 11},
		{input: float64(10.9), expected: 11},
		{input: []byte{'1', '0'}, expected: 10},
		{input: "10", expected: 10},
		{input: true, expected: 1},
		{input: false, expected: 0},
	}

	for _, unit := range units {
		output := Int64(unit.input)
		if output != unit.expected {
			fmt.Printf("input: %v, type:%T, expected: %v, output: %v\n", unit.input, unit.input, unit.expected, output)
			t.Fail()
		}
	}

	units = []struct {
		input    interface{}
		expected int64
	}{
		{input: 1 + 1i, expected: 0},
		{input: struct{}{}, expected: 0},
	}
	for _, unit := range units {
		func(input interface{}, expected int64) {
			defer func() {
				if err := recover(); err == nil {
					t.Fail()
				}

			}()
			output := Int64(input)
			fmt.Println(output)
			if output != expected {
				fmt.Printf("input: %v, type:%T, expected: %v, output: %v\n", input, input, expected, output)
				t.Fail()
			}
		}(unit.input, unit.expected)
	}

}
