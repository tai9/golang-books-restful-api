package service

import (
	"log"

	"github.com/mashingan/smapping"
	"github.com/tai9/golang_jwt/dto"
	"github.com/tai9/golang_jwt/entity"
	"github.com/tai9/golang_jwt/repository"
)

type BookService interface {
	Inser(b dto.BookCreateDTO) entity.Book
	Update(b dto.BookUpdateDTO) entity.Book
	Delete(b entity.Book)
	All() []entity.Book
	FindById(bookID uint64) entity.Book
	IsAllowedToEdit(userID string, bookID uint64) bool
}

type bookService struct {
	bookRepository repository.BookRepository
}

func NewBookService(bookRepository repository.BookRepository) BookService {
	return &bookService{
		bookRepository: bookRepository,
	}
}

func (service *bookService) Inser(b dto.BookCreateDTO) entity.Book {
	book := entity.Book{}
	err := smapping.FillStruct(&book, smapping.MapFields(&b))
	if err != nil {
		log.Fatalf("Failed map %v", err)
	}
	res := service.bookRepository.InserBook(book)
	return res
}
func (service *bookService) Update(b dto.BookUpdateDTO) entity.Book {
	book := entity.Book{}
	err := smapping.FillStruct(&book, smapping.MapFields(&b))
	if err != nil {
		log.Fatalf("Failed map %v", err)
	}
	res := service.bookRepository.UpdateBook(book)
	return res
}
func (service *bookService) Delete(b entity.Book) {
	service.bookRepository.DeleteBook(b)
}
func (service *bookService) All() []entity.Book {
	return service.bookRepository.AllBook()
}
func (service *bookService) FindById(bookID uint64) entity.Book {
	return service.bookRepository.FindBookByID(bookID)
}

func (service *bookService) IsAllowedToEdit(userID string, bookID uint64) bool {
	return true
}
