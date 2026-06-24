package outliers

import (
	"bytes"
	"os"
	"os/exec"
	"strconv"
	"testing"
)

func TestOutliers(t *testing.T) {
	var tests []*exec.Cmd
	p := "./cmd/outliers"
	for i := 1; i <= 5; i++ {
		f := "test" + strconv.Itoa(i) + ".txt"
		test := exec.Command(p, f)
		tests = append(tests, test)
	}
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
