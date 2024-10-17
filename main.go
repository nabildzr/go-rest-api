package main

import (
	"encoding/json"
	"net/http"

	"github.com/golang-todos/database"
	"github.com/labstack/echo"
)

type CreateRequest struct {
	Title       string `json:"title"`
	Description string `json:"description"`
}

func main() {
	db := database.InitDb()
	defer db.Close()

	err := db.Ping()
	if err != nil {
		panic(err)
	}

	e := echo.New()

	e.POST("/todos", func(ctx echo.Context) error {
			rows, err := db.QueryRow("SELECT * FROM todos")
			if err != nil {
				return ctx.String(http.StatusInternalServerError, err.Error())
			}

			for rows.Next() {
				var id int
				var title string
				var description string
				var done int


				err := rows.Scan(&id, &title, &description, &done)
				if err != nil {
					return ctx.String(http.StatusInternalServerError, err.Error())
				}
	
	
			}
		},
	)

	e.POST("/todos", func(ctx echo.Context) error {

		var request CreateRequest
		json.NewDecoder(ctx.Request().Body).Decode(&request)

		_, err := db.Exec(
			"INSERT INTO todos (title, description, done) VALUES (?, ?, 0)",
			request.Title,
			request.Description,
		)
		if err != nil {

			return ctx.String(http.StatusInternalServerError, err.Error())
		}

		return ctx.String(http.StatusOK, "OK")
	})

	e.Start(":8080")
}
