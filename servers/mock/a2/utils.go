package mocka2

func intPtr(i int) *int {
	return &i
}

func getUserOrDriver(idOffset int, exdrivers ...int) (driver int) {
	defer func() {
		if len(exdrivers) > 0 && driver == exdrivers[0] {
			driver *= 1000
		}
	}()

	if idOffset%(PairLimit+1) == 0 {
		driver = idOffset / (PairLimit + 1)
	} else {
		driver = int(idOffset/(PairLimit+1)) + 1
	}

	return
}
