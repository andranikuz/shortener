package generator

import (
	"github.com/speps/go-hashids/v2"
	"math/rand"
)

func GenerateUniqueID() string {
	hd := hashids.NewData()
	hashids.NewWithData(hd)
	h, _ := hashids.NewWithData(hd)
	e, _ := h.Encode([]int{rand.Intn(1000), rand.Intn(1000), rand.Intn(1000)})

	return e
}
