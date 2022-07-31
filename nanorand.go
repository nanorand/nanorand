package nanorand

import (
	"errors"
	"fmt"
	"math"
	"strconv"
	"time"
)

// Gen generates pseudo-random number
// generation can take time for an iteration with length 100 or more
func Gen(length int) (string, error) {
	if length < 0 {
		return "", errors.New("invalid length")
	}

	var (
		i      int
		p      int
		output string
	)

	if length > 7 {
		i = int(math.Floor(float64(length / 7)))
		p = length % 7
	} else {
		return generateShort(length)
	}

	arr, err := generateArrayShort(7, i, 23)
	if err != nil {
		return "", err
	}

	for k, e := range arr {
		g := fmt.Sprint(e * (k + 1))
		if len(g) < 7 {
			g += g[:1]
		}

		output += g[:7]
	}

	if p != 0 {
		piece, err := generateShort(p)
		if err != nil {
			return "", err
		}
		output += piece
	}

	return output, nil
}

// generateShort return only small numbers up to 7
func generateShort(length int) (string, error) {
	if length > 7 || length <= 0 {
		return "", errors.New(fmt.Sprintf("invalid length %d. Try to use Gen() method instead", length))
	}

	timeNow := time.Now().Nanosecond()

	g := fmt.Sprint(timeNow)
	if len(g) < length {
		g += g[:1]
	}
	return g[:length], nil
}

// generateArrayShort Available length up to 7
func generateArrayShort(length int, amount int, offset int) ([]int, error) {
	if offset <= 0 || offset > 50 {
		offset = 20
	}
	arr := make([]int, amount)

	// store elements to slice of integer
	for i := range arr {
		t, err := generateShort(length)
		if err != nil {
			return nil, errors.New(fmt.Sprintf("invalid length %d. Try to use GenerateArray() method instead", length))
		}
		elem, err := strconv.Atoi(t)
		if err != nil {
			return nil, errors.New(t)
		}

		arr[i] = elem

		// calling sleep method
		sleepAmount, _ := strconv.Atoi(t[length-1:])
		sleepAmount++

		dur := time.Duration(sleepAmount*offset) * time.Nanosecond
		time.Sleep(dur)
	}

	return arr, nil
}
