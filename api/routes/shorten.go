package routes

import (
	"time"
	"github.com/gofiber/fiber/v2"
	"github.com/thesilentline/shorten-url-project/helpers"
)

type request struct {
	URL				string			`json:"url"`
	CustomShort		string			`json:"short"`
	Expiry			time.Duration	`json:"expiry"`
}

type response struct {
	URL					string			`json:"url"`
	CustomShort			string			`json:"short"`
	Expiry				time.Duration	`json:"expiry"`
	XRateRemaining		int				`json:"rate_limit"`
	XRateLimitReset		time.Duration	`json:"rate_limit_reset"`
}


func ShortenURL(c *fiber.Ctx) error {

	body := new(request)	//body now points to a newly created request struct in memory

	if err := c.BodyParser(&body); err!= nil{
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error":"cannot parse json"})
	}

	// check if the url entered is correct
	if !govalidator.IsURL(body.URL){
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error":"invalid URL"})
	}

	//check for domain error (user enters local host 3000 causing infinite loops)
	if !helpers.RemoveDomainError(body.URL){
		return c.Status(fiber.StatusServiceUnavailable).JSON(fiber.Map{"error":"can't hack the system"})
	}

	//enforce HTTP SSL
	body.URL = helpers.EnforceHTTP(body.URL)

}