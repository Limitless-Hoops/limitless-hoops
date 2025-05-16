package database

func PopulateDB() {
	Migrate()
	Seed()
}
