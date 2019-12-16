package service

func GetSum(inputs ...int64) (ret int64) {
	for _, v := range inputs {
		ret += v
	}

	return
}
