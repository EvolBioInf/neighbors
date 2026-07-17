// Package util provides utility functions for the programs
// indexNeighbors and neighbors.
package util

import (
	"bytes"
	"fmt"
	"github.com/evolbioinf/clio"
	"io"
	"log"
	"math"
	"mime/multipart"
	"net/http"
	"net/url"
	"os"
	"sort"
	"strconv"
	"strings"
)

type Quart struct {
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
	author := "EvolBioInf"
	email := "haubold|mourato@evolbio.mpg.de"
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

// The function Quartiles takes as argument a slice of floats and calculates the quartiles, including fences for outlier analysis.
func Quartiles(data []float64) *Quart {
	q := new(Quart)
	sort.Float64s(data)
	n := len(data)
	m := (n + 1) / 2
	q.Median = data[m-1]
	if n%2 == 0 {
		q.Median = (q.Median + data[m]) / 2.0
	}
	exactQ := float64(n+1) * 0.25
	f := math.Floor(exactQ)
	l := int(f)
	x := math.Remainder(exactQ, f)
	q.LowerQuartile = data[l-1] + (data[l]-data[l-1])*x
	exactQ = float64(n+1) * 0.75
	f = math.Floor(exactQ)
	l = int(f)
	x = math.Remainder(exactQ, f)
	q.UpperQuartile = data[l-1] + (data[l]-data[l-1])*x
	r := q.UpperQuartile - q.LowerQuartile
	q.LowerInnerFence = q.LowerQuartile - 1.5*r
	q.UpperInnerFence = q.UpperQuartile + 1.5*r
	q.LowerOuterFence = q.LowerQuartile - 3.0*r
	q.UpperOuterFence = q.UpperQuartile + 3.0*r
	return q
}

// The function SendGetRequest takes as argument a url path as a string and the program's options and extra arguments as string slices, as well as miscellaneous arguments as a map. It sends a get request using these values to never at neighbors.evolbio.mpg.de and returns the result.
func SendGetRequest(address string, options, extraArgs []string, miscArgs map[string]string) string {
	address = "http://localhost:8080" + address
	qb := new(strings.Builder)
	urlEncodeSlice(qb, options, "options")
	urlEncodeSlice(qb, extraArgs, "extra")
	for m, a := range miscArgs {
		if qb.Len() == 0 {
			qb.WriteRune('?')
		} else {
			qb.WriteRune('&')
		}
		qb.WriteString(url.QueryEscape(m))
		qb.WriteRune('=')
		qb.WriteString(url.QueryEscape(a))
	}
	req, err := http.NewRequest(http.MethodGet, address+qb.String(), nil)
	Check(err)
	resp, err := http.DefaultClient.Do(req)
	Check(err)
	body, err := io.ReadAll(resp.Body)
	Check(err)
	return string(body)
}
func urlEncodeSlice(qb *strings.Builder, slc []string, paramName string) {
	for _, v := range slc {
		if qb.Len() == 0 {
			qb.WriteRune('?')
		} else {
			qb.WriteRune('&')
		}
		qb.WriteString(paramName)
		qb.WriteRune('=')
		qb.WriteString(url.QueryEscape(v))
	}
}

// The function SendPostRequest takes as argument a url path as a string, program options and extra arguments as a slice of strings, as well as files and stdin. It sends a post request to never at neighbors.evolbio.mpg.de using these values and returns the result.
func SendPostRequest(address string, options, extraArgs []string, miscArgs map[string]string, files []*os.File, stdin *os.File) string {
	address = "http://localhost:8080/" + address
	qb := new(strings.Builder)
	urlEncodeSlice(qb, options, "options")
	urlEncodeSlice(qb, extraArgs, "extra")
	for m, a := range miscArgs {
		if qb.Len() == 0 {
			qb.WriteRune('?')
		} else {
			qb.WriteRune('&')
		}
		qb.WriteString(url.QueryEscape(m))
		qb.WriteRune('=')
		qb.WriteString(url.QueryEscape(a))
	}
	var b bytes.Buffer
	w := multipart.NewWriter(&b)
	for i, file := range files {
		fw, err := w.CreateFormFile(strconv.Itoa(i), file.Name())
		Check(err)
		_, err = io.Copy(fw, file)
		Check(err)
	}
	if stdin != nil {
		fw, err := w.CreateFormFile("stdin", "stdin")
		Check(err)
		_, err = io.Copy(fw, stdin)
		Check(err)
	}
	w.Close()
	req, err := http.NewRequest(http.MethodPost, address+qb.String(), &b)
	Check(err)
	req.Header.Set("Content-Type", w.FormDataContentType())
	resp, err := http.DefaultClient.Do(req)
	Check(err)
	body, err := io.ReadAll(resp.Body)
	Check(err)
	return string(body)
}
