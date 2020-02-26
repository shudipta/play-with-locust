package main

import (
	mock_a2 "github.com/shudipta/play-with-locust/mock/a2"
	mock_core "github.com/shudipta/play-with-locust/mock/core"
)

func main() {
	go mock_core.RunMockCore()
	mock_a2.RunMockA2()
}

