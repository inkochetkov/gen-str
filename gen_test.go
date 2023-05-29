package gen

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"golang.org/x/crypto/bcrypt"
)

func TestBcrypt(t *testing.T) {
	want := "test"

	have, err := GenPassword(want)
	assert.NoError(t, err)

	err = bcrypt.CompareHashAndPassword([]byte(have), []byte(want))
	assert.NoError(t, err)
}

func TestNumber(t *testing.T) {
	i := GenNamber()
	t.Log(i)

	j := GenFloat()
	t.Log(j)

	n := GenIntDiapason(1, 2)
	t.Log(n)
}

func TestString(t *testing.T) {
	str := GenRandomStrRune(1)
	t.Log(str)

	str = GenRandomStrSpec(2)
	t.Log(str)
}

func TestUUID(t *testing.T) {
	str, err := GenUUID()
	assert.NoError(t, err)

	assert.Equal(t, len(str), 36)

	t.Log(str)
}
