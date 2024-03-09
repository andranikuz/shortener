package generator

import (
	"github.com/speps/go-hashids/v2"
	"math/rand/v2"
)

func GenerateUniqueId() string {
	hd := hashids.NewData()
	hashids.NewWithData(hd)
	h, _ := hashids.NewWithData(hd)
	e, _ := h.Encode([]int{rand.IntN(99999), rand.IntN(99999)})

	return e
}
