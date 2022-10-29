package config

import (
	_ "github.com/go-sql-driver/mysql"
)

// func InitDB(config Config) (*sql.DB, error) {

// 	url := fmt.Sprintf("%s:%s@tcp(%s)/%s",
// 		config.UserDB,
// 		config.PasswordDB,
// 		config.HostDB,
// 		config.NameDB,
// 	)

// 	db, err := sql.Open("mysql", url)

// 	if err != nil {
// 		log.Panic(err)
// 	}

// 	return db, err
// }
