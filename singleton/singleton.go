package singleton

import "sync"

type Database struct {
	connection string
}

var instance *Database
var once sync.Once

func GetInstance() *Database {
	once.Do(func() {
		instance = &Database{
			connection: "database connection established",
		}
	})
	return instance
}

func (db *Database) Query(sql string) string {
	return "Executing: " + sql + " on " + db.connection
}