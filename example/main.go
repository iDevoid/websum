package main

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/iDevoid/websum"
)

func main() {
	server := fiber.New()
	server.Static("/", "./static")

	server.Add(http.MethodPost, "/summarize", handler)

	err := server.Listen(":9000")
	if err != nil {
		panic(err)
	}
}

func handler(ctx *fiber.Ctx) error {
	url := string(ctx.Context().FormValue("url"))
	if url == "" {
		ctx.Status(http.StatusBadRequest)
		return errors.New(http.StatusText(http.StatusBadRequest))
	}

	data, err := websum.SummarizeWeb(url)
	if err != nil {
		return err
	}

	raw, err := json.Marshal(data)
	ctx.Response().SetBody(raw)
	return err
}
