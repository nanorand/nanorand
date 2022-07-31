package nanorand

import (
	"errors"
	"fmt"
	"math"
	"strconv"
	"time"
)

// Generate can take time for an iteration with length 100 or more
func Generate(length int) (string, error) {
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
		return GenerateShort(length)
	}

	arr, err := GenerateArray(7, i, 23)
	if err != nil {
		return "", err
	}

	for k, e := range arr {
		g := fmt.Sprint(e * (k + 1))
		if len(g) < 7 {
			g += g[:1]
		}

		row := g[:7]
		output += row
	}

	if p != 0 {
		piece, err := GenerateShort(p)
		if err != nil {
			return "", err
		}
		output += piece
	}

	return output, nil
}

// GenerateShort return only small numbers up to 7
func GenerateShort(length int) (string, error) {
	if length > 7 || length <= 0 {
		return fmt.Sprintf("invalid length %d", length), nil
	}

	timeNow := time.Now().Nanosecond()
	g := fmt.Sprint(timeNow)
	if len(g) < length {
		g += g[:1]
	}
	return g[:length], nil
}

// GenerateArray parameters length of one element, amount of elements, some value between 1 and 50 for sleep offset.
// if offset not in the range, it constantly set to 20
func GenerateArray(length int, amount int, offset int) ([]int, error) {
	if offset <= 0 || offset > 50 {
		offset = 20
	}
	arr := make([]int, amount)

	// store elements to slice of integer
	for i := range arr {
		t, err := GenerateShort(length)
		if err != nil {
			return nil, err
		}
		elem, err := strconv.Atoi(t)
		if err != nil {
			return nil, errors.New(t)
		}
		arr[i] = elem

		// calling sleep method
		sleepAmount, _ := strconv.Atoi(t[:1])
		sleepAmount++
		dur := time.Duration(sleepAmount*offset) * time.Nanosecond
		time.Sleep(dur)
	}

	return arr, nil
}
