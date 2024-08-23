package routes

import (
	"os"
	"strconv"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/thesilentline/shorten-url-project/database"
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

	// body now points to a newly created request struct in memory
	body := new(request)	

	if err := c.BodyParser(&body); err!= nil{
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error":"cannot parse json"},)
	}

	// impleament rate limiting
	r2 := database.CreateClient(1)	// redis database client
	defer r2.Close()

	val, err := r2.GET(database.Ctx, c.IP()).Result()
	if err == redis.Nil {
		_ = r2.SET(database.Ctx, c.IP(), os.Getenv("API_QUOTA"), 30*60*time.Second).Err()
	} else {
		val, _ = r2.GET(database.Ctx, c.IP()).Result()
		valInt, _ := strconv.Atoi(val)

		if valInt <= 0 {
			limit, _ := r2.TTL(database.Ctx, c.IP()).Result()
			return c.Status(fiber.StatusServiceUnavailable).JSON(fiber.Map{
				"error": "rate limit exceeded",
				"rate limit reset": limit / time.Nanosecond / time.Minute,
			})

		}
	}

	// check if the url entered is correct
	if !govalidator.IsURL(body.URL){
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error":"invalid URL"},)
	}

	//check for domain error (user enters local host 3000 causing infinite loops)
	if !helpers.RemoveDomainError(body.URL){
		return c.Status(fiber.StatusServiceUnavailable).JSON(fiber.Map{"error":"can't hack the system"},)
	}

	//enforce HTTP SSL
	body.URL = helpers.EnforceHTTP(body.URL)

	// decrement the counter
	r2.Decr(database.Ctx, c.IP())

}