package Tunnel

// Just a callback function for the password request.
func PasswordCallback() (string, error) {
	return callbackPassword, nil
}

func SetPassword4Callback(password string) {
	callbackPassword = password
}
