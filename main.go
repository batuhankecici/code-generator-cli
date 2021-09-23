package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"strings"

	"github.com/iancoleman/strcase"
	"github.com/urfave/cli"
)

func main() {
	app := cli.NewApp()
	app.Name = "Code Generator"
	app.Usage = "Generates code according to DDD"

	restTestFlags := []cli.Flag{
		cli.StringFlag{
			Name:  "path",
			Value: "../cmd/server/rest/",
		},
		cli.StringFlag{
			Name: "rest-test-name",
		},
	}
	restFlags := []cli.Flag{
		cli.StringFlag{
			Name:  "path",
			Value: "../cmd/server/rest/",
		},
		cli.StringFlag{
			Name: "rest-name",
		},
	}
	createFlags := []cli.Flag{
		cli.StringFlag{
			Name: "create-name",
		},
	}
	modelFlags := []cli.Flag{
		cli.StringFlag{
			Name:  "path",
			Value: "../models/",
		},
		cli.StringFlag{
			Name: "model-name",
		},
	}
	managerFlags := []cli.Flag{
		cli.StringFlag{
			Name:  "path",
			Value: "../models/managers/",
		},
		cli.StringFlag{
			Name: "manager-name",
		},
	}
	exchangeFlags := []cli.Flag{
		cli.StringFlag{
			Name:  "path",
			Value: "../exchanges/",
		},
		cli.StringFlag{
			Name: "exchange-name",
		},
	}
	controllerFlags := []cli.Flag{
		cli.StringFlag{
			Name:  "path",
			Value: "../controller/",
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
					fmt.Println("model name cannot be empty")
				} else {

					modelString := strcase.ToCamel(value)
					if _, err := os.Stat(filepath + strings.ToLower(modelString) + ".go"); err == nil {
						fmt.Printf("%s already exists in models", modelString)
					} else {
						file, err := os.OpenFile(filepath+strings.ToLower(modelString)+".go", os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0644)
						if err != nil {
							log.Fatalf("models directory does not find %s", err.Error())
						}
						defer file.Close()
						len, err := ioutil.ReadFile("source/modelTemplate.txt")
						if err != nil {
							log.Fatal(err)
						}
						if strings.Contains(string(len), "$ModelNameUpperCase") || strings.Contains(string(len), "$model_name_snake_case") {
							len = []byte(strings.Replace(string(len), "$ModelNameUpperCase", modelString, -1))
							len = []byte(strings.Replace(string(len), "$model_name_snake_case", strcase.ToSnake(modelString), -1))
						}
						file.WriteString(string(len))
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
					if _, err := os.Stat(filepath + strings.ToLower(managerString) + ".go"); err == nil {
						fmt.Printf("%s already exists in managers", managerString)
					} else {
						file, err := os.OpenFile(filepath+strings.ToLower(managerString)+".go", os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0644)
						if err != nil {
							log.Fatalf("managers directory does not find %s", err.Error())
						}
						defer file.Close()
						len, err := ioutil.ReadFile("source/managerTemplate.txt")
						if err != nil {
							log.Fatal(err)
						}
						if strings.Contains(string(len), "${UpperCaseName}") || strings.Contains(string(len), "${camelCaseName}") {
							len = []byte(strings.Replace(string(len), "${UpperCaseName}", managerString, -1))
							len = []byte(strings.Replace(string(len), "${camelCaseName}", strcase.ToLowerCamel(managerString), -1))
						}
						file.WriteString(string(len))
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
					if _, err := os.Stat(filepath + strings.ToLower(exchangeString) + ".go"); err == nil {
						fmt.Printf("%s already exists in exchange", exchangeString)
					} else {
						file, err := os.OpenFile(filepath+strings.ToLower(exchangeString)+".go", os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0644)
						if err != nil {
							log.Fatalf("exchanges directory does not find %s", err.Error())
						}
						defer file.Close()
						len, err := ioutil.ReadFile("source/exchangeTemplate.txt")
						if err != nil {
							log.Fatal(err)
						}
						if strings.Contains(string(len), "${Name}") || strings.Contains(string(len), "${name_snake_case}") {
							len = []byte(strings.Replace(string(len), "${Name}", exchangeString, -1))
							len = []byte(strings.Replace(string(len), "${name_snake_case}", strcase.ToSnake(exchangeString), -1))
						}
						file.WriteString(string(len))
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
					if _, err := os.Stat(filepath + strings.ToLower(controllerString) + ".go"); err == nil {
						fmt.Printf("%s already exists in controller", controllerString)
					} else {
						file, err := os.OpenFile(filepath+strings.ToLower(controllerString)+".go", os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0644)
						if err != nil {
							log.Fatalf("controller directory does not find %s", err.Error())
						}
						defer file.Close()
						len, err := ioutil.ReadFile("source/controllerTemplate.txt")
						if err != nil {
							log.Fatal(err)
						}
						if strings.Contains(string(len), "${Name}") || strings.Contains(string(len), "${name}") {
							len = []byte(strings.Replace(string(len), "${Name}", controllerString, -1))
							len = []byte(strings.Replace(string(len), "${name}", strcase.ToLowerCamel(controllerString), -1))
						}
						file.WriteString(string(len))
					}
				}
				return nil
			},
		},
		{
			Name:  "create",
			Usage: "Create models,managers,exchanges,controller,rest and rest_test template",
			Flags: createFlags,
			Action: func(c *cli.Context) error {

				createName := c.String("create-name")
				bin, err := exec.LookPath("go")
				if err != nil {
					fmt.Println(err)
				}
				cmdModel := exec.Command(bin, "run", "main.go", "model", "--model-name", createName)
				cmdModel.Stdout = os.Stdout
				cmdModel.Stderr = os.Stderr
				cmdModel.Run()

				cmdManager := exec.Command(bin, "run", "main.go", "manager", "--manager-name", createName)
				cmdManager.Stdout = os.Stdout
				cmdManager.Stderr = os.Stderr
				cmdManager.Run()

				cmdExchange := exec.Command(bin, "run", "main.go", "exchange", "--exchange-name", createName)
				cmdExchange.Stdout = os.Stdout
				cmdExchange.Stderr = os.Stderr
				cmdExchange.Run()

				cmdController := exec.Command(bin, "run", "main.go", "controller", "--controller-name", createName)
				cmdController.Stdout = os.Stdout
				cmdController.Stderr = os.Stderr
				cmdController.Run()

				cmdRest := exec.Command(bin, "run", "main.go", "rest", "--rest-name", createName)
				cmdRest.Stdout = os.Stdout
				cmdRest.Stderr = os.Stderr
				cmdRest.Run()

				cmdRestTest := exec.Command(bin, "run", "main.go", "restTest", "--rest-test-name", createName)
				cmdRestTest.Stdout = os.Stdout
				cmdRestTest.Stderr = os.Stderr
				cmdRestTest.Run()

				return nil
			},
		},
		{
			Name:  "rest",
			Usage: "Create rest template in rest folder",
			Flags: restFlags,
			Action: func(c *cli.Context) error {
				filepath := c.String("path")
				value := c.String("rest-name")
				if value == "" {
					fmt.Println("rest name cannot be empty")
				} else {

					restString := strcase.ToCamel(value)
					if _, err := os.Stat(filepath + strings.ToLower(restString) + ".go"); err == nil {
						fmt.Printf("%s already exists in rest", restString)
					} else {
						file, err := os.OpenFile(filepath+strings.ToLower(restString)+".go", os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0644)
						if err != nil {
							log.Fatalf("rest directory does not find %s", err.Error())
						}
						defer file.Close()
						len, err := ioutil.ReadFile("source/restTemplate.txt")
						if err != nil {
							log.Fatal(err)
						}
						if strings.Contains(string(len), "${Name}") || strings.Contains(string(len), "${name}") || strings.Contains(string(len), "${nameCamelCase}") {
							len = []byte(strings.Replace(string(len), "${Name}", restString, -1))
							len = []byte(strings.Replace(string(len), "${nameCamelCase}", strcase.ToLowerCamel(restString), -1))
						}
						file.WriteString(string(len))
					}
				}
				return nil
			},
		},
		{
			Name:  "restTest",
			Usage: "Create rest test template in rest folder",
			Flags: restTestFlags,
			Action: func(c *cli.Context) error {
				filepath := c.String("path")
				value := c.String("rest-test-name")
				if value == "" {
					fmt.Println("rest test name cannot be empty")
				} else {

					restTestString := strcase.ToCamel(value)
					if _, err := os.Stat(filepath + strings.ToLower(restTestString) + "_test" + ".go"); err == nil {
						fmt.Printf("%s already exists in rest", restTestString)
					} else {
						file, err := os.OpenFile(filepath+strings.ToLower(restTestString)+"_test"+".go", os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0644)
						if err != nil {
							log.Fatalf("rest directory does not find %s", err.Error())
						}
						defer file.Close()
						lens, err := ioutil.ReadFile("source/restTestTemplate.txt")
						if err != nil {
							log.Fatal(err)
						}
						if strings.Contains(string(lens), "${ModelName}") || strings.Contains(string(lens), "${modelName}") || strings.Contains(string(lens), "${model_name_snake_case}") || strings.Contains(string(lens), "${routeSuffix}") {
							lens = []byte(strings.Replace(string(lens), "${ModelName}", restTestString, -1))
							lens = []byte(strings.Replace(string(lens), "${modelName}", strcase.ToLowerCamel(restTestString), -1))
							lens = []byte(strings.Replace(string(lens), "${model_name_snake_case}", strcase.ToSnake(restTestString), -1))
							substring := restTestString[len(restTestString)-1:]
							generated := restTestString[0 : len(restTestString)-1]
							if substring == "y" {
								lens = []byte(strings.Replace(string(lens), "${routeSuffix}", strcase.ToKebab(generated)+"ies", -1))
							} else {
								lens = []byte(strings.Replace(string(lens), "${routeSuffix}", strcase.ToKebab(restTestString)+"s", -1))
							}

						}
						file.WriteString(string(lens))
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
