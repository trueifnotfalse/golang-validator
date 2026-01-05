package locale

type Locale map[string]string

func (r Locale) Translate(v string) string {
	t, ok := r[v]
	if ok {
		return t
	}
	return v
}
