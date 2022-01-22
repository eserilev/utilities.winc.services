package winc_csv

import (
	"encoding/csv"
	"log"
	"os"
)

func ReadCsv(filePath string) [][]string {
	f, err := os.Open(filePath)
	if err != nil {
		log.Fatal("Unable to read input file "+filePath, err)
		panic(err)
	}
	defer f.Close()

	csvReader := csv.NewReader(f)
	records, err := csvReader.ReadAll()
	if err != nil {
		log.Fatal("Unable to parse file as CSV for "+filePath, err)
		panic(err)
	}
	return records
}

func WriteCsv(content [][]string, fileName string) string {
	f, err := os.Create(fileName)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	csvWriter := csv.NewWriter(f)
	err = csvWriter.WriteAll(content)
	if err != nil {
		panic(err)
	}
	return f.Name()
}
