package ownrsa

import "math/rand"

func randInt(min int, max int) int {
	r := rand.Intn(max-min) + min
	return r
}
func randPrime(min int, max int) int {
	primes := sieveOfEratosthenes(max)

	randN := rand.Intn(len(primes)-0) + 0

	return primes[randN]

}

// return list of primes less than N
func sieveOfEratosthenes(N int) (primes []int) {
	b := make([]bool, N)
	for i := 2; i < N; i++ {
		if b[i] == true {
			continue
		}
		primes = append(primes, i)
		for k := i * i; k < N; k += i {
			b[k] = true
		}
	}
	return
}

func gcd(a, b int) int {
	var bgcd func(a, b, res int) int

	bgcd = func(a, b, res int) int {
		switch {
		case a == b:
			return res * a
		case a%2 == 0 && b%2 == 0:
			return bgcd(a/2, b/2, 2*res)
		case a%2 == 0:
			return bgcd(a/2, b, res)
		case b%2 == 0:
			return bgcd(a, b/2, res)
		case a > b:
			return bgcd(a-b, b, res)
		default:
			return bgcd(a, b-a, res)
		}
	}

	return bgcd(a, b, 1)
}
