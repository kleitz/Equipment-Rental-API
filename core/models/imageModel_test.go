package models

import (
	"testing"
	. "github.com/franela/goblin"
	"github.com/remony/Equipment-Rental-API/core/router"
	"github.com/remony/Equipment-Rental-API/core/config"
)

const CONF_FILE = "./../../config.json"

func TestImageModel(t *testing.T) {
	g := Goblin(t)
	router := router.API{Context: config.Connection(config.LoadConfig(CONF_FILE, true).Production.DbUrl)}

	g.Describe("Image Available", func() {
		// TODO when database is complete change to check is true
		g.It("Should be true", func() {
			g.Assert(IsImageAvailable(router, "image.jpg")).IsFalse()
		})
		g.It("Should be false", func() {
			g.Assert(IsImageAvailable(router, "image.jpg")).IsFalse()
		})
	})
}
