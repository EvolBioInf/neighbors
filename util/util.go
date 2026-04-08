// Package util provides utility functions for the programs
// indexNeighbors and neighbors.
package util

import (
	"fmt"
	"github.com/evolbioinf/clio"
	"log"
	"math"
	"os"
	"sort"
)

type OutlierStats struct {
	LowerOuterFence float64
	LowerInnerFence float64
	LowerQuartile   float64
	Median          float64
	UpperQuartile   float64
	UpperInnerFence float64
	UpperOuterFence float64
}

var version, date string
var name string
var assemblyLevels = []string{"complete",
	"chromosome",
	"scaffold",
	"contig"}

// PrintInfo prints program information and exits.
func PrintInfo(program string) {
	author := "Bernhard Haubold"
	email := "haubold@evolbio.mpg.de"
	license := "Gnu General Public License, " +
		"https://www.gnu.org/licenses/gpl.html"
	clio.PrintInfo(program, version, date,
		author, email, license)
	os.Exit(0)
}

// Open opens a file with error checking.
func Open(file string) *os.File {
	f, err := os.Open(file)
	if err != nil {
		fmt.Fprintf(os.Stderr, "couldn't open %s\n",
			file)
		os.Exit(1)
	}
	return f
}

// Check checks an error and aborts if it isn't nil.
func Check(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

// The function SetName sets the name of the program.
func SetName(n string) {
	name = n
	s := fmt.Sprintf("%s: ", n)
	log.SetPrefix(s)
	log.SetFlags(0)
}

// The function Version prints the version and other information about the program and exits.
func Version() {
	PrintInfo(name)
}

// The function LevelMsg prints the user message for the option directing the assembly level.
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

// The function OutlierStatistics takes as argument a slice of floats and calculates outlier statistics as defined by the National Institute of Standards and Technology,  https://www.itl.nist.gov/div898/handbook/prc/section1/prc16.htm
func OutlierStatistics(data []float64) *OutlierStats {
	os := new(OutlierStats)
	sort.Float64s(data)
	n := len(data)
	m := (n + 1) / 2
	os.Median = data[m-1]
	if n%2 == 0 {
		os.Median = (os.Median + data[m]) / 2.0
	}
	exactQ := float64(n+1) * 0.25
	f := math.Floor(exactQ)
	l := int(f)
	x := math.Remainder(exactQ, f)
	os.LowerQuartile = data[l-1] + (data[l]-data[l-1])*x
	exactQ = float64(n+1) * 0.75
	f = math.Floor(exactQ)
	l = int(f)
	x = math.Remainder(exactQ, f)
	os.UpperQuartile = data[l-1] + (data[l]-data[l-1])*x
	r := os.UpperQuartile - os.LowerQuartile
	os.LowerInnerFence = os.LowerQuartile - 1.5*r
	os.UpperInnerFence = os.UpperQuartile + 1.5*r
	os.LowerOuterFence = os.LowerQuartile - 3.0*r
	os.UpperOuterFence = os.UpperQuartile + 3.0*r
	return os
}
