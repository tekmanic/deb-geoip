package handlers

import (
	_ "embed"
	"fmt"
	"net"

	"github.com/gofiber/fiber/v2"
	geoip2 "github.com/oschwald/geoip2-golang"
)

// See https://pkg.go.dev/github.com/oschwald/geoip2-golang#City for a full list of options you can use here to modify
// what data is returned for a specific IP.
type ipLookup struct {
	GeoNameID      uint   `json:"geoname_id"`
	City           string `json:"city"`
	Country        string `json:"country"`
	IsoCode        string `json:"iso_code"`
	AccuracyRadius uint16 `json:"accuracy_radius"`
}

//go:embed GeoLite2-City.mmdb
var mmdb []byte

var geoIPDb *geoip2.Reader

func init() {
	// Load MaxMind DB
	var err error
	geoIPDb, err = geoip2.FromBytes(mmdb)
	if err != nil {
		fmt.Println("Unable to load 'GeoLite2-City.mmdb'.")
		panic(err)
	}
}

// GeoIP is a handler for IP address lookups
func GeoIP() fiber.Handler {

	// Return handler
	return func(c *fiber.Ctx) error {
		ipAddr := c.Params("ip", c.IP())

		// Check IP address format
		ip := net.ParseIP(ipAddr)
		if ip == nil {
			return c.Status(400).JSON(map[string]string{"status": "error", "message": "Invalid IP address"})
		}

		// Perform lookup
		record, err := geoIPDb.City(ip)
		if err != nil {
			return err
		}

		// Send response
		return c.JSON(ipLookup{
			GeoNameID:      record.City.GeoNameID,
			City:           record.City.Names["en"],
			IsoCode:        record.Country.IsoCode,
			AccuracyRadius: record.Location.AccuracyRadius,
			Country:        record.Country.Names["en"],
		})
	}
}
