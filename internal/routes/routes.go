package routes

import (
	"github.com/aprianfirlanda/go-log-producer/internal/config"
	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"
	"math/rand"
	"strconv"
	"strings"
	"time"
)

func Register(httpServer *fiber.App) {
	log := config.Log

	logs := httpServer.Group("/logs")

	// GET /logs/exception?message=Simulated%20failure
	logs.Get("/exception", func(c *fiber.Ctx) error {
		message := c.Query("message", "Simulated failure")
		// Simulate an error and log it (like catching an exception)
		err := fiber.NewError(fiber.StatusInternalServerError, message)
		log.WithError(err).Error("An exception occurred")
		return c.SendString("Exception logged")
	})

	// GET /logs/bulk?level=INFO&count=10&message=Bulk%20message&sleepMs=0
	logs.Get("/bulk", func(c *fiber.Ctx) error {
		level := c.Query("level", "INFO")
		count := c.QueryInt("count", 10)
		message := c.Query("message", "Bulk message")
		sleepMs := c.QueryInt("sleepMs", 0)

		for i := 1; i <= count; i++ {
			msg := "[bulk " + strconv.Itoa(i) + "/" + strconv.Itoa(count) + "] " + message
			logBuilder(log, level, msg)
			if sleepMs > 0 {
				time.Sleep(time.Duration(sleepMs) * time.Millisecond)
			}
		}
		return c.SendString("Bulk logged: " + strconv.Itoa(count))
	})

	// GET /logs/random
	logs.Get("/random", func(c *fiber.Ctx) error {
		lvls := []string{"TRACE", "DEBUG", "INFO", "WARN", "ERROR"}
		lvl := lvls[rand.Intn(len(lvls))]
		logBuilder(log, lvl, "Random log")
		return c.SendString("OK")
	})

	// GET /logs/:level?message=Ping
	logs.Get("/:level", func(c *fiber.Ctx) error {
		level := c.Params("level")
		message := c.Query("message", "Ping")
		logBuilder(log, level, "[manual] "+message)
		return c.SendString("OK")
	})
}

func logBuilder(log *logrus.Logger, level, msg string) {
	switch strings.ToUpper(level) {
	case "TRACE":
		log.Trace(msg)
	case "DEBUG":
		log.Debug(msg)
	case "INFO":
		log.Info(msg)
	case "WARN", "WARNING":
		log.Warn(msg)
	case "ERROR":
		log.Error(msg)
	default:
		log.Infof("[unknown-level:%s] %s", level, msg)
	}
}
