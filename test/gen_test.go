package test

import (
	"testing"

	"github.com/inkochetkov/gen-str/pkg/gen"
	"github.com/stretchr/testify/assert"
	"golang.org/x/crypto/bcrypt"
)

func TestBcrypt(t *testing.T) {
	want := "test"

	have, err := gen.GenPassword(want)
	assert.NoError(t, err)

	err = bcrypt.CompareHashAndPassword([]byte(have), []byte(want))
	assert.NoError(t, err)
}

func TestNumber(t *testing.T) {
	i := gen.GenNamber()
	t.Log(i)

	j := gen.GenFloat()
	t.Log(j)

	n := gen.GenIntDiapason(1, 2)
	t.Log(n)
}

func TestString(t *testing.T) {
	str := gen.GenRandomStrRune(1)
	t.Log(str)

	str = gen.GenRandomStrSpec(2)
	t.Log(str)
}

func TestUUID(t *testing.T) {
	str, err := gen.GenUUID()
	assert.NoError(t, err)

	assert.Equal(t, len(str), 36)

	t.Log(str)
}
