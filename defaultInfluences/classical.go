package defaultInfluences

import (
	"github.com/zond/godip"
	"github.com/zond/godip/variants/classical"
)

// Classical variant
var Classical = Default{
	Influences: map[godip.Nation][]godip.Province{
		godip.Russia:  {"stp", "fin", "mos", "lvn", "war", "ukr", "sev"},
		godip.Turkey:  {"con", "ank", "smy", "arm", "syr"},
		godip.Austria: {"tri", "bud", "gal", "vie", "boh", "tyr"},
		godip.Italy:   {"pic", "ven", "tus", "rom", "apu", "nap"},
		godip.France:  {"mar", "gas", "bur", "par", "bre", "pic"},
		godip.England: {"lon", "wal", "lvp", "yor", "edi", "cly"},
		godip.Germany: {"pru", "ber", "kie", "ruh", "mun", "sil"},
	},
	Variant: classical.ClassicalVariant,
}
