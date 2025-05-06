package rule

type Interface interface {
	Valid(key string, values map[string]any) error
}
