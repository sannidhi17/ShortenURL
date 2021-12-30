package storage

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
)

type PostgresService struct {
	postgresdb *sql.DB
}

var postgresService = &PostgresService{}

const (
	host     = "localhost"
	port     = 5432
	user     = "defaultuser"
	password = "password"
	dbname   = "shorturldb"
)

func InitDB() *PostgresService {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
	host, port, user, password, dbname)
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err) //TODO update
	}
	err = db.Ping()
	if err != nil {
		panic(err) //TODO update
	}
	fmt.Println("Successfully connected to the database!")

	postgresService.postgresdb = db
	return postgresService
}

func AddURLToDB(originalUrl string, shortUrl string) {
	sqlStmt := `INSERT INTO url_tbl VALUES ($1, $2)`
	_, err := postgresService.postgresdb.Exec(sqlStmt, shortUrl, originalUrl)
	if err != nil {
		panic(fmt.Sprintf("Failed to save to the databse | Error: %v - shortUrl: %s - originalUrl: %s\n", err, shortUrl, originalUrl))
	}
}

func GetURLFromDB(shortUrl string) string {
	sqlStmt := `SELECT original FROM url_tbl WHERE short=$1`
	row := postgresService.postgresdb.QueryRow(sqlStmt, shortUrl)
	var res string
	err := row.Scan(&res)
	if err == sql.ErrNoRows {
		return "No Result" //should this be an exception?
	} else if err != nil {
		panic(fmt.Sprintf("Failed to retrieve URL from the database | Error: %v - shortUrl: %s\n", err, shortUrl))
	}
	return res
}