package models

import (
	"fmt"
	"strings"
)

type Index struct{
	data	map[string][]int
}

func GetNewIndex() *Index {
	return &Index{
		data: make(map[string][]int),
	}
}


func (i *Index) Create(docs []Doc){
	// take a line only single time
	hashset_util := make(map[string]map[int]struct{})
	
	for idx, doc := range docs{
		tokens := strings.SplitSeq(doc.Description, " ")
		for token := range tokens{
			if _, ok := hashset_util[token]; !ok {
				hashset_util[token] = make(map[int]struct{})
			}

			if _, ok := hashset_util[token][idx]; ok{
				continue
			}
			i.data[token] = append(i.data[token], idx)
			hashset_util[token][idx] = struct{}{}
		}
	}

	fmt.Printf("Total tokens indexed: %d\n", len(i.data))
}

func (i *Index) Search(docs []Doc, tokens []string) []Doc {
	result := make([]Doc, 0)
	var idices []int

	for token_id := range len(tokens){
		if token_id == 0{
			idices = i.data[tokens[0]]
			continue
		}
		idices = i.intersection(idices, i.data[tokens[token_id]])
	}

	for _, idx := range idices{
		result = append(result, docs[idx])
	}
	return result
}


func (i *Index) intersection(arr1 []int, arr2 []int) []int {
	result := make([]int, 0)

	arr1_elements := make(map[int]struct{})
	for _, element := range arr1{
		arr1_elements[element] = struct{}{}
	}

	for _, element := range arr2{
		_, ok := arr1_elements[element]
		if ok{
			result = append(result, element)
		}
	}

	return result
}
