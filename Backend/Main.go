package main

import (
	"encoding/json"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"os"
	"sawu-monitor/config"
	"sawu-monitor/connector"
	"sawu-monitor/entities"
	"sawu-monitor/kafka"
)

var app *fiber.App

func main() {
	var fiberConfig = fiber.Config{}
	fiberConfig.DisableStartupMessage = true
	app = fiber.New(fiberConfig)

	go kafka.DoKafkaConsumerStuff()
	// Load config.yml
	var defaults config.Conf
	defaults.GetDefaults()

	connector.ConnectDB()
	AddMainController()
	AddEventController()
	AddSearchController()

	//Set default port if not set
	port, isPresent := os.LookupEnv("fiber_port")
	if isPresent == false {
		port = defaults.Port
	}

	app.Listen(":" + port)
}

// AddEventController creates the controller for events
func AddEventController() {
	// Catches new events and sends them to kafka
	app.Get("/event/new", func(c *fiber.Ctx) error {
		c.Accepts("application/json")

		event := new(entities.NextStepEvent)
		if err := c.BodyParser(event); err != nil {
			return err
		}

		fmt.Println(event)
		topic := fmt.Sprintf("%s-%s", event.ProcessName, event.ProcessStep)
		go kafka.SendNextStepEvent(topic, *event)
		return c.SendString("sent.")
	})

	app.Get("/event/search/:processInstanceID", func(c *fiber.Ctx) error {
		processInstanceID := c.Params("processInstanceID")
		events := connector.FindProcessEventsByProcessInstanceID(processInstanceID)

		jsonString, _ := json.Marshal(events)
		fmt.Println(string(jsonString))

		return c.SendString(string(jsonString))
	})

}

// AddMainController creates the controller for all sorts of things
func AddMainController() {
	app.Static("/", "./public")
}

// AddSearchController creates the controller for search requests
func AddSearchController() {
}
