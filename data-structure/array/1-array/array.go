package __array

import "errors"

func Find(arr []string, x string) (int, error) {
	for i, v := range arr {
		if v == x {
			return i, nil
		}
	}

	return 0, errors.New("not found")
}

func Counter(arr []string, x string) (int, error) {
	count := 0
	for _, v := range arr {
		if v == x {
			count++
		}
	}
	if count == 0 {
		return 0, errors.New("not found")
	}
	return count, nil
}
