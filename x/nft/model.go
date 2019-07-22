package nft

import (
	"regexp"
	"strings"

	"github.com/iov-one/weave"
	"github.com/iov-one/weave/errors"
	"github.com/iov-one/weave/migration"
	"github.com/iov-one/weave/orm"
)

func init() {
	migration.MustRegister(1, &Token{}, migration.NoModification)
}

var _ orm.CloneableData = (*Token)(nil)

func (n *Token) Validate() error {
	var errs error
	errs = errors.AppendField(errs, "Metadata", n.Metadata.Validate())
	errs = errors.AppendField(errs, "Owner", n.Owner.Validate())
	errs = errors.AppendField(errs, "Address", n.Address.Validate())

	errors.AppendField(errs, "KeyedValues", KeyedValueSet(n.KeyedValues).Validate())

	return errs
}

const ReservedKeyNamespace = "_weave.nft"

const maxKeysCount = 10

type KeyedValueSet []KeyedValue

func (s KeyedValueSet) Validate() error {
	if len(s) > maxKeysCount {
		return errors.Wrapf(errors.ErrLimit, "max count %d", maxKeysCount)
	}
	var errs error
	index := make(map[string]struct{})
	for _, v := range s {
		if _, exists := index[v.Key]; exists {
			errs = errors.AppendField(errs, "KeyedValues", errors.Wrapf(errors.ErrDuplicate, "key exists: %q", v.Key))
		}
	}
	for _, v := range s {
		errors.Append(errs, v.Validate())
	}
	return errs
}

func (s KeyedValueSet) ContainsReservedKeys() bool {
	for _, v := range s {
		if strings.HasPrefix(v.Key, ReservedKeyNamespace) {
			return true
		}
	}
	return false
}

func (s KeyedValueSet) Copy() KeyedValueSet {
	kv := make(KeyedValueSet, len(s))
	for i, v := range s {
		kv[i] = v
	}
	return kv
}

func (s KeyedValueSet) Get(key string) ([]byte, bool) {
	for _, v := range s {
		if v.Key == key {
			return v.Value, true
		}
	}
	return nil, false

}

const maxValueLength = 100

var validKey = regexp.MustCompile(`^[a-z-9_.-]{1,64}$`).MatchString

func (n KeyedValue) Validate() error {
	if len(n.Value) > maxValueLength {
		return errors.Wrapf(errors.ErrLimit, "max length %d", maxValueLength)
	}
	if !validKey(n.Key) {
		return errors.Wrap(errors.ErrInput, "invalid key")
	}
	return nil
}

func (n *Token) Copy() orm.CloneableData {
	return &Token{
		Metadata:    n.Metadata.Copy(),
		ID:          n.ID,
		Address:     n.Address.Clone(),
		Owner:       n.Owner.Clone(),
		Description: n.Description,
		Image:       n.Image,
		TokenURI:    n.TokenURI,
		KeyedValues: KeyedValueSet(n.KeyedValues).Copy(),
	}
}

func TokenCondition(key []byte) weave.Condition {
	return weave.NewCondition("nft", "token", key)
}

func ExtensionCondition(ext []byte) weave.Condition {
	return weave.NewCondition("nft", "custom", ext)
}
