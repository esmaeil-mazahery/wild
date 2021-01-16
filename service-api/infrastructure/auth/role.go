package auth

//AccessibleRoles ...
func AccessibleRoles() map[string][]string {
	const AuthService = "/pb_auth.AuthService/"

	return map[string][]string{
		AuthService + "Register":             {"anonymous"},
		AuthService + "ForgetPassword":       {"anonymous"},
		AuthService + "ForgetPasswordChange": {"anonymous"},
		AuthService + "Login":                {"anonymous"},
	}
}
