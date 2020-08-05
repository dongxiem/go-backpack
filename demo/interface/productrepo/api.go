package productrepo

// 定义ProductRepository接口，它代表的就是存储库
type ProductRepository interface {
	// StoreProduct()方法用于存储产品信息
	StoreProduct(name string, id int)
	// FindProductByID()方法通过产品ID查找产品信息
	FindProductByID(id int)
}

// 基于环境值返回ProductRepository接口的正确实现对象。
func New(environment string) ProductRepository {
	// 根据传递进来的env选择哪种存储方式
	switch environment {
	case "aliCloud":
		return aliCloudProductRepo{}
	case "local-mysql":
		return mysqlProductRepo{}
	}
	// 返回默认的存储库
	return mockProductRepo{}
}
