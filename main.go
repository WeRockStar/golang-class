package main

import (
	"log"
	"net/http"
	"strconv"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

// Handler
func hello(c echo.Context) error {
	return c.String(http.StatusOK, "Hello, World!")
}

func add(c echo.Context) error {
	var a, b int
	var err error
	if a, err = strconv.Atoi(c.Param("a")); err != nil {
		return c.String(http.StatusBadRequest, "invalid value a")
	}
	if b, err = strconv.Atoi(c.Param("b")); err != nil {
		return c.String(http.StatusBadRequest, "invalid value b")
	}

	return c.String(http.StatusOK, strconv.Itoa(a+b))
}

func main() {
	// Echo instance
	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.CORS())

	// Routes
	e.GET("/", hello)
	e.GET("/add/:a/:b", add)
	e.GET("/fizzbuzz/:number", func(c echo.Context) error {
		number, err := strconv.Atoi(c.Param("number"))
		if err != nil {
			log.Fatal("Cannot convert string to int")
		}
		return c.JSON(http.StatusOK, map[string]string{"answer": fizzbuzz(number)})
	})

	// Start server
	e.Logger.Fatal(e.Start(":1323"))
}

func fizzbuzz(number int) string {
	if number%3 == 0 && number%5 == 0 {
		return "FizzBuzz"
	} else if number%3 == 0 {
		return "Fizz"
	} else if number%5 == 0 {
		return "Buzz"
	}
	return strconv.Itoa(number)
}
