package main

import (
	"log"
	"math/rand"
	"regexp"
	"strconv"
)

func rollText(text string) int {
	var re = regexp.MustCompile(`(?mi)([\d]+)?(?:d|к)([\d]+)`)
	totalRoll := 0
	for _, match := range re.FindAllStringSubmatch(text, -1) {
		num, err := strconv.Atoi(match[1])
		if err != nil {
			num = 1
		}

		dice, err := strconv.Atoi(match[2])
		if err != nil {
			log.Println(err)
			return 0
		}
		totalRoll += roll(num, dice)
	}

	return totalRoll
}

func roll(num int, dice int) int {
	sum := 0
	for num > 0 {
		sum += 1 + rand.Intn(dice-1)
		num--
	}

	return sum
}
