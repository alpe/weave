package nft

import "github.com/iov-one/weave/errors"

func (m *CreateMsg) Path() string {
	return "multisig/create"
}

func (m *CreateMsg) Validate() error {
	var errs error
	errs = errors.AppendField(errs, "Metadata", m.Metadata.Validate())
	if len(m.Owner) != 0 {
		errs = errors.AppendField(errs, "Owner", m.Owner.Validate())
	}
	valueSet := KeyedValueSet(m.KeyedValues)
	errors.AppendField(errs, "KeyedValues", valueSet.Validate())
	if valueSet.ContainsReservedKeys() {
		errors.AppendField(errs, "KeyedValues", errors.Wrapf(errors.ErrInput, "reserved key namespace %q used", ReservedKeyNamespace))

	}
	return errs

}
