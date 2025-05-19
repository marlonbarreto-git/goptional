package goptional

type (
	Optional interface {
		IsPresent() bool
		IfPresent(function IfFunc)
		Or(alternative any, err ...any) Optional
		OrElse(alternative any) any
		OrElseGet(supplier OrElseSupplier) any
		OrElsePanic(panicMsg string) any
	}

	optional struct {
		value       any
		alternative any
		error       any
	}

	IfFunc         func(value any)
	OrElseSupplier func() any
)

func NewOptional(args ...any) Optional {
	switch len(args) {
	case 0:
		return &optional{}
	case 1:
		return &optional{value: args[0]}
	case 2:
		return &optional{
			value: args[0],
			error: args[1],
		}
	default:
		return &optional{
			value: args[:len(args)-1],
			error: args[len(args)-1],
		}
	}
}

func (optional *optional) IsPresent() bool {
	return optional.value != nil
}

func (optional *optional) Get() (any, any) {
	return optional.value, optional.error
}

func (optional *optional) IfPresent(function IfFunc) {
	if !optional.IsPresent() || optional.error != nil {
		return
	}

	function(optional.value)
}

func (optional *optional) Or(alternative any, err ...any) Optional {
	if !optional.IsPresent() && optional.alternative == nil && len(err) == 0 {
		optional.alternative = alternative
	}

	if len(err) > 0 {
		optional.error = err[0]
	}

	return optional
}

func (optional *optional) OrElse(alternative any) any {
	if optional.IsPresent() {
		return optional.value
	}

	if optional.alternative != nil {
		return optional.alternative
	}

	return alternative
}

func (optional *optional) OrElseGet(supplier OrElseSupplier) any {
	if optional.IsPresent() {
		return optional.value
	}

	return supplier()
}

func (optional *optional) OrElsePanic(panicMsg string) any {
	if !optional.IsPresent() {
		panic(panicMsg)
	}

	return optional.value
}
