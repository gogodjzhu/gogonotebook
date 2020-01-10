package abstract

import (
	"fmt"
	"testing"
)
import back "gogonotebook/algorithm/abstract/backtracking"
import dynamic "gogonotebook/algorithm/abstract/dynamic_programming"
import dp "gogonotebook/algorithm/abstract/double_pointer"
import gree "gogonotebook/algorithm/abstract/greedy"

func TestClimbingStairs(t *testing.T) {
	backtrackSolver := back.ClimbingStairsSolver{}
	dynamicSolver := dynamic.ClimbingStairsSolver{}
	for i := 0; i < 10; i++ {
		if backtrackSolver.ClimbingStairsSolve(i) != dynamicSolver.ClimbingStairsSolve(i) {
			t.Fatal()
		}
	}

}

func TestThreeNumberSum(t *testing.T) {
	backSolver := back.ThreeNumberSumSolve{}
	dpSolver := dp.ThreeNumberSumSolve{}

	arr := []int{-1, 0, 1, 2, -1, -4}

	res1 := backSolver.ThreeNumberSumSolve(arr)
	fmt.Println(res1)
	res2 := dpSolver.ThreeNumberSumSolve(arr)
	fmt.Println(res2)
}

func TestMaximumSubarray(t *testing.T) {
	greedySolver := gree.MaximumSubarraySolver{}
	result := greedySolver.MaximumSubarraySolve([]int{-2, 1})
	fmt.Println(result)
}

func TestThreeSumClosest(t *testing.T) {
	arr := []int{0, 2, 1, -3}
	backSolver := back.ThreeSumClosestSolver{}
	result1 := backSolver.ThreeSumClosestSolve(arr, 1)
	dpSolver := dp.ThreeSumClosestSolver{}
	result2 := dpSolver.ThreeSumClosestSolve(arr, 1)
	fmt.Println(result1)
	fmt.Println(result2)
}

func TestGenerateParentheses(t *testing.T) {
	dynamicSolver := dynamic.GenerateParenthesesSolver{}
	fmt.Println(dynamicSolver.GenerateParenthesesSolve(3))
	backSolver := back.GenerateParenthesesSolver{}
	fmt.Println(backSolver.GenerateParenthesesSolve(3))
}

func TestLetterCombinationsOfAPhoneNumber(t *testing.T) {
	dynamicSolver := dynamic.LetterCombinationsOfAPhoneNumberSolver{}
	fmt.Println(dynamicSolver.LetterCombinationsOfAPhoneNumberSolve("2"))
	backSolver := back.LetterCombinationsOfAPhoneNumberSolver{}
	fmt.Println(backSolver.LetterCombinationsOfAPhoneNumberSolve("2"))
}

func TestFourNumberSum(t *testing.T) {
	dpSolver := dp.FourNumberSumSolver{}
	arr := []int{1, 0, -1, 0, -2, 2}
	target := 0
	fmt.Println(dpSolver.FourNumberSumSolve(arr, target))
}

func TestUniquePaths(t *testing.T) {
	dynamicSolver := dynamic.UniquePathsSolver{}
	fmt.Println(dynamicSolver.UniquePathsSolve(23, 16))
}

func TestUniquePathsII(t *testing.T) {
	dynamicSolver := dynamic.UniquePathsIISolver{}
	fmt.Println(dynamicSolver.UniquePathsIISolve([][]int{
		{1},
		{0},
	}))
}

func TestMinimumPathSum(t *testing.T) {
	dynamicSolver := dynamic.MinimumPathSumSolver{}
	fmt.Println(dynamicSolver.MinimumPathSumSolve([][]int{
		{1, 3, 1, 1},
		{1, 3, 1, 1},
		{1, 3, 1, 1},
	}))
}

func TestBestTimeToBuyAndSellStock(t *testing.T) {
	dynamicSolver := dynamic.BestTimeToBuyAndSellStockSolver{}
	fmt.Println(dynamicSolver.BestTimeToBuyAndSellStockSolve([]int{7, 9, 1, 0, 2}))
}

func TestTriangle(t *testing.T) {
	arr := [][]int{
		{2},
	}
	dynamicSolver := dynamic.TriangleSolver{}
	fmt.Println(dynamicSolver.TriangleSolve(arr))
}

func TestDecodeWays(t *testing.T) {
	dynamicSolver := dynamic.DecodeWaysSolver{}
	testCase := map[string]int{
		"110":   1,
		"27":    1,
		"100":   0,
		"10":    1,
		"101":   1,
		"1001":  0,
		"10101": 1,
		"301":   0,
		"401":   0,
		"444":   1,
		"1111":  5,
		"17":    2,
	}
	for str, ways := range testCase {
		cnt := dynamicSolver.DecodeWaysSolve(str)
		if cnt != ways {
			t.Fatal(str, " expect:", ways, ", actual:", cnt)
		}
	}
}
