package main

import (
	"log"
	"os"
	"strconv"
	"strings"

	"github.com/urfave/cli"
)

func main() {
	app := cli.NewApp()
	app.Name = "Code Generator"
	app.Usage = "text"

	myFlags := []cli.Flag{
		cli.StringFlag{
			Name:  "path",
			Value: "models/user.go",
		},
	}

	app.Commands = []cli.Command{
		{
			Name:  "User",
			Usage: "Create user models",
			Flags: myFlags,
			Action: func(c *cli.Context) error {
				filepath := c.String("path")
				userString := "User"
				file, err := os.OpenFile(filepath, os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0644)
				if err != nil {
					log.Fatal(err)
				}

				// Fonksiyon sonunda dosyayı kapat
				defer file.Close()

				// Dosyaya yaz
				len, err := file.WriteString("package models\n\n//" + userString + " is struct for db. \n type " + userString + " struct{\n\tBase\n}\n\n// TableName custom table name function\n func(m " + userString + ") TableName()string { \n return " + "\"" + "hs_" + strings.ToLower(userString) + "\"" + "\n}")
				if err != nil {
					log.Fatal(err)
				} else {
					log.Println("Yazılan byte boyutu: " + strconv.Itoa(len))
				}
				return nil
			},
		},
	}

	// start our application
	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}

}

/*
package models

// $ModelNameUpperCase is struct for db.
type $ModelNameUpperCase struct {
   Base
}

// TableName custom table name function
func (m $ModelNameUpperCase) TableName() string {
   return "hs_$model_name_snake_case"
}

*/
