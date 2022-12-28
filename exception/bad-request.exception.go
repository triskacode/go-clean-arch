package exception

type ErrorValidation struct {
	Field string `json:"field"`
	Value string `json:"value"`
	Tag   string `json:"tag"`
}
