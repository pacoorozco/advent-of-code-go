package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	file, err := os.Open("../input.txt")
	if err != nil {
		log.Fatalf("could not read input data from file, err=%v", err)
	}
	defer file.Close()

	var nums []int

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		change, err := strconv.Atoi(scanner.Text())
		if err != nil {
			log.Fatalf("could not convert line to integer, err=%v", err)
		}
		nums = append(nums, change)
	}
	if err := scanner.Err(); err != nil {
		log.Fatalf("could not read line on frequencies, err=%v", err)
	}

	freq := 0
	seen := map[int]bool{0: true}

	for {
		for _, v := range nums {
			freq += v
			if seen[freq] {
				fmt.Println(freq)
				return
			}
			seen[freq] = true
		}
	}
}
