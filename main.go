package main

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo"
)

type User struct {
	Name  string `jsonL:"name`
	Email string `json:"email`
}

func getUser(c echo.Context) error {
	id := c.Param("id")
	return c.String(http.StatusOK, id)
}
func show(c echo.Context) error {
	// 从请求参数里获取 team 和 member 的值
	team := c.QueryParam("team")
	member := c.QueryParam("member")
	return c.String(http.StatusOK, "team:"+team+", member:"+member)
}
func save(c echo.Context) error {
	// 获取 name 和 email 的值
	name := c.FormValue("name")
	email := c.FormValue("email")
	return c.String(http.StatusOK, "name:"+name+", email:"+email)
}
func main() {
	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello world")
	})

	e.GET("/users/:id", getUser)
	e.GET("/show", show)
	e.POST("/save", save)
	e.POST("/users", func(c echo.Context) (err error) {
		fmt.Println(err)
		u := new(User)
		if err = c.Bind(u); err != nil {
			return
		}
		return c.JSON(http.StatusCreated, u)
	})
	e.Logger.Fatal(e.Start(":1323"))
}
