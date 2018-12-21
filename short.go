/*
short is a Go port of http://code.activestate.com/recipes/576918/, which is a
Python implementation for generating Tiny URL- and bit.ly-like URLs.
*/
package short

import (
	"fmt"
	"math"
	"strings"
)

const (
	alphabet  = "mn6j2c4rv8bpygw95z7hsdaetxuk3fq"
	blockSize = 24
	minLength = 5
	mask      = (1 << blockSize) - 1
)

var (
	mapping = initMapping()
)

// Encode generates a minimum 5-characters long representation of given n.
// Given number could be auto-incremented integers used in relational
// databases.
func Encode(n int) string {
	return enbase(encode(n))
}

// Decode returns the number represented by the short string s.
func Decode(s string) int {
	return decode(debase(s))
}

func _enbase(x int) string {
	n := len(alphabet)
	if x < n {
		return string(alphabet[x])
	}
	div := x / n
	mod := x % n
	return _enbase(div) + string(alphabet[mod])
}

func enbase(x int) string {
	r := _enbase(x)

	repeat := minLength - len(r)
	if repeat < 0 {
		repeat = 0
	}

	padding := strings.Repeat(string(alphabet[0]), repeat)
	return fmt.Sprintf("%v%v", padding, r)
}

func encode(n int) int {
	x := n & ^mask
	y := _encode(n & mask)

	r := x | y
	return r
}

func _encode(n int) int {
	var result int
	for i, v := range mapping {
		if n&(1<<uint(i)) != 0 {
			result |= (1 << uint(v))
		}
	}
	return result
}

func decode(n int) int {
	x := n & ^mask
	y := _decode(n & mask)
	return x | y
}

func _decode(n int) int {
	var result int
	for i, v := range mapping {
		if n&(1<<uint(v)) > 0 {
			result |= (1 << uint(i))
		}
	}
	return result
}

func debase(s string) int {
	n := len(alphabet)
	var result int
	for i, r := range reverse(s) {
		result += strings.IndexRune(alphabet, r) * int(math.Pow(float64(n), float64(i)))
	}
	return result
}

func initMapping() [blockSize]int {
	var mapping [blockSize]int
	for i := 0; i < blockSize; i++ {
		mapping[i] = blockSize - i - 1
	}
	return mapping
}

// reverse returns its argument string reversed rune-wise left to right.
func reverse(s string) string {
	r := []rune(s)
	for i, j := 0, len(r)-1; i < len(r)/2; i, j = i+1, j-1 {
		r[i], r[j] = r[j], r[i]
	}
	return string(r)
}
