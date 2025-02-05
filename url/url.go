package url

import (
	"github.com/FarhanRiziq01/tagih/controller"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/websocket/v2"
)

func Web(page *fiber.App) {
	page.Post("/api/whatsauth/request", controller.PostWhatsAuthRequest)  //API from user whatsapp message from iteung gowa
	page.Get("/ws/whatsauth/qr", websocket.New(controller.WsWhatsAuthQR)) //websocket whatsauth
	page.Get("/", controller.Homepage)
	page.Get("/pelanggan", controller.GetPelanggan)   //API from user whatsapp message from iteung gowa
	page.Get("/tagihan", controller.GetTagihan)       //API from user whatsapp message from iteung gowa
	page.Get("/pembayaran", controller.GetPembayaran) //API from user whatsapp message from iteung gowa
	page.Get("/produk", controller.GetProduk)         //API from user whatsapp message from iteung gowa
	page.Get("/about", controller.GetAbout)           //API from user whatsapp message from iteung gowa
}
