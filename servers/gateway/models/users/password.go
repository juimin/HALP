package users

// PasswordUpdate represents requirements for changing the user's password
type PasswordUpdate struct {
	NewPassword     string `json:"newPassword"`
	NewPasswordConf string `json:"newPasswordConf"`
}

// PassUpdate holds the hashed password for the new pass
type PassUpdate struct {
	PassHash []byte
}
