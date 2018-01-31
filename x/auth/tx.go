package auth

import (
	"bytes"
	"fmt"
)

// SignedTx represents a transaction that contains signatures,
// which can be verified by the auth.Decorator
type SignedTx interface {
	// GetSignBytes returns the canonical byte representation of the Msg.
	// Equivalent to weave.MustMarshal(tx.GetMsg()) if Msg has a deterministic
	// serialization.
	//
	// Helpful to store original, unparsed bytes here, just in case.
	GetSignBytes() []byte

	// Signatures returns the signature of signers who signed the Msg.
	GetSignatures() []*StdSignature
}

// Validate ensures the StdSignature meets basic standards
func (s *StdSignature) Validate() error {
	seq := s.GetSequence()
	if seq < 0 {
		// TODO: ErrInvalidSequence
		return fmt.Errorf("Sequence is negative")
	}
	if seq == 0 && s.PubKey == nil {
		// TODO: ErrInvalidSignature
		return fmt.Errorf("PubKey missing for sequence 0")
	}
	if s.PubKey == nil && s.Address == nil {
		// TODO: ErrInvalidSignature
		return fmt.Errorf("PubKey or Address is required")
	}
	if s.PubKey != nil && s.Address != nil {
		if !bytes.Equal(s.Address, s.PubKey.Address()) {
			// TODO: ErrInvalidSignature
			return fmt.Errorf("PubKey and Address do not match")
		}
	}

	if s.Signature == nil {
		// TODO: ErrInvalidSignature
		return fmt.Errorf("Signature is missing")
	}

	return nil
}