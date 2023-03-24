package gen

import (
	"math/rand"
	"time"
)

// GenNamber random number
func GenNamber() int {

	r := rand.New(rand.NewSource(time.Now().UnixNano()))

	return r.Int()
}

// GenNamber random float
func GenFloat() float64 {

	r := rand.New(rand.NewSource(time.Now().UnixNano()))

	return r.Float64()
}

// GenIntDiapason random number between two
func GenIntDiapason(a, b int) int {

	r := rand.New(rand.NewSource(time.Now().UnixNano()))

	return a + r.Intn(b-a+1)

}
