package controllers

import(
	"fmt"
	"Library_Management_System/models"
	"Library_Management_System/services"

	
)

type LibraryController struct{
	libraryService services.LibraryManager
	
}

func NewLibraryController() *LibraryController{
	return &LibraryController{
		libraryService: services.NewLibrary(),
	}
}


func (lc *LibraryController) AddMember(){
	var id int
	var name string
	fmt.Println("Enter Member ID:")
	fmt.Scanln(&id)
	fmt.Println("Enter Member Name:")
	fmt.Scanln(&name)

	member := models.Member{ID: id, Name: name}
	lc.libraryService.AddMember(member)
	fmt.Println("Member added successfully")
}

func (lc *LibraryController) AddBook(){
	var id int
	var title  , author string
	fmt.Println("Enter Book ID:")
	fmt.Scanln(&id)
	fmt.Println("Enter Book Title:")
	fmt.Scanln(&title)
	fmt.Println("Enter Book Author:")
	fmt.Scanln(&author)

	book := models.Book{ID: id, Title: title, Author: author , Status: "Available"}
	lc.libraryService.AddBook(book)
	fmt.Println("Book added successfully")
}



func (lc *LibraryController) RemoveBook(){
	var id int
	fmt.Println("Enter Book ID:")
	fmt.Scanln(&id)

	lc.libraryService.RemoveBook(id)
	fmt.Println("Book removed successfully")
}

func (lc *LibraryController) BorrowBook(){
	var bookID , memberID int
	fmt.Println("Enter Book ID:")
	fmt.Scanln(&bookID)
	fmt.Println("Enter Member ID:")
	fmt.Scanln(&memberID)

	err := lc.libraryService.BorrowBooks(bookID, memberID)
	if err != nil{
		fmt.Println(err)
	}else{
		fmt.Println("Book borrowed successfully")
	}
}


func (lc *LibraryController) ReturnBook(){
	var bookID , memberID int
	fmt.Println("Enter Book ID:")
	fmt.Scanln(&bookID)
	fmt.Println("Enter Member ID:")
	fmt.Scanln(&memberID)

	err := lc.libraryService.ReturnBook(bookID, memberID)
	if err != nil{
		fmt.Println(err)
	}else{
		fmt.Println("Book returned successfully")
	}
}

func (lc *LibraryController) ListAvailableBooks(){
	books := lc.libraryService.ListAvailableBooks()
	for _ , book := range books{
		fmt.Printf("ID: %d\nTitle: %s\nAuthor: %s\n\n" , book.ID, book.Title, book.Author)
	}
}


func (lc *LibraryController) ListBorrowedBooks(){
	var memberID int
	fmt.Println("Enter Member ID:")
	fmt.Scanln(&memberID)

	books := lc.libraryService.ListBorrowedBooks(memberID)
	if len(books) == 0{
		fmt.Println("No books borrowed by this member")
	}else{
        fmt.Println("Books borrowed by this member:")
		for _ , book := range books{
			fmt.Printf("ID: %d\nTitle: %s\nAuthor: %s\n\n" , book.ID, book.Title, book.Author)
		}
	}
}

