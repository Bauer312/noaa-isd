package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"time"

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

	var oldRec isd.Record
	count := 0
	scanner := bufio.NewScanner(infp)
	for scanner.Scan() {
		line := scanner.Text()

		rec := isd.Parse(line)
		if rec.SurfaceObservationCode == "FM-15" {
			if count > 0 {
				processRecords(oldRec, rec, outfp)
			}
			count++
			oldRec = rec
		}

	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Processed %d FM-15 records\n", count)
}

func processRecords(old, new isd.Record, outfp *os.File) {
	oldDateString := fmt.Sprintf("%s %s", old.Date, old.Time)
	newDateString := fmt.Sprintf("%s %s", new.Date, new.Time)

	oldDate, err := time.Parse("20060102 1504", oldDateString)
	if err != nil {
		log.Println(old.Date, old.Time, oldDateString)
		log.Fatal(err)
	}
	newDate, err := time.Parse("20060102 1504", newDateString)
	if err != nil {
		log.Println(new.Date, new.Time, newDateString)
		log.Fatal(err)
	}

	oldDateRounded := oldDate.Round(10 * time.Minute)
	newDateRounded := newDate.Round(10 * time.Minute)

	for oldDateRounded.Before(newDateRounded) {
		old.Date = oldDateRounded.Format("20060102")
		old.Time = oldDateRounded.Format("150405")
		fmt.Fprintln(outfp, old.RecordString(","))
		oldDateRounded = oldDateRounded.Add(10 * time.Minute)
	}
}
