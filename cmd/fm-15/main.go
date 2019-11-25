package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/bauer312/noaa-isd/pkg/isd"
)

var (
	in  = flag.String("in", "", "raw ISD file to parse")
	out = flag.String("out", "", "destination CSV file")
)

func main() {
	flag.Parse()

	if len(*in) == 0 || len(*out) == 0 {
		flag.Usage()
		return
	}

	infp, err := os.Open(*in)
	if err != nil {
		log.Fatal(err)
	}
	defer infp.Close()

	outfp, err := os.Create(*out)
	if err != nil {
		log.Fatal(err)
	}
	defer outfp.Close()

	fmt.Fprintln(outfp, isd.BasicHeader(","))

	count := 0
	scanner := bufio.NewScanner(infp)
	for scanner.Scan() {
		line := scanner.Text()

		rec := isd.Parse(line)
		if rec.SurfaceObservationCode == "FM-15" {
			fmt.Fprintln(outfp, rec.RecordString(","))
			count++
		}

	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Processed %d FM-15 records in %s\n", count, *in)
}
