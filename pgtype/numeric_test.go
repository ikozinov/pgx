package pgtype_test

import (
	"math/big"
	"testing"

	"github.com/jackc/pgx/pgtype"
)

func TestNumericNormalize(t *testing.T) {
	testSuccessfulNormalize(t, []normalizeTest{
		{
			sql:   "select '1'::numeric",
			value: pgtype.GenericBinary{Bytes: nil, Status: pgtype.Present},
		},
		{
			sql:   "select '10.00'::numeric",
			value: pgtype.GenericBinary{Bytes: nil, Status: pgtype.Present},
		},
		{
			sql:   "select '1e-3'::numeric",
			value: pgtype.GenericBinary{Bytes: nil, Status: pgtype.Present},
		},
		{
			sql:   "select '-1'::numeric",
			value: pgtype.GenericBinary{Bytes: nil, Status: pgtype.Present},
		},
		{
			sql:   "select '10000'::numeric",
			value: pgtype.GenericBinary{Bytes: nil, Status: pgtype.Present},
		},
		{
			sql:   "select '3.14'::numeric",
			value: pgtype.GenericBinary{Bytes: nil, Status: pgtype.Present},
		},
	})
}

func TestNumericTranscode(t *testing.T) {
	testSuccessfulTranscode(t, "numeric", []interface{}{
		pgtype.Numeric{Int: *big.NewInt(0), Exp: 0, Status: pgtype.Present},
		pgtype.Numeric{Int: *big.NewInt(1), Exp: 0, Status: pgtype.Present},
		pgtype.Numeric{Int: *big.NewInt(314), Exp: -2, Status: pgtype.Present},
		pgtype.Numeric{Int: *big.NewInt(100), Exp: -2, Status: pgtype.Present},
		pgtype.Numeric{Int: *big.NewInt(123), Exp: -1500, Status: pgtype.Present},
		pgtype.Numeric{Status: pgtype.Null},
	})
}

// func TestNumericSet(t *testing.T) {
// 	successfulTests := []struct {
// 		source interface{}
// 		result pgtype.Numeric
// 	}{
// 		{source: float32(1), result: pgtype.Numeric{Float: 1, Status: pgtype.Present}},
// 		{source: float64(1), result: pgtype.Numeric{Float: 1, Status: pgtype.Present}},
// 		{source: int8(1), result: pgtype.Numeric{Float: 1, Status: pgtype.Present}},
// 		{source: int16(1), result: pgtype.Numeric{Float: 1, Status: pgtype.Present}},
// 		{source: int32(1), result: pgtype.Numeric{Float: 1, Status: pgtype.Present}},
// 		{source: int64(1), result: pgtype.Numeric{Float: 1, Status: pgtype.Present}},
// 		{source: int8(-1), result: pgtype.Numeric{Float: -1, Status: pgtype.Present}},
// 		{source: int16(-1), result: pgtype.Numeric{Float: -1, Status: pgtype.Present}},
// 		{source: int32(-1), result: pgtype.Numeric{Float: -1, Status: pgtype.Present}},
// 		{source: int64(-1), result: pgtype.Numeric{Float: -1, Status: pgtype.Present}},
// 		{source: uint8(1), result: pgtype.Numeric{Float: 1, Status: pgtype.Present}},
// 		{source: uint16(1), result: pgtype.Numeric{Float: 1, Status: pgtype.Present}},
// 		{source: uint32(1), result: pgtype.Numeric{Float: 1, Status: pgtype.Present}},
// 		{source: uint64(1), result: pgtype.Numeric{Float: 1, Status: pgtype.Present}},
// 		{source: "1", result: pgtype.Numeric{Float: 1, Status: pgtype.Present}},
// 		{source: _int8(1), result: pgtype.Numeric{Float: 1, Status: pgtype.Present}},
// 	}

// 	for i, tt := range successfulTests {
// 		var r pgtype.Numeric
// 		err := r.Set(tt.source)
// 		if err != nil {
// 			t.Errorf("%d: %v", i, err)
// 		}

// 		if r != tt.result {
// 			t.Errorf("%d: expected %v to convert to %v, but it was %v", i, tt.source, tt.result, r)
// 		}
// 	}
// }

// func TestNumericAssignTo(t *testing.T) {
// 	var i8 int8
// 	var i16 int16
// 	var i32 int32
// 	var i64 int64
// 	var i int
// 	var ui8 uint8
// 	var ui16 uint16
// 	var ui32 uint32
// 	var ui64 uint64
// 	var ui uint
// 	var pi8 *int8
// 	var _i8 _int8
// 	var _pi8 *_int8
// 	var f32 float32
// 	var f64 float64
// 	var pf32 *float32
// 	var pf64 *float64

// 	simpleTests := []struct {
// 		src      pgtype.Numeric
// 		dst      interface{}
// 		expected interface{}
// 	}{
// 		{src: pgtype.Numeric{Float: 42, Status: pgtype.Present}, dst: &f32, expected: float32(42)},
// 		{src: pgtype.Numeric{Float: 42, Status: pgtype.Present}, dst: &f64, expected: float64(42)},
// 		{src: pgtype.Numeric{Float: 42, Status: pgtype.Present}, dst: &i16, expected: int16(42)},
// 		{src: pgtype.Numeric{Float: 42, Status: pgtype.Present}, dst: &i32, expected: int32(42)},
// 		{src: pgtype.Numeric{Float: 42, Status: pgtype.Present}, dst: &i64, expected: int64(42)},
// 		{src: pgtype.Numeric{Float: 42, Status: pgtype.Present}, dst: &i, expected: int(42)},
// 		{src: pgtype.Numeric{Float: 42, Status: pgtype.Present}, dst: &ui8, expected: uint8(42)},
// 		{src: pgtype.Numeric{Float: 42, Status: pgtype.Present}, dst: &ui16, expected: uint16(42)},
// 		{src: pgtype.Numeric{Float: 42, Status: pgtype.Present}, dst: &ui32, expected: uint32(42)},
// 		{src: pgtype.Numeric{Float: 42, Status: pgtype.Present}, dst: &ui64, expected: uint64(42)},
// 		{src: pgtype.Numeric{Float: 42, Status: pgtype.Present}, dst: &ui, expected: uint(42)},
// 		{src: pgtype.Numeric{Float: 42, Status: pgtype.Present}, dst: &_i8, expected: _int8(42)},
// 		{src: pgtype.Numeric{Float: 0, Status: pgtype.Null}, dst: &pi8, expected: ((*int8)(nil))},
// 		{src: pgtype.Numeric{Float: 0, Status: pgtype.Null}, dst: &_pi8, expected: ((*_int8)(nil))},
// 	}

// 	for i, tt := range simpleTests {
// 		err := tt.src.AssignTo(tt.dst)
// 		if err != nil {
// 			t.Errorf("%d: %v", i, err)
// 		}

// 		if dst := reflect.ValueOf(tt.dst).Elem().Interface(); dst != tt.expected {
// 			t.Errorf("%d: expected %v to assign %v, but result was %v", i, tt.src, tt.expected, dst)
// 		}
// 	}

// 	pointerAllocTests := []struct {
// 		src      pgtype.Numeric
// 		dst      interface{}
// 		expected interface{}
// 	}{
// 		{src: pgtype.Numeric{Float: 42, Status: pgtype.Present}, dst: &pf32, expected: float32(42)},
// 		{src: pgtype.Numeric{Float: 42, Status: pgtype.Present}, dst: &pf64, expected: float64(42)},
// 	}

// 	for i, tt := range pointerAllocTests {
// 		err := tt.src.AssignTo(tt.dst)
// 		if err != nil {
// 			t.Errorf("%d: %v", i, err)
// 		}

// 		if dst := reflect.ValueOf(tt.dst).Elem().Elem().Interface(); dst != tt.expected {
// 			t.Errorf("%d: expected %v to assign %v, but result was %v", i, tt.src, tt.expected, dst)
// 		}
// 	}

// 	errorTests := []struct {
// 		src pgtype.Numeric
// 		dst interface{}
// 	}{
// 		{src: pgtype.Numeric{Float: 150, Status: pgtype.Present}, dst: &i8},
// 		{src: pgtype.Numeric{Float: 40000, Status: pgtype.Present}, dst: &i16},
// 		{src: pgtype.Numeric{Float: -1, Status: pgtype.Present}, dst: &ui8},
// 		{src: pgtype.Numeric{Float: -1, Status: pgtype.Present}, dst: &ui16},
// 		{src: pgtype.Numeric{Float: -1, Status: pgtype.Present}, dst: &ui32},
// 		{src: pgtype.Numeric{Float: -1, Status: pgtype.Present}, dst: &ui64},
// 		{src: pgtype.Numeric{Float: -1, Status: pgtype.Present}, dst: &ui},
// 		{src: pgtype.Numeric{Float: 0, Status: pgtype.Null}, dst: &i32},
// 	}

// 	for i, tt := range errorTests {
// 		err := tt.src.AssignTo(tt.dst)
// 		if err == nil {
// 			t.Errorf("%d: expected error but none was returned (%v -> %v)", i, tt.src, tt.dst)
// 		}
// 	}
// }
