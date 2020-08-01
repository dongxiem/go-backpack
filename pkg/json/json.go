package main

import (
	"encoding/json"
	"fmt"
)

// User 模型
type User struct {
	ID       int    `json:"id"` // 在结构字段上使用标签来自定义JSON密钥
	Username string `json:"username"`
	Email    string `json:"email"`
	Bio      string `json:"bio,omitempty"` // “ omitempty”选项指定如果字段具有空值（零值），则应从编码中省略该字段。
}

func main() {
	// 将JSON数据解组到结构体中
	userData := []byte(`{"id":1, "username":"Bob", "email":"bob@gmail.com"}`)
	var user User
	if err := json.Unmarshal(userData, &user); err != nil {
		panic(err)
	}
	fmt.Println(user) // 你也可以使用: fmt.Printf("%#v", user) %#v 使用Go语法输出变量。

	// 将JSON解组到map [string] interface {}中
	var userMap map[string]interface{}
	if err := json.Unmarshal(userData, &userMap); err != nil {
		panic(err)
	}
	userID := userMap["id"].(float64) // 为了使用解码后的map中的值，我们需要将其转换为适当的类型
	fmt.Println(userID)
	fmt.Println(user)

	// 将元数据结构Struct转换为JSON
	user = User{ID: 1, Username: "John", Email: "johny@foo.bar"}
	userData, err := json.Marshal(user)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(userData))

	// 将JSON数组解组为结构体切片
	usersData := []byte(`[{"id":1, "username":"Bob", "email":"bob@gmail.com"}, {"id":1, "username":"Bob", "email":"bob@gmail.com"}]`)
	var users []User
	if err = json.Unmarshal(usersData, &users); err != nil {
		panic(err)
	}
	fmt.Println(users)

	// 嵌入式对象
	type Err struct {
		Code    int    `json:"error"`
		Message string `json:"message"`
	}

	type AppError struct {
		Error Err `json:"error"`
	}

	// 将JSON解组到结构体Struct中
	errData := []byte(`{"error":{"code":200, "message":"oops, something went wrong"}}`)
	var appErr AppError
	if err := json.Unmarshal(errData, &appErr); err != nil {
		panic(err)
	}
	fmt.Println(appErr)

	// 将元数据结构Struct转换为JSON
	appErr = AppError{
		Error: Err{
			Code:    200,
			Message: "Some error message",
		},
	}
	errData, err = json.Marshal(appErr)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(errData))
}
