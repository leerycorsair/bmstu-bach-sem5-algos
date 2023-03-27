package main

import (
	"fmt"
	"local/dict"
	"reflect"
	"runtime"
)

func GetFunctionName(i interface{}) string {
	return runtime.FuncForPC(reflect.ValueOf(i).Pointer()).Name()
}

func main() {
	fmt.Println("DICTIONARY SEARCH")
	var filename string
	fmt.Printf("Enter the dictionary filename:")
	fmt.Scanln(&filename)
	var dictionary dict.Dict
	dict.DictInit(&dictionary, filename)
	fmt.Printf("DICTIONARY SIZE = %v\n", len(dictionary))
	var searchWord string
	fmt.Printf("Enter the word to search:")
	fmt.Scanln(&searchWord)
	f := [3]func(dict.Dict, string) (int, int, bool){dict.BruteS, dict.BinS, dict.SegmS}
	for _, currFunc := range f {
		val, count, exists := currFunc(dictionary, searchWord)
		fmt.Println(GetFunctionName(currFunc))
		if !exists {
			fmt.Printf("The key was not found. COMPARES = %d\n", count)
		} else {
			fmt.Printf("VALUE = %d COMPARES = %d\n", val, count)
		}
	}
	dict.EfficiencyCheck(dictionary)
}
