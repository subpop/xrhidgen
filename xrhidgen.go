// package xrhidgen generates X-Rh-Identity JSON records suitable for passing
// into HTTP requests to console.redhat.com services.
package xrhidgen

import "github.com/pioz/faker"

// SetSeed uses the provided seed value to initialize the generator to a
// deterministic state (see rand.Seed).
func SetSeed(seed int64) {
	faker.SetSeed(seed)
}
