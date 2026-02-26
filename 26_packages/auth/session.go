package auth

func GetSession() string {
	return "logged in session"
}

// private function that cannot be accessed outside the package
func extractSession() string {
	return extractSession()
}