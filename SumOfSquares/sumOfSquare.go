package sumofsquares

import (
	"fmt"
	"strconv"
	"strings"
)

var n int

func sumOfSquares(lengthOfSide []string, iterate int) int {
	number, _ := strconv.Atoi(lengthOfSide[iterate])
	if iterate == n {
		return number * number
	}

	return ((number * number) + sumOfSquares(lengthOfSide, iterate+1))
}

func main() {
	lengthOfSide := "89, 5, 0.0067, 1.35, 143"
	strArray := strings.Fields(lengthOfSide)
	n = len(strArray) - 1
	result := sumOfSquares(strArray, 0)
	fmt.Println("The sum of the given squares are:", result)
}
