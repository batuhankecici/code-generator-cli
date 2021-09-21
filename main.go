package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"strconv"
	"strings"

	"github.com/iancoleman/strcase"
	"github.com/urfave/cli"
)

func main() {
	app := cli.NewApp()
	app.Name = "Code Generator"
	app.Usage = "Generates code according to DDD"

	createFlags := []cli.Flag{
		cli.StringFlag{
			Name: "create-name",
		},
	}
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
	exchangeFlags := []cli.Flag{
		cli.StringFlag{
			Name:  "path",
			Value: "exchanges/",
		},
		cli.StringFlag{
			Name: "exchange-name",
		},
	}
	controllerFlags := []cli.Flag{
		cli.StringFlag{
			Name:  "path",
			Value: "controller/",
		},
		cli.StringFlag{
			Name: "controller-name",
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

					modelString := strcase.ToCamel(value)
					file, err := os.OpenFile(filepath+strings.ToLower(modelString)+".go", os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0644)
					if err != nil {
						log.Fatal(err)
					}

					// Fonksiyon sonunda dosyayı kapat
					defer file.Close()

					// Dosyaya yaz
					len, err := file.WriteString("package models\n\n//" + modelString + " is struct for db. \n type " + modelString + " struct{\n\tBase\n}\n\n// TableName custom table name function\n func(m " + modelString + ") TableName()string { \n return " + "\"" + "hs_" + strcase.ToSnake(modelString) + "\"" + "\n}")
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
					managerString := strcase.ToCamel(value)
					file, err := os.OpenFile(filepath+strings.ToLower(managerString)+".go", os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0644)
					if err != nil {
						log.Fatal(err)
					}
					// Fonksiyon sonunda dosyayı kapat
					defer file.Close()

					// Dosyaya yaz
					len, err := file.WriteString("package managers\n\nimport (\n\t" + "\"" + "github.com/biges/hybrone-sentinel-backend/models" + "\"\n\t" + "\"github.com/biges/hybrone-sentinel-backend/helpers" + "\"\n\t" + "\"" + "gorm.io/gorm" + "\"\n)\n\n//" + managerString + "Storage defines the behaviors required to perform CRUD operations on a Device model.\ntype " + managerString + "Storage interface {\n\tCreate(" + strcase.ToLowerCamel(managerString) + " *models." + managerString + ") error\n\tUpdate(" + strcase.ToLowerCamel(managerString) + " *models." + managerString + ") error\n\tDelete(id uint) error\n\tGet(" + strcase.ToLowerCamel(managerString) + " *models." + managerString + ") error\n\tList(" + strcase.ToLowerCamel(managerString) + " []*models." + managerString + ")error\n}\n\n// " + managerString + "Manager represents Device related CRUD operations.\ntype " + managerString + "Manager struct {\n\tdb *gorm.DB\n}\n\n// New" + managerString + "Manager returns DeviceManager instance.\nfunc New" + managerString + "Manager(db *gorm.DB) " + managerString + "Storage {\n\treturn &" + managerString + "Manager{db: db}\n}\n\n// Create creates new " + managerString + " record.\nfunc (m " + managerString + "Manager) Create(" + strcase.ToLowerCamel(managerString) + " *models." + managerString + ") error {\n\tresult := m.db.Create(" + strcase.ToLowerCamel(managerString) + ")\n\treturn helpers.GORMErrConverter(result.Error)\n}\n\n// Update updates given " + managerString + " record.\nfunc (m " + managerString + "Manager) Update(" + strcase.ToLowerCamel(managerString) + " *models." + managerString + ") error {\n\tresult := m.db.Save(" + strcase.ToLowerCamel(managerString) + ")\n\treturn helpers.GORMErrConverter(result.Error)\n}\n\n// Delete deletes " + managerString + " record.\nfunc (m " + managerString + "Manager) Delete(id uint) error {\n\tresult := m.db.Delete(&models." + managerString + "{Base: models.Base{ID: id}})\n\treturn result.Error\n}\n\n// Get get " + managerString + " by given id.\nfunc (m " + managerString + "Manager) Get(" + strcase.ToLowerCamel(managerString) + " *models." + managerString + ") error {\n\tresult := m.db.First(" + strcase.ToLowerCamel(managerString) + ")\n\treturn helpers.GORMErrConverter(result.Error)\n}\n\n// List get " + managerString + "s list.\nfunc (m " + managerString + "Manager) List(" + strcase.ToLowerCamel(managerString) + "s *[]models." + managerString + ") error {\n\tresult := m.db.Find(" + strcase.ToLowerCamel(managerString) + "s)\n\treturn helpers.GORMErrConverter(result.Error)\n}")
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
			Name:  "exchange",
			Usage: "Create new exchange",
			Flags: exchangeFlags,
			Action: func(c *cli.Context) error {
				filepath := c.String("path")
				value := c.String("exchange-name")
				if value == "" {
					fmt.Println("exchange name cannot be empty")
				} else {

					exchangeString := strcase.ToCamel(value)
					file, err := os.OpenFile(filepath+strings.ToLower(exchangeString)+".go", os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0644)
					if err != nil {
						log.Fatal(err)
					}

					// Fonksiyon sonunda dosyayı kapat
					defer file.Close()

					// Dosyaya yaz
					len, err := file.WriteString("package exchanges\n\nimport \"github.com/biges/hybrone-sentinel-backend/models\"\n\ntype " + exchangeString + "Service interface {\n\t" + exchangeString + "Create(" + exchangeString + "CreateRequest, " + exchangeString + "CreateResponse) error\n\t" + exchangeString + "Update(" + exchangeString + "UpdateRequest, " + exchangeString + "UpdateResponse) error\n\t" + exchangeString + "Delete(" + exchangeString + "DeleteRequest, " + exchangeString + "DeleteResponse) error\n\t" + exchangeString + "Get(" + exchangeString + "GetRequest, " + exchangeString + "GetResponse) error\n\t" + exchangeString + "List(" + exchangeString + "ListRequest, " + exchangeString + "ListResponse) error\n}\n\n// " + exchangeString + "CreateRequest holds data for creating new " + exchangeString + ".\ntype " + exchangeString + "CreateRequest struct {\n\t" + exchangeString + " *models." + exchangeString + " `json:\"" + strcase.ToSnake(exchangeString) + "\" validate:\"required\"`\n}\n\n// " + exchangeString + "CreateResponse holds newly created " + exchangeString + " instance data.\ntype " + exchangeString + "CreateResponse struct {\n\tBaseResponse\n\t" + exchangeString + " *models." + exchangeString + " `json:\"" + strcase.ToSnake(exchangeString) + "\"`\n}\n\n// " + exchangeString + "UpdateRequest holds data for updating a " + exchangeString + ".\ntype " + exchangeString + "UpdateRequest struct {\n\t" + exchangeString + " *models." + exchangeString + " `json:\"" + strcase.ToSnake(exchangeString) + "\" validate:\"required\"`\n}\n\n// " + exchangeString + "UpdateResponse holds newly updated " + exchangeString + " instance data.\ntype " + exchangeString + "UpdateResponse struct {\n\tBaseResponse\n\t" + exchangeString + " *models." + exchangeString + " `json:\"" + strcase.ToSnake(exchangeString) + "\"`\n}\n\n// " + exchangeString + "DeleteRequest holds data for deleting new " + exchangeString + ".\ntype " + exchangeString + "DeleteRequest struct {\n\tID uint `json:\"id\"`\n}\n\n// " + exchangeString + "DeleteResponse holds newly deleted " + exchangeString + " query error.\ntype " + exchangeString + "DeleteResponse struct {\n\tBaseResponse\n\tID uint `json:\"id\"`\n}\n\n// " + exchangeString + "GetRequest holds data for get " + exchangeString + ".\ntype " + exchangeString + "GetRequest struct {\n\t" + exchangeString + " *models." + exchangeString + " `json:\"" + strcase.ToSnake(exchangeString) + "\" validate:\"required\"`\n}\n\n// " + exchangeString + "GetResponse holds " + exchangeString + " instance data.\ntype " + exchangeString + "GetResponse struct {\n\tBaseResponse\n\t" + exchangeString + " *models." + exchangeString + " `json:\"" + strcase.ToSnake(exchangeString) + "\"`\n}\n\n// " + exchangeString + "ListRequest holds filters for " + exchangeString + " list.\ntype " + exchangeString + "ListRequest struct {\n\t// filters\n}\n\n// " + exchangeString + "ListResponse holds " + exchangeString + "s data.\ntype " + exchangeString + "ListResponse struct {\n\tBaseResponse\n\t" + exchangeString + "s *[]models." + exchangeString + " `json:\"" + strcase.ToSnake(exchangeString) + "s\"`\n}")
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
			Name:  "controller",
			Usage: "Create new controller",
			Flags: controllerFlags,
			Action: func(c *cli.Context) error {
				filepath := c.String("path")
				value := c.String("controller-name")
				if value == "" {
					fmt.Println("controller name cannot be empty")
				} else {

					controllerString := strcase.ToCamel(value)
					file, err := os.OpenFile(filepath+strings.ToLower(controllerString)+".go", os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0644)
					if err != nil {
						log.Fatal(err)
					}

					// Fonksiyon sonunda dosyayı kapat
					defer file.Close()

					// Dosyaya yaz
					len, err := file.WriteString("package controller\n\nimport (\n\t\"github.com/biges/hybrone-sentinel-backend/exchanges\"\n\t\"github.com/biges/hybrone-sentinel-backend/models\"\n)\n\n// " + controllerString + "Create is a service method of " + controllerString + "Service interfaces.\nfunc (s Service) " + controllerString + "Create(req exchanges." + controllerString + "CreateRequest) (exchanges." + controllerString + "CreateResponse, error) {\nres := exchanges." + controllerString + "CreateResponse{}\nerr := s." + controllerString + "Manager.Create(req." + controllerString + ")\nif err != nil {\n\tres.Message = \"Oluşturulamadı.\"\n\treturn res,err\n}\nres." + controllerString + " = req." + controllerString + "\nreturn res,err\n}\n\n// " + controllerString + "Update is a service method of " + controllerString + "Service interfaces.\nfunc (s Service) " + controllerString + "Update(req exchanges." + controllerString + "UpdateRequest) (exchanges." + controllerString + "UpdateResponse, error) {\nres := exchanges." + controllerString + "UpdateResponse{}\nerr := s." + controllerString + "Manager.Update(req." + controllerString + ")\nif err != nil {\n\tres.Message = \"Oluşturulamadı.\"\n\treturn res,err\n}\nres." + controllerString + " = req." + controllerString + "\nreturn res,err\n}\n\n// " + controllerString + "Delete is a service method of " + controllerString + "Service interfaces.\nfunc (s Service) " + controllerString + "Delete(req exchanges." + controllerString + "DeleteRequest) (exchanges." + controllerString + "DeleteResponse, error) {\nres := exchanges." + controllerString + "DeleteResponse{}\nerr := s." + controllerString + "Manager.Delete(req.ID)\nif err != nil {\n\tres.Message = \"Silinemedi.\"\n\treturn res,err\n}\nres.ID = req.ID\nreturn res,err\n}\n\n// " + controllerString + "Get is a service method of " + controllerString + "Service interfaces.\nfunc (s Service) " + controllerString + "Get(req exchanges." + controllerString + "GetRequest) (exchanges." + controllerString + "GetResponse, error) {\nres := exchanges." + controllerString + "GetResponse{}\nerr := s." + controllerString + "Manager.Get(req." + controllerString + ")\nif err != nil {\n\tres.Message = \"Bulunamadı.\"\n\treturn res,err\nres." + controllerString + " = req." + controllerString + "\nreturn res,err\n}\n\n// " + controllerString + "List is a service method of " + controllerString + "Service interfaces.\nfunc (s Service) " + controllerString + "List(req exchanges." + controllerString + "ListRequest) (exchanges." + controllerString + "ListResponse, error) {\nres := exchanges." + controllerString + "ListResponse{}\n" + strcase.ToLowerCamel(controllerString) + "s := new([]models." + controllerString + ")\nerr := s." + controllerString + "Manager.List(" + strcase.ToLowerCamel(controllerString) + "s)\nif err != nil {\n\tres.Message = \"Bulunamadı.\"\n\treturn res,err\n}\nres." + controllerString + "s = " + strcase.ToLowerCamel(controllerString) + "s\nreturn res,err\n}")
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
			Name:  "create",
			Usage: "Create models,managers,exchanges,controller template",
			Flags: createFlags,
			Action: func(c *cli.Context) error {

				createName := c.String("create-name")
				bin, err := exec.LookPath("go")
				if err != nil {
					fmt.Println(err)
				}
				cmdModel := exec.Command(bin, "run", "main.go", "model", "--model-name", createName)
				err = cmdModel.Run()
				if err != nil {
					return err
				}
				cmdManager := exec.Command(bin, "run", "main.go", "manager", "--manager-name", createName)
				err = cmdManager.Run()
				if err != nil {
					return err
				}
				cmdExchange := exec.Command(bin, "run", "main.go", "exchange", "--exchange-name", createName)
				err = cmdExchange.Run()
				if err != nil {
					return err
				}
				cmdController := exec.Command(bin, "run", "main.go", "controller", "--controller-name", createName)
				err = cmdController.Run()
				if err != nil {
					return err
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
