package helpers

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"github.com/zond/godip"
	"github.com/zond/godip/state"
	"github.com/zond/godip/variants"
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

func assertAllInArray(t *testing.T, arr []godip.Province, expected bool, g godip.Graph,  f func(graph godip.Graph, n godip.Province) bool){
	for _, province := range arr {
		assert.Equal(t, expected, f(g, province), fmt.Sprintf("Expected %s province to be %v", province, expected))
	}
}
func scaffoldVariant(t *testing.T, variantName string) (s *state.State){
	variant, found := variants.Variants[variantName]
	if !found {
		t.Fatal("Variant", variantName, "not found")
	}
	s, err := variant.Start()
	if err != nil {
		t.Fatal(err.Error())
	}
	return
}
