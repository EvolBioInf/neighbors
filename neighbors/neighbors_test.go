package neighbors

import (
	"bytes"
	"os"
	"os/exec"
	"strconv"
	"testing"
)

type Test struct {
	t *exec.Cmd
	r int
}

func TestNeighbors(t *testing.T) {
	var tests []Test
	db := "../data/test.db"
	p := "./cmd/neighbors"
	for i := 1; i <= 4; i++ {
		in := "tid" + strconv.Itoa(i) + ".txt"
		test := Test{t: exec.Command(p, db, in), r: i}
		tests = append(tests, test)
		test = Test{t: exec.Command(p, "-D", "test", db, in), r: i}
		tests = append(tests, test)
	}
	test := Test{t: exec.Command(p, "-l", db, "tid4.txt"), r: 5}
	tests = append(tests, test)
	test = Test{t: exec.Command(p, "-D", "test", "-l", db, "tid4.txt"), r: 5}
	tests = append(tests, test)
	test = Test{t: exec.Command(p, "-g", db, "tid4.txt"), r: 6}
	tests = append(tests, test)
	test = Test{t: exec.Command(p, "-D", "test", "-g", db, "tid4.txt"), r: 6}
	tests = append(tests, test)
	test = Test{t: exec.Command(p, "-T", db, "tid4.txt"), r: 7}
	tests = append(tests, test)
	test = Test{t: exec.Command(p, "-D", "test", "-T", db, "tid4.txt"), r: 7}
	tests = append(tests, test)
	test = Test{t: exec.Command(p, "-t", "9606", db), r: 8}
	tests = append(tests, test)
	test = Test{t: exec.Command(p, "-D", "test", "-t", "9606", db), r: 8}
	tests = append(tests, test)
	test = Test{t: exec.Command(p, "-t",
		"9606,9605", db), r: 9}
	tests = append(tests, test)
	test = Test{t: exec.Command(p, "-D", "test", "-t",
		"9606,9605", db), r: 9}
	tests = append(tests, test)
	test = Test{t: exec.Command(p, "-t", "9606",
		"-L", "complete", db), r: 10}
	tests = append(tests, test)
	test = Test{t: exec.Command(p, "-D", "test", "-t", "9606",
		"-L", "complete", db), r: 10}
	tests = append(tests, test)
	test = Test{t: exec.Command(p, "-t", "9606",
		"-L", "complete,chromosome", db), r: 11}
	tests = append(tests, test)
	test = Test{t: exec.Command(p, "-D", "test", "-t", "9606",
		"-L", "complete,chromosome", db), r: 11}
	tests = append(tests, test)
	test = Test{t: exec.Command(p,
		"-o", db, "tid4.txt"), r: 12}
	tests = append(tests, test)
	test = Test{t: exec.Command(p, "-D", "test",
		"-o", db, "tid4.txt"), r: 12}
	tests = append(tests, test)
	test = Test{t: exec.Command(p,
		"-l", "-o", db, "tid4.txt"), r: 13}
	tests = append(tests, test)
	test = Test{t: exec.Command(p, "-D", "test",
		"-l", "-o", db, "tid4.txt"), r: 13}
	tests = append(tests, test)
	for i, test := range tests {
		if i < 16 {
			get, err := test.t.Output()
			if err != nil {
				t.Errorf("couldn't run %q", test.t)
			}
			f := "r" + strconv.Itoa(test.r) + ".txt"
			want, err := os.ReadFile(f)
			if err != nil {
				t.Errorf("couldn't open %q", f)
			}
			if !bytes.Equal(get, want) {
				t.Errorf("get:\n%s\nwant:\n%s\n", get, want)
			}
		}
	}
}
