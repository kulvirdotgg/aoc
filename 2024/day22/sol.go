package day22

import (
	"aoc/stl"
	"fmt"
)

func Solution() {
	ip := stl.ReadFile("input.txt")

	var secretSum int64 = 0
	for _, num := range ip {
		secret := int64(stl.IntsFromString(num)[0])
		secretSum += calculateFinal(secret)
	}
	fmt.Println(secretSum)
}

const (
	MOD = 16777216
)

func nextSecret(secret int64) int64 {
	one := secret * 64
	one = (one ^ secret) % MOD

	two := one / 32
	two = (two ^ one) % MOD

	three := two * 2048
	three = (three ^ two) % MOD

	return three
}

func calculateFinal(secret int64) int64 {
	for range 2000 {
		secret = nextSecret(secret)
	}
	return secret
}
