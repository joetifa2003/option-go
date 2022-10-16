package option

// Option a generic option type
type Option[T any] struct {
	value  T
	isSome bool
}

// NewSome creates an option with a value (Some)
func NewSome[T any](value T) Option[T] {
	return Option[T]{
		value:  value,
		isSome: true,
	}
}

// NewNone creates an option without a value (None)
func NewNone[T any]() Option[T] {
	return Option[T]{
		isSome: false,
	}
}

// SetSome sets an option to a value (Some)
func (o *Option[T]) SetSome(value T) {
	o.value = value
	o.isSome = true
}

// SetNone sets an option to none
func (o *Option[T]) SetNone() {
	var zeroValue T

	o.value = zeroValue
	o.isSome = false
}

// Unwrap returns the value of an option if it is some and the zero value if it is none
func (o *Option[T]) Unwrap() (value T, isSome bool) {
	return o.value, o.isSome
}

// IsSome returns if an option is Some (Has a value)
func (o *Option[T]) IsSome() bool {
	return o.isSome
}

// IsNone returns if an option is Some (Has no value)
func (o *Option[T]) IsNone() bool {
	return !o.isSome
}

// Gets the value from an option and returns it if it's some and returns other if it's none
func (o *Option[T]) OrElse(other T) T {
	if o.isSome {
		return o.value
	}

	return other
}
