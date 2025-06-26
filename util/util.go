// Package util provides utility functions for the programs
// indexNeighbors and neighbors.
package util

import (
	"fmt"
	"github.com/evolbioinf/clio"
	"log"
	"os"
)

var version, date string
var name string
var assemblyLevels = []string{"complete",
	"chromosome",
	"scaffold",
	"contig"}

func PrintInfo(program string) {
	author := "Bernhard Haubold"
	email := "haubold@evolbio.mpg.de"
	license := "Gnu General Public License, " +
		"https://www.gnu.org/licenses/gpl.html"
	clio.PrintInfo(program, version, date,
		author, email, license)
	os.Exit(0)
}
func Open(file string) *os.File {
	f, err := os.Open(file)
	if err != nil {
		fmt.Fprintf(os.Stderr, "couldn't open %s\n",
			file)
		os.Exit(1)
	}
	return f
}
func Check(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
func SetName(n string) {
	name = n
	s := fmt.Sprintf("%s: ", n)
	log.SetPrefix(s)
	log.SetFlags(0)
}
func Version() {
	PrintInfo(name)
}
func LevelMsg() string {
	levels := assemblyLevels
	m := "assembly-level: comma-delimited combination " +
		"of " + levels[0]
	for i := 1; i < len(levels)-1; i++ {
		m += ", " + levels[i]
	}
	m += ", or " + levels[len(levels)-1]
	m += " (default any)"
	return m
}
