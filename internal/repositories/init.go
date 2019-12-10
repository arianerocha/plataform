package repositories

func init() {
	go NewAccounts().AutoMigrate()
}
