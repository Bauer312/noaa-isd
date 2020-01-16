package main

import (
	"bufio"
	"compress/gzip"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/bauer312/noaa-isd/pkg/isd"
)

var (
	inDir  = flag.String("inputDir", "", "directory containing gzipped isd files")
	inFile = flag.String("inputFile", "", "a specific gzipped isd file")
	outDir = flag.String("outputDir", "", "destination directory for CSV output files")
)

func main() {
	flag.Parse()

	if (len(*inDir) == 0 && len(*inFile) == 0) || len(*outDir) == 0 {
		flag.Usage()
		return
	}

	if len(*inFile) != 0 {
		processGzipFile(*inFile,
			filepath.Join(*outDir, strings.Replace(strings.ToLower(filepath.Base(*inFile)), ".gz", ".csv", 1)))
	}

	if len(*inDir) != 0 {
		files, err := ioutil.ReadDir(*inDir)
		if err != nil {
			log.Fatal(err)
		}

		for _, file := range files {
			if file.IsDir() == false && filepath.Ext(strings.ToLower(file.Name())) == ".gz" {
				processGzipFile(filepath.Join(*inDir, file.Name()),
					filepath.Join(*outDir, strings.Replace(strings.ToLower(file.Name()), ".gz", ".csv", 1)))
			}
		}
	}
}

func processGzipFile(inPath, outPath string) {
	fmt.Println(inPath + " -> " + outPath)
	fp, err := os.Open(inPath)
	if err != nil {
		log.Fatal(err)
	}
	defer fp.Close()

	zr, err := gzip.NewReader(fp)
	if err != nil {
		log.Fatal(err)
	}

	outfp, err := os.Create(outPath)
	if err != nil {
		log.Fatal(err)
	}
	defer outfp.Close()

	fmt.Fprintln(outfp, isd.BasicHeader(","))

	count := 0

	scanner := bufio.NewScanner(zr)
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
	fmt.Printf("Processed %d lines from %s\n", count, inPath)
}
