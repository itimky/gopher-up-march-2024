package ordersdomain

import (
	"context"
	"math/rand"
	"time"

	"github.com/itimky/gopher-up-march-2024/pkg"
)

type RandomSeedKey struct{}

const RandomIDDigitCount int = 10

func GenerateRandomID(ctx context.Context) pkg.OrderID {
	const charset = "0123456789"

	// for tests
	seed, ok := ctx.Value(RandomSeedKey{}).(int64)
	if !ok {
		seed = time.Now().UnixNano()
	}

	rng := rand.New(rand.NewSource(seed))

	b := make([]byte, 10)
	for i := range b {
		b[i] = charset[rng.Intn(len(charset))]
	}

	return pkg.OrderID(b)
}
