package constant

const (
	InvalidCredentialErrorMessage        = "email or password is wrong"
	UserAlreadyExistsErrorMessage        = "user already exists"
	InvalidRefreshTokenErrorMessage      = "invalid refresh token"
	UserAlreadyVerifiedErrorMessage      = "user already verified"
	VerificationTokenExpiredErrorMessage = "verification token expired"
	InvalidVerificationTokenErrorMessage = "invalid verification token"
	VerificationCooldownErrorMessage     = "A verification email has already been sent. Please wait before trying again."
	EmailNotVerifiedErrorMessage         = "email not verified"
	ForgotPasswordCooldErrorMessage      = "A reset password email has already been sent. Please wait before trying again."
	ResetPasswordErrorMessage            = "invalid email or token"
)
