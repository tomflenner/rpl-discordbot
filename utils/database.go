package utils

import "fmt"

func GetDatabaseDNS(user string, pwd string, host string, port string, dbName string) string {
	return fmt.Sprintf("%s:%s@tcp(%s:%s)/%s",
		user,
		pwd,
		host,
		port, dbName)
}
