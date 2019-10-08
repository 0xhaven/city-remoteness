package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"math"
	"os"
	"sort"

	"github.com/jacobhaven/city-remoteness/lib"
)

func readCities(csvFile string) (map[*lib.City]struct{}, error) {
	f, err := os.Open(csvFile)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	cityData := csv.NewReader(f)
	// Discard header row
	if _, err := cityData.Read(); err != nil {
		return nil, err
	}
	cities := make(map[*lib.City]struct{})
	for {
		record, err := cityData.Read()
		switch err {
		case nil:
		case io.EOF:
			return cities, nil
		default:
			return nil, err
		}

		if len(record) != 11 {
			continue
		}
		city, err := lib.NewCity(record[0], record[2], record[3], record[9])
		if err != nil {
			continue
		}

		cities[city] = struct{}{}
	}
}

func main() {
	cities, err := readCities("worldcities.csv")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Processing %d cities\n", len(cities))
	type result struct {
		score float64
		city  *lib.City
	}

	const numWorkers = 50
	jobCh := make(chan *lib.City, numWorkers)
	resultCh := make(chan result, numWorkers)
	results := make([]result, len(cities))

	go func() {
		for city := range cities {
			jobCh <- city
		}
	}()

	for i := 0; i < numWorkers; i++ {
		go func() {
			for city := range jobCh {
				var score float64
				for other := range cities {
					distance := city.Distance(other)
					score += float64(other.Population) * math.Pow(1.0002, -distance)
				}
				resultCh <- result{score, city}
			}
		}()
	}

	i := 0
	for result := range resultCh {
		results[i] = result
		if i++; i >= len(results) {
			close(jobCh)
			close(resultCh)
		}
	}
	sort.Slice(results, func(i, j int) bool { return results[i].score > results[j].score })

	const outFileName = "output.csv"
	outFile, err := os.Create(outFileName)
	if err != nil {
		log.Fatal(err)
	}
	defer outFile.Close()

	out := csv.NewWriter(outFile)
	defer out.Flush()

	writeLine := func(columns ...string) {
		if err := out.Write(columns); err != nil {
			log.Fatal(err)
		}
	}

	writeLine("City", "Density", "Population", "Latitude", "Longitude")

	for _, result := range results {
		writeLine(
			result.city.Name,
			fmt.Sprintf("%f", result.score),
			fmt.Sprintf("%d", result.city.Population),
			fmt.Sprintf("%f", result.city.Location.Lat),
			fmt.Sprintf("%f", result.city.Location.Lon),
		)
	}
	fmt.Printf("Wrote %d results to %s\n", len(results), outFileName)
}
