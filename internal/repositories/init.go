package repositories

import "log"

func init() {
	log.Println("Auto Migrate:", "Accounts:", "(Error?)", NewAccounts().AutoMigrate())
}
