package main

import (
	"bytes"
	"io/ioutil"
	"os/exec"
	"strconv"
	"testing"
)

func TestPickle(t *testing.T) {
	var tests []*exec.Cmd
	f := "test.nwk"
	test := exec.Command("./pickle", "7", f)
	tests = append(tests, test)
	test = exec.Command("./pickle", "7,3", f)
	tests = append(tests, test)
	test = exec.Command("./pickle", "9", f)
	tests = append(tests, test)
	test = exec.Command("./pickle", "4", f)
	tests = append(tests, test)
	test = exec.Command("./pickle", "-c", "7", f)
	tests = append(tests, test)
	test = exec.Command("./pickle", "-c", "7,3", f)
	tests = append(tests, test)
	test = exec.Command("./pickle", "-c", "9", f)
	tests = append(tests, test)
	test = exec.Command("./pickle", "-c", "4", f)
	tests = append(tests, test)
	test = exec.Command("./pickle", "-t", "7", f)
	tests = append(tests, test)
	test = exec.Command("./pickle", "-t", "7,3", f)
	tests = append(tests, test)
	test = exec.Command("./pickle", "-t", "9", f)
	tests = append(tests, test)
	test = exec.Command("./pickle", "-t", "4", f)
	tests = append(tests, test)
	test = exec.Command("./pickle", "-c", "-t", "7", f)
	tests = append(tests, test)
	test = exec.Command("./pickle", "-c", "-t", "7,3", f)
	tests = append(tests, test)
	test = exec.Command("./pickle", "-c", "-t", "9", f)
	tests = append(tests, test)
	test = exec.Command("./pickle", "-c", "-t", "4", f)
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
