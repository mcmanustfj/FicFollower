// // File to handle database calls

package main

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/mattn/go-sqlite3"
)

func Recreate_db() (error) {
	//if foo.db exists, delete it
	if _, err := os.Stat("./ff.db"); err == nil {
		fmt.Println("Removing old db")
		os.Remove("./ff.db")
	}
	db, err := sql.Open("sqlite3", "./ff.db")
	if err != nil {
		fmt.Println(err)
		return err 
	}
	sqlFile := "./migrations/0.sql"
	sqlStmt, err := os.ReadFile(sqlFile)
	fmt.Println("Writing " + sqlFile + " to db\n")
	fmt.Println(string(sqlStmt))
	_, err = db.Exec(string(sqlStmt))
	if err != nil {
		fmt.Println(err)
		db.Close()
		return err
	}

	db.Close()

	return nil
}

func AddUser(username string) (error) {
	db, err := sql.Open("sqlite3", "./ff.db")
	if err != nil {
		fmt.Println(err)
		return err 
	}
	defer db.Close()

	sqlStmt := `INSERT INTO users (username) VALUES (?)`
	_, err = db.Exec(string(sqlStmt), username)
	if err != nil {
		fmt.Println(err)
		return err
	}

	return nil
}

func GetUser(username string) (User, error) {
	db, err := sql.Open("sqlite3", "./ff.db")
	if err != nil {
		fmt.Println(err)
		return User{}, err 
	}
	defer db.Close()

	sqlStmt := `SELECT username FROM users WHERE username = ?`
	row := db.QueryRow(string(sqlStmt), username)
	var u User
	err = row.Scan(&u.username)
	if err != nil {
		fmt.Println(err)
		return User{}, err
	}

	return u, nil
}

func AddStory(story Story) (error) {
	db, err := sql.Open("sqlite3", "./ff.db")
	if err != nil {
		fmt.Println(err)
		return err 
	}
	defer db.Close()

	sqlStmt := `INSERT INTO stories (id, title, site, link, updated, updatedlink, author, authorlink, checked) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?)`
	_, err = db.Exec(string(sqlStmt), story.id, story.title, story.site, story.link, story.updated, story.updatedlink, story.author, story.authorlink, story.checked)
	if err != nil {
		fmt.Println(err)
		return err
	}

	return nil
}

func AddUserFollows(userFollow UserFollows) (error) {
	db, err := sql.Open("sqlite3", "./ff.db")
	if err != nil {
		fmt.Println(err)
		return err 
	}
	defer db.Close()

	sqlStmt := `INSERT INTO userfollows (username, storyId, readDate, readLink) VALUES (?, ?, ?, ?)`
	_, err = db.Exec(string(sqlStmt), userFollow.username, userFollow.storyId, userFollow.readDate, userFollow.readLink)
	if err != nil {
		fmt.Println(err)
		return err
	}

	return nil
}

