package local_types

import (
	"go4.org/netipx"
)

type lRange struct {
	*netipx.IPSet
}

type AddressSet struct {
	Subnet, GW, BC, First, Last string
}
