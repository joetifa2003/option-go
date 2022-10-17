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

// SetSome sets an option to a value
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

// Unwrap returns the value of an option if it is Some and the zero value if it is None
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

// Or returns the contained Some value or a provided default
func (o *Option[T]) Or(other T) T {
	if o.isSome {
		return o.value
	}

	return other
}

// OrElse returns the contained Some value or computes it from a closure.
func (o *Option[T]) OrElse(f func() T) T {
	if o.isSome {
		return o.value
	}

	return f()
}

// IsSomeAnd returns true if the option is a Some and the value inside of it matches a predicate.
func (o *Option[T]) IsSomeAnd(f func(T) bool) bool {
	return o.isSome && f(o.value)
}

// Replace replaces the actual value in the option by the value given in parameter, returning the old value if present
func (o *Option[T]) Replace(value T) Option[T] {
	old := Option[T]{value: o.value, isSome: o.isSome}

	o.SetSome(value)

	return old
}
