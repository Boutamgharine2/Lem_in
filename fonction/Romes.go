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
		
		rome += string(s[i])

	}
	return rome
}
func ContienNumber(str string) bool {
	for i:=0;i<len(str);i++ {
		if str[i]>='0' && str[i]<='9' {
			return true
		}
	}
	return false
}