package services

import(
	"errors"
	"Library_Management_System/models"
	
)


type Library struct{
	Books map[int]models.Book
	Members map[int]models.Member
}

type LibraryManager interface{
	AddBook(book models.Book) 
	AddMember(member models.Member)
	RemoveBook(bookID int) 
	BorrowBooks(bookID int, memberID int) error
	ReturnBook(bookID int, memberID int) error
	ListAvailableBooks() []models.Book
	ListBorrowedBooks(memberID int) []models.Book

}

func NewLibrary() LibraryManager{
	return &Library{
		Books: make(map[int]models.Book),
		Members: make(map[int]models.Member),
	}
}

func (l *Library) AddMember(member models.Member){
	l.Members[member.ID] = member
	
}

func (l *Library) AddBook(book models.Book){
	l.Books[book.ID] = book
}

func (l *Library) RemoveBook(bookID int){
	delete(l.Books, bookID)
}

func (l *Library) BorrowBooks(bookID int , memberID int)error{
	book , exists := l.Books[bookID]
	if !exists{
		return errors.New("book not found")
	}
	if book.Status == "Borrowed"{
		return errors.New("book already borrowed")
	}

	member , exists := l.Members[memberID]
	if !exists{
		return errors.New("book not found")
	}

	book.Status = "Borrowed"
	l.Books[bookID] = book
	member.BorrowedBooks = append(member.BorrowedBooks , book)
	l.Members[memberID] = member

	return nil

}

func (l *Library) ReturnBook(bookID int, memberID int) error{
	

	member , exists := l.Members[memberID]
	if !exists{
		return errors.New("member not found")
	}

	bookIndex := -1
	for i , book := range member.BorrowedBooks{
		if book.ID == bookID{
			bookIndex = i
			break
		}
	}

	if bookIndex == -1{
		return errors.New("book not found in member's borrowed books")
	}

	member.BorrowedBooks = append(member.BorrowedBooks[:bookIndex], member.BorrowedBooks[bookIndex+1:]...)
	l.Members[memberID] = member

	book := l.Books[bookID]
	book.Status = "Available"
	l.Books[bookID] = book

	return nil
}

func (l *Library) ListAvailableBooks() []models.Book{
	var availableBooks []models.Book
	for _, book := range l.Books{
		if book.Status == "Available"{
			availableBooks = append(availableBooks, book)
		}
	}
	return availableBooks
}

func (l *Library) ListBorrowedBooks(memberID int) []models.Book{

	member , exists := l.Members[memberID]
	if !exists{
		return []models.Book{}
	}

	return member.BorrowedBooks
}