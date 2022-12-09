package main

import (
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	calories, err := loadReindeerCalories("./calories.txt")
	if err != nil {
		panic(err)
	}

	totals := calculateTotals(calories)
	highest := sortHighestCalories(totals)

	fmt.Printf("Highest calories is %d\n", highest[0])

	top3Total := sumOfCalories(highest[0], highest[1], highest[2])

	fmt.Printf("Top 3 highest calories total is %d\n", top3Total)
}

func sumOfCalories(calories ...int) int {
	total := 0

	for _, c := range calories {
		total += c
	}

	return total
}

func sortHighestCalories(calories []int) []int {
	sort.Slice(calories, func(i, j int) bool {
		return calories[i] > calories[j]
	})

	return calories
}

func calculateTotals(calories [][]int) []int {
	totals := make([]int, 0)
	for _, c := range calories {
		total := sumOfCalories(c...)

		totals = append(totals, total)
	}

	return totals
}

func loadReindeerCalories(filename string) ([][]int, error) {
	data, err := os.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	str := strings.Split(
		strings.Replace(
			string(data),
			"\r\n",
			"\n",
			-1,
		),
		"\n",
	)

	cals := make([][]int, 0)
	reindeerCal := make([]int, 0)

	for _, s := range str {
		if s == "" {
			cals = append(cals, reindeerCal)
			reindeerCal = make([]int, 0)
			continue
		}

		converted, err := strconv.Atoi(s)
		if err != nil {
			return nil, err
		}

		reindeerCal = append(reindeerCal, converted)
	}

	return cals, nil
}
