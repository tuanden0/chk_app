package handlers

import (
	"encoding/csv"
	"log"
	"os"
	"reflect"
	"strconv"
)

type Rows struct {
	UNIX   int64
	SYMBOL string
	OPEN   float64
	HIGH   float64
	LOW    float64
	CLOSE  float64
}

var (
	ROW_HEADER []string = []string{"UNIX", "SYMBOL", "OPEN", "HIGH", "LOW", "CLOSE"}
)

func createCSVData(data [][]string) (out []*Rows, err error) {

	for i, fields := range data {
		if i == 0 {
			if !reflect.DeepEqual(fields, ROW_HEADER) {
				err = errCSVHeader
				return
			}
		} else {

			unix, err := strconv.ParseInt(fields[0], 10, 64)
			if err != nil {
				log.Println(err)
				return nil, errCSVWrongDataType
			}

			open, err := strconv.ParseFloat(fields[2], 64)
			if err != nil {
				log.Println(err)
				return nil, errCSVWrongDataType
			}

			high, err := strconv.ParseFloat(fields[2], 64)
			if err != nil {
				log.Println(err)
				return nil, errCSVWrongDataType
			}

			low, err := strconv.ParseFloat(fields[2], 64)
			if err != nil {
				log.Println(err)
				return nil, errCSVWrongDataType
			}

			close, err := strconv.ParseFloat(fields[2], 64)
			if err != nil {
				log.Println(err)
				return nil, errCSVWrongDataType
			}

			out = append(out, &Rows{
				UNIX:   unix,
				SYMBOL: fields[1],
				OPEN:   open,
				HIGH:   high,
				LOW:    low,
				CLOSE:  close,
			})
		}
	}

	return
}

func HandleCSVFile(fPath string) (out []*Rows, err error) {

	f, err := os.Open(fPath)
	if err != nil {
		return
	}
	defer f.Close()

	csvReader := csv.NewReader(f)
	data, err := csvReader.ReadAll()
	if err != nil {
		return
	}

	out, err = createCSVData(data)
	if err != nil {
		return
	}

	return
}
