package util

import (
	"bufio"
	"strings"
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
	m := "user message does not match level %q"
	msg := LevelMsg()
	for _, level := range levels {
		if strings.Index(msg, level) == -1 {
			t.Errorf(m, level)
		}
	}
	w := len(levels)
	g := len(AssemblyLevels())
	if w != g {
		t.Errorf("want:\n%d levels\nget:\n%d levels", w, g)
	}
}
