package utils

import "fmt"

var historyList []string

func AddHistory(conversion string) []string {
	historyList = append(historyList, conversion)
	return historyList
}

func ShowHistory() {
	if len(historyList) == 0 {
		fmt.Println("No conversion history available.")
	} else {
		for i, conversion := range historyList {
			fmt.Printf("%d. %s\n", i+1, conversion)
		}
	}
}
