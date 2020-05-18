package controllers

import (
	"github.com/Kamva/mgm/v2"
	"github.com/gofiber/fiber"
	"github.com/hl-service/go-api/models"
	"go.mongodb.org/mongo-driver/bson"
)

// IndexArticles - GET /api/articles
func IndexArticles(ctx *fiber.Ctx) {
	collection := mgm.Coll(&models.Article{})
	articles := []models.Article{}

	err := collection.SimpleFind(&articles, bson.D{})

	if err != nil {
		ctx.Status(500).JSON(fiber.Map{
			"success": false,
			"message": err.Error(),
		})
		return
	}

	ctx.JSON(fiber.Map{
		"success": true,
		"data":    articles,
	})
}

// ShowArticle - GET /api/articles/:id
func ShowArticle(ctx *fiber.Ctx) {
	id := ctx.Params("id")

	article := &models.Article{}
	collection := mgm.Coll(article)

	err := collection.FindByID(id, article)

	if err != nil {
		ctx.Status(404).JSON(fiber.Map{
			"success": false,
			"message": "Article not found.",
		})

		return
	}

	ctx.JSON(fiber.Map{
		"success": true,
		"data":    article,
	})
}

// StoreArticle - POST /api/articles
func StoreArticle(ctx *fiber.Ctx) {
	params := new(struct {
		Title       string
		Description string
	})

	ctx.BodyParser(&params)

	if len(params.Title) == 0 || len(params.Description) == 0 {
		ctx.Status(400).JSON(fiber.Map{
			"success": false,
			"message": "Title or description not specified.",
		})
		return
	}

	article := models.CreateArticle(params.Title, params.Description)

	err := mgm.Coll(article).Create(article)

	if err != nil {
		ctx.Status(500).JSON(fiber.Map{
			"success": false,
			"message": err.Error(),
		})
		return
	}

	ctx.JSON(fiber.Map{
		"success": true,
		"data":    article,
	})
}

// UpdateArticle - PUT /api/articles/:id
func UpdateArticle(ctx *fiber.Ctx) {
	id := ctx.Params("id")

	params := new(struct {
		Title       string
		Description string
	})

	ctx.BodyParser(&params)

	if len(params.Title) == 0 || len(params.Description) == 0 {
		ctx.Status(400).JSON(fiber.Map{
			"success": false,
			"message": "Title or description not specified.",
		})
		return
	}

	article := &models.Article{}
	collection := mgm.Coll(article)

	err := collection.FindByID(id, article)

	if err != nil {
		ctx.Status(404).JSON(fiber.Map{
			"success": false,
			"message": "Article not found.",
		})

		return
	}

	article.Title = params.Title
	article.Description = params.Description

	err = collection.Update(article)

	if err != nil {
		ctx.Status(500).JSON(fiber.Map{
			"ok":    false,
			"error": err.Error(),
		})
		return
	}

	ctx.JSON(fiber.Map{
		"success": true,
		"data":    article,
	})
}

// DeleteArticle - DELETE /api/articles/:id
func DeleteArticle(ctx *fiber.Ctx) {
	id := ctx.Params("id")

	article := &models.Article{}
	collection := mgm.Coll(article)

	err := collection.FindByID(id, article)

	if err != nil {
		ctx.Status(404).JSON(fiber.Map{
			"success": false,
			"message": "Article not found.",
		})
		return
	}

	err = collection.Delete(article)
	if err != nil {
		ctx.Status(500).JSON(fiber.Map{
			"success": false,
			"message": err.Error(),
		})
		return
	}

	ctx.JSON(fiber.Map{
		"success": true,
		"data":    article,
	})
}
