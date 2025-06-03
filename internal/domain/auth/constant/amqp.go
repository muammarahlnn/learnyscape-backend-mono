package constant

const (
	AccountVerifiedExchange = "notifications"
	AccountVerifiedKey      = "account.verified"
	AccountVerifiedQueue    = "account-verified-email"

	ForgotPasswordExchange = "notifications"
	ForgotPasswordKey      = "account.forgot-password"
	ForgotPasswordQueue    = "account-forgot-password-email"
)

const (
	AMQPRetryDelay = 3
	AMQPRetryLimit = 3
)
