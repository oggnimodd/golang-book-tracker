package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/oggnimodd/golang-clerk-practice/context"
	"github.com/oggnimodd/golang-clerk-practice/feature/book"
	note "github.com/oggnimodd/golang-clerk-practice/feature/notes"
	"github.com/oggnimodd/golang-clerk-practice/feature/review"
	"github.com/oggnimodd/golang-clerk-practice/feature/shelf"
	"github.com/oggnimodd/golang-clerk-practice/middleware"
)

func SetupRoutes(router *gin.Engine) {

	v1 := router.Group("/api/v1")

	// Book
	bookRepo := book.NewBookRepository(context.DB)
	bookService := book.NewBookService(bookRepo)
	bookHandler := book.NewBookHandler(bookService)

	book := v1.Group("/books/")

	book.Use(middleware.AuthMiddleware())

	book.GET("", bookHandler.GetAllBooks)
	book.POST("", bookHandler.CreateBook)
	book.GET("/:id", bookHandler.GetBookById)
	book.GET("/shelf/:shelfId", bookHandler.GetBooksByShelfId)
	book.PUT("/:id", bookHandler.UpdateBook)
	book.DELETE("/:id", bookHandler.DeleteBook)

	// Shelf
	shelfRepo := shelf.NewRepositoryShelf(context.DB)
	shelfService := shelf.NewServiceShelf(shelfRepo)
	shelfHandler := shelf.NewHandlerShelf(shelfService)

	shelf := v1.Group("/shelves")

	shelf.Use(middleware.AuthMiddleware())

	shelf.GET("", shelfHandler.GetAllShelves)
	shelf.POST("", shelfHandler.CreateShelf)
	shelf.GET("/:id", shelfHandler.GetShelfById)
	shelf.PUT("/:id", shelfHandler.UpdateShelf)
	shelf.DELETE("/:id", shelfHandler.DeleteShelf)
	shelf.GET("/initial", shelfHandler.GenerateInitialShelves)

	// Notes
	notesRepo := note.NewNoteRepository(context.DB)
	notesService := note.NewNoteService(notesRepo)
	notesHandler := note.NewNoteHandler(notesService)

	notes := v1.Group("/notes")

	notes.Use(middleware.AuthMiddleware())

	notes.GET("", notesHandler.GetAllNotes)
	notes.GET("/:id", notesHandler.GetNoteById)
	notes.POST("", notesHandler.CreateNote)
	notes.PUT("/:id", notesHandler.UpdateNote)
	notes.DELETE("/:id", notesHandler.DeleteNote)

	// Review
	reviewRepo := review.NewReviewRepository(context.DB)
	reviewService := review.NewReviewService(reviewRepo)
	reviewHandler := review.NewReviewHandler(reviewService)

	review := v1.Group("/reviews")

	review.Use(middleware.AuthMiddleware())

	review.GET("/:book_id", reviewHandler.GetReviewByBookId)
	review.POST("/:book_id", reviewHandler.CreateReview)
	review.PUT("/:id", reviewHandler.UpdateReview)
	review.DELETE("/:id", reviewHandler.DeleteReview)
}
