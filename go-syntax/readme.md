

# Từ khóa trong Go
Cũng như các ngôn ngữ khác, trong golang cũng có những từ khóa mà cũng ta không được phép sử dụng để định danh (dùng khai báo tên biến, tên hàm.. ).
Trong golang có 25 từ khóa:

    break        default      func         interface    select
    case         defer        go           map          struct
    chan         else         goto         package      switch
    const        fallthrough  if           range        type
    continue     for          import       return       var

khá ít nếu so với các ngôn ngữ khác như: Java(50 từ khóa), Python(35 từ khóa)
# Những cái cơ bản cần biết trong Go
##### 1. Cú pháp Go đơn giản , dễ dàng cho người mới bắt đầu
##### 2. Go không có OOP mà là functional
##### 3. Hàm main là duy nhất và nó chỉ nằm trong  package main<br>
```Go
    // func <function name>(paremeter input ) <return type>{ // something }

    func addNumber(number int) int{
        return number*2
    }
```
##### 4. Import trong Go: khi sử dụng 1 thư viện nào đấy(kể cả build-in) khác package với package hiện tại, cần import vào để sử dụng nằm trong dấu nháy kép ""
```Go
package main
import "fmt"
func main(){
    fmt.Println("Hello World")
}
```
import thư viện trong go có 2 cách<br>
Cách 1, import riêng từng package
```Go
package main
import "fmt"
import "math"
func main(){
      fmt.Println("Hello World - ", math.Sqrt(4))
}
```
Cách 2 - import các package lồng trong ngoặc tròn đơn 
```Go
package main
import (
    "fmt"
    "math"
)
func main(){
    fmt.Println("Hello World - ", math.Sqrt(4))
}
```
##### 5. Trong Go không sử dụng dấu chấm phẩy (;) khi kết thúc một dòng 
##### Khai báo kiểu dữ liệu trong Go, kiểu dữ liệu sẽ viết sau tên định danh<br>

        Go : <tên đại diện> <kiểu dữ liệu>
        Java : <Kiểu dữ liệu> <tên đại diện>
Ví dụ: 
```Go
 var number int = 10
 func addNumber(number int){
    
}
```
Trong Java
```Java
int number = 10;
void addNumber(int number){

}
```
##### 6. Trong Go nếu viết thường chữ cái đầu tiên trong tên(thuộc tính, hàm, sruct) thì scope của nó chỉ nằm trong phạm vi package nó được khai báo
Muốn truy cập bên ngoài package , viết hoa chữ cái đầu tiên của nó lên
```Go
package main
// chỉ cho phép truy cập từ bên ngoài package
func GetA()string{
    return "A"
}
// chỉ cho phép truy cập từ bên trong package
func getB()string{
    return "B"
}
```
##### 7. Khai báo 1 biến mà không sử dụng là không được chấp nhận trong go
chương trình sẽ không chạy và báo lỗi nếu bạn làm điều này
```Go
func main() {
	 var number int = 10
	 chuoi:="ABC" // khai báo nhưng không sử dụng
	 fmt.Println(number)
}
```
nó nói ràng, "chuoi" đã được khao báo nhưng không được sử dụng
> ./prog.go:7:3: chuoi declared but not used <br> Go build failed.

##### 8. Các thư viện của Go nếu không phải build-in đều được cài đặt từ github

        go get <github repository>
##### 9. Kể từ 1.16 (hiện tại 1.67) Go sử dụng go module(go.mod) để quán lý các package cài đặt (tương tự như pom.xml hoặc gradle trong java )

# Khai báo kiểu dữ liệu trong Go
## Với kiểu dữ liệu nguyên thủy(primitive type)
Cơ bản có 3 cách khai báo với loại dữ liệu này:
```Go
  tên biến   toán tử gán 
      |        |
var number int = 10 -- giá trị
 |          |
 |      kiểu dữ liệu
variable

var number = 10 

number:= 10 -- giá trị
       |
    kiểu dữ liệu ngầm định(implicit type)
```
## Với kiểu dữ liệu cấu trúc(struct)
sử dụng từ khóa 
> <strong style="color:white">type</strong> <Tên định danh> <strong style="color:white">struct</strong>{
    // thuộc tính 
}

``` Go
type User struct {
	UserName string
	Age int
}

func main() {
    // Cách 1:  khai báo và cấp phát bộ nhớ với kiểu dữ liệu là User,
    // là 1 giá trị(mặc định {0})
	var user User 
    // Cách 2: khai báo và cấp phát bộ nhớ với kiểu dữ liệu là User ,
    //  giá trị trả ra một con trỏ là giá trị của vùng nhớ được trỏ,
    //  mặc định là &{0}
	user2:=new(User)

    // 
    fmt.Println(user)  //  {0}
	fmt.Println(user2) //  &{0}
}
```
## Với kiểu dữ liệu tập hợp(slice)
Có 4 cách khai báo như sau: 
```Go
var primes [] int  // []
primes2:= []int{} // []
primes3:= [6]int{} // [0 0 0 0 0 0]
primes4:=make([]int,6) // [0 0 0 0 0 0]
```
> cho dù ta có quy định trước kích thước trước như kiểu số 2, 3 ta vẫn có thể add thêm value bằng từ khóa <strong style="color: white;">append</strong>
cho nên không nhất thiết là khai báo 6 là cố định giá số phần tử 6 cho nên khi dùng cũng không quan trọng lắm về việc dùng kiểu nào, 
Nên dùng kiểu số 1 code sẽ được tường minh hơn
``` Go
var primes [] int
primes = append(primes,1)
primes = append(primes,2)
fmt.Println(primes) // [1 2]
```


# Các kiểu dữ liệu cơ bản trong Go
> Tham khảo: https://golang.org/ref/spec#Predeclared_identifiers
## Số Nguyên: 
- Giá trị mặc định: 0

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
- Giá trị mặc định: 0

Kiểu dữ liệu | Số bit biểu diễn  | Giới hạn biểu diễn 
--------------|------------------ |-------------------
int           | ?                 | ?
uint          | ?                 | ?
uintptr       | như int, uint, nhưng lưu trữ giá trị của 1 con trỏ | ?


    3 kiểu dữ liệu này phụ thụộc vào hệ điều hành của ban, nếu hệ điều hành là 32
    bit là nó có thể biểu diễn số nguyên có chiều dài tối đa là 32 bit, nếu hệ
    điều hành 64 bit, thì nó sẽ là 64 bit

## Số thực( số hữu tỉ, số thập phân):
- Giá trị mặc định: 0.0

Kiểu dữ liệu | Số bit biểu diễn | Giới hạn biểu diễn 
-------------|------------------|--------------------------------------------
float32      | 32 bit           |  số dấu phẩy động 32-bit theo chuẩn IEEE-754 
float64      | 64 bit           |  số dấu phẩy động 64-bit theo chuẩn IEEE-754 

## Số phức:
- Giá trị mặc định: (0+0i)

Kiểu dữ liệu | Số bit biểu diễn | Giới hạn biểu diễn 
-------------|------------------|----------------------------------------------
complex64    | 64 bit(32 + 32)  |  phần thực và phần ảo mỗi phần là 1 số float32
complex128   | 128 bit(64 + 64) |  phần thực và phần ảo mỗi phần là 1 số float64

## Kiểu boolean:
- Giá trị mặc định: false

Kiểu dữ liệu | Số bit biểu diễn | Giới hạn biểu diễn 
-------------|------------------|----------------------------------------------
bool         |1 bit             | chỉ biểu diễn 2 giá trị là true(1) hoặc false (0)  

## Kiểu chuỗi:
- Giá trị mặc định: "".

Kiểu dữ liệu   | Giới hạn biểu diễn
---------------|----------------------------------------------------------------
string         | là kiểu dữ liệu được định nghĩa là 1 mảng các byte(không có số âm) theo chuẩn UTF-8 và không có giá trị nil( là null trong Go ), và kiểu dữ liệu bất biến (immutable, không thể thay đổi nội dung của chuỗi sau khi được khởi tạo)

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
<b>Lưu ý</b>: kiểu dữ liệu <b>interface{}</b> có thể đại diện cho hầu hết mọi kiểu dữ liệu trong Go, nhưng ngược lại thì không

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

nghĩa là ở hàm 
```Go
test2(data int)
```
nhận vào 1 kiểu dữ liệu là int, nhưng ở lời gọi hàm lại truyền vào 1 kiểu interface {}

## Kiểu struct 
Tập hợp các trường có kiểu dữ liệu có thể giống và khác nhau trong Go (tương tự như class trong Java)
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
## Kiểu dữ liệu Arrays
- Một kiểu dữ liệu tương tự như slice nhưng cố định trước kích thước(fixed-sized)
- Không thể thêm phần tử nhiều hơn giá trị size ban đầu đã cố định
- Là tập hợp các giá trị có cùng kiểu dữ liệu<br>
```Go
var numbers [2] int
numbers[0] = 10
numbers[1] = 11
fmt.Println(numbers) // [0 1]
```

## Kiểu dữ liệu slice 
- Light weight data structure
- Tương tự như mảng trong nhưng không cần quy định trước kích thước(dynamically-sized)
- Bản chất là 1 con trỏ mảng trỏ đến 1 mảng bên dưới
- Có thể thêm phần tử nhiều hơn giá trị size ban đầu đã cố định
- Là tập hợp các giá trị có cùng kiểu dữ liệu<br>

Trong 1 slice gồm 3 thành phần:

    1. pointer: trỏ đến 1 mảng 
    2. length: chiêu dài, là số phần tử hiện có trong slice
    3. capacity: tối đa số phần tử mà 1 slice có thể chứa

Ví dụ khai báo một slice số nguyên như sau:  

```Go
var primes [] int
```
- Thêm giá trị vào  slice không cần theo index 
> Sử dụng từ khóa <strong style="color: white">append</strong><br>

Ví dụ dưới đây khai báo 1 slice chưa số nguyên, và chỉ thêm vào slice nếu số > 10
``` Go
func addNumber(number int){
    var slice [] int
    if number > 10{
        slice = append(slice,number)
    }
}
```
- Gộp 2 slice lại, sử dụng từ khóa append
```Go
slice1:=[]int{1,2,3}
slice2:=[]int{4}
slice3:=append(slice1,slice2...)
fmt.Println(slice3) // [1 2 3 4]
```
- Slice literals - một cách mà khi khai báo, các giá trị là không đổi
 > chỉ nên sử dụng trong trường hợp các giá trị là cố định và không đổi , chỉ việc lấy ra sử dụng, như các hằng số , các giá trị không đổi
```Go
slice:=[]int{1,2,3,4,5,6}
fmt.Println(slice) // [1 2 3 4 5 6]
```
- Lấy chiều dài 1 slice 
> sử dụng từ khóa len(slice)
``` Go 
slice:=[]int{1,2,3,4,5,6}
fmt.Println(len(slice)) // 6
```

## Kiểu dữ liệu map trong golang
- Là tập hợp ánh xạ các giá trị(value) theo khóa(key)
- Mỗi khóa trong map là duy nhất
- Nếu thêm 2 khóa trùng nhau, sẽ lấy giá trị sau

Khai báo map với dynamically-sized
```Go
maps:=make(map[string]int)
maps["A"] = 1
maps["B"] = 2
maps["C"] = 20
fmt.Println(maps) // map[A:1 B:20 C:20]
```
- Map literals - Một cách khai báo map khác<br>
cách này tạo ra và gán giá trị cố định cho nó luôn
trong trường hợp các giá trị là hằng cố định và không thay đổi, chỉ việc lấy ra sử dụng

```Go
// Một cách khai báo cồng kềnh 
var maps map[string] int = map[string]int{
    "A": 1,
    "B": 2,
    "C": 3,
}
fmt.Println(maps) // map[A:1 B:2 C:3]
```
```Go
// Bớt cồng kềnh hơn bên trên
var maps = map[string]int{
    "A": 1,
    "B": 2,
    "C": 3,
}
fmt.Println(maps) // map[A:1 B:2 C:3]
```
```Go
// Bớt cồng kềnh hơn bên trên
maps:= map[string]int{
    "A": 1,
    "B": 2,
    "C": 3,
}
fmt.Println(maps) // map[A:1 B:2 C:3]
```

- Xóa một phẩn tử trong map
```Go
maps:=make(map[string]int)
maps["A"] = 1
maps["B"] = 2
maps["C"] = 20
delete(maps, "B")
fmt.Println(maps) // map[A:1 C:20]
```

Nguồn tham khảo:
- https://developer.mozilla.org/en-US/docs/Web/JavaScript/Reference/Lexical_grammar
- https://realpython.com/python-keywords/
- https://docs.oracle.com/javase/tutorial/java/nutsandbolts/_keywords.html
- https://golang.org/ref/spec
- https://golang.org/ref/spec#Predeclared_identifiers
- https://tour.golang.org/basics/11