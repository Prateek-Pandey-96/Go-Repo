package utils

import (
	"encoding/csv"
	"io"
	"os"

	"github.com/prateek69/searchIndex/models"
)

func ReadFile() []models.Doc {
	file, err := os.Open("output.tsv")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	reader := csv.NewReader(file)
	reader.Comma = '\t'
	
	docs := make([]models.Doc, 0)
	for {
		record, err := reader.Read()
		if err != nil{
			if err == io.EOF{
				break
			}
			continue
		}
		
		title := record[0]
		description := record[1]

		docs = append(docs, models.Doc{Title: title, Description: description})
	}

	return docs
}
