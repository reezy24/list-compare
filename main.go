package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

const usage = "Usage: list-compare path/to/listA path/to/listB"

func main() {
	if len(os.Args) < 3 {
		log.Fatal(usage)
	}

	pathA := os.Args[1]
	pathB := os.Args[2]

	if pathA == "" || pathB == "" {
		log.Fatal(usage)
	}

	listA, err := readLines(pathA)
	if err != nil {
		panic(err)
	}

	listB, err := readLines(pathB)
	if err != nil {
		panic(err)
	}

	setA := sliceToSet(listA)
	fmt.Println("\nSet A:")
	printSet(setA)

	setB := sliceToSet(listB)
	fmt.Println("\nSet B:")
	printSet(setB)

	fmt.Println("\nIntersection:")
	printSet(intersection(setA, setB))

	fmt.Println("\nA - B:")
	printSet(difference(setA, setB))

	fmt.Println("\nB - A:")
	printSet(difference(setB, setA))
}

func sliceToSet(slice []string) []string {
	dic := map[string]bool{}
	for _, v := range slice {
		dic[v] = true
	}
	ret := []string{}
	for k := range dic {
		ret = append(ret, k)
	}
	return ret
}

func intersection(setA, setB []string) []string {
	intersect := []string{}
	for _, v := range setA {
		for _, vb := range setB {
			if v == vb {
				intersect = append(intersect, v)
			}
		}
	}
	return intersect
}

func difference(setA, setB []string) []string {
	diff := []string{}
Loop:
	for _, v := range setA {
		for _, vb := range setB {
			if v == vb {
				continue Loop
			} 
		}
		diff = append(diff, v)
	}
	return diff
}

func readLines(path string) ([]string, error) {
    file, err := os.Open(path)
    if err != nil {
        return nil, err
    }
    defer file.Close()

    var lines []string
    scanner := bufio.NewScanner(file)
    for scanner.Scan() {
        lines = append(lines, scanner.Text())
    }
    return lines, scanner.Err()
}

func printSet(set []string) {
	for _, v := range set {
		fmt.Println(v)
	}
}