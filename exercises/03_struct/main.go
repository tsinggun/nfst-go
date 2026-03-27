package main

import (
	"encoding/json"
	"fmt"
)

// ============================================================
// 练习 1：基础结构体
// ============================================================

// TODO: 定义 Book 结构体
// 字段：Title (书名), Author (作者), Price (价格), ISBN
type Book struct {
	// 在这里添加字段
}

// TODO: 为 Book 实现 String() 方法，返回书籍的详细信息
// func (b Book) String() string {
//     ...
// }

// TODO: 为 Book 实现 ApplyDiscount(percent float64) 方法，应用折扣
// 使用指针接收者，因为需要修改价格
// func (b *Book) ApplyDiscount(percent float64) {
//     ...
// }

// ============================================================
// 练习 2：几何图形
// ============================================================

// TODO: 定义 Rectangle 结构体（长 Length, 宽 Width）
type Rectangle struct {
	// 在这里添加字段
}

// TODO: 实现 Area() 方法 - 返回面积
// func (r Rectangle) Area() float64 {
//     ...
// }

// TODO: 实现 Perimeter() 方法 - 返回周长
// func (r Rectangle) Perimeter() float64 {
//     ...
// }

// TODO: 实现 Scale(factor float64) 方法 - 按比例缩放（需要指针接收者）
// func (r *Rectangle) Scale(factor float64) {
//     ...
// }

// ============================================================
// 练习 3：结构体嵌入
// ============================================================

// 基础 Person 结构体（已完成）
type Person struct {
	Name  string
	Age   int
	Phone string
}

// Person 的方法
func (p Person) Introduce() string {
	return fmt.Sprintf("我是 %s，今年 %d 岁", p.Name, p.Age)
}

// TODO: 定义 Employee 结构体
// - 嵌入 Person（匿名嵌入）
// - 添加字段：EmployeeID (员工号), Department (部门), Salary (薪资)
type Employee struct {
	// 在这里添加字段
}

// TODO: 为 Employee 实现 Work() 方法
// 返回："{姓名} 正在 {部门} 努力工作"
// func (e Employee) Work() string {
//     ...
// }

// TODO: 为 Employee 实现 GiveRaise(percent float64) 方法
// 给员工加薪，percent 是加薪百分比（如 10 表示加薪 10%）
// func (e *Employee) GiveRaise(percent float64) {
//     ...
// }

// ============================================================
// 练习 4：JSON 序列化
// ============================================================

// TODO: 定义 User 结构体，使用 JSON 标签
// 字段和标签：
// - ID int `json:"id"`
// - Username string `json:"username"`
// - Password string `json:"-"` (忽略，不序列化)
// - Email string `json:"email,omitempty"` (空值时省略)
// - IsActive bool `json:"is_active"`
type User struct {
	// 在这里添加字段
}

// ============================================================
// 主函数 - 测试你的实现
// ============================================================

func main() {
	fmt.Println("=== Go 结构体练习 ===\n")

	// 测试练习 1：Book
	fmt.Println("--- 练习 1: Book ---")
	// 取消下面注释来测试
	// book := Book{
	// 	Title:  "Go 程序设计语言",
	// 	Author: "Alan Donovan",
	// 	Price:  99.00,
	// 	ISBN:   "978-7-111-55842-2",
	// }
	// fmt.Println(book.String())
	// book.ApplyDiscount(20) // 打 8 折
	// fmt.Printf("折后价格: %.2f\n", book.Price)

	// 测试练习 2：Rectangle
	fmt.Println("\n--- 练习 2: Rectangle ---")
	// 取消下面注释来测试
	// rect := Rectangle{Length: 10, Width: 5}
	// fmt.Printf("面积: %.2f\n", rect.Area())
	// fmt.Printf("周长: %.2f\n", rect.Perimeter())
	// rect.Scale(2)
	// fmt.Printf("放大 2 倍后: 长=%.2f, 宽=%.2f\n", rect.Length, rect.Width)

	// 测试练习 3：Employee
	fmt.Println("\n--- 练习 3: Employee ---")
	// 取消下面注释来测试
	// emp := Employee{
	// 	Person:     Person{Name: "张三", Age: 28, Phone: "13800138000"},
	// 	EmployeeID: "EMP001",
	// 	Department: "技术部",
	// 	Salary:     15000,
	// }
	// fmt.Println(emp.Introduce())  // 继承自 Person 的方法
	// fmt.Println(emp.Work())
	// fmt.Printf("当前薪资: %.2f\n", emp.Salary)
	// emp.GiveRaise(10)  // 加薪 10%
	// fmt.Printf("加薪后: %.2f\n", emp.Salary)

	// 测试练习 4：JSON
	fmt.Println("\n--- 练习 4: User JSON ---")
	// 取消下面注释来测试
	// user := User{
	// 	ID:       1,
	// 	Username: "alice",
	// 	Password: "secret123",
	// 	Email:    "alice@example.com",
	// 	IsActive: true,
	// }
	// jsonData, _ := json.MarshalIndent(user, "", "  ")
	// fmt.Println(string(jsonData))
	//
	// // 测试空 Email
	// user2 := User{
	// 	ID:       2,
	// 	Username: "bob",
	// 	Password: "password",
	// 	IsActive: false,
	// }
	// jsonData2, _ := json.MarshalIndent(user2, "", "  ")
	// fmt.Println(string(jsonData2))

	fmt.Println("\n=== 完成所有练习后，取消注释运行测试 ===")

	// 防止 unused import 错误
	_ = json.Marshal
}
