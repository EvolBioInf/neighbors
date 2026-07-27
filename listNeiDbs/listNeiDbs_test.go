package main

import (
	"bytes"
	"os/exec"
	"testing"
)

func TestListNeiDbs(t *testing.T) {
	test := exec.Command("./listNeiDbs")
	g, e := test.Output()
	if e != nil {
		t.Error(e)
	}
	tab := bytes.Split(g, []byte{'\n'})
	n := len(tab) - 1
	if n < 2 {
		t.Error("listNeiDbs: get no databases\n")
	}
}
