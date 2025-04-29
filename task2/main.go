package main

import (
	"fmt"
	"sync"
	"sync/atomic"
	"time"
)

// 指针 题目1 ：编写一个Go程序，定义一个函数，该函数接收一个整数指针作为参数，在函数内部将该指针指向的值增加10，然后在主函数中调用该函数并输出修改后的值。
func addNum(num *int) {
	*num += 10
	fmt.Println(*num)
}

// 指针 题目2 ：实现一个函数，接收一个整数切片的指针，将切片中的每个元素乘以2。
func multSlice(nums *[]int) {
	for i := 0; i < len(*nums); i++ {
		(*nums)[i] *= 2
	}
	fmt.Println(*nums)
}

// Goroutine 题目1 ：编写一个程序，使用 go 关键字启动两个协程，一个协程打印从1到10的奇数，另一个协程打印从2到10的偶数。
func printJiOu() {
	go func() {
		for i := 1; i <= 10; i++ {
			if i%2 == 1 {
				fmt.Println("奇数：", i)
			}
		}
	}()
	go func() {
		for i := 1; i <= 10; i++ {
			if i%2 == 0 {
				fmt.Println("偶数", i)
			}
		}
	}()
	time.Sleep(time.Second * 2)

}

// 任务调度
func exeTask(tasks []func()) {
	for _, task := range tasks {
		go task()
	}
}

// 面向对象 题目1 ：定义一个 Shape 接口，包含 Area() 和 Perimeter() 两个方法。然后创建 Rectangle 和 Circle 结构体，实现 Shape 接口。在主函数中，创建这两个结构体的实例，并调用它们的 Area() 和 Perimeter() 方法。
type Shape interface {
	Area()
	Perimeter()
}

type Rectangle struct {
}
type Cicle struct {
}

func (r Rectangle) Area() {
	fmt.Println("Rectangle Area")
}
func (r Rectangle) Perimeter() {
	fmt.Println("Rectangle Perimeter")
}
func (r Cicle) Area() {
	fmt.Println("Cicle Area")
}
func (r Cicle) Perimeter() {
	fmt.Println("Cicle Perimeter")
}
func doFaceObj() {
	var shape Shape
	shape = Rectangle{}
	shape.Area()
	shape.Perimeter()
	shape = Cicle{}
	shape.Area()
	shape.Perimeter()

}

// 面向对象 题目2 ：使用组合的方式创建一个 Person 结构体，包含 Name 和 Age 字段，再创建一个 Employee 结构体，组合 Person 结构体并添加 EmployeeID 字段。为 Employee 结构体实现一个 PrintInfo() 方法，输出员工的信息。
type Person struct {
	Name string
	Age  int
}

type Employee struct {
	Person
	EmployeeID int
}

func (e Employee) PrintInfo() {
	fmt.Println("Name:", e.Name)
	fmt.Println("Age:", e.Age)
	fmt.Println("EmployeeID:", e.EmployeeID)
}

// Channel 1题目 ：编写一个程序，使用通道实现两个协程之间的通信。一个协程生成从1到10的整数，
// 并将这些整数发送到通道中，另一个协程从通道中接收这些整数并打印出来。
func printChannel() {
	ch := make(chan int)
	go func() {
		for i := 1; i <= 10; i++ {
			ch <- i
			fmt.Println("发送到channel:", i)
		}
		close(ch)
	}()
	go func() {
		for num := range ch {
			fmt.Println("获取channel:", num)
		}
	}()
	time.Sleep(time.Second * 2)
}

// Channel 2题目:实现一个带有缓冲的通道，生产者协程向通道中发送100个整数，消费者协程从通道中接收这些整数并打印。
func printChannel2() {
	ch := make(chan int, 100)
	go func() {
		for i := 1; i <= 100; i++ {
			ch <- i
			fmt.Println("发送到channel:", i)
			time.Sleep(10 * time.Microsecond)
		}
		close(ch)
	}()
	go func() {
		for i := 1; i <= 100; i++ {
			fmt.Println("获取channel:", <-ch)
			time.Sleep(10 * time.Microsecond)
		}
	}()
	time.Sleep(time.Second * 5)
}

// 锁机制 题目1 ：编写一个程序，使用 sync.Mutex 来保护一个共享的计数器。启动10个协程，每个协程对计数器进行1000次递增操作，最后输出计数器的值。
func syncLock() {
	var sc = sync.Mutex{}
	var count int
	for i := 0; i < 10; i++ {
		go func() {
			sc.Lock()
			for i := 0; i < 1000; i++ {
				count++
			}
			sc.Unlock()
		}()
	}
	time.Sleep(time.Second * 2)
	fmt.Println("count:", count)
}

// 锁机制 题目2 ：使用原子操作（ sync/atomic 包）实现一个无锁的计数器。启动10个协程，每个协程对计数器进行1000次递增操作，最后输出计数器的值。
func atmicLock() {
	fmt.Println("使用原子操作实现计数器")
	var count = atomic.Int64{}
	for i := 0; i < 10; i++ {
		go func() {
			for i := 0; i < 1000; i++ {
				count.Add(1)
			}
		}()
	}
	time.Sleep(time.Second * 2)
	fmt.Println("count:", count.Load())
}

func main() {
	// num := 10
	// fmt.Println(addNum(&num))
	// numSlice := []int{1, 2, 3, 4, 5}
	// fmt.Println(multSlice(&numSlice))

	// printJiOu()
	// doFaceObj()
	// person := Employee{
	// 	Person: Person{
	// 		Name: "张三",
	// 		Age:  18,
	// 	},
	// 	EmployeeID: 1,
	// }
	// person.PrintInfo()
	atmicLock()
}
