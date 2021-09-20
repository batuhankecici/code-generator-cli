package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"

	"github.com/mingrammer/casec"
	"github.com/urfave/cli"
)

func main() {
	app := cli.NewApp()
	app.Name = "Code Generator"
	app.Usage = "Generates code according to DDD"

	modelFlags := []cli.Flag{
		cli.StringFlag{
			Name:  "path",
			Value: "models/",
		},
		cli.StringFlag{
			Name: "model-name",
		},
	}
	managerFlags := []cli.Flag{
		cli.StringFlag{
			Name:  "path",
			Value: "managers/",
		},
		cli.StringFlag{
			Name: "manager-name",
		},
	}
	app.Commands = []cli.Command{
		{
			Name:  "model",
			Usage: "Create new models",
			Flags: modelFlags,
			Action: func(c *cli.Context) error {
				filepath := c.String("path")
				value := c.String("model-name")
				if value == "" {
					fmt.Println("model cannot be empty")
				} else {
					modelString := strings.Title(strings.ToLower(value))
					file, err := os.OpenFile(filepath+strings.ToLower(modelString)+".go", os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0644)
					if err != nil {
						log.Fatal(err)
					}

					// Fonksiyon sonunda dosyayı kapat
					defer file.Close()

					// Dosyaya yaz
					len, err := file.WriteString("package models\n\n//" + modelString + " is struct for db. \n type " + modelString + " struct{\n\tBase\n}\n\n// TableName custom table name function\n func(m " + modelString + ") TableName()string { \n return " + "\"" + "hs_" + casec.ToSnake(modelString) + "\"" + "\n}")
					if err != nil {
						log.Fatal(err)
					} else {
						log.Println("Yazılan byte boyutu: " + strconv.Itoa(len))
					}
				}

				return nil
			},
		},
		{
			Name:  "manager",
			Usage: "Create a new manager",
			Flags: managerFlags,
			Action: func(c *cli.Context) error {
				filepath := c.String("path")
				value := c.String("manager-name")
				if value == "" {
					fmt.Println("manager name cannot be empty")
				} else {
					managerString := strings.Title(strings.ToLower(value))
					file, err := os.OpenFile(filepath+strings.ToLower(managerString)+".go", os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0644)
					if err != nil {
						log.Fatal(err)
					}
					// Fonksiyon sonunda dosyayı kapat
					defer file.Close()

					// Dosyaya yaz
					len, err := file.WriteString("package managers\n\nimport (\n\t" + "\"" + "github.com/biges/hybrone-sentinel-backend/models" + "\"\n\t" + "\"github.com/biges/hybrone-sentinel-backend/helpers" + "\"\n\t" + "\"" + "gorm.io/gorm" + "\"\n)\n\n//" + managerString + "Storage defines the behaviors required to perform CRUD operations on a Device model.\ntype " + managerString + "Storage interface {\n\tCreate(" + casec.ToCamel(managerString) + " *models." + managerString + ") error\n\tUpdate(" + casec.ToCamel(managerString) + " *models." + managerString + ") error\n\tDelete(id uint) error\n\tGet(" + casec.ToCamel(managerString) + " *models." + managerString + ") error\n\tList(" + casec.ToCamel(managerString) + " []*models." + managerString + ")error\n}\n\n// " + managerString + "Manager represents Device related CRUD operations.\ntype " + managerString + "Manager struct {\n\tdb *gorm.DB\n}\n\n// New" + managerString + "Manager returns DeviceManager instance.\nfunc New" + managerString + "Manager(db *gorm.DB) " + managerString + "Storage {\n\treturn &" + managerString + "Manager{db: db}\n}\n\n// Create creates new " + managerString + " record.\nfunc (m " + managerString + "Manager) Create(" + casec.ToCamel(managerString) + " *models." + managerString + ") error {\n\tresult := m.db.Create(" + casec.ToCamel(managerString) + ")\n\treturn helpers.GORMErrConverter(result.Error)\n}\n\n// Update updates given " + managerString + " record.\nfunc (m " + managerString + "Manager) Update(" + casec.ToCamel(managerString) + " *models." + managerString + ") error {\n\tresult := m.db.Save(" + casec.ToCamel(managerString) + ")\n\treturn helpers.GORMErrConverter(result.Error)\n}\n\n// Delete deletes " + managerString + " record.\nfunc (m " + managerString + "Manager) Delete(id uint) error {\n\tresult := m.db.Delete(&models." + managerString + "{Base: models.Base{ID: id}})\n\treturn result.Error\n}\n\n// Get get " + managerString + " by given id.\nfunc (m " + managerString + "Manager) Get(" + casec.ToCamel(managerString) + " *models." + managerString + ") error {\n\tresult := m.db.First(" + casec.ToCamel(managerString) + ")\n\treturn helpers.GORMErrConverter(result.Error)\n}\n\n// List get " + managerString + "s list.\nfunc (m " + managerString + "Manager) List(" + casec.ToCamel(managerString) + "s *[]models." + managerString + ") error {\n\tresult := m.db.Find(" + casec.ToCamel(managerString) + "s)\n\treturn helpers.GORMErrConverter(result.Error)\n}")
					if err != nil {
						log.Fatal(err)
					} else {
						log.Println("Yazılan byte boyutu: " + strconv.Itoa(len))
					}
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
