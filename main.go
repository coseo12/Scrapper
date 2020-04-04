package main

import (
	"os"
	"strings"

	"github.com/example/learngo/scrapper"
	"github.com/labstack/echo"
)

const fileName string = "jobs.csv"

func handelHome(c echo.Context) error {
	return c.File("home.html")
}

func handelScrape(c echo.Context) error {
	defer os.Remove(fileName)
	term := strings.ToLower(scrapper.CleanString(c.FormValue("term")))
	scrapper.Scrape(term)
	return c.Attachment(fileName, fileName)
}

func main() {
	// scrapper.Scrape("vue.js")
	e := echo.New()
	e.GET("/", handelHome)
	e.POST("/scrape", handelScrape)
	e.Logger.Fatal(e.Start(":1323"))
}
