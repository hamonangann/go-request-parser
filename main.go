package main

import (
	"fmt"
	"github.com/alecthomas/kingpin/v2"
	"github.com/fsnotify/fsnotify"
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/spf13/viper"
	"net/http"
	"os"
)

// User role additional info: 1 for Admin, 2 for registered user, empty/0 for guest/visitor temp account
type User struct {
	Name  string `json:"name" form:"name" query:"name" validate:"required"`
	Email string `json:"email" form:"email" query:"email" validate:"required,email"`
	Role  int    `json:"role" form:"role" query:"role" validate:"gte=0,lte=2"`
}

type CustomValidator struct {
	validator *validator.Validate
}

func (cv *CustomValidator) Validate(i any) error {
	return cv.validator.Struct(i)
}

func main() {
	app := kingpin.New("App", "Simple app")
	portFlag := app.Flag("port", "Server port").Short('p').Int()

	kingpin.MustParse(app.Parse(os.Args[1:]))

	r := echo.New()

	// Send config
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")
	viper.SetConfigName("config")

	err := viper.ReadInConfig()
	if err != nil {
		r.Logger.Fatal(err)
	}

	viper.WatchConfig()
	viper.OnConfigChange(func(e fsnotify.Event) {
		fmt.Printf("Config changed: %s", e.Name)
	})

	r.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Skipper:          nil,
		Format:           "method=${method}, uri=${uri}, status=${status}\n",
		CustomTimeFormat: "",
		Output:           nil,
	}))

	r.Validator = &CustomValidator{validator: validator.New()}

	r.HTTPErrorHandler = func(err error, c echo.Context) {
		report, ok := err.(*echo.HTTPError)
		if !ok {
			report = echo.NewHTTPError(http.StatusInternalServerError, err.Error())
		}

		if castedObj, ok := err.(validator.ValidationErrors); ok {
			report.Code = http.StatusBadRequest
			for _, err := range castedObj {
				switch err.Tag() {
				case "required":
					report.Message = fmt.Sprintf("%s is required", err.Field())
				case "email":
					report.Message = fmt.Sprintf("%s is not valid email", err.Value())
				case "gte":
					report.Message = fmt.Sprintf("%s enum is invalid", err.Field())
				case "lte":
					report.Message = fmt.Sprintf("%s enum is invalid", err.Field())
				}
				break
			}
		}

		c.JSON(report.Code, report)
	}

	r.Any("/user", func(c echo.Context) error {
		u := new(User)
		if err := c.Bind(u); err != nil {
			return err
		}
		if err := c.Validate(u); err != nil {
			return err
		}
		return c.JSON(http.StatusOK, u)
	})

	port := fmt.Sprintf(":%d", *portFlag)

	if port == ":0" {
		port = fmt.Sprintf(":%d", viper.GetInt("server.port"))
	}
	if port == ":0" {
		port = ":9000"
	}

	fmt.Printf("server started at %s", port)
	r.Logger.Fatal(r.Start(port))
}
