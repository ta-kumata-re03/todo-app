package main

import (
	"net/http"
	"os"
	"strconv"

	"database/sql"
	"fmt"

	"log"

	_ "github.com/go-sql-driver/mysql"
	"github.com/labstack/echo"
)

type Todos struct {
	Todos []Todo `json:"todos"`
}

type Todo struct {
	ID          string `json:"id"`
	Title       string `json:"title"`
	Detail      string `json:"detail"`
	Expire_date string `json:"expire_date"`
}

func main() {
	// [ユーザ名]:[パスワード]@tcp([ホスト名]:[ポート番号])/[データベース名]?charset=[文字コード]
	datasourceUser := os.Getenv("DATASOURCE_USER")
	datasourcePassword := os.Getenv("DATASOURCE_PASSWORD")
	datasourceHost := os.Getenv("DATASOURCE_HOST")
	datasourcePort := os.Getenv("DATASOURCE_PORT")
	datasourceDatabase := os.Getenv("DATASOURCE_DATABASE")
	dbconf := fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s?charset=utf8mb4",
		datasourceUser,
		datasourcePassword,
		datasourceHost,
		datasourcePort,
		datasourceDatabase,
	)

	db, err := sql.Open("mysql", dbconf)

	// // 接続が終了したらクローズする
	// defer db.Close()

	if err != nil {
		fmt.Println(err.Error())
	}

	err = db.Ping()

	if err != nil {
		fmt.Println("データベース接続失敗")
		return
	} else {
		fmt.Println("データベース接続成功")
	}
	defer db.Close()

	e := echo.New()
	//Select
	e.GET("/todos", func(c echo.Context) error {
		rows, err := db.Query("SELECT id, title FROM todo")
		if err != nil {
			log.Fatal(err)
		}
		var result []Todo
		for rows.Next() {
			todo := Todo{}
			if err := rows.Scan(&todo.ID, &todo.Title); err != nil {
				log.Fatal(err)
			}
			result = append(result, todo)
		}
		todos := Todos{result}
		return c.JSON(http.StatusOK, todos)
	})
	//Select detail
	e.GET("/todos/:id", func(c echo.Context) error {
		id, _ := strconv.Atoi(c.Param("id"))
		row := db.QueryRow("SELECT * FROM todo WHERE id=?", id)

		var todo Todo
		if err := row.Scan(&todo.ID, &todo.Title, &todo.Detail, &todo.Expire_date); err != nil {
			log.Fatal(err)
		}

		fmt.Println("id: ", todo.ID, ", title: ", todo.Title, ", detail: ", todo.Detail, "expire_date: ", todo.Expire_date)
		return c.JSON(http.StatusOK, todo)
	})
	//Insert
	e.POST("/create", func(c echo.Context) error {
		title := c.FormValue("title")
		detail := c.FormValue("detil")
		expireDate := c.FormValue("expire_date")
		insert, err := db.Prepare("INSERT INTO todo(title,detail,expire_date) VALUES(?,?,?)")
		if err != nil {
			log.Fatal(err)
		}
		defer insert.Close()
		insert.Exec(title, detail, expireDate)
		return c.String(http.StatusOK, "title:"+title+", detail:"+detail+", expire_date:"+expireDate)
	})
	//Update
	e.POST("todos/:id", func(c echo.Context) error {
		id, _ := strconv.Atoi(c.Param("id"))
		title := c.FormValue("title")
		detail := c.FormValue("detil")
		expireDate := c.FormValue("expire_date")
		update, err := db.Prepare("UPDATE todo SET title = ?,detail = ?,expire_date = ? WHERE id = ?")
		if err != nil {
			log.Fatal(err)
		}
		defer update.Close()
		update.Exec(title, detail, expireDate, id)
		return c.String(http.StatusOK, "title:"+title+", detail:"+detail+", expire_date:"+expireDate)
	})
	e.Logger.Fatal(e.Start(":1323"))
}
