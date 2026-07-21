package ants

import (
	"bytes"
	"io/ioutil"
	"os/exec"
	"strconv"
	"testing"
)

type Test struct {
	t *exec.Cmd
	r int
}

func TestAnts(t *testing.T) {
	var tests []Test
	tid := "9606"
	db := "../data/test.db"
	test := Test{t: exec.Command("./cmd/ants", tid, db), r: 1}
	tests = append(tests, test)
	test = Test{t: exec.Command("./cmd/ants", "-D", "test", tid, db), r: 1}
	tests = append(tests, test)
	for _, test := range tests {
		get, err := test.t.Output()
		if err != nil {
			t.Error(err)
		}
		f := "r" + strconv.Itoa(test.r) + ".txt"
		want, err := ioutil.ReadFile(f)
		if err != nil {
			t.Error(err)
		}
		if !bytes.Equal(get, want) {
			t.Errorf("get:\n%s\nwant:\n%s\n", get, want)
		}
	}
}
