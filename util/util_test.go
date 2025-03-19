package util

import (
	"bufio"
	"testing"
)

func TestUtil(t *testing.T) {
	f := Open("r.txt")
	defer f.Close()
	sc := bufio.NewScanner(f)
	if !sc.Scan() {
		t.Error("scan failed")
	}
	get := sc.Text()
	want := "success"
	if get != want {
		t.Errorf("get:\n%s\nwant:\n%s\n", get, want)
	}
}
