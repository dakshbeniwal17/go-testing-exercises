package exercise2

var db map[string]bool

func init() {
	db = make(map[string]bool)
	db["daksh@repozitory.com"] = true
	db["daksh.beniwal@repozitory.com"] = true
}

func UserExists(email string) bool {
	if _, ok := db[email]; !ok {
		return false
	}
	return true
}
