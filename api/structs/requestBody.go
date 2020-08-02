package structs

type RequestBody struct {
	Identifier string
	Customer   string
}

func (r RequestBody) ToMap() map[string]interface{} {
	return map[string]interface{}{
		"identifier": string(r.Identifier),
		"customer":  string(r.Customer),
	}
}