package xrhidgen

import "github.com/pioz/faker"

// SetSeed uses the provided seed value to initialize the generator to a
// deterministic state (see rand.Seed).
func SetSeed(seed int64) {
	faker.SetSeed(seed)
}
