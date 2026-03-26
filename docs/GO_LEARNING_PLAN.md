# Go 语言学习计划

## 目录

1. [环境搭建](#1-环境搭建)
2. [基础语法](#2-基础语法)
3. [核心特性](#3-核心特性)
4. [标准库](#4-标准库)
5. [进阶主题](#5-进阶主题)
6. [实战项目](#6-实战项目)
7. [学习资源](#7-学习资源)

---

## 1. 环境搭建

### 1.1 安装 Go

```bash
# macOS (使用 Homebrew)
brew install go

# 验证安装
go version
```

### 1.2 配置环境变量

```bash
# 在 ~/.zshrc 或 ~/.bash_profile 中添加
export GOPATH=$HOME/go
export PATH=$PATH:$GOPATH/bin
```

### 1.3 IDE/编辑器推荐

- **VS Code** + Go 插件（推荐初学者）
- **GoLand**（JetBrains 出品，功能强大）
- **Cursor**（你正在使用的 AI 编辑器）

### 1.4 第一个程序

```go
// main.go
package main

import "fmt"

func main() {
    fmt.Println("Hello, Go!")
}
```

```bash
# 运行
go run main.go

# 编译
go build main.go
```

---

## 2. 基础语法

### 2.1 变量与常量

```go
// 变量声明
var name string = "Go"
var age int = 10
var isAwesome = true  // 类型推断

// 短声明（函数内部）
count := 42

// 常量
const Pi = 3.14159
const (
    StatusOK = 200
    StatusNotFound = 404
)
```

### 2.2 基本数据类型

| 类型 | 说明 | 示例 |
|------|------|------|
| `bool` | 布尔值 | `true`, `false` |
| `string` | 字符串 | `"Hello"` |
| `int`, `int8`, `int16`, `int32`, `int64` | 有符号整数 | `42` |
| `uint`, `uint8`, `uint16`, `uint32`, `uint64` | 无符号整数 | `42` |
| `float32`, `float64` | 浮点数 | `3.14` |
| `complex64`, `complex128` | 复数 | `1+2i` |
| `byte` | `uint8` 的别名 | |
| `rune` | `int32` 的别名，表示 Unicode 码点 | |

### 2.3 控制流

```go
// if-else
if x > 0 {
    fmt.Println("正数")
} else if x < 0 {
    fmt.Println("负数")
} else {
    fmt.Println("零")
}

// if 带初始化语句
if err := doSomething(); err != nil {
    fmt.Println("出错了:", err)
}

// for 循环（Go 只有 for，没有 while）
for i := 0; i < 10; i++ {
    fmt.Println(i)
}

// while 风格
for count > 0 {
    count--
}

// 无限循环
for {
    // ...
    break
}

// range 遍历
for index, value := range slice {
    fmt.Printf("索引: %d, 值: %v\n", index, value)
}

// switch
switch day {
case "Monday":
    fmt.Println("周一")
case "Tuesday", "Wednesday":
    fmt.Println("周二或周三")
default:
    fmt.Println("其他")
}
```

### 2.4 函数

```go
// 基本函数
func add(a, b int) int {
    return a + b
}

// 多返回值
func divide(a, b float64) (float64, error) {
    if b == 0 {
        return 0, errors.New("除数不能为零")
    }
    return a / b, nil
}

// 命名返回值
func split(sum int) (x, y int) {
    x = sum * 4 / 9
    y = sum - x
    return // 裸返回
}

// 可变参数
func sum(nums ...int) int {
    total := 0
    for _, n := range nums {
        total += n
    }
    return total
}

// 匿名函数和闭包
func counter() func() int {
    count := 0
    return func() int {
        count++
        return count
    }
}
```

### 2.5 复合数据类型

```go
// 数组（固定长度）
var arr [5]int
arr := [3]int{1, 2, 3}
arr := [...]int{1, 2, 3, 4, 5} // 编译器推断长度

// 切片（动态数组）
slice := []int{1, 2, 3}
slice = append(slice, 4, 5)
subSlice := slice[1:3] // 切片操作

// make 创建切片
slice := make([]int, 5)     // 长度为5
slice := make([]int, 0, 10) // 长度0，容量10

// Map
m := make(map[string]int)
m["one"] = 1
m["two"] = 2

// Map 字面量
m := map[string]int{
    "one": 1,
    "two": 2,
}

// 检查 key 是否存在
value, exists := m["three"]
if exists {
    fmt.Println(value)
}

// 删除 key
delete(m, "one")
```

---

## 3. 核心特性

### 3.1 结构体 (Struct)

```go
// 定义结构体
type Person struct {
    Name string
    Age  int
    Email string
}

// 创建实例
p1 := Person{Name: "Alice", Age: 30}
p2 := Person{"Bob", 25, "bob@example.com"}
p3 := new(Person) // 返回指针

// 访问字段
fmt.Println(p1.Name)

// 结构体嵌入（组合）
type Employee struct {
    Person  // 匿名嵌入
    Company string
    Salary  float64
}

emp := Employee{
    Person:  Person{Name: "Charlie", Age: 35},
    Company: "Tech Corp",
    Salary:  100000,
}
fmt.Println(emp.Name) // 可直接访问嵌入字段
```

### 3.2 方法 (Methods)

```go
// 值接收者
func (p Person) Greet() string {
    return "Hello, I'm " + p.Name
}

// 指针接收者（可修改结构体）
func (p *Person) SetAge(age int) {
    p.Age = age
}

// 使用方法
p := Person{Name: "David", Age: 28}
fmt.Println(p.Greet())
p.SetAge(29)
```

### 3.3 接口 (Interface)

```go
// 定义接口
type Speaker interface {
    Speak() string
}

type Writer interface {
    Write(data []byte) (int, error)
}

// 组合接口
type ReadWriter interface {
    Reader
    Writer
}

// 实现接口（隐式）
type Dog struct {
    Name string
}

func (d Dog) Speak() string {
    return d.Name + " says: Woof!"
}

// 使用接口
func MakeSound(s Speaker) {
    fmt.Println(s.Speak())
}

dog := Dog{Name: "Buddy"}
MakeSound(dog) // Dog 实现了 Speaker 接口

// 空接口
var anything interface{}
anything = 42
anything = "hello"
anything = Dog{Name: "Rex"}

// 类型断言
str, ok := anything.(string)
if ok {
    fmt.Println("是字符串:", str)
}

// 类型 switch
switch v := anything.(type) {
case int:
    fmt.Println("整数:", v)
case string:
    fmt.Println("字符串:", v)
default:
    fmt.Printf("未知类型: %T\n", v)
}
```

### 3.4 错误处理

```go
import "errors"

// 自定义错误
var ErrNotFound = errors.New("not found")

// 使用 fmt.Errorf
func findUser(id int) (*User, error) {
    if id <= 0 {
        return nil, fmt.Errorf("invalid user id: %d", id)
    }
    // ...
}

// 错误包装 (Go 1.13+)
if err != nil {
    return fmt.Errorf("failed to process: %w", err)
}

// 错误检查
if errors.Is(err, ErrNotFound) {
    // 处理特定错误
}

// 自定义错误类型
type ValidationError struct {
    Field   string
    Message string
}

func (e *ValidationError) Error() string {
    return fmt.Sprintf("%s: %s", e.Field, e.Message)
}
```

### 3.5 并发编程

#### Goroutine

```go
// 启动 goroutine
go func() {
    fmt.Println("在新的 goroutine 中执行")
}()

// 带参数
go func(msg string) {
    fmt.Println(msg)
}("Hello")
```

#### Channel

```go
// 创建 channel
ch := make(chan int)        // 无缓冲
ch := make(chan int, 10)    // 有缓冲，容量10

// 发送和接收
ch <- 42        // 发送
value := <-ch   // 接收

// 关闭 channel
close(ch)

// 检查是否关闭
value, ok := <-ch
if !ok {
    fmt.Println("channel 已关闭")
}

// range 遍历 channel
for v := range ch {
    fmt.Println(v)
}

// select 多路复用
select {
case msg := <-ch1:
    fmt.Println("收到 ch1:", msg)
case msg := <-ch2:
    fmt.Println("收到 ch2:", msg)
case <-time.After(time.Second):
    fmt.Println("超时")
default:
    fmt.Println("没有数据")
}
```

#### 常用并发模式

```go
// Worker Pool
func worker(id int, jobs <-chan int, results chan<- int) {
    for j := range jobs {
        results <- j * 2
    }
}

func main() {
    jobs := make(chan int, 100)
    results := make(chan int, 100)

    // 启动 3 个 worker
    for w := 1; w <= 3; w++ {
        go worker(w, jobs, results)
    }

    // 发送任务
    for j := 1; j <= 9; j++ {
        jobs <- j
    }
    close(jobs)

    // 收集结果
    for a := 1; a <= 9; a++ {
        <-results
    }
}

// 使用 sync.WaitGroup
var wg sync.WaitGroup

for i := 0; i < 5; i++ {
    wg.Add(1)
    go func(n int) {
        defer wg.Done()
        // 做一些事情
    }(i)
}

wg.Wait() // 等待所有 goroutine 完成

// 使用 sync.Mutex
var mu sync.Mutex
var count int

func increment() {
    mu.Lock()
    defer mu.Unlock()
    count++
}
```

---

## 4. 标准库

### 4.1 常用包速览

| 包 | 用途 |
|-----|------|
| `fmt` | 格式化 I/O |
| `strings` | 字符串操作 |
| `strconv` | 字符串转换 |
| `os` | 操作系统功能 |
| `io` | I/O 原语 |
| `bufio` | 缓冲 I/O |
| `net/http` | HTTP 客户端和服务器 |
| `encoding/json` | JSON 编解码 |
| `time` | 时间处理 |
| `sync` | 同步原语 |
| `context` | 上下文控制 |
| `testing` | 测试框架 |
| `log` | 日志 |
| `regexp` | 正则表达式 |

### 4.2 文件操作

```go
// 读取整个文件
data, err := os.ReadFile("file.txt")

// 写入文件
err := os.WriteFile("file.txt", []byte("Hello"), 0644)

// 逐行读取
file, _ := os.Open("file.txt")
defer file.Close()

scanner := bufio.NewScanner(file)
for scanner.Scan() {
    fmt.Println(scanner.Text())
}
```

### 4.3 JSON 处理

```go
type User struct {
    Name  string `json:"name"`
    Email string `json:"email"`
    Age   int    `json:"age,omitempty"`
}

// 编码
user := User{Name: "Alice", Email: "alice@example.com"}
data, _ := json.Marshal(user)
fmt.Println(string(data))

// 解码
var user2 User
json.Unmarshal(data, &user2)

// 格式化输出
data, _ := json.MarshalIndent(user, "", "  ")
```

### 4.4 HTTP 服务

```go
package main

import (
    "encoding/json"
    "net/http"
)

func main() {
    http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        w.Write([]byte("Hello, World!"))
    })

    http.HandleFunc("/api/users", func(w http.ResponseWriter, r *http.Request) {
        w.Header().Set("Content-Type", "application/json")
        json.NewEncoder(w).Encode(map[string]string{
            "message": "success",
        })
    })

    http.ListenAndServe(":8080", nil)
}
```

### 4.5 HTTP 客户端

```go
// GET 请求
resp, err := http.Get("https://api.example.com/data")
if err != nil {
    log.Fatal(err)
}
defer resp.Body.Close()

body, _ := io.ReadAll(resp.Body)

// POST 请求
data := map[string]string{"name": "Alice"}
jsonData, _ := json.Marshal(data)

resp, err := http.Post(
    "https://api.example.com/users",
    "application/json",
    bytes.NewBuffer(jsonData),
)
```

---

## 5. 进阶主题

### 5.1 模块管理 (Go Modules)

```bash
# 初始化模块
go mod init github.com/username/project

# 添加依赖
go get github.com/gin-gonic/gin

# 整理依赖
go mod tidy

# 查看依赖
go list -m all
```

### 5.2 测试

```go
// math_test.go
package math

import "testing"

func TestAdd(t *testing.T) {
    result := Add(2, 3)
    if result != 5 {
        t.Errorf("Add(2, 3) = %d; want 5", result)
    }
}

// 表驱动测试
func TestAddTable(t *testing.T) {
    tests := []struct {
        name     string
        a, b     int
        expected int
    }{
        {"positive", 2, 3, 5},
        {"negative", -1, -1, -2},
        {"zero", 0, 0, 0},
    }

    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            result := Add(tt.a, tt.b)
            if result != tt.expected {
                t.Errorf("got %d, want %d", result, tt.expected)
            }
        })
    }
}

// 基准测试
func BenchmarkAdd(b *testing.B) {
    for i := 0; i < b.N; i++ {
        Add(2, 3)
    }
}
```

```bash
# 运行测试
go test ./...

# 带覆盖率
go test -cover ./...

# 运行基准测试
go test -bench=. ./...
```

### 5.3 Context

```go
import "context"

// 创建带超时的 context
ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
defer cancel()

// 创建可取消的 context
ctx, cancel := context.WithCancel(context.Background())

// 在 goroutine 中检查取消
go func(ctx context.Context) {
    select {
    case <-ctx.Done():
        fmt.Println("被取消了:", ctx.Err())
        return
    case <-time.After(10 * time.Second):
        fmt.Println("完成")
    }
}(ctx)

// HTTP 请求使用 context
req, _ := http.NewRequestWithContext(ctx, "GET", url, nil)
resp, err := client.Do(req)
```

### 5.4 泛型 (Go 1.18+)

```go
// 泛型函数
func Min[T constraints.Ordered](a, b T) T {
    if a < b {
        return a
    }
    return b
}

// 泛型类型
type Stack[T any] struct {
    items []T
}

func (s *Stack[T]) Push(item T) {
    s.items = append(s.items, item)
}

func (s *Stack[T]) Pop() (T, bool) {
    if len(s.items) == 0 {
        var zero T
        return zero, false
    }
    item := s.items[len(s.items)-1]
    s.items = s.items[:len(s.items)-1]
    return item, true
}

// 使用
stack := Stack[int]{}
stack.Push(1)
stack.Push(2)
```

### 5.5 反射 (Reflection)

```go
import "reflect"

// 获取类型信息
t := reflect.TypeOf(someValue)
fmt.Println(t.Name(), t.Kind())

// 获取值信息
v := reflect.ValueOf(someValue)
fmt.Println(v.Interface())

// 遍历结构体字段
t := reflect.TypeOf(person)
for i := 0; i < t.NumField(); i++ {
    field := t.Field(i)
    fmt.Printf("%s: %s\n", field.Name, field.Type)
}
```

---

## 6. 实战项目

### 6.1 入门项目

1. **命令行工具**
   - TODO 列表管理器
   - 文件批量重命名工具
   - 简单的爬虫

2. **Web API**
   - RESTful API 服务
   - URL 短链接服务
   - 简单的博客后端

### 6.2 中级项目

1. **网络编程**
   - 聊天服务器
   - 代理服务器
   - WebSocket 实时应用

2. **数据处理**
   - 日志分析工具
   - CSV/Excel 处理器
   - 数据管道

### 6.3 高级项目

1. **分布式系统**
   - 分布式缓存
   - 消息队列
   - 任务调度器

2. **云原生**
   - Kubernetes Operator
   - CLI 工具（类似 kubectl）
   - 微服务架构

### 6.4 推荐的项目练手

根据你的项目名 `nsfw-go`，这里有一些相关项目建议：

1. **图像分类服务** - 使用 Go 调用机器学习模型 API
2. **内容审核 API** - 构建 RESTful API 处理内容审核
3. **媒体处理服务** - 图片/视频处理管道

---

## 7. 学习资源

### 7.1 官方资源

- [Go 官网](https://go.dev/)
- [A Tour of Go](https://go.dev/tour/) - 官方互动教程
- [Effective Go](https://go.dev/doc/effective_go) - 最佳实践
- [Go by Example](https://gobyexample.com/) - 实例学习

### 7.2 书籍推荐

| 书名 | 难度 | 说明 |
|------|------|------|
| 《Go 程序设计语言》 | ⭐⭐ | Go 圣经，必读 |
| 《Go 语言实战》 | ⭐⭐ | 实战导向 |
| 《Go Web 编程》 | ⭐⭐ | Web 开发 |
| 《Go 语言高级编程》 | ⭐⭐⭐ | 深入理解 |
| 《Go 语言并发之道》 | ⭐⭐⭐ | 并发专精 |

### 7.3 常用框架/库

| 类别 | 库 | 说明 |
|------|-----|------|
| Web 框架 | Gin, Echo, Fiber | 高性能 HTTP 框架 |
| ORM | GORM, Ent | 数据库 ORM |
| CLI | Cobra, urfave/cli | 命令行工具框架 |
| 配置 | Viper | 配置管理 |
| 日志 | Zap, Logrus | 结构化日志 |
| 测试 | Testify, GoMock | 测试辅助 |
| 验证 | Validator | 数据验证 |

### 7.4 学习路线总结

```
第一阶段：基础 (1-2周)
├── 环境搭建
├── 基础语法
├── 数据类型
└── 控制流、函数

第二阶段：核心 (2-3周)
├── 结构体和方法
├── 接口
├── 错误处理
└── 包和模块

第三阶段：并发 (2周)
├── Goroutine
├── Channel
├── sync 包
└── 并发模式

第四阶段：标准库 (2周)
├── 文件 I/O
├── 网络编程
├── JSON 处理
└── HTTP 服务

第五阶段：进阶 (持续)
├── 测试
├── 性能优化
├── 泛型
└── 设计模式

第六阶段：实战 (持续)
├── 完成项目
├── 阅读开源代码
└── 贡献社区
```

---

## 快速参考

### 常用命令

```bash
go run main.go       # 运行
go build             # 编译
go test ./...        # 测试
go mod init <name>   # 初始化模块
go mod tidy          # 整理依赖
go fmt ./...         # 格式化代码
go vet ./...         # 静态检查
go doc <package>     # 查看文档
```

### 编码规范

1. 使用 `gofmt` 格式化代码
2. 导出的名称首字母大写
3. 简短的变量名（i, n, err 等）
4. 接口名通常以 -er 结尾（Reader, Writer）
5. 错误变量以 Err 开头
6. 先处理错误，再处理正常逻辑

---

祝你学习顺利！🚀
