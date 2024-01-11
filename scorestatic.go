package main

import "fmt"

func scoreavg(scores []int) int{
	sum := 0
	for i := 0; i < len(scores); i++ {
		sum += scores[i]
	}
	return sum / len(scores)
}

func main() { 
	mathscores := []int{40,89,77,68,72,39,58,87,93,48,65,74,60}
	fmt.Printf("%dによる数学試験の結果：\n", len(mathscores))
	fmt.Printf("平均点%d", scoreavg(mathscores))
}