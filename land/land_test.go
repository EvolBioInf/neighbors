package main

import (
	"bytes"
	"io/ioutil"
	"os/exec"
	"strconv"
	"testing"
)

func TestLand(t *testing.T) {
	var tests []*exec.Cmd
	f := "test.nwk"
	test := exec.Command("./land", f)
	tests = append(tests, test)
	test = exec.Command("./land", "-p", "p", f)
	tests = append(tests, test)
	test = exec.Command("./land", "-s", "s", f)
	tests = append(tests, test)
	for i, test := range tests {
		get, err := test.Output()
		if err != nil {
			t.Error(err)
		}
		f := "r" + strconv.Itoa(i+1) + ".txt"
		want, err := ioutil.ReadFile(f)
		if err != nil {
			t.Error(err)
		}
		if !bytes.Equal(get, want) {
			t.Errorf("get:\n%s\nwant:\n%s\n", get, want)
		}
	}
}
