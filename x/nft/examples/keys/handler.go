package keys

import (
	"github.com/iov-one/weave"
	"github.com/iov-one/weave/errors"
	"github.com/iov-one/weave/migration"
	"github.com/iov-one/weave/x"
	"github.com/iov-one/weave/x/nft"
)

// all handlers in this package
func RegisterRoutes(r weave.Registry, auth x.Authenticator) {
	r = migration.SchemaMigratingRegistry("nft", r)
	bucket := nft.NewBucket()
	r.Handle(&nft.CreateMsg{}, MintNftHandler{auth, bucket})
}

// RegisterQuery register queries from buckets in this package
//func RegisterQuery(qr weave.QueryRouter) {
//	NewContractBucket().Register("contracts", qr)
//}

const MintingCost = 100

type MintNftHandler struct {
	auth   x.Authenticator
	bucket *nft.Bucket
}

func (h MintNftHandler) Check(ctx weave.Context, db weave.KVStore, tx weave.Tx) (*weave.CheckResult, error) {
	if _, _, err := h.validate(ctx, db, tx); err != nil {
		return nil, err
	}
	return &weave.CheckResult{GasAllocated: MintingCost}, nil
}

func (h MintNftHandler) Deliver(ctx weave.Context, db weave.KVStore, tx weave.Tx) (*weave.DeliverResult, error) {
	msg, owner, err := h.validate(ctx, db, tx)
	if err != nil {
		return nil, err
	}

	key, err := nft.IDSequence.NextVal(db)
	if err != nil {
		return nil, errors.Wrap(err, "cannot acquire ID")
	}
	_, err = h.bucket.Put(db, key, &nft.Token{
		Metadata:    &weave.Metadata{Schema: 1},
		ID:          key,
		Address:     nft.TokenCondition(key).Address(),
		Owner:       owner,
		Description: msg.Description,
		Image:       msg.Image,
		TokenURI:    msg.TokenURI,
		KeyedValues: msg.KeyedValues,
	})
	if err != nil {
		return nil, errors.Wrap(err, "cannot store revenue")
	}
	return &weave.DeliverResult{Data: key}, nil
}

func (h MintNftHandler) validate(ctx weave.Context, db weave.KVStore, tx weave.Tx) (*nft.CreateMsg, weave.Address, error) {
	var msg nft.CreateMsg
	if err := weave.LoadMsg(tx, &msg); err != nil {
		return nil, nil, errors.Wrap(err, "load msg")
	}
	owner := msg.Owner
	if len(owner) == 0 {
		owner = x.MainSigner(ctx, h.auth).Address()
	} else {
		if !h.auth.HasAddress(ctx, owner) {
			return nil, nil, errors.Wrap(errors.ErrUnauthorized, "must be signed by owner")
		}
	}
	return &msg, owner, nil
}
