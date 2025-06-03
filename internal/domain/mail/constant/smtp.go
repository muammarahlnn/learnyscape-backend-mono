package constant

const (
	SendVerificationSubject = "[Learnyscape] Email Verification"
	AccountVerifiedSubject  = "[Learnyscape] Account Verified"
	ResetPasswordSubject    = "[Learnyscape] Reset Password"
)

const (
	SendVerificationTemplate = `
<!DOCTYPE html>
	<html lang="en">
	<head>
		<meta charset="UTF-8">
	</head>
	<body>
		<h1>Email Verification</h1>
		<p>Dear %s,</p>
		<p>Thank you for registering. Please use the following OTP code to verify your email address:</p>
		<h2>%s</h2>
		<p>This code will expire in 10 minutes.</p>
		<p>If you did not request this, please ignore this email.</p>
		<p>Warm regards,</p>
		<p>Learnyscape Team</p>
	</body>
</html>
	`
	AccountVerifiedTemplate = `
<!DOCTYPE html>
	<html lang="en">
	<head>
		<meta charset="UTF-8">
	</head>
	<body>
		<h1>Account Verified</h1>
		<p>Dear %s,</p>
		<p>We are pleased to inform you that your account has been successfully verified.</p>
		<p>You can now login using your email.</p>
		<p>If you have any questions or need further assistance, please do not hesitate to contact us.</p>
		<p>Warm regards,</p>
		<p>Learnyscape Team</p>
	</body>
</html>
	`

	ResetPasswordTemplate = `
<!DOCTYPE html>
	<html lang="en">
	<head>
		<meta charset="UTF-8">
	</head>
	<body>
		<h1>Reset Password</h1>
		<p>Dear %s,</p>
		<p>You requested to reset your password. Here's your reset password token:.</p>
		<h2>%s</h2>
		<p>This code will expire in 5 minutes.</p>
		<p>If you did not request this, please ignore this email.</p>
		<p>Warm regards,</p>
		<p>Learnyscape Team</p>
	</body>
</html>
	`
)
