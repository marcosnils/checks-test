package main

import (
	"context"
	"fmt"

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
	// abc
	if m.Secret != nil {
		fmt.Println(m.Secret.Plaintext(ctx))
	} else {
		fmt.Println("Secret is nil")
	}
	return &CheckStatus{}
}
