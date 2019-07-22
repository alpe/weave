package nft

import (
	"github.com/iov-one/weave"
	"github.com/iov-one/weave/errors"
	"github.com/iov-one/weave/migration"
	"github.com/iov-one/weave/orm"
	"github.com/iov-one/weave/store"
)

type Bucket struct {
	orm.ModelBucket
}

func NewBucket() *Bucket {
	b := orm.NewModelBucket("nft", &Token{},
		orm.WithIDSequence(IDSequence),
		orm.WithIndex("owner", ownerIndexer, false),
		orm.WithIndex("authN", authNIndexer, false),
	)
	return &Bucket{migration.NewModelBucket("nft", b)}
}

func (b Bucket) ByAuthNIndex(db store.KVStore, owner weave.Address) ([]Token, error) {
	var result []Token
	_, err := b.ByIndex(db, "authN", owner, result)
	return result, errors.Wrap(err, "failed to find by authN index")
}

func ownerIndexer(obj orm.Object) ([]byte, error) {
	nft, err := toNonFungibleToken(obj)
	if err != nil {
		return nil, err
	}
	return nft.Owner, nil
}

// authNIndexer is a indexes only owners for nfts that contains authMarker data
func authNIndexer(obj orm.Object) ([]byte, error) {
	nft, err := toNonFungibleToken(obj)
	if err != nil {
		return nil, err
	}
	if _, exists := KeyedValueSet(nft.KeyedValues).Get(authMarker); !exists {
		return nil, nil
	}
	return nft.Owner, nil
}

func toNonFungibleToken(obj orm.Object) (*Token, error) {
	if obj == nil {
		return nil, errors.Wrap(errors.ErrHuman, "can not take index of nil")
	}
	nft, ok := obj.Value().(*Token)
	if !ok {
		return nil, errors.Wrap(errors.ErrHuman, "can only take index of Token")
	}
	return nft, nil
}

var IDSequence = orm.NewSequence("nft", "id")
