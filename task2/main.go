package main

import (
	"fmt"
	"math"
	"math/rand"
	"sync"
	"sync/atomic"
	"time"
)

// 函数接收一个整数指针作为参数，在函数内部将该指针指向的值增加10，在返回
func pointerTen(num *int) int {
	*num += 10
	return *num
}

// 接收一个整数切片的指针，将切片中的每个元素乘以2
func sliceMultiplyTow(num []int) []int {
	for i := range num {
		num[i] *= 2

	}
	return num
}

// 输出偶数
func inc() {
	for i := 1; i <= 10; i++ {
		if i%2 == 0 {
			fmt.Printf("inc 输出偶数：%v \n", i)
			time.Sleep(30 * time.Millisecond)
		}
	}
}

// 输出奇数
func dec() {
	for i := 1; i <= 10; i++ {
		if i%2 == 1 {
			fmt.Printf("dec 输出奇数：%v \n", i)
			time.Sleep(30 * time.Millisecond)
		}
	}

}

type task func() time.Duration
type taskContent struct {
	id        int
	startTime time.Time
	endTime   time.Time
	duration  time.Duration
}

// 执行任务并记录开始时间，结束时间和任务的执行时间
func performTasks(tasks []task) []taskContent {
	results := make([]taskContent, len(tasks))
	for i, v := range tasks {
		go func(i int, t task) {
			start := time.Now().Truncate(time.Second)
			time.Sleep(v())
			end := time.Now().Truncate(time.Second)
			results[i] = taskContent{
				id:        i,
				duration:  end.Sub(start),
				startTime: start,
				endTime:   end,
			}
		}(i, v)
		time.Sleep(v())
	}
	return results
}

type Shape interface {
	Perimeter()
	Area()
}

// 长方形
type Rectangle struct {
	length float64
	width  float64
}

// 圆形
type Circle struct {
	radius float64
}

// 计算长方形面积
func (r Rectangle) Perimeter() float64 {
	area := 2 * (r.length + r.width)
	return area
}

// 计算长方形周长

func (r Rectangle) Area() float64 {
	area := r.length * r.width
	return area
}

// 计算圆形周长
func (c Circle) Area() float64 {
	return math.Pi * c.radius * c.radius
}

// 计算圆形面积
func (c Circle) Perimeter() float64 {
	return 2 * math.Pi * c.radius
}

type Person struct {
	name string
	Age  int
}
type Employee struct {
	data     Person
	id       int
	position string
}

// 输出员工信息
func (e Employee) PrintInfo() {
	fmt.Printf("员工档案：{工号：%03v,姓名：%v,年龄：%v,职位：%v}", e.id, e.data.name, e.data.Age, e.position)
}

// 只接收channel的函数
func receiveOnly(ch <-chan int) {
	for v := range ch {
		fmt.Printf("通道打印整数: %d\n", v)
	}
}

// 只发送channel的函数
func sendOnly(ch chan<- int) {
	for i := 1; i <= 10; i++ {
		ch <- i
	}
	close(ch)
}

// 生产者
func producer(ch chan<- int) {
	for i := 1; i <= 100; i++ {
		ch <- i
	}
	close(ch)
}

// 消费者
func consumer(ch <-chan int) {
	for v := range ch {
		fmt.Printf("消费者打印整数: %d\n", v)
	}
}

// 加1000
func AddOneThousand() {
	defer wg.Done()
	for i := 1; i <= 1000; i++ {
		mu.Lock()
		num++
		mu.Unlock()
	}
}
func AddOneThousand2() {
	defer wg.Done()
	for i := 1; i <= 1000; i++ {
		atomic.AddInt64(&num1, 1)
	}
}

var wg sync.WaitGroup
var mu sync.Mutex
var num int
var num1 int64

func main() {
	//——————————2.10
	//使用原子操作（ sync/atomic 包）实现一个无锁的计数器。启动10个协程，每个协程对计数器进行1000次递增操作，最后输出计数器的值。
	//考察点 ：原子操作、并发数据安全。
	wg.Add(10)
	for i := 0; i < 10; i++ {
		go AddOneThousand2()
	}
	wg.Wait()
	fmt.Printf("原子操作:%v\n", num1)

	//——————————2.9
	//编写一个程序，使用 sync.Mutex 来保护一个共享的计数器。启动10个协程，每个协程对计数器进行1000次递增操作，最后输出计数器的值。
	//考察点 ： sync.Mutex 的使用、并发数据安全。
	wg.Add(10)
	for i := 0; i < 10; i++ {
		go AddOneThousand()
	}
	wg.Wait()
	fmt.Printf("sync.Mutex 的使用:%v\n", num)

	//——————————2.8
	//实现一个带有缓冲的通道，生产者协程向通道中发送100个整数，消费者协程从通道中接收这些整数并打印。
	//考察点 ：通道的缓冲机制。
	var intChan1 = make(chan int, 100) //定义有缓冲的通道
	go producer(intChan1)
	go consumer(intChan1)
	time.Sleep(500 * time.Millisecond)

	//——————————2.7
	//编写一个程序，使用通道实现两个协程之间的通信。一个协程生成从1到10的整数，并将这些整数发送到通道中，另一个协程从通道中接收这些整数并打印出来。
	//考察点 ：通道的基本使用、协程间通信。
	var intChan = make(chan int) //定义通道
	go sendOnly(intChan)
	go receiveOnly(intChan)
	time.Sleep(500 * time.Millisecond)

	//——————————2.6
	//使用组合的方式创建一个 Person 结构体，包含 Name 和 Age 字段，再创建一个 Employee 结构体，组合 Person 结构体并添加 EmployeeID 字段。为 Employee 结构体实现一个 PrintInfo() 方法，输出员工的信息。
	//考察点 ：组合的使用、方法接收者。
	info := Employee{Person{"李四", 28}, 2, "经理"}
	info.PrintInfo()

	//——————————2.5
	//题目 ：定义一个 Shape 接口，包含 Area() 和 Perimeter() 两个方法。然后创建 Rectangle 和 Circle 结构体，实现 Shape 接口。在主函数中，创建这两个结构体的实例，并调用它们的 Area() 和 Perimeter() 方法。
	//考察点 ：接口的定义与实现、面向对象编程风格。
	//perimeter := Rectangle{10, 20}
	perimeter := Rectangle{}
	perimeter.length = 23
	perimeter.width = 44
	//circle := Circle{10}
	circle := Circle{}
	circle.radius = 18
	fmt.Printf("圆形面积: %.2f\n圆形周长: %.2f\n\n", circle.Area(), circle.Perimeter())
	fmt.Printf("长方形面积: %.2f\n长方形周长: %.2f\n\n", perimeter.Area(), perimeter.Perimeter())

	//——————————2.4
	//设计一个任务调度器，接收一组任务（可以用函数表示），并使用协程并发执行这些任务，同时统计每个任务的执行时间。
	//考察点 ：协程原理、并发任务调度。
	tasks := []task{
		func() time.Duration { return 1 * time.Second },
		func() time.Duration { return 2 * time.Second },
		func() time.Duration { return 3 * time.Second },
	}
	for _, r := range performTasks(tasks) {
		fmt.Printf("任务 %d: 执行时间=%v, 开始时间=%v, 结束时间=%v\n", r.id+1, r.duration, r.startTime.Format("15:04:05.000"), r.endTime.Format("15:04:05.000"))
	}

	//——————————2.3
	//编写一个程序，使用 go 关键字启动两个协程，一个协程打印从1到10的奇数，另一个协程打印从2到10的偶数。
	//考察点 ： go 关键字的使用、协程的并发执行
	go func() {
		inc()
	}()
	go func() {
		dec()
	}()
	time.Sleep(2 * time.Second)

	//——————————2.2
	//题目 ：实现一个函数，接收一个整数切片的指针，将切片中的每个元素乘以2。
	//考察点 ：指针运算、切片操作。
	slice := make([]int, 10)
	for i := 0; i < len(slice); i++ {
		slice[i] = getRand()
	}
	//打印赋值后的值
	fmt.Println("切片的值是：", slice)
	//调用每个元素乘以2的函数
	sliceMultiplyTow(slice)
	//打印每个元素乘以2的值
	fmt.Println("切片的值*2后的结果是：", slice)

	//——————————2.1
	//题目 ：编写一个Go程序，定义一个函数，该函数接收一个整数指针作为参数，在函数内部将该指针指向的值增加10，然后在主函数中调用该函数并输出修改后的值。
	//考察点 ：指针的使用、值传递与引用传递的区别
	num := 10
	fmt.Printf("+10后的值：%v", pointerTen(&num))

}

// 获取随机数
func getRand() int {
	rand.Seed(time.Now().UnixNano()) // 设置随机数种子
	return rand.Intn(100)
}
