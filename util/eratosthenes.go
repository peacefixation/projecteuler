package util

// SieveOfEratosthenes generate an array of booleans where each index is flagged true if it is prime
func SieveOfEratosthenes(max uint64) []bool {
	// initialise an array of bool set to true
	a := make([]bool, max)
	for i := uint64(2); i <= max; i++ {
		a[i] = true
	}

	// unflag non-primes
	for i := uint64(2); i <= max; i++ {
		if a[i] == true {
			// start at i^2, lower numbers will already be unflagged
			for j := i * i; j <= max; j += i {
				a[j] = false
			}
		}
	}

	return a
}
