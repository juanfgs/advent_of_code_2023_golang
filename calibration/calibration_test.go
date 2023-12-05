package calibration_test

import (
	"fmt"
	"github.com/juanfgs/advent_2023/calibration"
	"github.com/juanfgs/advent_2023/test-helper"
	"testing"
)

func TestCalibration(t *testing.T) {
	input := map[string]int{"1abc2": 12,
		"pqr3stu8vwx":      38,
		"a1b2c3d4e5f":      15,
		"treb7uchet":       77,
		"two1nine":         29,
		"eightwothree":     83,
		"abcone2threexyz":  13,
		"xtwone3four":      24,
		"4nineeightseven2": 42,
		"zoneight234":      14,
		"7pqrstsixteen":    76,
	}

	for key, result := range input {
		if result != calibration.Calibration(key) {
			t.Error(fmt.Sprintf("%s should return %d", key, result))
		}
	}

}

func TestSumCalibrations(t *testing.T) {
	inputData := []string{"two1nine",
		"eightwothree",
		"abcone2threexyz",
		"xtwone3four",
		"4nineeightseven2",
		"zoneight234",
		"7pqrstsixteen",
	}
	testData := test_helper.SetupTestDataString("calibration", "trebuchet")

	if 53868 != calibration.SumCalibrations(testData) {
		t.Error("That is not the correct solution to the problem")
	}
	

	if 281 != calibration.SumCalibrations(inputData) {

		t.Error("it should return 281, take into account written numbers")
	}

}
