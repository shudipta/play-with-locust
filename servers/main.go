package main

import (
	"flag"
	"fmt"

	mock_a2 "github.com/shudipta/play-with-locust/mock/a2"
	mock_core "github.com/shudipta/play-with-locust/mock/core"
	mock_others "github.com/shudipta/play-with-locust/mock/others"
)

func init() {
	flag.BoolVar(&mock_a2.AllocateDriver, "allocate-driver", false,
		"whether A2 will privide driver, default: false.")
	flag.IntVar(&mock_a2.PairLimit, "pair-limit", 5,
		"the max number of orders a driver can accept within a specified time, default: 5.")

	flag.Parse()
	fmt.Printf("--allocate-driver=%v --pair-limit=%v\n", mock_a2.AllocateDriver, mock_a2.PairLimit)
}

func main() {
	go mock_others.RunMockKuego()
	go mock_others.RunMockSocketServer()
	go mock_others.RunMockMeta()
	go mock_a2.RunMockA2()
	mock_core.RunMockCore()
}
