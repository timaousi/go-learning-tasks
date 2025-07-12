package main

import (
	"fmt"
	"math"
	"sync"
	"sync/atomic"
	"time"
)

func main() {
	// ========== 指针 ==========
	fmt.Println("===== 指针示例 =====")
	x := 5
	addTen(&x)
	fmt.Println("加10后的值:", x)

	arr := []int{1, 2, 3}
	doubleSlice(&arr)
	fmt.Println("切片元素乘2:", arr)

	// ========== Goroutine ==========
	fmt.Println("\n===== Goroutine示例 =====")
	go printOdd()
	go printEven()
	time.Sleep(2 * time.Second) // 等待打印完成

	// ========== 任务调度器 ==========
	fmt.Println("\n===== 任务调度器示例 =====")
	tasks := []func(){
		func() { time.Sleep(500 * time.Millisecond); fmt.Println("任务1完成") },
		func() { time.Sleep(300 * time.Millisecond); fmt.Println("任务2完成") },
		func() { time.Sleep(700 * time.Millisecond); fmt.Println("任务3完成") },
	}
	var wg sync.WaitGroup
	scheduler(tasks, &wg)
	wg.Wait()
	fmt.Println("所有任务完成")

	// ========== 面向对象 ==========
	fmt.Println("\n===== 面向对象示例 =====")
	r := Rectangle{Width: 10, Height: 20}
	c := Circle{Radius: 10}
	fmt.Printf("矩形面积: %.2f, 周长: %.2f\n", r.Area(), r.Perimeter())
	fmt.Printf("圆面积: %.2f, 周长: %.2f\n", c.Area(), c.Perimeter())

	p := Employee{Person: Person{Name: "wang", Age: 12}, EmployeeID: 123}
	p.PrintInfo()

	// ========== Channel ==========
	fmt.Println("\n===== Channel示例 =====")
	channelExample()

	fmt.Println("\n===== 缓冲通道示例 =====")
	bufferedChannelExample()

	// ========== 锁机制 ==========
	fmt.Println("\n===== 锁机制示例 =====")
	mutexCounter()
	atomicCounter()
}

// ---------- 指针 ----------
func addTen(num *int) {
	*num += 10
}

func doubleSlice(nums *[]int) {
	for i := range *nums {
		(*nums)[i] *= 2
	}
}

// ---------- Goroutine ----------
func printOdd() {
	for i := 1; i <= 10; i += 2 {
		fmt.Println("奇数:", i)
		time.Sleep(100 * time.Millisecond)
	}
}

func printEven() {
	for i := 2; i <= 10; i += 2 {
		fmt.Println("偶数:", i)
		time.Sleep(100 * time.Millisecond)
	}
}

// ---------- 任务调度器 ----------
func scheduler(tasks []func(), wg *sync.WaitGroup) {
	for i, task := range tasks {
		wg.Add(1)
		go func(id int, t func()) {
			defer wg.Done()
			start := time.Now()
			t()
			fmt.Printf("任务 %d 执行时间: %v\n", id+1, time.Since(start))
		}(i, task)
	}
}

// ---------- 面向对象 ----------
type Shape interface {
	Area() float64
	Perimeter() float64
}

type Rectangle struct {
	Width, Height float64
}

func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

func (r Rectangle) Perimeter() float64 {
	return 2 * (r.Width + r.Height)
}

type Circle struct {
	Radius float64
}

func (c Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}

func (c Circle) Perimeter() float64 {
	return 2 * math.Pi * c.Radius
}

type Person struct {
	Name string
	Age  int
}

type Employee struct {
	Person
	EmployeeID int
}

func (e Employee) PrintInfo() {
	fmt.Printf("员工姓名: %s, 年龄: %d, 员工ID: %d\n", e.Name, e.Age, e.EmployeeID)
}

// ---------- Channel ----------
func channelExample() {
	ch := make(chan int)
	go func() {
		for i := 1; i <= 10; i++ {
			ch <- i
		}
		close(ch)
	}()
	for v := range ch {
		fmt.Println(v)
	}
}

func bufferedChannelExample() {
	ch := make(chan int, 10)
	go func() {
		for i := 1; i <= 100; i++ {
			ch <- i
		}
		close(ch)
	}()

	count := 0
	for v := range ch {
		fmt.Println(v)
		count++
	}
}

// ---------- 锁机制 ----------
func mutexCounter() {
	var mu sync.Mutex
	var wg sync.WaitGroup
	counter := 0

	for i := 1; i <= 10; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for j := 0; j < 1000; j++ {
				mu.Lock()
				counter++
				mu.Unlock()
			}
		}()
	}
	wg.Wait()
	fmt.Println("使用 Mutex 计数器结果:", counter)
}

func atomicCounter() {
	var counter int64
	var wg sync.WaitGroup
	for i := 1; i <= 10; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for j := 1; j <= 1000; j++ {
				atomic.AddInt64(&counter, 1)
			}
		}()
	}
	wg.Wait()
	fmt.Println("使用原子操作计数器结果:", counter)
}
