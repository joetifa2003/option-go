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
	assert.Equal(15, optional.Or(15))
	optional.SetSome(60)
	assert.Equal(60, optional.Or(15))
}

func TestOptionOrElse(t *testing.T) {
	assert := assert.New(t)

	optional := option.NewNone[int]()
	assert.Equal(15, optional.OrElse(func() int { return 15 }))
	optional.SetSome(60)
	assert.Equal(60, optional.OrElse(func() int { return 15 }))
}

func TestOptionIsSomeAnd(t *testing.T) {
	assert := assert.New(t)
	optional := option.NewNone[int]()
	assert.Equal(false, optional.IsSomeAnd(func(i int) bool { return i > 1 }))
	optional.SetSome(-1)
	assert.Equal(false, optional.IsSomeAnd(func(i int) bool { return i > 1 }))
	optional.SetSome(10)
	assert.Equal(true, optional.IsSomeAnd(func(i int) bool { return i > 1 }))
}

func TestOptionReplace(t *testing.T) {
	assert := assert.New(t)
	optional := option.NewNone[int]()
	old := optional.Replace(15)
	assert.Equal(true, old.IsNone())
	assert.Equal(true, optional.IsSome())
	val, _ := optional.Unwrap()
	assert.Equal(15, val)

	old = optional.Replace(30)
	assert.Equal(true, old.IsSome())
	assert.Equal(true, optional.IsSome())
	val, _ = optional.Unwrap()
	assert.Equal(30, val)
}
