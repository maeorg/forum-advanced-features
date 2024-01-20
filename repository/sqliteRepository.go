package repository

import (
	"database/sql"
)

// database pointer
var Database *sql.DB

var tables = []string{users, posts, comments, likes, categories, categoryPost}

const users = `CREATE TABLE IF NOT EXISTS users (
	id INTEGER PRIMARY KEY AUTOINCREMENT UNIQUE NOT NULL,
  	username TEXT NOT NULL UNIQUE,
  	email TEXT NOT NULL UNIQUE,
  	password TEXT NOT NULL
	);`

const posts = `CREATE TABLE IF NOT EXISTS posts (
	id INTEGER PRIMARY KEY AUTOINCREMENT UNIQUE NOT NULL,
	title TEXT NOT NULL,
	content TEXT NOT NULL,
	created_at DATETIME NOT NULL,
	image_url TEXT,
	user_id INTEGER NOT NULL,
	FOREIGN KEY (user_id) REFERENCES users(id)
	);`

const comments = `CREATE TABLE IF NOT EXISTS comments (
	id INTEGER PRIMARY KEY AUTOINCREMENT UNIQUE NOT NULL,
	content TEXT NOT NULL,
	created_at DATETIME NOT NULL,
	user_id INTEGER NOT NULL,
	post_id INTEGER NOT NULL,
	FOREIGN KEY (user_id) REFERENCES users(id),
	FOREIGN KEY (post_id) REFERENCES posts(id)
	);`

const likes = `CREATE TABLE IF NOT EXISTS likes (
	id INTEGER PRIMARY KEY AUTOINCREMENT UNIQUE NOT NULL,
	type TEXT NOT NULL,
	user_id INTEGER NOT NULL,
	post_id INTEGER,
	comment_id INTEGER,
	FOREIGN KEY (user_id) REFERENCES users(id),
	FOREIGN KEY (post_id) REFERENCES posts(id),
	FOREIGN KEY (comment_id) REFERENCES comments(id)
	);`

const categories = `CREATE TABLE IF NOT EXISTS categories (
	id INTEGER PRIMARY KEY AUTOINCREMENT UNIQUE NOT NULL,
	name TEXT UNIQUE NOT NULL
	);`

const categoryPost = `CREATE TABLE IF NOT EXISTS categoryPost (
	id INTEGER PRIMARY KEY AUTOINCREMENT UNIQUE NOT NULL,
	post_id INTEGER NOT NULL,
	category_id INTEGER NOT NULL,
	FOREIGN KEY (post_id) REFERENCES posts(id),
	FOREIGN KEY (category_id) REFERENCES categories(id)
	);`

func InitializeDatabase() (*sql.DB, error) {
	database, err := sql.Open("sqlite3", "database/forumDatabase.db")
	if err != nil {
		return nil, err
	}
	
	if err := database.Ping(); err != nil {
		return nil, err
	}

	return database, nil
}

func CreateTables(database *sql.DB) error {
	for _, table := range tables {
		_, err := database.Exec(table)
		if err != nil {
			return err
		}
	}

	return nil
}
