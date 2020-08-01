package structs

type RequestBody struct {
	Identifier string
	Customer   string
}

func (u *RequestBody) ToMap() map[string]interface{} {
	return map[string]interface{}{
		"identifier": string(u.Identifier),
		"customer":  string(u.Customer),
	}
}