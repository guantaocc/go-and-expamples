package main

import (
	"reflect"
	"testing"
)

func TestSplit(t *testing.T) {
	got := Split("我爱你", "爱")
	want := []string{"我", "你"}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("want:%v, got:%v", want, got)
	}
}
