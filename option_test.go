package option_test

import (
	"testing"

	"github.com/joetifa2003/option-go"
	"github.com/stretchr/testify/assert"
)

func TestOptionNone(t *testing.T) {
	assert := assert.New(t)

	optional := option.NewNone[int]()
	assert.Equal(optional.IsSome(), false)
	assert.Equal(optional.IsNone(), true)

	value, isSome := optional.Unwrap()
	assert.Equal(false, isSome)
	assert.Equal(0, value) // Zero value
}

func TestOptionSome(t *testing.T) {
	assert := assert.New(t)

	optional := option.NewSome(15)
	assert.Equal(optional.IsSome(), true)
	assert.Equal(optional.IsNone(), false)

	value, isSome := optional.Unwrap()
	assert.Equal(true, isSome)
	assert.Equal(15, value)
}

func TestOptionSetSome(t *testing.T) {
	assert := assert.New(t)

	optional := option.NewNone[int]()
	optional.SetSome(15)
	assert.Equal(true, optional.IsSome())
	assert.Equal(false, optional.IsNone())

	value, isSome := optional.Unwrap()
	assert.Equal(true, isSome)
	assert.Equal(15, value)
}

func TestOptionSetNone(t *testing.T) {
	assert := assert.New(t)

	optional := option.NewSome(15)
	optional.SetNone()
	assert.Equal(optional.IsSome(), false)
	assert.Equal(optional.IsNone(), true)

	value, isSome := optional.Unwrap()
	assert.Equal(false, isSome)
	assert.Equal(0, value)
}

func TestOptionOr(t *testing.T) {
	assert := assert.New(t)

	optional := option.NewNone[int]()
	assert.Equal(15, optional.OrElse(15))
	optional.SetSome(60)
	assert.Equal(60, optional.OrElse(15))
}
