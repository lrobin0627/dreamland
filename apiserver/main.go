// Copyright 2023 Leo <lipf160627@163.com>. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package main

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()
	e.GET("/hello", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello World Golang, I am coming")
	})
	e.Logger.Fatal(e.Start(":8080"))
}
