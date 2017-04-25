package validation

type ErrorForm struct {
	Key     string `json:"key"`
	Message string `json:"message"`
}

func Validate(request) (errors []ErrorForm) {

	if request.Username == "" {
		errors = append(errors, ErrorForm{
			key:     "username",
			Message: "Username is required.",
		})
	}

	if request.Email == "" {
		errors = append(errors, ErrorForm{
			key:     "email",
			Message: "Email is required.",
		})
	}

	...

}
