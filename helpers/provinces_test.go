package helpers

import (
	"github.com/zond/godip"
	"testing"
)

func TestIsLand(t *testing.T) {
	s := scaffoldVariant(t, "Classical")
	var assertTrue = []godip.Province{"gal", "bud", "war", "mun", "mos"}
	var assertFalse = []godip.Province{"nth", "tys", "ion", "eas", "nwy", "spa", "stp"}
	assertAllInArray(t, assertTrue, true, s.Graph(), IsLand)
	assertAllInArray(t, assertFalse, false, s.Graph(), IsLand)

}

func TestIsSea(t *testing.T) {
s := scaffoldVariant(t, "Classical")
	var assertTrue = []godip.Province{"nth", "tys", "ion", "eas"}
	var assertFalse = []godip.Province{"spa", "stp", "nwy", "gal", "bud", "kie"}
	assertAllInArray(t, assertTrue, true, s.Graph(), IsSea)
	assertAllInArray(t, assertFalse, false, s.Graph(), IsSea)
}

func TestIsCoast(t *testing.T) {
s := scaffoldVariant(t, "Classical")
	var assertTrue = []godip.Province{"spa", "stp", "fin", "por", "smy", "tri", "bul", "naf", "tun"}
	var assertFalse = []godip.Province{"nth", "tys", "ion", "eas", "gal", "bud", "war", "muc", "mos"}
	assertAllInArray(t, assertTrue, true, s.Graph(), IsCoast)
	assertAllInArray(t, assertFalse, false, s.Graph(), IsCoast)
}

