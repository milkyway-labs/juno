package hd

import fmt "fmt"

// String returns the full absolute HD path of the BIP44 (https://github.com/bitcoin/bips/blob/master/bip-0044.mediawiki) params:
// m / purpose' / coin_type' / account' / change / address_index
func (p BIP44Params) String() string {
	var changeStr string
	if p.Change {
		changeStr = "1"
	} else {
		changeStr = "0"
	}
	return fmt.Sprintf("m/%d'/%d'/%d'/%s/%d",
		p.Purpose,
		p.CoinType,
		p.Account,
		changeStr,
		p.AddressIndex)
}
