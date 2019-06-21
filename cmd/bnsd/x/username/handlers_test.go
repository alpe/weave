package username

import (
	"context"
	"testing"

	"github.com/iov-one/weave"
	"github.com/iov-one/weave/errors"
	"github.com/iov-one/weave/migration"
	"github.com/iov-one/weave/store"
	"github.com/iov-one/weave/weavetest"
	"github.com/iov-one/weave/weavetest/assert"
	"github.com/iov-one/weave/x"
)

func TestRegisterUsernameTokenHandler(t *testing.T) {
	var (
		aliceCond = weavetest.NewCondition()
		bobbyCond = weavetest.NewCondition()
	)

	cases := map[string]struct {
		Tx             weave.Tx
		Auth           x.Authenticator
		WantCheckErr   *errors.Error
		WantDeliverErr *errors.Error
	}{
		"success": {
			Tx: &weavetest.Tx{
				Msg: &RegisterUsernameTokenMsg{
					Metadata: &weave.Metadata{Schema: 1},
					Username: "bobby*iov",
					Targets: []BlockchainAddress{
						{BlockchainID: "bc_1", Address: []byte("addr1")},
						{BlockchainID: "bc_2", Address: []byte("addr2")},
					},
				},
			},
			Auth: &weavetest.Auth{Signer: bobbyCond},
		},
		"username must be unique": {
			Tx: &weavetest.Tx{
				Msg: &RegisterUsernameTokenMsg{
					Metadata: &weave.Metadata{Schema: 1},
					Username: "alice*iov",
					Targets: []BlockchainAddress{
						{BlockchainID: "bc_1", Address: []byte("addr1")},
					},
				},
			},
			Auth:           &weavetest.Auth{Signer: aliceCond},
			WantCheckErr:   errors.ErrDuplicate,
			WantDeliverErr: errors.ErrDuplicate,
		},
		"target cannot be empty": {
			Tx: &weavetest.Tx{
				Msg: &RegisterUsernameTokenMsg{
					Metadata: &weave.Metadata{Schema: 1},
					Username: "alice*iov",
					Targets:  []BlockchainAddress{},
				},
			},
			Auth:           &weavetest.Auth{Signer: aliceCond},
			WantCheckErr:   errors.ErrEmpty,
			WantDeliverErr: errors.ErrEmpty,
		},
		"username must be provided": {
			Tx: &weavetest.Tx{
				Msg: &RegisterUsernameTokenMsg{
					Metadata: &weave.Metadata{Schema: 1},
					Username: "",
					Targets: []BlockchainAddress{
						{BlockchainID: "bc_1", Address: []byte("addr1")},
					},
				},
			},
			Auth:           &weavetest.Auth{Signer: aliceCond},
			WantCheckErr:   errors.ErrInput,
			WantDeliverErr: errors.ErrInput,
		},
	}

	for testName, tc := range cases {
		t.Run(testName, func(t *testing.T) {
			db := store.MemStore(179)
			migration.MustInitPkg(db, "username")

			b := NewUsernameTokenBucket()
			_, err := b.Put(db, []byte("alice*iov"), &UsernameToken{
				Metadata: &weave.Metadata{Schema: 1},
				Targets: []BlockchainAddress{
					{BlockchainID: "unichain", Address: []byte("some-unichain-address")},
				},
				Owner: aliceCond.Address(),
			})
			assert.Nil(t, err)

			h := registerUsernameTokenHandler{
				auth:   tc.Auth,
				bucket: b,
			}

			cache := db.CacheWrap()
			if _, err := h.Check(context.TODO(), cache, tc.Tx); !tc.WantCheckErr.Is(err) {
				t.Fatalf("unexpected check error: %s", err)
			}
			cache.Discard()
			if _, err := h.Deliver(context.TODO(), db, tc.Tx); !tc.WantDeliverErr.Is(err) {
				t.Fatalf("unexpected deliver error: %s", err)
			}
		})
	}
}

func TestChangeUsernameTokenOwnerHandler(t *testing.T) {
	var (
		aliceCond = weavetest.NewCondition()
		bobbyCond = weavetest.NewCondition()
	)

	cases := map[string]struct {
		Tx             weave.Tx
		Auth           x.Authenticator
		WantCheckErr   *errors.Error
		WantDeliverErr *errors.Error
	}{
		"success": {
			Tx: &weavetest.Tx{
				Msg: &TransferUsernameTokenMsg{
					Metadata: &weave.Metadata{Schema: 1},
					Username: "alice*iov",
					NewOwner: bobbyCond.Address(),
				},
			},
			Auth: &weavetest.Auth{Signer: aliceCond},
		},
		"only the owner can change the token": {
			Tx: &weavetest.Tx{
				Msg: &TransferUsernameTokenMsg{
					Metadata: &weave.Metadata{Schema: 1},
					Username: "alice*iov",
					NewOwner: bobbyCond.Address(),
				},
			},
			WantCheckErr:   errors.ErrUnauthorized,
			WantDeliverErr: errors.ErrUnauthorized,
			Auth:           &weavetest.Auth{Signer: bobbyCond},
		},
		"token must exist": {
			Tx: &weavetest.Tx{
				Msg: &TransferUsernameTokenMsg{
					Metadata: &weave.Metadata{Schema: 1},
					Username: "does-not-exist*iov",
					NewOwner: bobbyCond.Address(),
				},
			},
			WantCheckErr:   errors.ErrNotFound,
			WantDeliverErr: errors.ErrNotFound,
			Auth:           &weavetest.Auth{Signer: bobbyCond},
		},
		"change to the same owner (no change) is allowed": {
			Tx: &weavetest.Tx{
				Msg: &TransferUsernameTokenMsg{
					Metadata: &weave.Metadata{Schema: 1},
					Username: "alice*iov",
					NewOwner: aliceCond.Address(),
				},
			},
			Auth: &weavetest.Auth{Signer: aliceCond},
		},
	}
	for testName, tc := range cases {
		t.Run(testName, func(t *testing.T) {
			db := store.MemStore(179)
			migration.MustInitPkg(db, "username")

			b := NewUsernameTokenBucket()
			_, err := b.Put(db, []byte("alice*iov"), &UsernameToken{
				Metadata: &weave.Metadata{Schema: 1},
				Targets: []BlockchainAddress{
					{BlockchainID: "unichain", Address: []byte("some-unichain-address")},
				},
				Owner: aliceCond.Address(),
			})
			assert.Nil(t, err)

			h := transferUsernameTokenHandler{
				auth:   tc.Auth,
				bucket: b,
			}

			cache := db.CacheWrap()
			if _, err := h.Check(context.TODO(), cache, tc.Tx); !tc.WantCheckErr.Is(err) {
				t.Fatalf("unexpected check error: %s", err)
			}
			cache.Discard()
			if _, err := h.Deliver(context.TODO(), db, tc.Tx); !tc.WantDeliverErr.Is(err) {
				t.Fatalf("unexpected deliver error: %s", err)
			}
		})
	}
}

func TestChangeUsernameTokenTargetHandler(t *testing.T) {
	var (
		aliceCond = weavetest.NewCondition()
		bobbyCond = weavetest.NewCondition()
	)
	cases := map[string]struct {
		Tx             weave.Tx
		Auth           x.Authenticator
		WantCheckErr   *errors.Error
		WantDeliverErr *errors.Error
	}{
		"success": {
			Tx: &weavetest.Tx{
				Msg: &ChangeUsernameTokenTargetsMsg{
					Metadata: &weave.Metadata{Schema: 1},
					Username: "alice*iov",
					NewTargets: []BlockchainAddress{
						{BlockchainID: "hydracoin", Address: []byte("some-hydra-address")},
						{BlockchainID: "pegasuscoin", Address: []byte("some-pagasus-address")},
					},
				},
			},
			Auth: &weavetest.Auth{Signer: aliceCond},
		},
		"can change target to the same value (no change)": {
			Tx: &weavetest.Tx{
				Msg: &ChangeUsernameTokenTargetsMsg{
					Metadata: &weave.Metadata{Schema: 1},
					Username: "alice*iov",
					NewTargets: []BlockchainAddress{
						{BlockchainID: "unichain", Address: []byte("some-unicorn-address")},
					},
				},
			},
			Auth: &weavetest.Auth{Signer: aliceCond},
		},
		"invalid message": {
			Tx: &weavetest.Tx{
				Msg: &ChangeUsernameTokenTargetsMsg{
					Metadata:   nil,
					Username:   "",
					NewTargets: []BlockchainAddress{},
				},
			},
			Auth:           &weavetest.Auth{Signer: aliceCond},
			WantCheckErr:   errors.ErrMetadata,
			WantDeliverErr: errors.ErrMetadata,
		},
		"only the owner can change the token": {
			Tx: &weavetest.Tx{
				Msg: &ChangeUsernameTokenTargetsMsg{
					Metadata: &weave.Metadata{Schema: 1},
					Username: "alice*iov",
					NewTargets: []BlockchainAddress{
						{BlockchainID: "hydracoin", Address: []byte("some-hydra-address")},
					},
				},
			},
			WantCheckErr:   errors.ErrUnauthorized,
			WantDeliverErr: errors.ErrUnauthorized,
			Auth:           &weavetest.Auth{Signer: bobbyCond},
		},
		"token must exist": {
			Tx: &weavetest.Tx{
				Msg: &ChangeUsernameTokenTargetsMsg{
					Metadata: &weave.Metadata{Schema: 1},
					Username: "does-not-exist*iov",
					NewTargets: []BlockchainAddress{
						{BlockchainID: "hydracoin", Address: []byte("some-hydra-address")},
					},
				},
			},
			WantCheckErr:   errors.ErrNotFound,
			WantDeliverErr: errors.ErrNotFound,
			Auth:           &weavetest.Auth{Signer: bobbyCond},
		},
	}
	for testName, tc := range cases {
		t.Run(testName, func(t *testing.T) {
			db := store.MemStore(179)
			migration.MustInitPkg(db, "username")

			b := NewUsernameTokenBucket()
			_, err := b.Put(db, []byte("alice*iov"), &UsernameToken{
				Metadata: &weave.Metadata{Schema: 1},
				Targets: []BlockchainAddress{
					{BlockchainID: "unichain", Address: []byte("some-unicorn-address")},
				},
				Owner: aliceCond.Address(),
			})
			assert.Nil(t, err)

			h := changeUsernameTokenTargetsHandler{
				auth:   tc.Auth,
				bucket: b,
			}

			cache := db.CacheWrap()
			if _, err := h.Check(context.TODO(), cache, tc.Tx); !tc.WantCheckErr.Is(err) {
				t.Fatalf("unexpected check error: %s", err)
			}
			cache.Discard()
			if _, err := h.Deliver(context.TODO(), db, tc.Tx); !tc.WantDeliverErr.Is(err) {
				t.Fatalf("unexpected deliver error: %s", err)
			}
		})
	}
}
