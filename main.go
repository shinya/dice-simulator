package main

import (
	"fmt"
	"math"
	"os"
	"runtime"
	"strconv"
	"sync"
	"time"
)

func calculateDiceDistributionParallel(n int) map[int]int {
	// Create a thread-safe map to store the distribution
	var mu sync.Mutex
	distribution := make(map[int]int)
	
	numCPU := runtime.NumCPU()
	fmt.Printf("利用可能なCPUコア数: %d\n", numCPU)
	
	if n <= 2 {
		return calculateDiceDistributionSequential(n)
	}
	
	var wg sync.WaitGroup
	
	for firstDice := 1; firstDice <= 6; firstDice++ {
		wg.Add(1)
		go func(startValue int) {
			defer wg.Done()
			
			// Calculate partial distribution for this starting value
			partialDistribution := make(map[int]int)
			
			var calculateCombinations func(depth int, currentSum int)
			calculateCombinations = func(depth int, currentSum int) {
				if depth == n {
					partialDistribution[currentSum]++
					return
				}
				
				for i := 1; i <= 6; i++ {
					calculateCombinations(depth+1, currentSum+i)
				}
			}
			
			calculateCombinations(1, startValue)
			
			mu.Lock()
			for sum, count := range partialDistribution {
				distribution[sum] += count
			}
			mu.Unlock()
		}(firstDice)
	}
	
	wg.Wait()
	
	return distribution
}

func calculateDiceDistributionSequential(n int) map[int]int {
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

	runtime.GOMAXPROCS(runtime.NumCPU())
	
	startTime := time.Now()
	
	// Calculate the distribution using parallel processing
	distribution := calculateDiceDistributionParallel(n)
	
	// Calculate execution time
	executionTime := time.Since(startTime)

	fmt.Printf("サイコロを%d個振った時の情報：\n", n)
	fmt.Printf("計算時間: %v\n", executionTime)
	
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
