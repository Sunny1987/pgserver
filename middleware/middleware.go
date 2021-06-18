package middleware

import (
	"database/sql"
	"fmt"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"log"
	"pgserver/models"
)

func createConnection() *sql.DB {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	// Open the connection
	db, err := sql.Open("postgres", "postgres://idcyllburgvvbd:889cc4417bffce17643aa4619773cbccaa6efe0610882eb41da60581fd56241c@ec2-52-4-111-46.compute-1.amazonaws.com:5432/ddgju1afa6tggc")
	//db, err := sql.Open("postgres", "user=user password=pass dbname=postgres sslmode=disable")
	if err != nil {
		panic(err)
	}

	err = db.Ping()
	if err != nil {
		panic(err)
	}

	fmt.Println("Successfully connected!")
	// return the connection
	return db
}

func GetUsers() ([]models.User, error) {

	//create the postgres db connection
	db := createConnection()

	//close the db connection
	defer db.Close()

	var users []models.User

	//sql query
	sqlStatement := `select * from users`

	//execute the sql statement
	rows, err := db.Query(sqlStatement)
	if err != nil {
		log.Fatalf("Unable to execute the query. %v", err)
	}

	//close the statement
	defer rows.Close()

	// iterate over the rows
	for rows.Next() {
		var user models.User

		// unmarshal the row object to user
		err = rows.Scan(&user.ID, &user.Username, &user.Password)

		if err != nil {
			log.Fatalf("Unable to scan the row. %v", err)
		}

		// append the user in the users slice
		users = append(users, user)

	}

	// return empty user on error
	return users, err

}
