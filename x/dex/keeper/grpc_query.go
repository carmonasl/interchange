package keeper

import (
	"github.com/carmonasl/interchange/x/dex/types"
)

var _ types.QueryServer = Keeper{}
