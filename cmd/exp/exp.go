package main

import (
	"fmt"

	_ "github.com/jackc/pgx/v4/stdlib"
	"github.com/joncalhoun/lenslocked/models"
)

func must(err error) {
	if err != nil {
		panic(err)
	}
}

func main() {

	conf := models.DefaultPostgresConfig()
	db, err := models.OpenDB(conf)
	must(err)
	defer db.Close()

	must(db.Ping())
	fmt.Println("Connected")

	// email := "2@gmail.com"
	// password := "123abc"

	// hashedBytes, err := bcrypt.GenerateFromPassword(
	// 	[]byte(password), bcrypt.DefaultCost)
	// if err != nil {
	// 	log.Fatalf("1:%w", err)
	// 	panic(err)
	// }
	// passwordHash := string(hashedBytes)

	// var id int
	// err = db.QueryRow(`
	// insert into users(email, password_hash)
	// values ($1, $2) returning id`, email, passwordHash).Scan(&id)
	// if err != nil {
	// 	log.Fatalf("2:%w", err)
	// 	panic(err)
	// }
	// fmt.Printf("id: %d", id)

	// us := models.UserService{
	// 	DB: db,
	// }
	// user, err := us.Create("1@mail", "123abc")
	// must(err)
	// fmt.Printf("%+v", user)

	// _, err = db.Exec(`CREATE TABLE IF NOT EXISTS users (
	// 	id SERIAL PRIMARY KEY,
	// 	name TEXT,
	// 	email TEXT NOT NULL
	// 	);
	// 	CREATE TABLE IF NOT EXISTS orders (
	// 	id SERIAL PRIMARY KEY,
	// 	user_id INT NOT NULL,
	// 	amount INT,
	// 	description TEXT
	// 	);`)

	// must(err)
	// fmt.Println("Tables created.")

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
	// 	type Order struct {
	// 		ID          int
	// 		UserID      int
	// 		Amount      int
	// 		Description string
	// 	}
	// 	var orders []Order
	// 	userID := 1 // Use the same ID you used in the previous lesson
	// 	rows, err := db.Query(`
	// 	SELECT id, amount, description
	// 	FROM orders
	// 	WHERE user_id=$1`, userID)
	// 	must(err)
	// 	defer rows.Close()

	// 	for rows.Next() {
	// 		var order Order
	// 		order.UserID = userID
	// 		err := rows.Scan(&order.ID, &order.Amount, &order.Description)
	// 		must(err)
	// 		orders = append(orders, order)
	// 	}
	// 	must(rows.Err())
	// 	fmt.Printf("Orders:%+v", orders)
	// }
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
