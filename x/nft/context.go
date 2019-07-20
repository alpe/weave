package nft

import (
	"context"

	"github.com/iov-one/weave"
	"github.com/iov-one/weave/x"
)

// local to the nft module
type contextKey int

const (
	contextOwnedAuthNTokens contextKey = iota
)

// withTokens is a private method, as only this module
// can add a signer
func withTokens(ctx weave.Context, owners []weave.Condition) weave.Context {
	return context.WithValue(ctx, contextOwnedAuthNTokens, owners)
}

// Authenticate implements x.Authenticator and provides
// authentication based on nft ownership
type Authenticate struct{}

var _ x.Authenticator = Authenticate{}

func (a Authenticate) GetConditions(ctx weave.Context) []weave.Condition {
	// todo: revisit and test with big numbers of tokens

	// (val, ok) form to return nil instead of panic if unset
	val, _ := ctx.Value(contextOwnedAuthNTokens).([]weave.Condition)
	// if we were paranoid about our own code, we would deep-copy
	// the signers here
	return val
}

// HasAddress returns true if the current owner for the the given nft address
// has signed the transaction.
func (a Authenticate) HasAddress(ctx weave.Context, addr weave.Address) bool {
	conds := a.GetConditions(ctx)
	for _, s := range conds {
		if addr.Equals(s.Address()) {
			return true
		}
	}
	return false
}
