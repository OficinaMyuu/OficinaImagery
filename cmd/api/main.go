package main

import (
	ipdata "github.com/ipdata/go"
	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"github.com/playwright-community/playwright-go"
	"oficina-img/internal/routes"
	"oficina-img/internal/service"
	"os"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		panic(err)
	}

	err = playwright.Install()
	if err != nil {
		panic(err)
	}

	pw, err := playwright.Run()
	if err != nil {
		panic(err)
	}
	service.InitializePlaywrightService(pw)

	ipDataKey := os.Getenv("IP_DATA_KEY")
	client, err := ipdata.NewClient(ipDataKey)
	if err != nil {
		return
	}
	service.SetIpDataClient(&client)

	e := echo.New()
	e.Static("/static", "./static")

	e.POST("/api/levels/cards", routes.GetLevelCard)
	e.POST("/api/levels/roles", routes.GetLevelsRoles)

	e.GET("/api/external/ips", routes.GetIPData)

	if err := e.Start(":8080"); err != nil {
		e.Logger.Fatal(err)
	}
}
