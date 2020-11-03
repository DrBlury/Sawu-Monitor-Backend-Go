package main

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"os"
	"sawu-monitor/config"
	"sawu-monitor/entities"
	"sawu-monitor/kafka"
)

var app = fiber.New()

func main() {
	go kafka.DoKafkaConsumerStuff()
	// Load config.yml
	var defaults config.Conf
	defaults.GetDefaults()

	AddMainController()
	AddEventController()
	AddSearchController()

	//Set default port if not set
	port, isPresent := os.LookupEnv("fiber_port")
	if isPresent == false {
		port = defaults.Port
	}

	var fiberConfig = new(fiber.Config)
	fiberConfig.DisableStartupMessage = true
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
	app.Get("/event/search", func(c *fiber.Ctx) error {
		c.Accepts("application/json")
		return c.SendString("found.")
	})

}

// AddMainController creates the controller for all sorts of things
func AddMainController() {
	app.Static("/", "./public")
}

// AddSearchController creates the controller for search requests
func AddSearchController() {
}
