package godipInfluence

import "github.com/zond/godip"

type Influence map[godip.Province]godip.Nation

type InfluenceCalculator func (old Influence, new godip.State) Influence


