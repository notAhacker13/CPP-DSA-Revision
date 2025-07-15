package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/go-sql-driver/mysql"
)

var db *sql.DB

//defining a construct called Album

type Album struct {
	ID     int64
	Title  string
	Artist string
	Price  float32
}

func main() {

	//capturing connection properties

	cfg := mysql.NewConfig()
	cfg.User = os.Getenv("DBUSER")
	cfg.Passwd = os.Getenv("DBPASS")
	cfg.Net = "tcp"
	cfg.Addr = "127.0.0.1:3306"
	cfg.DBName = "recordings"

	//getting a databse handle

	var err error
	db, err = sql.Open("mysql", cfg.FormatDSN())

	if err != nil {
		log.Fatal(err)
	}

	pingErr := db.Ping()
	if pingErr != nil {
		log.Fatal(pingErr)
	}

	fmt.Println("Connected")

	res, err := albumsByArtist("John Coltrane")

	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Albums found: %v\n", res)

	//calling albumsByID function

	alb, err := albumsByID(2)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Album foundL %v\n", alb)

	//calling addAlbum function to try and add a row in the table

	albID, err := addAlbum(Album{
		Title:  "The Moder Sound of Success",
		Artist: "Aniket Kumar",
		Price:  75.00,
	})

	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("ID of added albumL %v\n", albID)

}

//function to query the database

func albumsByArtist(name string) ([]Album, error) {
	var albums []Album

	rows, err := db.Query("SELECT * FROM album WHERE artist =?", name)

	if err != nil {
		return nil, fmt.Errorf("albumsByArtist %q: %v", name, err)
	}
	defer rows.Close()

	for rows.Next() {
		var alb Album
		if err := rows.Scan(&alb.ID, &alb.Title, &alb.Artist, &alb.Price); err != nil {
			return nil, fmt.Errorf("albumsByArtist %q: %v", name, err)
		}

		albums = append(albums, alb)

	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("albumsByArtist %q: %v", name, err)
	}

	return albums, nil
}

func albumsByID(id int64) (Album, error) {

	var alb Album

	row := db.QueryRow("SELECT * FROM album where id =?", id)

	if err := row.Scan(&alb.ID, &alb.Title, &alb.Artist, &alb.Price); err != nil {
		if err == sql.ErrNoRows {
			return alb, fmt.Errorf("albumsByID: no such album", id)
		}
		return alb, fmt.Errorf("albumsBYID %d: %v", id, err)
	}
	return alb, nil
}

//adding data to the table

func addAlbum(alb Album) (int64, error) {

	result, err := db.Exec("INSERT INTO album (title, artist, price) VALUES (?, ?, ?)", alb.Title, alb.Artist, alb.Price)

	if err != nil {
		return 0, fmt.Errorf("addAlbum: %v", err)
	}

	id, err := result.LastInsertId()
	if err != nil {
		return 0, fmt.Errorf("addAlbum: %v", err)
	}
	return id, nil
}
