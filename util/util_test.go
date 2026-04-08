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
	for _, level := range assemblyLevels {
		if strings.Index(msg, level) == -1 {
			t.Errorf(m, level)
		}
	}
	data := []float64{30, 171, 184, 201}
	os := OutlierStatistics(data)
	wr := []float64{-329.25, -132, 65.25, 177.5,
		196.75, 394, 591.25}
	gr := []float64{os.LowerOuterFence, os.LowerInnerFence,
		os.LowerQuartile, os.Median, os.UpperQuartile,
		os.UpperInnerFence, os.UpperOuterFence}
	for i, w := range wr {
		if w != gr[i] {
			t.Errorf("want: %f\nget: %f\n",
				w, gr[i])
		}
	}
}
