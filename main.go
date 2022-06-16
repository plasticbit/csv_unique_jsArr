package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type workingRecord struct {
	since, until float64
}

type workRecsWithCount struct {
	workingRecord
	days uint
}

func (r workRecsWithCount) String() string {
	return fmt.Sprintf("{ Since: %.2f, Until: %.2f, Days: %d }", r.since, r.until, r.days)
}

func main() {
	file, err := os.Open(os.Args[1])
	if err != nil {
		log.Fatalln(err)
	}

	reader := csv.NewReader(file)
	recs, err := reader.ReadAll()
	if err != nil {
		log.Fatalln(err)
	}

	var uniqueRecs []workRecsWithCount

	for _, record := range recs {
		since := strings.Split(record[1], ":")
		until := strings.Split(record[2], ":")

		sh, _ := strconv.ParseFloat(since[0], 32)
		sm, _ := strconv.ParseFloat(since[1], 32)
		uh, _ := strconv.ParseFloat(until[0], 32)
		um, _ := strconv.ParseFloat(until[1], 32)

		dayRec := workingRecord{sh + (sm / 60.), uh + (um / 60.)}
		someIndex := -1
		for idx, elem := range uniqueRecs {
			if dayRec.since == elem.since && dayRec.until == elem.until {
				someIndex = idx
				break
			}
		}

		if someIndex != -1 {
			uniqueRecs[someIndex].days++
		} else {
			uniqueRecs = append(uniqueRecs, workRecsWithCount{dayRec, 1})
		}
	}

	for _, v := range uniqueRecs {
		fmt.Printf("%s,\n", v)
	}
}
