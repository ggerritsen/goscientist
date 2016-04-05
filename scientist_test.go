package goscientist

import (
	"testing"
	"reflect"
)

func TestNewExperiment(t *testing.T) {
	f1 := func() int {return 1}
	f2 := func() int {return 2}

	exp, err := NewExperiment(f1, f2)
	if err != nil {
		t.Error(err)
	}

	if got, want := exp.currentFunc, reflect.ValueOf(f1); got != want {
		t.Errorf("got %+v want %+v", got, want)
	}
	if got, want := exp.improvedFunc, reflect.ValueOf(f1); got!= want {
		t.Errorf("got %+v want %+v", got, want)
	}
}
