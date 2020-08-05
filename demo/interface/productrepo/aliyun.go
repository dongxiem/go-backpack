package productrepo

import "fmt"

type aliCloudProductRepo struct {
}

func (m aliCloudProductRepo) StoreProduct(name string, id int) {
	fmt.Println("aliCloudProductRepo: mocking the StoreProduct func")
	// In a real world project you would query an ali Cloud database here.
}

func (m aliCloudProductRepo) FindProductByID(id int) {
	fmt.Println("aliCloudProductRepo: mocking the FindProductByID func")
	// In a real world project you would query an ali Cloud database here.
}
