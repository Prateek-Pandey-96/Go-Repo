package main

import (
	"fmt"
	"strings"
	"time"

	"github.com/prateek69/searchIndex/models"
	"github.com/prateek69/searchIndex/utils"
)

func main() {
	fmt.Println("Welcome to search index!")

	docs := utils.ReadFile()

	start := time.Now()
	index := models.GetNewIndex()
	index.Create(docs)
	end := time.Since(start)
	fmt.Printf("took %d ms time to index 100mb of data\n", end.Milliseconds())
	
	start = time.Now()
	query := []string{"word306", "word856", "word1", "word973"}
	result_docs := index.Search(docs, query)
	end = time.Since(start)
	fmt.Printf("took %d us to search for the query in the index\n", end.Microseconds())
	fmt.Println(result_docs)

	start = time.Now()
	temp := make([]models.Doc, 0)
	for _, doc := range(docs){
		count := 0
		for _, q := range(query){
			if strings.Contains(doc.Description, q){
				count += 1
			}
		}
		if count == len(query){
			temp = append(temp, doc)
		}
	}
	end = time.Since(start)
	fmt.Printf("took %d us to search for the word using string.contains\n", end.Microseconds())
}
