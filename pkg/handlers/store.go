package handlers

type memStore interface {
	InsertUrl(url, code string)
	GetUrl(code string) (string, error)
}
