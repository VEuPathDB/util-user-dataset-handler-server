package svc

type ValidationSet map[string][]string

func (s ValidationSet) Append(other ValidationSet) {
	for k, v := range other {
		if _, ok := s[k]; ok {
			s[k] = append(s[k], v...)
		} else {
			s[k] = v
		}
	}
}

type ValidationResult struct {
	Failed bool
	Result ValidationSet
}

func (r *ValidationResult) AddError(key, val string) {
	if r.Result == nil {
		r.Result = make(ValidationSet)
	}

	r.Failed = true
	r.Result[key] = append(r.Result[key], val)
}
