package main

import (
	"fmt"
	"github.com/labstack/echo"
	"net/http"
)

type User struct {
	Name  string `json:"name" form:"name" query:"name"`
	Email string `json:"email" form:"email" query:"email"`
}

func main() {
	r := echo.New()

	r.Any("/user", func(c echo.Context) error {
		u := new(User)
		if err := c.Bind(u); err != nil {
			return err
		}
		return c.JSON(http.StatusOK, u)
	})

	fmt.Println("server started at :9000")
	r.Start(":9000")
}
