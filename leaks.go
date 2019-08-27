package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/urfave/cli"
)

type Result struct {
	num_hints int
}

func search(query string) {

	file, err := os.Open("./leaks.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		if strings.Contains(scanner.Text(), query) {
			fmt.Println(scanner.Text())
			//Eventually these should not be printed ad hoc.
			//Result.append(scanner.Text())
		}
	}
	//return Results
}

func main() {
	var query string

	app := cli.NewApp()

	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:  "query, q",
			Value: "",
			Usage: "String to search for",
		},
		cli.BoolFlag{
			Name:  "fuzzy",
			Usage: "keeping it fuzzy",
		},
	}

	app.Name = "LeakSearch"
	app.Usage = "Looks like you've got a leak..."
	app.Action = func(c *cli.Context) error {
		fmt.Println("Can't fix a leak without the write tools...")
		return nil
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
	search(query)
}
