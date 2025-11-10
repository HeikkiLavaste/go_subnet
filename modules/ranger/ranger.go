package ranger

import (
	"net/netip"
	"regexp"

	local_types "github.com/heikkilavaste/go_subnet/modules/types"
	"go4.org/netipx"
)

// var b netipx.IPSetBuilder
//var pList []string

type lRange struct {
	*netipx.IPSet
}

func NewRange(strRange []string) lRange {
	var rt lRange
	for _, r := range strRange {
		b := netipx.IPSetBuilder{}
		b.AddPrefix(netip.MustParsePrefix(r))
		out, _ := b.IPSet()
		rt = lRange{out}
	}

	return rt
}

func (n *lRange) BreakDown(e []string, sn uint8) []string {
	//r, _ := regexp.Compile(`(\d+\.\d+\.\d+\.\d+)`)
	var pList []string
	if sn == 0 {
		sn = uint8(n.Prefixes()[0].Bits())
	}
	for {
		p, nw, ok := n.RemoveFreePrefix(sn)
		if ok {
			pList = append(pList, p.String())
			n = &lRange{nw}
		} else {
			return pList
		}

	}
}

func (n *lRange) Parse() local_types.AddressSet {
	r, _ := regexp.Compile(`(\d+\.\d+\.\d+\.\d+)`)
	Set := local_types.AddressSet{}
	Set.Subnet = n.Prefixes()[0].String()
	Set.GW = n.Ranges()[0].From().Next().String()
	Set.BC = n.Ranges()[0].To().String()
	Set.Last = n.Ranges()[0].To().Prev().String()

	all := n.BreakDown(nil, 32)
	switch {
	case n.Prefixes()[0].Bits() <= 24:
		Set.First = r.FindString(all[10])
	case n.Prefixes()[0].Bits() > 24:
		Set.First = r.FindString(all[7])
	}
	return Set
}

func (n *lRange) WriteToConsole(Set []local_types.AddressSet) {

}
