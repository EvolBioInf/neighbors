package main

import (
	"bytes"
	"os"
	"os/exec"
	"strconv"
	"testing"
)

func TestFintac(t *testing.T) {
	var tests []*exec.Cmd
	test := exec.Command("./fintac", "test.nwk")
	tests = append(tests, test)
	test = exec.Command("./fintac", "-a", "test.nwk")
	tests = append(tests, test)
	for i, test := range tests {
		get, err := test.Output()
		if err != nil {
			t.Error(err)
		}
		name := "r" + strconv.Itoa(i+1) + ".txt"
		want, err := os.ReadFile(name)
		if err != nil {
			t.Error(err)
		}
		if !bytes.Equal(get, want) {
			t.Errorf("get:\n%s\nwant:\n%s\n", get, want)
		}
	}
}
