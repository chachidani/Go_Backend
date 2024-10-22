package main
import(
	"fmt"
	"Library_Management_System/controllers"
)

func main(){
	libraryController := controllers.NewLibraryController()
	
	for{
		fmt.Println("1. Add Book")
		fmt.Println("2. Remove Book")
		fmt.Println("3. Borrow Book")
		fmt.Println("4. Return Book")
		fmt.Println("5. List Available Books")
		fmt.Println("6. List Borrowed Books")
		fmt.Println("7. Addmember")
		fmt.Println("8. Exit")
		fmt.Println("Enter your choice:")
		
		var choice int
		fmt.Scanln(&choice)
		
		switch choice{
			case 1:
				libraryController.AddBook()
			case 2:
				libraryController.RemoveBook()
			case 3:
				libraryController.BorrowBook()
			case 4:
				libraryController.ReturnBook()
			case 5:
				libraryController.ListAvailableBooks()
			case 6:
				libraryController.ListBorrowedBooks()
			case 7:
				libraryController.AddMember()
			case 8:
				return
			default:
				fmt.Println("Invalid choice")
		}
	}
}