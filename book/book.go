package book

import (
	"github.com/gofiber/fiber/v2"
	"strconv"
)

// Book struct to hold book data
type Book struct {
	ID     int    `json:"id"`
	Title  string `json:"title"`
	Author string `json:"author"`
}

// Initialize in-memory data
var books = []Book{
	{ID: 1, Title: "1984", Author: "George Orwell"},
	{ID: 2, Title: "The Great Gatsby", Author: "F. Scott Fitzgerald"},
}

// Handler functions

func GetBooks(c *fiber.Ctx) error {
	return c.JSON(books)
}

func GetBook(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.SendStatus(fiber.StatusBadRequest)
	}

	for _, book := range books {
		if book.ID == id {
			return c.JSON(book)
		}
	}

	return c.SendStatus(fiber.StatusNotFound)
}

func CreateBook(c *fiber.Ctx) error {
	book := new(Book)

	if err := c.BodyParser(book); err != nil {
		return c.Status(fiber.StatusBadRequest).SendString(err.Error())
	}

	book.ID = len(books) + 1
	books = append(books, *book)

	return c.JSON(book)
}

func UpdateBook(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.SendStatus(fiber.StatusBadRequest)
	}

	updatedBook := new(Book)
	if err := c.BodyParser(updatedBook); err != nil {
		return c.Status(fiber.StatusBadRequest).SendString(err.Error())
	}

	for i, book := range books {
		if book.ID == id {
			book.Title = updatedBook.Title
			book.Author = updatedBook.Author
			books[i] = book
			return c.JSON(book)
		}
	}

	return c.SendStatus(fiber.StatusNotFound)
}

func DeleteBook(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.SendStatus(fiber.StatusBadRequest)
	}

	for i, book := range books {
		if book.ID == id {
			books = append(books[:i], books[i+1:]...)
			return c.SendStatus(fiber.StatusNoContent)
		}
	}

	return c.SendStatus(fiber.StatusNotFound)
}

func Test(c *fiber.Ctx) error {
	book := new(Book)
	if err := c.BodyParser(book); err != nil {
		return c.Status(fiber.StatusBadRequest).SendString(err.Error())
	}
	return c.JSON(book)
}