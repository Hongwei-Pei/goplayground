package pipefilter

import (
	"testing"
)

func TestFilter(t *testing.T) {
	u := UselessFilter{}
	inp := "test"
	out, err := u.Process(inp)
	if err != nil {
		t.Error(err)
	} else {
		t.Logf("inp=%+v, out=%+v", inp, out)
	}
}

func TestPipeFilter(t *testing.T) {
	u := UselessFilter{}
	p := NewPipeFilter()
	p.RegisterFilter(&u)
	p.RegisterFilter(&u)
	p.RegisterFilter(&u)
	inp := "test"
	out, err := p.Process(inp)
	if err != nil {
		t.Error(err)
	} else {
		t.Logf("inp=%+v, out=%+v", inp, out)
	}
}
