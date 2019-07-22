package nft

import (
	"github.com/iov-one/weave"
	"github.com/iov-one/weave/x"
)

var _ weave.Decorator = Decorator{}

const (
	authMarker      = ReservedKeyNamespace + ".authenticateable"
	ownerVerifyCost = 1000
)

type Decorator struct {
	signers x.Authenticator
	bucket  *Bucket
}

func NewDecorator(auth x.Authenticator) Decorator {
	return Decorator{auth, NewBucket()}
}

// Check adds authN nfts of the signers before calling down the stack.
func (d Decorator) Check(ctx weave.Context, db weave.KVStore, tx weave.Tx, next weave.Checker) (*weave.CheckResult, error) {
	conds, err := d.authTokens(ctx, db)
	if err != nil {
		return nil, err
	}
	ctx = withTokens(ctx, conds)
	res, err := next.Check(ctx, db, tx)
	if err != nil {
		return nil, err
	}
	// We must charge gas proportionally to the effort. We only charge for the
	// nfts found by the signatures. Non matching nfts are ignored.
	res.GasPayment += int64(len(conds) * ownerVerifyCost)
	return res, nil
}

// Deliver adds authN nfts of the signers before calling down the stack.
func (d Decorator) Deliver(ctx weave.Context, db weave.KVStore, tx weave.Tx, next weave.Deliverer) (*weave.DeliverResult, error) {
	conds, err := d.authTokens(ctx, db)
	if err != nil {
		return nil, err
	}

	ctx = withTokens(ctx, conds)
	return next.Deliver(ctx, db, tx)
}

func (d Decorator) authTokens(ctx weave.Context, db weave.KVStore) ([]weave.Condition, error) {
	var conds []weave.Condition
	for _, v := range d.signers.GetConditions(ctx) {
		tokens, err := d.bucket.ByAuthNIndex(db, v.Address())
		if err != nil {
			return nil, err
		}
		for _, token := range tokens {
			conds = append(conds, TokenCondition(token.ID))
			ext, _ := KeyedValueSet(token.KeyedValues).Get(authMarker)
			if len(ext) != 0 {
				conds = append(conds, ExtensionCondition(ext))
			}
		}
	}
	return conds, nil
}
