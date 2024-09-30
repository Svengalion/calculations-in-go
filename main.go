package main

import (
	"bufio"
	"flag"
	"fmt"
	"math"
	"os"
	"sort"
	"strconv"
)

func Mean(numbers []int) float64 {
	if len(numbers) == 0 {
		return 0
	}
	var sum int64 = 0
	for _, num := range numbers {
		sum += int64(num)
	}
	return float64(sum) / float64(len(numbers))
}

func Sd(numbers []int, mean float64) float64 {
	if len(numbers) == 0 {
		return 0.0
	}

	var sumSquares float64
	for _, num := range numbers {
		diff := float64(num) - mean
		sumSquares += diff * diff
	}

	variance := sumSquares / float64(len(numbers)-1)
	return math.Sqrt(variance)
}

func Median(numbers []int) float64 {
	if len(numbers) == 0 {
		return 0
	}
	sorted := make([]int, len(numbers))
	copy(sorted, numbers)
	sort.Ints(sorted)
	n := len(sorted)
	mid := n / 2

	if n%2 == 1 {
		return float64(sorted[mid])
	}
	return float64(sorted[mid-1]+sorted[mid]) / 2.0
}

func Mode(numbers []int) int {
	if len(numbers) == 0 {
		return 0
	}
	freqMap := make(map[int]int)
	maxFreq := 0
	mode := numbers[0]
	for _, num := range numbers {
		freqMap[num]++
		if freqMap[num] > maxFreq {
			maxFreq = freqMap[num]
			mode = num
		} else if freqMap[num] == maxFreq {
			if num < mode {
				mode = num
			}
		}
	}
	return mode
}

func Parse() ([]int, error) {
	scanner := bufio.NewScanner(os.Stdin)
	var numbers []int

	for scanner.Scan() {
		line := scanner.Text()

		if line == "" {
			continue
		}

		num, err := strconv.Atoi(line)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Неверный ввод '%s', пропуск.\n", line)
			continue
		}

		if num <= -100000 || num >= 100000 {
			fmt.Fprintf(os.Stderr, "Число %d вне диапазона, пропуск.\n", num)
			continue
		}

		numbers = append(numbers, num)
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return numbers, nil
}

func main() {
	meanFlag := flag.Bool("mean", false, "Показать среднее (Mean)")
	medianFlag := flag.Bool("median", false, "Показать медиану (Median)")
	modeFlag := flag.Bool("mode", false, "Показать моду (Mode)")
	sdFlag := flag.Bool("sd", false, "Показать стандартное отклонение (SD)")

	flag.Parse()

	computeMean := *meanFlag
	computeMedian := *medianFlag
	computeMode := *modeFlag
	computeSd := *sdFlag

	if !(computeMean || computeMedian || computeMode || computeSd) {
		computeMean = true
		computeMedian = true
		computeMode = true
		computeSd = true
	}

	nums, err := Parse()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Ошибка при чтении входных данных: %v\n", err)
		os.Exit(1)
	}

	if len(nums) == 0 {
		fmt.Println("Нет валидных входных данных для обработки.")
		os.Exit(0)
	}

	if computeMean || computeSd {
		mean := Mean(nums)
		if computeMean {
			fmt.Printf("Mean: %.2f\n", mean)
		}
		if computeSd {
			sd := Sd(nums, mean)
			fmt.Printf("SD: %.2f\n", sd)
		}
	}

	if computeMedian {
		median := Median(nums)
		fmt.Printf("Median: %.2f\n", median)
	}

	if computeMode {
		mode := Mode(nums)
		fmt.Printf("Mode: %d\n", mode)
	}
}
