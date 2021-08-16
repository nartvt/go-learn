# Các câu lệnh điều kiện và vòng lặp trong Golang 
## Câu lênh điều kiện if trong Go
Thực thi khối các điều kiện trong dấu {...} nếu nếu quả của điều kiện if là đúng,
không bắt buộc điều kiện phải bao quanh dấu () trong câu điều kiện if 
```Go
  isTrue:=true
  func main(){
     if isTrue{
        fmt.Println("Hello World")
    }
      // OR  
    if(isTrue){
        fmt.Println("Hello World")
    }
  }

```
Dấu <b>{</b> bắt đầu khối lệnh trong điều kiện <b>if</b> phải nằm trên cùng hàng với <b>if</b>

```Go
    isTrue:=true
    // Đúng
    if isTrue{ // dấu { nằm cùng hàng với if  -> đúng
        // các điều kiện
    }
    // Sai
    if isTrue
    { // dấu { nằm lệch dưới với if -> sai
        // các điều kiện
    }
```
## Câu lệnh điều kiện if ... else
- Nếu điều kiện xảy ra trong <b>if</b> hoặc <b>else if</b> sai, thì sẽ chạy vào trong <b>else</b> 
- Câu lệnh <b>else</b> phải nằm trên cùng 1 hàng với dấu } kết thúc của câu lệnh <b>if</b> hoăc <b>else if </b>
``` Go
tuoi:=18
if tuoi<18 {
    fmt.Println("A") // câu else if nằm cùng hàng với } kết thúc của lệnh if 
}else { // câu else nằm cùng hàng với } kết thúc của lệnh if 
      fmt.Println("C")
}
```
- Có thể khởi tạo và so sánh trong cùng 1 câu điều kiện if 
```Go
func getNumber(a int, b int)int {
    return a*b
}
if number:=getNumber(2,3); numer==6{
    fmt.Println("number = 6")
}
```
## Câu lệnh điều kiện else if 
- Khi muốn so sánh nhiều điều kiện xảy ra bên cạnh nhưng khác với điều kiện if , ta có thể dùng else if
```Go
if a<0{
 fmt.Println("A")
}else if a==1{
 fmt.Println("B")
}else{
 fmt.Println("C")
}
``` 

## Toán tử so sánh trong Go
- <b>==</b> so sánh 2 giá trị cùng kiểu có bằng nhau hay không<br>
```Go
    if a==b{
        fmt.Println("A bang B)
    }
```
- <b><</b> nhỏ hơn<br>
```Go
    if a<b{
        fmt.Println("A nho hơn B)
    }
```
- <b>></b> lớn hơn<br>
```Go
    if a>b{
        fmt.Println("A lớn hơn B)
    }
```
- <b><=</b> nhỏ hơn hoặc bằng<br>
```Go
    if a<=b{
        fmt.Println("A nhỏ hơn hoặc bằng B)
    }
```
- <b>>=</b> lớn hơn hoặc bằng<br>
```Go
    if a>=b{
        fmt.Println("A lớn hơn hoặc bằng B)
    }
```
- <b>!=</b> khác nhau<br>
```Go
    if a!=b{
        fmt.Println("A khác B)
    }
```
## Toán tử logic trong Go
- Dấu && (đọc là và và) nối 2 hoặc nhiều điều kiện với nhau, chỉ đúng khi cả tất cả điều kiện đều đúng
```Go
if a==1 && b==2 && c ==3{
    fmt.Println("a=1, b=2, c=3)
}
```
- Dấu || (đọc là hoặc) cũng kết nối 2 hoặc nhiều điều kiện với nhau, chỉ cần 1 trong các điều kiện đúng là được
```Go
if a==1 || b==2 || c ==3{
    fmt.Println("a=1 hoặc b=2 hoặc c=3)
}
```
- Dấu ! (tạm gọi ngược lại) khi dùng dấu này, sẽ lấy giá trị ngược với giá trị đánh dấu
```Go
isTrue:=true
if !isTrue{  // khi so sánh với if có nghĩa rằng: nếu isTrue = false thì if thực thi
    fmt.Println("is false)
}
```
# Câu lệnh so sánh điều kiện Switch
- Giống như if và else if nhưng ngắn gọn hơn bằng cách so sánh trong switch và giá trị so sánh được bao bọc trong 
case
- Nếu tất cả case không đúng 1 cái nào, default sẽ xảy ra
- Khi kết thúc mỗi case hoặc default , chương trình sẽ tự thoát mà không cần break như các ngôn ngữ khác(Java)
```Go
func main() {
	chuoi:="A"
	switch chuoi {
	case "A":
		fmt.Println("A")
	case "B":
		fmt.Println("B")
	case "C":
		fmt.Println("C")
    default:
        fmt.Println("D")
	}
}
```
## Vòng lặp For có trong Go
- Vòng lặp trong Go sử dụng for 
cấu trúc 1 vòng lặp for gồm 3 phần cơ bản, phân cách bởi dấu chấm phẩy(;) như sau
    > <giá trị khỏi tạo>;<điều kiện lặp>;<điều kiện khi kết thúc 1 vòng lặp>
- Không cần sử dụng dấu () để bọc lại điều kiện trong for
```Go
for i:=0;i<=10;i++{
    fmt.Println(i)
}
 /*
 i:=0 - là giá trị khởi tạo 
 i<=10 - là điều kiện lặp
 i++ - là điều kiện xảy ra khi kết thúc 1 lần lặp
*/
```
- Trong Go không sử dụng vòng lặp while, mà sử dụng <b>for</b> thay thế
```Go
value:=10
for value>0{
    fmt.Println(value)
    value--
}
```

# Từ khóa defer trong Go
- Có thể đặt bất cứ đâu trong 1 hàm nhưng trước câu lệnh <b>return</b>
- Điều kiện gắn với <b>defer</b> sẽ thực thi khi hàm đó kết thúc
```Go
func getString() string{
	defer fmt.Println(" World")
	fmt.Print("Hello")
	return ""
}
func main() {
	getString()
}
```
- Câu lênh defer được push vào 1 stack, vì vậy khi kết thúc hàm, sẽ được lấy giá trị ra theo quy tắc las-int-first-out
```Go

func main() {
	fmt.Println("counting")

	for i := 0; i < 10; i++ {
		defer fmt.Println(i) // push vào stack
	}

	fmt.Println("done")
}
/* kết thúc hàm in ra : 
counting
done
9
8
7
6
5
4
3
2
1
0
*/
```

Nguồn: 
- https://tour.golang.org/
- https://golang.org/ref/spec