package godipInfluence

import (
	"github.com/zond/godip"
	"github.com/zond/godip/state"
)

type Influence map[godip.Province]godip.Nation

type InfluenceCalculator func (old Influence, new *state.State) Influence


