package productrepo

import "fmt"

type mysqlProductRepo struct {
}

func (m mysqlProductRepo) StoreProduct(name string, id int) {
	fmt.Println("mysqlProductRepo: mocking the StoreProduct func")
	// In a real world project you would query a MySQL database here.
}

func (m mysqlProductRepo) FindProductByID(id int) {
	fmt.Println("mysqlProductRepo: mocking the FindProductByID func")
	// In a real world project you would query a MySQL database here.
}
