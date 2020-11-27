package godip_influence

import "github.com/zond/godip"

type Influence map[godip.Province]godip.Nation

type InfluenceCalculator func (old godip.State, new godip.State) Influence
