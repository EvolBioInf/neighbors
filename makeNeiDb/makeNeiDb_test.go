package main

import (
	"bytes"
	"io/ioutil"
	"os/exec"
	"strconv"
	"testing"
)

func TestMakeNeiDb(t *testing.T) {
	var tests []*exec.Cmd
	test := exec.Command("./makeNeiDb",
		"-a", "../data/namesTest.dmp",
		"-d", "test.db",
		"-o", "../data/nodesTest.dmp",
		"-m", "../data/mergedTest.dmp",
		"-i", "../data/imagesTest.dmp",
		"-g", "../data/gbTest.txt",
		"-r", "../data/rsTest.txt")
	tests = append(tests, test)
	test = exec.Command("/usr/bin/sqlite3",
		"test.db",
		"select * from taxon order by taxid")
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
