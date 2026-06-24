package ranks

import (
	"bytes"
	"os"
	"os/exec"
	"strconv"
	"testing"
)

func TestRanks(t *testing.T) {
	tests := make([]*exec.Cmd, 0)
	pr := "./cmd/ranks"
	ta := "207598"
	db := "../data/test.db"
	test := exec.Command(pr, ta, db)
	tests = append(tests, test)
	test = exec.Command(pr, "-g", "genomes.txt", ta, db)
	tests = append(tests, test)
	test = exec.Command(pr, "-l", ta, db)
	tests = append(tests, test)
	test = exec.Command(pr, "-L", "complete", ta, db)
	tests = append(tests, test)
	test = exec.Command(pr, "-t", ta, db)
	tests = append(tests, test)
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
			t.Errorf("%s - get:\n%s\nwant:\n%s\n", f, get, want)
		}
	}
}
