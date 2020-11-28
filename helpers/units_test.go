package helpers

import (
	"github.com/zond/godip"
	"testing"
)

func TestCanHostArmy(t *testing.T) {
	s := scaffoldVariant(t, "Classical")
	var assertTrue = []godip.Province{"gal", "bud", "war", "mun", "mos", "spa", "stp", "stp/sc", "bul", "smy", "tun", "nwy"}
	var assertFalse = []godip.Province{"nth", "tys", "ion", "eas"}
	assertAllInArray(t, assertTrue, true, s.Graph(), CanHostArmy)
	assertAllInArray(t, assertFalse, false, s.Graph(), CanHostArmy)
}

func TestCanHostFleet(t *testing.T) {
	s := scaffoldVariant(t, "Classical")
	var assertTrue = []godip.Province{"spa", "stp", "stp/sc", "bul", "smy", "tun", "nth", "tys", "ion", "eas", "nwy"}
	var assertFalse = []godip.Province{"gal", "bud", "war", "mun", "mos"}
	assertAllInArray(t, assertTrue, true, s.Graph(), CanHostFleet)
	assertAllInArray(t, assertFalse, false, s.Graph(), CanHostFleet)
}