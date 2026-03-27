// +build ignore

package main

import (
	"encoding/json"
	"fmt"
)

// ============================================================
// 练习 1：基础结构体 - 参考答案
// ============================================================

type Book struct {
	Title  string
	Author string
	Price  float64
	ISBN   string
}

func (b Book) String() string {
	return fmt.Sprintf("《%s》\n  作者: %s\n  价格: ¥%.2f\n  ISBN: %s",
		b.Title, b.Author, b.Price, b.ISBN)
}

func (b *Book) ApplyDiscount(percent float64) {
	b.Price = b.Price * (1 - percent/100)
}

// ============================================================
// 练习 2：几何图形 - 参考答案
// ============================================================

type Rectangle struct {
	Length float64
	Width  float64
}

func (r Rectangle) Area() float64 {
	return r.Length * r.Width
}

func (r Rectangle) Perimeter() float64 {
	return 2 * (r.Length + r.Width)
}

func (r *Rectangle) Scale(factor float64) {
	r.Length *= factor
	r.Width *= factor
}

// ============================================================
// 练习 3：结构体嵌入 - 参考答案
// ============================================================

type Person struct {
	Name  string
	Age   int
	Phone string
}

func (p Person) Introduce() string {
	return fmt.Sprintf("我是 %s，今年 %d 岁", p.Name, p.Age)
}

type Employee struct {
	Person
	EmployeeID string
	Department string
	Salary     float64
}

func (e Employee) Work() string {
	return fmt.Sprintf("%s 正在 %s 努力工作", e.Name, e.Department)
}

func (e *Employee) GiveRaise(percent float64) {
	e.Salary = e.Salary * (1 + percent/100)
}

// ============================================================
// 练习 4：JSON 序列化 - 参考答案
// ============================================================

type User struct {
	ID       int    `json:"id"`
	Username string `json:"username"`
	Password string `json:"-"`
	Email    string `json:"email,omitempty"`
	IsActive bool   `json:"is_active"`
}

// ============================================================
// 主函数
// ============================================================

func main() {
	fmt.Println("=== Go 结构体练习 - 参考答案 ===\n")

	// 测试练习 1：Book
	fmt.Println("--- 练习 1: Book ---")
	book := Book{
		Title:  "Go 程序设计语言",
		Author: "Alan Donovan",
		Price:  99.00,
		ISBN:   "978-7-111-55842-2",
	}
	fmt.Println(book.String())
	book.ApplyDiscount(20)
	fmt.Printf("折后价格: ¥%.2f\n", book.Price)

	// 测试练习 2：Rectangle
	fmt.Println("\n--- 练习 2: Rectangle ---")
	rect := Rectangle{Length: 10, Width: 5}
	fmt.Printf("原始尺寸: 长=%.2f, 宽=%.2f\n", rect.Length, rect.Width)
	fmt.Printf("面积: %.2f\n", rect.Area())
	fmt.Printf("周长: %.2f\n", rect.Perimeter())
	rect.Scale(2)
	fmt.Printf("放大 2 倍后: 长=%.2f, 宽=%.2f\n", rect.Length, rect.Width)
	fmt.Printf("放大后面积: %.2f\n", rect.Area())

	// 测试练习 3：Employee
	fmt.Println("\n--- 练习 3: Employee ---")
	emp := Employee{
		Person:     Person{Name: "张三", Age: 28, Phone: "13800138000"},
		EmployeeID: "EMP001",
		Department: "技术部",
		Salary:     15000,
	}
	fmt.Println(emp.Introduce())
	fmt.Println(emp.Work())
	fmt.Printf("当前薪资: ¥%.2f\n", emp.Salary)
	emp.GiveRaise(10)
	fmt.Printf("加薪 10%% 后: ¥%.2f\n", emp.Salary)

	// 直接访问嵌入字段的属性
	fmt.Printf("员工姓名: %s (通过 emp.Name 直接访问)\n", emp.Name)
	fmt.Printf("员工年龄: %d (通过 emp.Age 直接访问)\n", emp.Age)

	// 测试练习 4：JSON
	fmt.Println("\n--- 练习 4: User JSON ---")
	user := User{
		ID:       1,
		Username: "alice",
		Password: "secret123",
		Email:    "alice@example.com",
		IsActive: true,
	}
	jsonData, _ := json.MarshalIndent(user, "", "  ")
	fmt.Println("User 1 JSON (注意 Password 被忽略):")
	fmt.Println(string(jsonData))

	// 测试空 Email 时 omitempty
	user2 := User{
		ID:       2,
		Username: "bob",
		Password: "password",
		IsActive: false,
	}
	jsonData2, _ := json.MarshalIndent(user2, "", "  ")
	fmt.Println("\nUser 2 JSON (注意空 Email 被省略):")
	fmt.Println(string(jsonData2))
}

type Teacher {
	name string
	age int
	subject string
	salary float64
}