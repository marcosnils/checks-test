package main

import (
	"context"
	"fmt"
	"time"

	"dagger/checks-test/internal/dagger"
)

type ChecksTest struct {
	// +private
	Secret *dagger.Secret
}

func New(
	// +optional
	secret *dagger.Secret,
) *ChecksTest {
	return &ChecksTest{Secret: secret}
}

type CheckStatus struct{}

// +check
// +cache=never
func (m *ChecksTest) CheckMatias(
	ctx context.Context,
) *CheckStatus {
	// one
	// two
	// three
	if m.Secret != nil {
		fmt.Println(m.Secret.Plaintext(ctx))
	} else {
		fmt.Println("Secret is nil")
	}

	time.Sleep(10 * time.Minute)
	return &CheckStatus{}
}
