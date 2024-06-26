package secp256k1

import fmt "fmt"

func (pubKey *PubKey) String() string {
	return fmt.Sprintf("PubKeySecp256k1{%X}", pubKey.Key)
}
