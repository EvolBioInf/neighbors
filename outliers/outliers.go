package main

import (
	"bufio"
	"flag"
	"fmt"
	"github.com/evolbioinf/clio"
	"github.com/evolbioinf/neighbors/util"
	"io"
	"math"
	"os"
	"sort"
	"strconv"
	"strings"
	"text/tabwriter"
)

func scan(r io.Reader, args ...interface{}) {
	data := make([]float64, 0)
	sc := bufio.NewScanner(r)
	for sc.Scan() {
		str := strings.Trim(sc.Text(), " ")
		if len(str) > 0 {
			d, e := strconv.ParseFloat(str, 64)
			util.Check(e)
			data = append(data, d)
		}
	}
	if len(data) >= 4 {
		sort.Float64s(data)
		n := len(data)
		m := (n + 1) / 2
		q2 := data[m-1]
		if n%2 == 0 {
			q2 = (q2 + data[m]) / 2.0
		}
		exactQ := float64(n+1) * 0.25
		f := math.Floor(exactQ)
		l := int(f)
		x := math.Remainder(exactQ, f)
		q1 := data[l-1] + (data[l]-data[l-1])*x
		exactQ = float64(n+1) * 0.75
		f = math.Floor(exactQ)
		l = int(f)
		x = math.Remainder(exactQ, f)
		q3 := data[l-1] + (data[l]-data[l-1])*x
		iq := q3 - q1
		lif := q1 - 1.5*iq
		uif := q3 + 1.5*iq
		lof := q1 - 3.0*iq
		uof := q3 + 3.0*iq
		w := tabwriter.NewWriter(os.Stdout, 0, 1, 2, ' ', 0)
		msg := "#Lower_outer_fence\tLower_inner_fence\t" +
			"Lower_quartile\tMedian\tUpper_quartile\t" +
			"Upper_inner_fence\tUpper_outer_fence"
		fmt.Fprintf(w, "%s\n", msg)
		fmt.Fprintf(w, "%g\t%g\t%g\t%g\t%g\t%g\t%g\n",
			lof, lif, q1, q2, q3, uif, uof)
		w.Flush()
		mouts := make([]float64, 0)
		for _, d := range data {
			if (d > lof && d < lif) ||
				d > uif && d < uof {
				mouts = append(mouts, d)
			}
		}
		eouts := make([]float64, 0)
		for _, d := range data {
			if d < lof || d > uof {
				eouts = append(eouts, d)
			}
		}
		printOutliers(mouts, "mild")
		printOutliers(eouts, "extreme")
	} else {
		m := "outliers - Need at least 4 data points " +
			"for an outlier analysis"
		fmt.Fprintf(os.Stderr, m)
	}
}
func printOutliers(data []float64, kind string) {
	n := len(data)
	if n == 0 {
		fmt.Printf("No_%s_outliers", kind)
	} else {
		s := strings.ToUpper(kind[0:1]) + kind[1:]
		fmt.Printf("%s_outlier", s)
	}
	if n > 1 {
		fmt.Printf("s")
	}
	if n > 0 {
		fmt.Printf(":")
	}
	for _, d := range data {
		fmt.Printf(" %g", d)
	}
	fmt.Printf("\n")
}
func main() {
	util.SetName("outliers")
	u := "outliers [option]... [file]..."
	p := "List outliers according to the quartile criterion."
	e := "outliers foo.dat"
	clio.Usage(u, p, e)
	optV := flag.Bool("v", false, "version")
	flag.Parse()
	if *optV {
		util.Version()
	}
	files := flag.Args()
	clio.ParseFiles(files, scan)
}
