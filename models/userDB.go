package models

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"fmt"
)

// const dbuser = "bccc34d20a6c7d"
// const dbpass = "ea5886d4"
// const dbhost = "us-cdbr-east-02.cleardb.com"
// const dbname = "heroku_69f5974c9ccf3a4"

// const dbconnect = dbuser+":"+dbpass+"@tcp("+dbhost+":3306)/"+dbname

func GetUsers() []User {
	db, err := sql.Open("mysql", dbconnect)

	// if there is an error opening the connection, handle it
	if err != nil {
		// simply print the error to the console
		fmt.Println("Err", err.Error())
		// returns nil on error
		return nil
	}

	defer db.Close()
	results, err := db.Query("SELECT * FROM user")

	if err != nil {
		fmt.Println("Err", err.Error())
		return nil
	}

	users := []User{}
	for results.Next() {
		var user User
        // for each row, scan into the Product struct
		err = results.Scan(&user.ID, &user.FirstName, &user.LastName, &user.Checked)
		if err != nil {
			panic(err.Error()) // proper error handling instead of panic in your app
		}
        // append the product into products array
		users = append(users, user)
	}

	return users

}

func DeleteUser(id string) string {

	db, err := sql.Open("mysql", dbconnect)
	status := ""
	user := &User{}
	if err != nil {
		// simply print the error to the console
		fmt.Println("Err", err.Error())
		// returns nil on error
		return ""
	}
    
	defer db.Close()

	results, err := db.Query("SELECT * FROM `user` where id=?", id)

	if err != nil {
		fmt.Println("Err", err.Error())
		return ""
	}

	if results.Next() {
		err = results.Scan(&user.ID, &user.FirstName, &user.LastName, &user.Checked)
		if err != nil {
			return ""
		}

		delete, err := db.Query("DELETE FROM `user` WHERE (`id` = '" + id + "');")
		print("DELETE FROM `user` WHERE (`id` = '?');", id)

		if err != nil {
			fmt.Println("Err", err.Error())
			return ""
		}

		print(delete)

		status = "Delete user " + user.FirstName + " success"
	} else {

		return ""
	}

	return status
}

func AddUser(user User) {

	db, err := sql.Open("mysql", dbconnect)

	if err != nil {
		panic(err.Error())
	}

	// defer the close till after this function has finished
	// executing
	defer db.Close()

	insert, err := db.Query(
		"INSERT INTO user (firstName,lastName,checked) VALUES (?,?,?)",
		user.FirstName, user.LastName, user.Checked)

	// if there is an error inserting, handle it
	if err != nil {
		panic(err.Error())
	}

	defer insert.Close()

}

func EditUser(user User, id string) {

	db, err := sql.Open("mysql", dbconnect)

	if err != nil {
		panic(err.Error())
	}

	// defer the close till after this function has finished
	// executing
	defer db.Close()

	var checked string
	if user.Checked == true {
		checked = "1"
	} else {
		checked = "0"
	}

	update, err := db.Query("UPDATE `user` SET `firstName` = '" + user.FirstName + "',`lastName` = '" + user.LastName + "',`checked` = '" + checked + "' WHERE (`id` = '" + id +  "');")

	// if there is an error inserting, handle it
	if err != nil {
		panic(err.Error())
	}

	defer update.Close()

}