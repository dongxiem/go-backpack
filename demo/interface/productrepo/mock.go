package productrepo

import "fmt"

type mockProductRepo struct {
}

func (m mockProductRepo) StoreProduct(name string, id int) {
	fmt.Println("mocking the StoreProduct func")
}

func (m mockProductRepo) FindProductByID(id int) {
	fmt.Println("mocking the FindProductByID func")
}
