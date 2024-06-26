package ed25519

import fmt "fmt"

// String returns Hex representation of a pubkey with it's type
func (pubKey *PubKey) String() string {
	return fmt.Sprintf("PubKeyEd25519{%X}", pubKey.Key)
}
