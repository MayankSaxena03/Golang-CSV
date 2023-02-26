package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"strconv"
	"text/tabwriter"
)

type student struct {
	name string
	age  int
	city string
}

func main() {

	csvFile, err := os.Open("students.csv")
	if err != nil {
		panic(err)
	}
	defer csvFile.Close()

	reader := csv.NewReader(csvFile)
	reader.FieldsPerRecord = -1

	csvData, err := reader.ReadAll()
	if err != nil {
		panic(err)
	}

	var students []student
	for i, row := range csvData {
		if len(row) != 3 {
			log.Fatalf("Error at line: %d", i+1)
			continue
		}

		studentAge, err := strconv.Atoi(row[1])
		if err != nil {
			log.Fatalf("Error converting age to int at line: %d", i+1)
			continue
		}
		students = append(students, student{
			name: row[0],
			age:  studentAge,
			city: row[2],
		})
	}

	tabwriter := tabwriter.NewWriter(os.Stdout, 0, 0, 1, ' ', tabwriter.Debug)
	fmt.Fprintln(tabwriter, "Name\tAge\tCity")
	for _, student := range students {
		fmt.Fprintf(tabwriter, "%s\t%d\t%s", student.name, student.age, student.city)
		fmt.Fprintln(tabwriter)
	}
	tabwriter.Flush()
}
