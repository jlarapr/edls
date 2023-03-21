package main

import (
	"flag"
	"fmt"
)

func main() {

	// filter patter
	flagPattern := flag.String("p", "", "filter by pattern")
	flagAll := flag.Bool("a", false, "all files include hide files")
	flagNumberRecord := flag.Int64("n", 0, "number of records")

	// order flags
	hasOrderByTime := flag.Bool("t", false, "sort by time, oldest first")
	hasOrderBySize := flag.Bool("s", false, "sort by size, smallest first")
	hasOrderByReverse := flag.Bool("r", false, "reverse order while sorting")

	flag.Parse()

	fmt.Println("flagPattern: ", *flagPattern)
	fmt.Println("flagAll: ", *flagAll)
	fmt.Println("flagNumberRecord: ", *flagNumberRecord)
	fmt.Println("hasOrderByTime: ", *hasOrderByTime)
	fmt.Println("hasOrderBySize: ", *hasOrderBySize)
	fmt.Println("hasOrderByReverse: ", *hasOrderByReverse)

}
