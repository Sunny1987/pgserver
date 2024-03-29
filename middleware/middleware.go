package middleware

import (
	"database/sql"
	"fmt"

	//"github.com/joho/godotenv"
	"log"

	_ "github.com/lib/pq"

	//"os"
	"pgserver/models"
)

func createConnection() *sql.DB {
	// err := godotenv.Load(".env")
	// if err != nil {
	// 	log.Fatalf("Error loading .env file")
	// }

	// Open the connection
	//db, err := sql.Open("postgres", os.Getenv("POSTGRES_URL"))
	//db, err := sql.Open("postgres", "user=postgres password=postgres dbname= sslmode=disable")
	//db, err := sql.Open("postgres", "user=postgres password=postgres host=ec2-44-202-11-141.compute-1.amazonaws.com port=5432 dbname=mydb sslmode=require")
	db, err := sql.Open("postgres", "postgres://user:FhsPJOoJ3oyYnpe32mEXvgHyvVaiOgNu@dpg-cfb3eqg2i3mjdukporg0-a.singapore-postgres.render.com/mydb_gg02")
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

func GetScans() ([]models.Scan, error) {
	//create the postgres db connection
	db := createConnection()

	//close the db connection
	defer db.Close()

	var scans []models.Scan

	//sql query
	sqlStatement := `select * from scans`

	//execute the sql statement
	rows, err := db.Query(sqlStatement)
	if err != nil {
		log.Fatalf("Unable to execute the query. %v", err)
	}

	//close the statement
	defer rows.Close()

	for rows.Next() {
		var scan models.Scan

		// unmarshal the row object to user
		err = rows.Scan(&scan.Id, &scan.Person, &scan.Url, &scan.Result)

		if err != nil {
			log.Fatalf("Unable to scan the row. %v", err)
		}

		// append the user in the users slice
		scans = append(scans, scan)

	}
	return scans, err
}

// InsertScan inserts one user in the DB
func InsertScan(scan models.Scan) int64 {

	// create the postgres db connection
	db := createConnection()

	// close the db connection
	defer func(db *sql.DB) {
		err := db.Close()
		if err != nil {

		}
	}(db)

	// create the insert sql query
	// returning userid will return the id of the inserted user
	sqlStatement := `INSERT INTO scans (url, person, result) VALUES ($1, $2, $3) RETURNING scan_id`
	//sqlStatement := fmt.Sprintf("INSERT INTO scans (url, person, result) VALUES ($1,$2,&3) RETURNING userid",)

	// the inserted id will store in this id
	var id int64

	// execute the sql statement
	// Scan function will save the insert id in the id
	err := db.QueryRow(sqlStatement, scan.Url, scan.Person, scan.Result).Scan(&id)

	if err != nil {
		log.Fatalf("Unable to execute the query. %v", err)
	}

	fmt.Printf("Inserted a single record %v", id)

	// return the inserted id
	return id
}

func DeleteScan(id int64) int64 {

	// create the postgres db connection
	db := createConnection()

	// close the db connection
	defer db.Close()

	// create the delete sql query
	sqlStatement := `DELETE FROM scans WHERE scan_id=$1`

	// execute the sql statement
	res, err := db.Exec(sqlStatement, id)

	if err != nil {
		log.Fatalf("Unable to execute the query. %v", err)
	}

	// check how many rows affected
	rowsAffected, err := res.RowsAffected()

	if err != nil {
		log.Fatalf("Error while checking the affected rows. %v", err)
	}

	return rowsAffected
}
