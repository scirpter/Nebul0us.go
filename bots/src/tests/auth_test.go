package tests

import (
	"fmt"
	"neb/src/auth"
	"testing"
)

func TestAuth(t *testing.T) {
	gateway := auth.NewAuthGateway()

	for _, account := range gateway.Accounts {
		fmt.Printf("Account: %v", account.Ranks)
	}
}
