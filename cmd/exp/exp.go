package main

import (
	"database/sql"
	"fmt"

	_ "github.com/jackc/pgx/v4/stdlib"
)

type PostgresConfig struct {
	Host     string
	Port     string
	User     string
	Password string
	Database string
	SSLMode  string
}

func (cfg PostgresConfig) String() string {
	return fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s", cfg.Host, cfg.Port, cfg.User, cfg.Password, cfg.Database, cfg.SSLMode)
}

func main() {
	cfg := PostgresConfig{
		Host:     "localhost",
		Port:     "5435",
		User:     "baloo",
		Password: "junglebook",
		Database: "lenslocked",
		SSLMode:  "disable",
	}
	db, err := sql.Open("pgx", cfg.String())
	must(err)
	defer db.Close()
	must(db.Ping())

	_, err = db.Exec(`CREATE TABLE IF NOT EXISTS users (
		id SERIAL PRIMARY KEY,
		name TEXT,
		email TEXT NOT NULL
		);
		CREATE TABLE IF NOT EXISTS orders (
		id SERIAL PRIMARY KEY,
		user_id INT NOT NULL,
		amount INT,
		description TEXT
		);`)

	must(err)
	fmt.Println("Tables created.")

	// name := "Jon Calhoun"
	// email := "jon@calhoun.io"
	// _, err = db.Exec(`
	// 	INSERT INTO users(name, email)
	// 	VALUES($1, $2);`, name, email)
	// must(err)
	// fmt.Println("User created.")

	// name := "New User"
	// email := "new@calhoun.io"
	// var id int
	// must(db.QueryRow(`
	// 	INSERT INTO users (name, email)
	// 	VALUES ($1, $2) RETURNING id;`, name, email).Scan(&id))
	// fmt.Println("User created. id =", id)

	// id := 1
	// var name, email string
	// must(db.QueryRow(`
	// SELECT name, email
	// FROM users
	// WHERE id=$1;`, id).Scan(&name, &email))
	// fmt.Printf("User information: name=%s, email=%s\n", name, email)

	// id := 200 // Choose an ID NOT in your users table
	// var name, email string
	// must(db.QueryRow(`
	// SELECT name, email
	// FROM users
	// WHERE id=$1;`, id).Scan(&name, &email))
	// if err == sql.ErrNoRows {
	// 	fmt.Println("Error, no rows!")
	// }
	// if err != nil {
	// 	panic(err)
	// }
	// fmt.Printf("User information: name=%s, email=%s\n", name, email)

	// userID := 1 // Pick an ID that exists in your DB
	// for i := 1; i <= 5; i++ {
	// 	amount := i * 100
	// 	desc := fmt.Sprintf("Fake order #%d", i)

	// 	_, err := db.Exec(`
	// 	INSERT INTO orders(user_id, amount, description)
	// 	VALUES($1, $2, $3)`, userID, amount, desc)
	// 	must(err)
	// }
	// fmt.Println("Created fake orders.")
	type Order struct {
		ID          int
		UserID      int
		Amount      int
		Description string
	}
	var orders []Order
	userID := 1 // Use the same ID you used in the previous lesson
	rows, err := db.Query(`
	SELECT id, amount, description
	FROM orders
	WHERE user_id=$1`, userID)
	must(err)
	defer rows.Close()

	for rows.Next() {
		var order Order
		order.UserID = userID
		err := rows.Scan(&order.ID, &order.Amount, &order.Description)
		must(err)
		orders = append(orders, order)
	}
	must(rows.Err())
	fmt.Printf("Orders:%+v", orders)

}

func must(err error) {
	if err != nil {
		panic(err)
	}
}

/*
statement := `CREATE TABLE IF NOT EXISTS users (
		id SERIAL PRIMARY KEY,
		name TEXT,
		email TEXT NOT NULL
		);
		CREATE TABLE IF NOT EXISTS orders (
		id SERIAL PRIMARY KEY,
		user_id INT NOT NULL,
		amount INT,
		description TEXT
		);`
	_, err = db.Exec(statement)
	must(err)
	userID := 1

	for i := 1; i <= 5; i++ {
		amount := i * 100
		desc := fmt.Sprintf("Fake order #%d", i)
		_, err := db.Exec(`
	INSERT INTO orders(user_id, amount, description)
	VALUES($1, $2, $3)`, userID, amount, desc)
		must(err)
	}
	fmt.Println("Created fake orders.")
	// id := 100
	// var name, email string
	// err = db.QueryRow(`
	// SELECT name, email
	// FROM users
	// WHERE id=$1;`, id).Scan(&name, &email)
	// if err == sql.ErrNoRows{
	// 	fmt.Println("Error, no rows")
	// }
	// if err != nil {
	// 	panic(err)
	// }
	// fmt.Printf("User information: name=%s, email=%s\n", name, email)
*/
