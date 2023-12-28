package frsProvider

import (
	"fmt"
	"math/rand"
	"time"
)

const DefaultLength = 4

type OTPProvider interface {
	Generate(digits int64) string
}

type GOTP struct{}

func (g *GOTP) Generate(digits int64) string {
	if digits < 1 {
		digits = DefaultLength
	}

	// Initialize the random number generator with the current time
	r := rand.New(rand.NewSource(time.Now().UnixNano()))

	// Calculate the minimum and maximum values based on the number of digits
	min, max := 1, 1
	for i := int64(0); i < digits; i++ {
		min *= 10
		max = max*10 + 9
	}

	// Generate a random OTP within the specified range
	ranges := max - min + 1
	otp := r.Intn(ranges)

	// Ensure that the OTP is exactly n digits
	otp = otp % ranges

	// Format the OTP with leading zeros if needed
	return fmt.Sprintf("%0*d", digits, otp)
}
