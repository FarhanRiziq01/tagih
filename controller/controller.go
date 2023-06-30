package controller

import (
	"github.com/FarhanRiziq01/inibackendriziq"
	"github.com/FarhanRiziq01/tagih/config"
	"github.com/aiteung/musik"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/websocket/v2"
	"github.com/whatsauth/whatsauth"
)

var DataPelanggan = "pelanggan"
var DataTagihan = "tagihan"
var DataPembayaran = "pembayaran"
var DataProduk = "produk"
var DataAbout = "about"

func WsWhatsAuthQR(c *websocket.Conn) {
	whatsauth.RunSocket(c, config.PublicKey, config.Usertables[:], config.Ulbimariaconn)
}

func PostWhatsAuthRequest(c *fiber.Ctx) error {
	if string(c.Request().Host()) == config.Internalhost {
		var req whatsauth.WhatsauthRequest
		err := c.BodyParser(&req)
		if err != nil {
			return err
		}
		ntfbtn := whatsauth.RunModuleLegacy(req, config.PrivateKey, config.Usertables[:], config.Ulbimariaconn)
		return c.JSON(ntfbtn)
	} else {
		var ws whatsauth.WhatsauthStatus
		ws.Status = string(c.Request().Host())
		return c.JSON(ws)
	}

}

func Homepage(c *fiber.Ctx) error {
	ipaddr := musik.GetIPaddress()
	return c.JSON(ipaddr)
}

func GetPelanggan(c *fiber.Ctx) error {
	getstatus := inibackendriziq.GetDataPelanggan("Farhan Riziq")
	return c.JSON(getstatus)
}

func GetTagihan(c *fiber.Ctx) error {
	getstatus := inibackendriziq.GetDataTagihan("08.01.2021")
	return c.JSON(getstatus)
}

func GetPembayaran(c *fiber.Ctx) error {
	getstatus := inibackendriziq.GetDataPembayaran("05.01.2021")
	return c.JSON(getstatus)
}

func GetProduk(c *fiber.Ctx) error {
	getstatus := inibackendriziq.GetDataProduk("Spotify")
	return c.JSON(getstatus)
}
func GetAbout(c *fiber.Ctx) error {
	getstatus := inibackendriziq.GetDataAbout("tes")
	return c.JSON(getstatus)
}
