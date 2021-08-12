

# Từ khóa trong Go
Cũng như các ngôn ngữ khác, trong golang cũng có những từ khóa mà cũng ta không được phép sử dụng để định danh (dùng khai báo tên biến, tên hàm.. ).
Trong golang có khoảng 25 từ khóa:

    break        default      func         interface    select
    case         defer        go           map          struct
    chan         else         goto         package      switch
    const        fallthrough  if           range        type
    continue     for          import       return       var

khá ít nếu so với các ngôn ngữ khác như: Java(50 từ khóa), Python(35 từ khóa)

# Các kiểu dữ liệu cơ bản trong Go
> Tham khảo: https://golang.org/ref/spec#Predeclared_identifiers
## Số Nguyên: 
Kiểu dữ liệu mặc định: 0
Kiểu dữ liệu | Số bit biểu diễn | Giới hạn biểu diễn 
-------------|------------------|-------------------
int8         | 8 bit            | -128 đến 127 
int16        | 16 bit           | -32768 đến 32767
int32        | 32 bit           | -2147483648 đến 2147483647 
int64        | 64 bit           | -9223372036854775808 đến 9223372036854775807
uint8        | 8 bit            | 0 đến 255
uint16       | 16 bit           | 0 đến 65535
uint32       | 32 bit           | 0 đến 4294967295    
uint64       | 64 bit           | 0 đến 18446744073709551615
byte         | đại diện cho kiểu uint8
rune         | đại diện cho kiểu int32


## Số Nguyên Động:
Kiểu dữ liệu mặc định: 0
 Kiểu dữ liệu | Số bit biểu diễn  | Giới hạn biểu diễn 
--------------|------------------ |-------------------
int           | ?                 | ?
uint          | ?                 | ?
uintptr       | như int, uint, nhưng lưu trữ giá trị của 1 con trỏ | ?


    3 kiểu dữ liệu này phụ thụộc vào hệ điều hành của ban, nếu hệ điều hành là 32
    bit là nó có thể biểu diễn số nguyên có chiều dài tối đa là 32 bit, nếu hệ
    điều hành 64 bit, thì nó sẽ là 64 bit

## Số thực( số hữu tỉ, số thập phân):
Kiểu dữ liệu mặc định: 0.0
Kiểu dữ liệu | Số bit biểu diễn | Giới hạn biểu diễn 
-------------|------------------|--------------------------------------------
float32      | 32 bit           |  số dấu phẩy động 32-bit theo chuẩn IEEE-754 
float64      | 64 bit           |  số dấu phẩy động 64-bit theo chuẩn IEEE-754 

## Số phức:
Kiểu dữ liệu mặc định: (0+0i)
Kiểu dữ liệu | Số bit biểu diễn | Giới hạn biểu diễn 
-------------|------------------|----------------------------------------------
complex64    | 64 bit(32 + 32)  |  phần thực và phần ảo mỗi phần là 1 số float32
complex128   | 128 bit(64 + 64) |  phần thực và phần ảo mỗi phần là 1 số float64

## Kiểu boolean:
Kiểu dữ liệu mặc định: false
Kiểu dữ liệu | Số bit biểu diễn | Giới hạn biểu diễn 
-------------|------------------|----------------------------------------------
bool         |1 bit             | chỉ biểu diễn 2 giá trị là true(1) hoặc false (0)  

## Kiểu chuỗi:
Kiểu dữ liệu mặc định: ""

Kiểu dữ liệu | biểu diễn
-------------|----------------------------------------------------------------
string       | là kiểu dữ liệu được định nghĩa là 1 mảng các byte(không có số âm) theo chuẩn UTF-8 và không có giá trị nil( là null trong Go ), và kiểu dữ liệu bất biến(immutable, không thể thay đổi nội dung của chuỗi sau khi được khởi tạo)


## Kiểu interface 
Đơn giản là mọi kiểu trong Go, có thể thay thế cho hầu hết kiểu dữ liệu trong Go(khác với kiểu <strong>interface struct</strong>)
```Go
package main

import "fmt"

func main() {
	// Khai báo kiểu dữ  liệu interface va gán giá trị là 10
	var interfaceTest []interface{} = 10
	fmt.Println(interfaceTest) // 10
	// gán lại giá trị bằng chuối
	interfaceTest = "abcd"
	fmt.Println(interfaceTest) // -> "abcd"
}
```
Lưu ý: kiểu dữ liệu interface{} có thể đại diện cho hầu hết mọi kiểu dữ liệu trong Go, nhương ngược lại thì không

Xem ví dụ dưới đây: 
```Go
func main() {
	// fmt.Println(getString("abc"))
	test("abc")  
	var data interface{} = 10
	test2(data)  
}

func test(data interface{}) {
	fmt.Println(data)
}
func test2(data int) {
	fmt.Println(data)
}
```
Chạy đoạn mã trên ta sẽ nhận được một ngoại lệ 
> cannot use data (type interface {}) as type int in argument to test2: need type assertion

tức là ở method 
```Go
test2(data int)
```
nhận vào 1 kiểu dữ liệu là int, nhưng ở lời gọi hàm lại truyền vào 1 kiểu interface {}

## Kiểu struct 
Là một tập hợp các trường có giá trị có thể giống và khác nhau trong Go (tương tự như class trong Java)
``` Go
package main

import "fmt"

// khai báo 1 struct tên là User ,
// có 2 trương dữ liệu là UserName kiểu string , và trường Age kiểu số 
type User struct {
    UserName string
    Age      int
}

func main() {
    user := new(User)
    user.UserName = "Tran van A"
    user.Age = 25
    fmt.Println(user.UserName) // "Tran van A"
    fmt.Println(user.Age)      // 25
}
```
## Kiểu slice 
Một kiểu dữ liệu tương tự như mảng trong java, là tập hợp các kiểu dữ liệu cùng loại 
Ví dụ: khai báo mảng số nguyên, thì mảng đó chỉ được chưa các số nguyên, không được chứa chuỗi hay số thực 
khai báo mảng cuối, thì mảng đó chỉ được chưa chuỗi,không chứa được số hay kiểu dữ liệu khác

Nguồn:
1. https://developer.mozilla.org/en-US/docs/Web/JavaScript/Reference/Lexical_grammar
2. https://realpython.com/python-keywords/
3. https://docs.oracle.com/javase/tutorial/java/nutsandbolts/_keywords.html
4. https://golang.org/ref/spec
5. https://golang.org/ref/spec#Predeclared_identifiers
6. https://tour.golang.org/basics/11
