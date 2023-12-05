package calibration

import (
	"fmt"
	"log"
	"regexp"
	"strconv"
)

func convertDigit(digit string) string {
	conversionTable := map[string]string{
		"one":   "1",
		"two":   "2",
		"three": "3",
		"four":  "4",
		"five":  "5",
		"six":   "6",
		"seven": "7",
		"eight": "8",
		"nine":  "9",
	}
	if converted, ok := conversionTable[digit]; ok {
		return converted
	} else {
		return digit
	}
}

func Calibration(value string) int {
	reg := regexp.MustCompile(`[0-9]|one|two|three|four|five|six|seven|eight|nine`)
	matches := reg.FindAllStringSubmatch(value, -1)
	if len(matches) > 0 {
		var lastMatch string
		firstMatch := convertDigit(matches[0][0])

		for i := len(value); i > 0; i-- {
			substring := value[len(value)-i : len(value)]
			match := reg.FindAllStringSubmatch(substring, 1)
			if len(match) == 1 {
				lastMatch = convertDigit(match[0][0])
			}
		}

		i, err := strconv.Atoi(fmt.Sprintf("%s%s", firstMatch, lastMatch))
		if err != nil {
			log.Fatal(err)
		}

		return i
	}
	return 0
}

func SumCalibrations(values []string) int {
	var acc int = 0
	for _, value := range values {
		acc = acc + Calibration(value)
	}
	return acc
}
