package file

import "testing"

func TestReadWithCustomStrategy(t *testing.T) {
	funcName := "Read"
	want := "test"
	params := []any{
		"custom",
		"./test/test.txt",
		customReader(0),
	}
	got, err := Read(
		params[0].(string),
		params[1].(string),
		params[2].(customReader),
	)
	if nil != err {
		t.Error(err)
	}
	if want != got {
		testError(t, funcName, want, got, params...)
	}
}

func TestReadWithBuiltinNumberStrategy(t *testing.T) {
	result, err := Read(
		"number",
		"./test/test.txt",
	)
	if nil != err {
		t.Error(err)
	}

	segmentFloat64, ok := result.([]float64)
	if !ok {
		t.Error("result should be []float64")
	}

	if 1 != segmentFloat64[0] {
		t.Error("result should be 1, got", segmentFloat64[0])
	}
	if 2 != segmentFloat64[1] {
		t.Error("result should be 2, got", segmentFloat64[1])
	}
	if 3 != segmentFloat64[2] {
		t.Error("result should be 3, got", segmentFloat64[2])
	}
	if 3 != segmentFloat64[3] {
		t.Error("result should be 3, got ", segmentFloat64[3])
	}
	if 2 != segmentFloat64[4] {
		t.Error("result should be 2, got ", segmentFloat64[4])
	}
	if 1 != segmentFloat64[5] {
		t.Error("result should be 1, got", segmentFloat64[5])
	}
}

func testError(t *testing.T, funcName string, want any, got any, params ...any) {
	t.Errorf(
		"%s(%v) = %#v, want %#v",
		funcName,
		params,
		got,
		want,
	)
}
