package main

import (
	"fmt"
	"math"
	"os"
	"strconv"
)

func calculateDiceDistribution(n int) map[int]int {
	distribution := make(map[int]int)

	var calculateCombinations func(depth int, currentSum int)
	calculateCombinations = func(depth int, currentSum int) {
		if depth == n {
			distribution[currentSum]++
			return
		}

		for i := 1; i <= 6; i++ {
			calculateCombinations(depth+1, currentSum+i)
		}
	}

	calculateCombinations(0, 0)
	return distribution
}

func calculateDoublesProb(n int) float64 {
	return 6.0 / math.Pow(6.0, float64(n))
}

func gcd(a, b int) int {
	for b != 0 {
		a, b = b, a%b
	}
	return a
}

func simplifyFraction(numerator, denominator int) (int, int) {
	if numerator == 0 {
		return 0, 1
	}
	
	divisor := gcd(numerator, denominator)
	return numerator / divisor, denominator / divisor
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("使用方法: go run main.go <サイコロの数>")
		os.Exit(1)
	}

	n, err := strconv.Atoi(os.Args[1])
	if err != nil || n <= 0 {
		fmt.Println("エラー: 正の整数を入力してください")
		os.Exit(1)
	}

	// Calculate the distribution
	distribution := calculateDiceDistribution(n)

	fmt.Printf("サイコロを%d個振った時の情報：\n", n)
	
	minSum := n       // Minimum sum (all 1s)
	maxSum := n * 6   // Maximum sum (all 6s)
	
	// Calculate total possible outcomes
	totalOutcomes := int(math.Pow(6.0, float64(n)))
	
	// Display the distribution with probabilities
	for sum := minSum; sum <= maxSum; sum++ {
		ways := distribution[sum]
		probability := float64(ways) / float64(totalOutcomes) * 100
		fmt.Printf("- %d: %d通り（%.2f%%）\n", sum, ways, probability)
	}
	
	doublesProb := calculateDoublesProb(n)
	percentage := doublesProb * 100
	
	numerator := 6
	denominator := int(math.Pow(6.0, float64(n)))
	simplifiedNum, simplifiedDenom := simplifyFraction(numerator, denominator)
	
	fmt.Printf("\nゾロ目が出る確率： %d/%d（%.2f%%）\n", 
		simplifiedNum, simplifiedDenom, percentage)
}
