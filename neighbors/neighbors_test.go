package main

import (
	"bytes"
	"io/ioutil"
	"os/exec"
	"strconv"
	"testing"
)

func TestNeighbors(t *testing.T) {
	var tests []*exec.Cmd
	db := "../data/neidb"
	for i := 1; i <= 4; i++ {
		in := "tid" + strconv.Itoa(i) + ".txt"
		test := exec.Command("./neighbors", db, in)
		tests = append(tests, test)
	}
	test := exec.Command("./neighbors", "-l", db, "tid4.txt")
	tests = append(tests, test)
	test = exec.Command("./neighbors", "-g", db, "tid4.txt")
	tests = append(tests, test)
	test = exec.Command("./neighbors", "-T", db, "tid4.txt")
	tests = append(tests, test)
	test = exec.Command("./neighbors", "-t", "866775", db)
	tests = append(tests, test)
	test = exec.Command("./neighbors", "-t",
		"866775,1216979", db)
	tests = append(tests, test)
	for i, test := range tests {
		get, err := test.Output()
		if err != nil {
			t.Errorf("couldn't run %q", test)
		}
		f := "r" + strconv.Itoa(i+1) + ".txt"
		want, err := ioutil.ReadFile(f)
		if err != nil {
			t.Errorf("couldn't open %q", f)
		}
		if !bytes.Equal(get, want) {
			t.Errorf("get:\n%s\nwant:\n%s\n", get, want)
		}
	}
}
