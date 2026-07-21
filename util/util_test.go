package util

import (
	"bufio"
	"slices"
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
	for _, level := range assemblyLevels {
		if !strings.Contains(msg, level) {
			t.Errorf(m, level)
		}
	}
	data := []float64{30, 171, 184, 201}
	q := Quartiles(data)
	wr := []float64{-329.25, -132, 65.25, 177.5,
		196.75, 394, 591.25}
	gr := []float64{q.LowerOuterFence, q.LowerInnerFence,
		q.LowerQuartile, q.Median, q.UpperQuartile,
		q.UpperInnerFence, q.UpperOuterFence}
	for i, w := range wr {
		if w != gr[i] {
			t.Errorf("want: %f\nget: %f\n",
				w, gr[i])
		}
	}
	args := []string{"first", "1234", "-t", "-m", "xyz", "-f", "r.txt", "-r",
		"-p", "1233", "&4982k", "-D", "smth", "fake1.db", "sdlfjk", "fake2.db"}
	cArgs := slices.Clone(args)

	res := SanitizeArguments(args, []Option{
		{Name: "r", WithValue: false},
		{Name: "D", WithValue: true},
		{Name: "p", WithValue: true},
	})
	e := []string{"first", "1234", "-t", "-m", "xyz", "-f", "r.txt", "&4982k",
		"sdlfjk"}

	if !slices.Equal(res, e) {
		t.Errorf("want:\n%s\nget:\n%s\n", e, res)
	}

	if !slices.Equal(args, cArgs) {
		t.Errorf("original args changed during processing\n")
	}
}
