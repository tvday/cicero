package models

func IsValidCategory(category string) bool {
	switch category {
	case "basic",
		"spirit",
		"liqueur",
		"bitters":
		return true
	default:
		return false
	}
}
