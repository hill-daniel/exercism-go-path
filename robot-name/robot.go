// Package robotname manages robot factory settings.
package robotname

import (
	"math/rand"
	"strconv"
	"time"
)

const charSet = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"

var rnd = rand.New(rand.NewSource(time.Now().UnixNano()))

var issuedNames = make(map[string]bool)

// Robot represents a robot.
type Robot struct {
	name string
}

// Name returns the name of the robot, or, if it is the first boot, generates a new one
// in the format of two uppercase letters followed by three digits, such as RX837 or BC811.
func (r *Robot) Name() string {
	if len(r.name) == 0 {
		r.name = generate()
	}
	return r.name
}

// Reset wipes the robots name. I.e. calling Name() the next time, will generate a new name.
func (r *Robot) Reset() {
	r.name = ""
}

func generate() string {
	name := createRndLetterStr(2) + createRndDigitsStr(3)
	if issuedNames[name] {
		return generate()
	}
	issuedNames[name] = true
	return name
}

func createRndLetterStr(len int) string {
	var result string
	for i := 0; i < len; i++ {
		result += createRndChar()
	}
	return result
}

func createRndChar() string {
	randomInt := rnd.Intn(len(charSet))
	return charSet[randomInt : randomInt+1]
}

func createRndDigitsStr(len int) string {
	var result string
	for i := 0; i < len; i++ {
		result += strconv.Itoa(rnd.Intn(10))
	}
	return result
}
