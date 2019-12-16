package service

func GetPrime(inputs ...int64) (ret []int64) {
	for _, v := range inputs {
		if isPrime(v) {
			ret = append(ret, v)
		}
	}

	return
}

func isPrime(v int64) bool {
	var j int64 = 2
	for j = 2; j <= v/2; j++ {
		if v%j == 0 {
			return false
		}
	}

	return true
}
