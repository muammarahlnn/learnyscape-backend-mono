package constant

const (
	SendVerificationSubject = "Learnyscape | Email Verification"
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
		<p>Learnyscape</p>
	</body>
</html>
	`
)
