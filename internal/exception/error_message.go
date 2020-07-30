package exception

type ErrorMessage string

const (
	LOGIN01 ErrorMessage = "Login model bind error"
	LOGIN02 ErrorMessage = "Email not found"
	LOGIN03 ErrorMessage = "Incorrect password"
	REGIS01 ErrorMessage = "Registration model bind error"
	REGIS02 ErrorMessage = "Registration password hash error"
	REGIS03 ErrorMessage = "Registration ORM error"
)
