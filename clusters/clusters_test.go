package main

import (
	"bytes"
	"os"
	"os/exec"
	"strconv"
	"testing"
)

func TestClusters(t *testing.T) {
	var tests []*exec.Cmd
	f := "test.nwk"
	test := exec.Command("./clusters", f)
	tests = append(tests, test)
	test = exec.Command("./clusters", "-b", "3090", f)
	tests = append(tests, test)
	test = exec.Command("./clusters", "-c", f)
	tests = append(tests, test)
	test = exec.Command("./clusters", "-m", f)
	tests = append(tests, test)
	test = exec.Command("./clusters", "-T", f)
	tests = append(tests, test)
	test = exec.Command("./clusters", "-s", "4", f)
	tests = append(tests, test)
	test = exec.Command("./clusters", "-t", f)
	tests = append(tests, test)
	for i, test := range tests {
		get, err := test.Output()
		if err != nil {
			t.Error(err)
		}
		f := "r" + strconv.Itoa(i+1) + ".txt"
		want, err := os.ReadFile(f)
		if err != nil {
			t.Error(err)
		}
		if !bytes.Equal(get, want) {
			t.Errorf("get:\n%s\nwant:\n%s\n", get, want)
		}
	}
}
