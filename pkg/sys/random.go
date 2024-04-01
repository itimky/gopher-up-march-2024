package sys

import (
	"context"
	"crypto/rand"
	"fmt"
)

type randomIDKey struct{}

const RandomIDDigitCount int = 10

func GenerateRandomID(ctx context.Context) (string, error) {
	randomID, ok := ctx.Value(randomIDKey{}).(string)
	if !ok {
		randomBytes := make([]byte, RandomIDDigitCount)
		_, err := rand.Read(randomBytes)
		if err != nil {
			return "", fmt.Errorf("rand read: %w", err)
		}

		// Convert each byte to its decimal representation and append to the result
		var result string
		for _, b := range randomBytes {
			// Convert byte to decimal (0-255) and take modulus to get a single decimal digit (0-9)
			decimalDigit := b % 10
			result += fmt.Sprintf("%d", decimalDigit)
		}

		randomID = result
	}

	return randomID, nil
}

func RandomIDToCtx(ctx context.Context, randomID string) context.Context {
	return context.WithValue(ctx, randomIDKey{}, randomID)
}
