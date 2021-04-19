package main

import (
	"fmt"
	"log"
	"net/url"
	"os"
	"sawu-monitor/config"
	"sawu-monitor/connector"
	"sawu-monitor/entities"
	"sawu-monitor/kafka"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

var app *fiber.App

func main() {
    var fiberConfig = fiber.Config{}
    fiberConfig.DisableStartupMessage = true
    app = fiber.New(fiberConfig)
    app.Use(cors.New())

    go kafka.DoKafkaConsumerStuff()
    // Load config.yml
    var defaults config.Conf
    defaults.GetDefaults()

    connector.ConnectDB()
    AddEventController()
    AddSearchController()

    //Set default port if not set
    port, isPresent := os.LookupEnv("fiber_port")
    if !isPresent {
        port = defaults.Port
    }

    app.Listen(":" + port)
}

// AddEventController creates the controller for events
func AddEventController() {
    // Catches new events and sends them to kafka
    app.Post("/events", func(c *fiber.Ctx) error {
        c.Accepts("application/json")

        event := new(entities.KafkaNextStepEvent)
        if err := c.BodyParser(event); err != nil {
            return err
        }

        fmt.Println(event)
        topic := fmt.Sprintf("%s-%s", event.ProcessName, event.ProcessStep)
        go kafka.SendNextStepEvent(topic, *event)
        return c.SendString("sent.")
    })

    app.Get("/events", func(c *fiber.Ctx) error {
        // TODO add pagination?
        processInstanceID := c.Query("processInstanceID")
        if processInstanceID == "" {
            return c.Status(400).JSON(&fiber.Map{
                "error": "Query 'processInstanceID' is required",
            })
        }

        events := connector.FindProcessEventsByProcessInstanceID(processInstanceID)
        return c.JSON(&events)
    })

    app.Get("/processes", func(c *fiber.Ctx) error {
        value, err := url.QueryUnescape(c.Query("value"))
        if err != nil {
            log.Fatal(err)
            return c.Status(400).JSON(&fiber.Map{
                "error": fmt.Sprintf("Failed to decode query param 'value': %s", c.Params("value")),
            })
        }

        if value == "" {
            processInstanceInfos := connector.FindAllProcessesInstanceInfo()
            return c.JSON(&processInstanceInfos)
        }

        processInstanceInfos := connector.FindProcessInstanceInfoByDataValue(value)
        return c.JSON(&processInstanceInfos)
    })

}

// AddSearchController creates the controller for search requests
func AddSearchController() {
}
