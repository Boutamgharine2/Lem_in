package Lemin

func Roms(s string) string {
	if s == "" {
		return ""
	}
	rome := ""
	for i := 0; i < len(s); i++ {

		if s[i] == ' ' {
			break
		}
		if s[i] < '0' || s[i] > '9' {
			return ""
		}
		rome += string(s[i])

	}
	return rome
}
