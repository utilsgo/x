package types

import (
	"testing"

	. "github.com/onsi/gomega"
)

func TestTypeFor(t *testing.T) {
	cases := []string{
		"string",
		"int",
		"map[int]int",
		"[]int",
		"[2]int",
		"error",

		"github.com/utilsgo/x/types/testdata/typ.String",
		"github.com/utilsgo/x/types/testdata/typ.AnyMap[int,string]",
	}

	for i := range cases {
		c := cases[i]
		NewWithT(t).Expect(FromTType(TypeFor(c)).String()).To(Equal(c))
	}
}
