package services

import (
	"context"
	"fmt"
	"io"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/google/generative-ai-go/genai"
	"google.golang.org/api/option"
)

func GeminiHandler(c *fiber.Ctx) error {
	content := c.FormValue("content")
	if content == "" {
		chatWithImage(c)
	} else {
		chat(c)
	}
	return nil
}

func chat(c *fiber.Ctx) error {
	system_instruction := c.FormValue("system_instruction")
	token := c.FormValue("token")
	content := c.FormValue("content")

	ctx := context.Background()
	client, err := genai.NewClient(ctx, option.WithAPIKey(token))
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}
	defer client.Close()

	model := client.GenerativeModel("gemini-2.0-flash-exp")
	model.SystemInstruction = genai.NewUserContent(genai.Text(system_instruction))
	cs := model.StartChat()

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	resp, err := cs.SendMessage(ctx, genai.Text(content))
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.JSON(fiber.Map{
		"content": resp.Candidates[0].Content.Parts[0],
	})
}

func chatWithImage(c *fiber.Ctx) error {
	file, err := c.FormFile("file")
	system_instruction := c.FormValue("system_instruction")
	token := c.FormValue("token")

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	fmt.Println(file.Filename)

	ctx := context.Background()
	client, err := genai.NewClient(ctx, option.WithAPIKey(token))
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}
	defer client.Close()

	model := client.GenerativeModel("gemini-2.0-flash-exp")
	model.SystemInstruction = genai.NewUserContent(genai.Text(system_instruction))
	cs := model.StartChat()

	f, err := file.Open()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	defer f.Close()
	// write the file f
	fileBytes, err := io.ReadAll(f)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	resp, err := cs.SendMessage(ctx, genai.Text(file.Filename), genai.ImageData("jpeg", fileBytes))
	if err != nil {
		log.Fatalln(err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.JSON(fiber.Map{
		"content": resp.Candidates[0].Content.Parts[0],
	})
}
