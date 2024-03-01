package exception

func Validation(errors ...map[string]string) error {
	httpError := HttpError{
		Code:    422,
		Message: "Invalid payload",
	}

	if len(errors) > 0 {
		httpError.Errors = errors[0]
	}

	return httpError
}

func PasswordRequired() error {
	return Validation(map[string]string{
		"password": "Password wajib diisi",
	})
}

func UserNotRegistered() error {
	return Validation(map[string]string{
		"email": "Email tidak terdaftar",
	})
}

func PasswordIncorrect() error {
	return Validation(map[string]string{
		"password": "Email tidak terdaftar",
	})
}

func EmailUsed() error {
	return Validation(map[string]string{
		"email": "Email telah digunakan",
	})
}
