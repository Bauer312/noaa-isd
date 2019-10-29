package main

import (
	"flag"
	"fmt"
	"log"
	"os/exec"
	"path/filepath"
	"time"
)

/*
Super-simple method for getting data, one step above a shell script...
*/

type isdFile struct {
	team      string
	stationID string
}

/*
Not all teams are here.  I am intentionally ignoring teams that play home games
	in stadiums with retractable/fixed roofs.  Maybe someday I can get
	the info I need to know if the roof is open during the game.
*/
var isdFiles = [...]isdFile{
	{"angels", "722976-03166"},
	{"athletics", "724930-23230"},
	{"braves", "722270-13864"},
	{"cardinals", "725314-03960"},
	{"dodgers", "722874-93134"},
	{"giants", "724940-23234"},
	{"indians", "725245-04853"},
	{"mariners", "727935-24234"},
	{"mets", "725030-14732"},
	{"nationals", "724050-13743"},
	{"orioles", "745944-93784"},
	{"padres", "722900-23188"},
	{"phillies", "724080-13739"},
	{"pirates", "725205-14762"},
	{"rangers", "722479-53907"},
	{"reds", "724297-93812"},
	{"redsox", "725090-14739"},
	{"rockies", "724666-93067"},
	{"royals", "724463-13988"},
	{"tigers", "725375-14822"},
	{"twins", "726575-94960"},
	{"whitesox", "725340-14819"},
	{"yankees", "725053-94728"},
}

func main() {
	url := flag.String("url", "ftp://ftp.ncdc.noaa.gov/pub/data/noaa", "The base URL used for constructing a request")
	year := flag.Int("year", time.Now().Year(), "The desired year")
	team := flag.String("team", "all", "The desired team")
	output := flag.String("output", "", "Directory for storing output file(s)")

	flag.Parse()

	if len(*output) == 0 {
		flag.Usage()
		return
	}

	for _, rec := range isdFiles {
		if *team == "all" || *team == rec.team {
			downloadFile(*url, rec.stationID, *output, *year)
		}
	}
}

/*
Go standard library doesn't support ftp.  There is a library available on GitHub, but...
	We can always just call out to curl, at least for now.
*/
func downloadFile(url, station, target string, year int) {
	filename := fmt.Sprintf("%s-%d.gz", station, year)
	fullURL := filepath.Join(filepath.Join(url, fmt.Sprintf("%d", year)), filename)
	fullTarget := filepath.Join(target, filename)

	fmt.Println(fullURL)
	cmd := exec.Command("curl", fullURL, "-o", fullTarget)
	err := cmd.Run()
	if err != nil {
		log.Fatal(err)
	}
}
