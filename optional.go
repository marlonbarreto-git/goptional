package goptional

type (
	Optional interface {
		IsPresent() bool
		IfPresent(function IfFunc)
		Or(alternative Any, error ...Any) Optional
		OrElse(alternative Any) Any
		OrElseGet(supplier OrElseSupplier) Any
		OrElsePanic(panicMsg string) Any
	}

	optional struct {
		value       Any
		alternative Any
		error       Any
	}

	Any interface{}

	IfFunc         func(value Any)
	OrElseSupplier func() Any
)

func NewOptional(args ...Any) Optional {
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

func (optional *optional) Get() (Any, Any) {
	return optional.value, optional.error
}

func (optional *optional) IfPresent(function IfFunc) {
	if !optional.IsPresent() || optional.error != nil {
		return
	}

	function(optional.value)
}

func (optional *optional) Or(alternative Any, error ...Any) Optional {
	if !optional.IsPresent() && optional.alternative == nil && error == nil {
		optional.alternative = alternative
	}

	return optional
}

func (optional *optional) OrElse(alternative Any) Any {
	if optional.IsPresent() {
		return optional.value
	}

	if optional.alternative != nil {
		return optional.alternative
	}

	return alternative
}

func (optional *optional) OrElseGet(supplier OrElseSupplier) Any {
	if optional.IsPresent() {
		return optional.value
	}

	return supplier()
}

func (optional *optional) OrElsePanic(panicMsg string) Any {
	if !optional.IsPresent() {
		panic(panicMsg)
	}

	return optional.value
}
