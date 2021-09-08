## Single thread 
- Hệ thông phần cứng chỉ có single core
- Chạy từng tác vụ cho đến khi xong tác vụ này mới đến tác vụ khác 

<img src="single_thread.png" style="width: 900px"/>

## Concurency
- Là việc 1 thread có thể thực hiện nhiều tác vụ trong 1 khoảng thời gian, nhưng tại 1 thời điểm chỉ có thể thực hiện 1 tác vụ
- Mất 1 khoảng thời gian cũng như tài nguyên để thực hiện việc chuyển đổi qua lại giữa các tác vụ ( gọi là cost)
- Quá trình chuyển qua lại giữa các tác vụ gọi là switch context
- Việc trở lại tiến trình đã switch context (chuyển trở lại tác vụ trước đó) gọi là restore context

## Parallelism
- Số lượng tác vụ có thể chạy parallelism được quyết định bởi số lượng thread mà hệ thống có(số lượng core)
- Là việc nhiều thread chạy cùng 1 lúc và điều phối nhiều tác vụ trong 1 khoảng thời gian, có thể giải quyết nhiều tác vụ trong cùng 1 thời điểm
- Không mất thời gian chuyển đổi tác vụ vì 1 tác vụ sẽ chạy từ đầu đến khi kết thúc mà không chuyển đổi hay dừng lại(trừ khi bị block bởi 1 tác vụ khác, như tranh chấp tài nguyên, block tiến trình bởi hệ điều hành .. )



<img src="con_para.jpg" style="width: 900px"/>

## Tham khảo 
1. https://devopedia.org/concurrency-vs-parallelism