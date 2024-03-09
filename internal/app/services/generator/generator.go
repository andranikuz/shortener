package generator

import (
	"github.com/speps/go-hashids/v2"
	"math/rand"
)

func GenerateUniqueID() string {
	hd := hashids.NewData()
	hashids.NewWithData(hd)
	h, _ := hashids.NewWithData(hd)
	e, _ := h.Encode([]int{rand.Int(), rand.Int()})

	return e
}
