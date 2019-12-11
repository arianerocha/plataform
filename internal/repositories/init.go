package repositories

func init() {
	go NewAccounts().AutoMigrate()
	go NewContacts().AutoMigrate()
}
