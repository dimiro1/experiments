package main

import (
	"errors"
	"net/http"
	"strings"

	"github.com/asaskevich/govalidator"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

type todo struct {
	ID    uint64 `query:"id"    json:"id"`
	Title string `query:"title" json:"title"`
	Done  bool   `query:"done"  json:"done"`
}

type user struct {
	Email         string `query:"email" json:"email" valid:"email~invalid,required~missing_field"`
	FavoriteColor string `query:"favoriteColor" json:"favoriteColor" valid:"hexcolor~invalid,required~missing_field"`
}

type validationFailed struct {
	Message string            `json:"message"`
	Errors  []validationError `json:"errors"`
}

func (v validationFailed) Error() string {
	var err string

	for _, e := range v.Errors {
		err += e.Error() + ";"
	}
	return err
}

type validationError struct {
	Field  string `json:"field"`
	Reason string `json:"reason"`
}

func (v validationError) Error() string {
	return v.Field + ": " + v.Reason
}

// Validator implements the echo.Validator interface
type Validator struct{}

// Validate validates the given struct
func (v Validator) Validate(i interface{}) error {
	_, err := govalidator.ValidateStruct(i)
	if err == nil {
		return nil
	}

	if errors, ok := err.(govalidator.Errors); ok {
		var (
			validationFailed = validationFailed{
				Message: "Validation Failed",
			}
			other []string
		)

		for _, e := range errors {
			if validatorErr, ok := e.(govalidator.Error); ok {
				validationFailed.Errors = append(validationFailed.Errors, validationError{
					Field:  validatorErr.Name,
					Reason: validatorErr.Err.Error(),
				})
			} else {
				// Otherwise add the error on unkown field
				other = append(other, e.Error())
			}
		}

		if len(other) > 0 {
			validationFailed.Errors = append(validationFailed.Errors, validationError{
				Field:  "unknown",
				Reason: strings.Join(other, ";"),
			})
		}

		return validationFailed
	}

	return err
}

func main() {
	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Validator = Validator{}

	e.GET("/", func(ctx echo.Context) error {
		return ctx.String(200, "Hello World")
	})

	e.GET("/panic", func(ctx echo.Context) error {
		panic("Panic")
	})

	e.GET("/error", func(ctx echo.Context) error {
		return echo.NewHTTPError(http.StatusInternalServerError)
	})

	e.GET("/genericError", func(ctx echo.Context) error {
		return errors.New("Some error")
	})

	e.GET("/unauthorized", func(ctx echo.Context) error {
		return echo.ErrUnauthorized
	})

	e.GET("/todos", func(ctx echo.Context) error {
		todos := []todo{
			{
				ID:    1,
				Title: "Example",
				Done:  false,
			},
		}
		ctx.Response().Header().Set("Link", "<http://www.google.com>; rel=\"next\"")
		return ctx.JSON(http.StatusOK, todos)
	})

	e.GET("/map", func(ctx echo.Context) error {
		return ctx.JSON(http.StatusOK, echo.Map{
			"Hello": "World",
			"One":   1,
		})
	})

	e.GET("/hello/:name", func(ctx echo.Context) error {
		return ctx.JSON(http.StatusOK, ctx.Param("name"))
	})

	e.GET("/query", func(ctx echo.Context) error {
		var todo todo

		if err := ctx.Bind(&todo); err != nil {
			return err
		}

		return ctx.JSON(http.StatusOK, todo)
	})

	e.GET("/validator", func(ctx echo.Context) error {
		var user user

		if err := ctx.Bind(&user); err != nil {
			return err
		}

		if err := ctx.Validate(user); err != nil {
			if vf, ok := err.(validationFailed); ok {
				return echo.NewHTTPError(http.StatusBadRequest, vf)
			}

			return err
		}

		return ctx.JSON(http.StatusOK, user)
	})

	e.Start(":9000")
}
