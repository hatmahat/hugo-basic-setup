package main

import (
	"encoding/json"
	"net/http"

	"github.com/labstack/echo/v4"
)

type User struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func main() {
	e := echo.New()

	//  Serve files from the 'public' directory under the root path '/'
	e.Static("/blog", "first_site/public")

	e.GET("/user", func(c echo.Context) error {
		user := User{
			Name: "John Doe",
			Age:  15,
		}

		c.Response().Header().Set(echo.HeaderContentType, echo.MIMEApplicationJSONCharsetUTF8)
		c.Response().WriteHeader(http.StatusOK)
		json := json.NewEncoder(c.Response()).Encode(user)

		return json
	})

	// Start the server
	e.Logger.Fatal(e.Start(":8080"))
}
