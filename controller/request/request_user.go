package request

import "fmt"

func errParamIsRequired(name, typ string) error {
	return fmt.Errorf("param: %s (type: %s) is required", name, typ)
}

type RequestUpdateUser struct {
	Name     string `json:"name"`
	Password string `json:"password"`
	Email    string `json:"email"`
}

func (r *RequestUpdateUser) Validate() error {
	if r.Name != "" || r.Email != "" || r.Password != "" {
		return nil
	}
	return fmt.Errorf("at least one valid filed must be provided")
}
