package wannabe_test

import (
	"testing"
	"time"

	"github.com/izumin5210/wannabe"
)

func TestBool(t *testing.T) {
	type Case struct {
		in  interface{}
		out bool
	}
	groups := []struct {
		test  string
		cases []Case
	}{
		{
			test: "bool",
			cases: []Case{
				// bool
				{in: true, out: true},
				{in: false, out: false},
			},
		},
		{
			test: "int",
			cases: []Case{
				{in: 1, out: true},
				{in: -1, out: true},
				{in: 0, out: false},
				// 8
				{in: int8(1), out: true},
				{in: int8(-1), out: true},
				{in: int8(0), out: false},
				// 16
				{in: int16(1), out: true},
				{in: int16(-1), out: true},
				{in: int16(0), out: false},
				// 32
				{in: int32(1), out: true},
				{in: int32(-1), out: true},
				{in: int32(0), out: false},
				// 64
				{in: int64(1), out: true},
				{in: int64(-1), out: true},
				{in: int64(0), out: false},
			},
		},
		{
			test: "uint",
			cases: []Case{
				{in: uint(1), out: true},
				{in: uint(0), out: false},
				// 8
				{in: uint8(1), out: true},
				{in: uint8(0), out: false},
				// 16
				{in: uint16(1), out: true},
				{in: uint16(0), out: false},
				// 32
				{in: uint32(1), out: true},
				{in: uint32(0), out: false},
				// 64
				{in: uint64(1), out: true},
				{in: uint64(0), out: false},
			},
		},
		{
			test: "float",
			cases: []Case{
				{in: 1.0, out: true},
				{in: -1.0, out: true},
				{in: 0.0, out: false},
				// 32
				{in: float32(1.0), out: true},
				{in: float32(-1.0), out: true},
				{in: float32(0.0), out: false},
				// 64
				{in: float64(1.0), out: true},
				{in: float64(-1.0), out: true},
				{in: float64(0.0), out: false},
			},
		},
		{
			test: "string",
			cases: []Case{
				// Truthy
				// https://github.com/prodis/wannabe_bool/blob/v0.7.1/spec/wannabe_bool/string_spec.rb#L4-L16
				{in: "1", out: true},
				{in: "1 ", out: true},
				{in: " 1", out: true},
				{in: " 1 ", out: true},
				{in: "t", out: true},
				{in: "t ", out: true},
				{in: " t", out: true},
				{in: " t ", out: true},
				{in: "T", out: true},
				{in: "T ", out: true},
				{in: " T", out: true},
				{in: " T ", out: true},
				{in: "true", out: true},
				{in: "true ", out: true},
				{in: " true", out: true},
				{in: " true ", out: true},
				{in: "TRUE", out: true},
				{in: "TRUE ", out: true},
				{in: " TRUE", out: true},
				{in: " TRUE ", out: true},
				{in: "on", out: true},
				{in: "on ", out: true},
				{in: " on", out: true},
				{in: " on ", out: true},
				{in: "ON", out: true},
				{in: "ON ", out: true},
				{in: " ON ", out: true},
				{in: " ON ", out: true},
				{in: "y", out: true},
				{in: "y ", out: true},
				{in: " y", out: true},
				{in: " y ", out: true},
				{in: "Y", out: true},
				{in: "Y ", out: true},
				{in: " Y", out: true},
				{in: " Y ", out: true},
				{in: "yes", out: true},
				{in: "yes ", out: true},
				{in: " yes", out: true},
				{in: " yes ", out: true},
				{in: "YES", out: true},
				{in: "YES ", out: true},
				{in: " YES", out: true},
				{in: " YES ", out: true},
				// Falsey
				// https://github.com/prodis/wannabe_bool/blob/v0.7.1/spec/wannabe_bool/string_spec.rb#L18-L30
				{in: "0", out: false},
				{in: "0 ", out: false},
				{in: " 0", out: false},
				{in: " 0 ", out: false},
				{in: "f", out: false},
				{in: "f ", out: false},
				{in: " f", out: false},
				{in: " f ", out: false},
				{in: "F", out: false},
				{in: "F ", out: false},
				{in: " F", out: false},
				{in: " F ", out: false},
				{in: "false", out: false},
				{in: "false ", out: false},
				{in: " false", out: false},
				{in: " false ", out: false},
				{in: "FALSE", out: false},
				{in: "FALSE ", out: false},
				{in: " FALSE", out: false},
				{in: " FALSE ", out: false},
				{in: "off", out: false},
				{in: "off ", out: false},
				{in: " off", out: false},
				{in: " off ", out: false},
				{in: "OFF", out: false},
				{in: "OFF ", out: false},
				{in: " OFF ", out: false},
				{in: " OFF ", out: false},
				{in: "n", out: false},
				{in: "n ", out: false},
				{in: " n", out: false},
				{in: " n ", out: false},
				{in: "N", out: false},
				{in: "N ", out: false},
				{in: " N", out: false},
				{in: " N ", out: false},
				{in: "no", out: false},
				{in: "no ", out: false},
				{in: " no", out: false},
				{in: " no ", out: false},
				{in: "NO", out: false},
				{in: "NO ", out: false},
				{in: " NO", out: false},
				{in: " NO ", out: false},
			},
		},
		{
			test: "nil",
			cases: []Case{
				{in: nil, out: false},
				{in: ([]int)(nil), out: false},
				{in: (map[string]int)(nil), out: false},
				{in: (*struct{})(nil), out: false},
			},
		},
		{
			test: "container",
			cases: []Case{
				{in: []int{}, out: false},
				{in: map[string]int{}, out: false},
				{in: struct{ Foo int }{}, out: false},
				{in: []int{1}, out: true},
				{in: map[string]int{"foo": 3}, out: true},
				{in: struct{ Foo int }{Foo: 5}, out: true},
			},
		},
		{
			test: "time.Duration",
			cases: []Case{
				{in: 1 * time.Second, out: true},
				{in: -1 * time.Second, out: true},
				{in: 0 * time.Second, out: false},
			},
		},
	}

	for _, tg := range groups {
		t.Run(tg.test, func(t *testing.T) {
			for _, tc := range tg.cases {
				if got, want := wannabe.Bool(tc.in), tc.out; got != want {
					t.Errorf("wannabe.Bool(%v) returned %t, want %t", tc.in, got, want)
				}
			}
		})
	}
}
