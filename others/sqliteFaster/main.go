package main

import (
	"database/sql"
	"github.com/gofiber/fiber/v2"
	_ "github.com/mattn/go-sqlite3"
	"log"
	"os"
)

func base() {
	app := fiber.New()
	app.Get("/", func(c *fiber.Ctx) error {
		file, err := os.Open("test.html")
		if err != nil {
			return err
		}
		stat, err := file.Stat()
		if err != nil {
			return err
		}
		ln := int(stat.Size())
		c.Type(".html")
		c.Response().SetBodyStream(file, ln)
		return nil
	})
	app.Listen(":3000")
}

func standard() {
	app := fiber.New()
	db, err := sql.Open("sqlite3", "file.db?cache=shared&mode=rwc")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	app.Get("/", func(c *fiber.Ctx) error {
		row := db.QueryRow("select body from pages where id=1")
		var body []byte
		err = row.Scan(&body)
		if err != nil {
			return err
		}
		c.Type(".html")
		c.Response().SetBody(body)
		return nil
	})
	app.Listen(":3000")
}

func prepareStatement() { //预处理
	app := fiber.New()
	db, err := sql.Open("sqlite3", "file.db?cache=shared&mode=rwc")
	if err != nil {
		log.Fatal(err)
	}
	bodyGet, err := db.Prepare("select body from pages where id=?")
	if err != nil {
		log.Fatal(err)
	}
	defer bodyGet.Close()
	app.Get("/", func(c *fiber.Ctx) error {
		row := bodyGet.QueryRow("1")
		var body []byte
		err = row.Scan(&body)
		if err != nil {
			return err
		}
		c.Type(".html")
		c.Response().SetBody(body)
		return nil
	})
	app.Listen(":3000")
}

func connectPool() {
	app := fiber.New()
	maxCons := 5
	cons := make(chan *sql.DB, maxCons)

	for i := 0; i < maxCons; i++ {
		conn, err := sql.Open("sqlite3", "file.db?cache=shared&mode=ro")
		if err != nil {
			log.Fatal(err)
		}
		defer func() {
			conn.Close()
		}()
		cons <- conn
	}

	checkout := func() *sql.DB {
		return <-cons
	}

	checkin := func(c *sql.DB) {
		cons <- c
	}
	app.Get("/", func(c *fiber.Ctx) error {
		db := checkout()
		defer checkin(db)
		row := db.QueryRow("select body from pages where id=1")

		var body []byte
		err := row.Scan(&body)
		if err != nil {
			return err
		}
		c.Type(".html")
		c.Response().SetBody(body)
		return nil
	})
	app.Listen(":3000")
}

func poolAndPrepareStatement() {
	app := fiber.New()
	maxCons := 5
	cons := make(chan *sql.Stmt, maxCons)

	for i := 0; i < maxCons; i++ {
		conn, err := sql.Open("sqlite3", "file.db?cache=shared&mode=ro")
		if err != nil {
			log.Fatal(err)
		}
		stmt, err := conn.Prepare("select body from pages where id=?")
		if err != nil {
			log.Fatal(err)
		}

		defer func() {
			stmt.Close()
			conn.Close()
		}()
		cons <- stmt
	}

	checkout := func() *sql.Stmt {
		return <-cons
	}

	checkin := func(c *sql.Stmt) {
		cons <- c
	}

	app.Get("/", func(c *fiber.Ctx) error {
		stmt := checkout()
		defer checkin(stmt)
		row := stmt.QueryRow(1)

		var body []byte
		err := row.Scan(&body)
		if err != nil {
			return err
		}
		c.Type(".html")
		c.Response().SetBody(body)
		return nil
	})
	app.Listen(":3000")
}

func main() {
	poolAndPrepareStatement()
}

// 测试
// hey -z 5s -c 50 http://localhost:3000/

// ps: hey安装https://github.com/rakyll/hey

/**
造数据
CREATE TABLE pages (
   id int  PRIMARY KEY,
   body TEXT NOT NULL
);

insert into pages values (1,"abc"),(2,"bcd"),(3,"cde");
*/
