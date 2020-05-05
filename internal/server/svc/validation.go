package svc

type ValidationSet map[string][]string

func (V ValidationSet) Append(other ValidationSet) {
	for k, v := range other {
		if _, ok := V[k]; ok {
			V[k] = append(V[k], v...)
		} else {
			V[k] = v
		}
	}
}

type ValidationResult struct {
	Ok     bool
	Result ValidationSet
}

func (V *ValidationResult) AddError(key, val string) {
	if V.Result == nil {
		V.Result = make(ValidationSet)
	}
	V.Ok = false
	if _, ok := V.Result[key]; ok {
		V.Result[key] = append(V.Result[key], val)
	} else {
		V.Result[key] = []string{val}
	}
}
