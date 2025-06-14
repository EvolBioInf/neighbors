package main

import (
	"bytes"
	"io/ioutil"
	"os/exec"
	"strconv"
	"testing"
)

func TestTaxi(t *testing.T) {
	var tests []*exec.Cmd
	db := "../data/test.db"
	taxa := []string{"homo sapiens",
		"haemophilus ducreyi",
		"pseudomonas fluorescens ATCC 17400"}
	for _, taxon := range taxa {
		test := exec.Command("./taxi", taxon, db)
		tests = append(tests, test)
		test = exec.Command("./taxi", "-e", taxon, db)
		tests = append(tests, test)
	}
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
