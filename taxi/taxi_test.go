package taxi

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
	p := "./cmd/taxi"
	test := exec.Command(p, "-t", "9606", db)
	tests = append(tests, test)
	taxa := []string{"homo sapiens",
		"homo  sapiens",
		"haemophilus ducreyi",
		"pseudomonas fluorescens ATCC 17400"}
	for _, taxon := range taxa {
		test := exec.Command(p, taxon, db)
		tests = append(tests, test)
		test = exec.Command(p, "-e", taxon, db)
		tests = append(tests, test)
	}
	test = exec.Command(p, "-D", "test", "-t", "9606")
	tests = append(tests, test)
	for _, taxon := range taxa {
		test := exec.Command(p, "-D", "test", taxon)
		tests = append(tests, test)
		test = exec.Command(p, "-D", "test", "-e", taxon)
		tests = append(tests, test)
	}
	for i, test := range tests {
		get, err := test.Output()
		if err != nil {
			t.Error(err)
		}
		f := "r" + strconv.Itoa((i%9)+1) + ".txt"
		want, err := ioutil.ReadFile(f)
		if err != nil {
			t.Error(err)
		}
		if !bytes.Equal(get, want) {
			t.Errorf("%s - get:\n%s\nwant:\n%s\n", f, get, want)
		}
	}
}
