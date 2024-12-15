package main

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

// const (
// 	host     = "localhost"
// 	port     = 5433
// 	user     = "postgres"
// 	password = "postgres"
// 	dbname   = "db_go_sql"
// )

const (
	host     = "10.8.135.107"
	port     = 5432
	user     = "postgres"
	password = "P@ssw0rd"
	dbname   = "rheldc"
)

var (
	db  *sql.DB
	err error
)

func main() {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	db, err = sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		panic(err)
	}

	fmt.Println("Successfully connected to Database "+dbname, "PORT =", port)

	// CreateCatalogue()
	// GetCatalogue()
	// UpdateCatalogue()

}

type Employee struct {
	ID        int
	Full_name string
	Email     int
	Age       string
	Division  int
}

func CreateEmployee() {
	var employee = Employee{}

	sqlStatement := `
	INSERT INTO employees (full_name, email, age, division)
	VALUES ($1, $2, $3, $4)
	Returning *
	`

	err = db.QueryRow(sqlStatement, "Dump1", "dump1@gmail.com", 27, "ITServices").
		Scan(&employee.ID, &employee.Full_name, &employee.Email, &employee.Age, &employee.Division)

	if err != nil {
		panic(err)
	}

	fmt.Printf("New Employee Data : %+v\n", employee)
}

func GetCatalogue() {
	var results = []Catalogue{}

	sqlStatament := `SELECT * FROM catalog`

	rows, err := db.Query(sqlStatament)

	if err != nil {
		panic(err)
	}
	defer rows.Close()

	for rows.Next() {
		var catalogue = Catalogue{}

		err = rows.Scan(&catalogue.id, &catalogue.nama_barang, &catalogue.jumlah, &catalogue.type_barang, &catalogue.harga)

		if err != nil {
			panic(err)
		}

		results = append(results, catalogue)
	}

	fmt.Println("Catalogue data ; ", results)
}

func UpdateCatalogue() {
	sqlStatement := `
	UPDATE catalog
	SET nama_barang = $2, jumlah = $3, type_barang = $4, harga = $5
	WHERE id = $1;`
	res, err := db.Exec(sqlStatement, 1, "Celana Cutbray", 100, "Dewasa", 200000)
	if err != nil {
		panic(err)
	}
	count, err := res.RowsAffected()
	if err != nil {
		panic(err)
	}
	fmt.Println("Update Data harga", count)
}
