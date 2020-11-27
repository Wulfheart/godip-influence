package defaultInfluences

import (
	"github.com/zond/godip"
	"github.com/zond/godip/variants/common"
)

type Default struct {
	Influences map[godip.Nation][]godip.Province
	Variant common.Variant
}
