package main

import (
	"fmt"
	"math/rand"
	"runtime"
	"sync"
	"sync/atomic"
	"time"
)

func main() {
	//  example_1_raceCondition()
	//  example_1_raceCondition_mutex()
	//  example_2_readWriteLock()
	// example_3_conditionalVariable()
	// example_4_executeOnce()
	// example_5_pool()
	// example_6_waitGroup()
	example_7_atomicOperation()
}

func example_1_raceCondition() {
	var data = []int{} // 두 고루틴이 동시에 접근하기 때문에 append가 정확히 처리 되지 않음

	go func() {
		for i := 0; i < 1000; i++ {
			data = append(data, 1)
			runtime.Gosched() // 다른 고루틴이 CPU를 사용할 수 있도록 양보
		}
	}()

	go func() {
		for i := 0; i < 1000; i++ {
			data = append(data, 1)
			runtime.Gosched() // 다른 고루틴이 CPU를 사용할 수 있도록 양보
		}
	}()

	time.Sleep(2 * time.Second)

	fmt.Println(len(data))
}

func example_1_raceCondition_mutex() {
	var data = []int{} // 두 고루틴이 동시에 접근하기 때문에 append가 정확히 처리 되지 않음
	var mutex = new(sync.Mutex)

	go func() {
		for i := 0; i < 1000; i++ {
			mutex.Lock()
			data = append(data, 1)
			mutex.Unlock()
			runtime.Gosched() // 다른 고루틴이 CPU를 사용할 수 있도록 양보
		}
	}()

	go func() {
		for i := 0; i < 1000; i++ {
			mutex.Lock()
			data = append(data, 1)
			mutex.Unlock()

			runtime.Gosched() // 다른 고루틴이 CPU를 사용할 수 있도록 양보
		}
	}()

	time.Sleep(2 * time.Second)

	fmt.Println(len(data))
}

/*
readLock : 읽을 때는 서로 막지 않으나, 읽기 중에 쓰기 락은 막는다.
writeLock : 읽기 / 쓰기 락 모두 막는다.
*/
func example_2_readWriteLock() {
	var data int = 0
	var rwMutex = new(sync.RWMutex)

	go func() { // write
		for i := 0; i < 3; i++ {
			rwMutex.Lock() // 쓰기 잠금 시작
			data += 1
			fmt.Println("write :", data)
			time.Sleep(10 * time.Millisecond)
			rwMutex.Unlock() // 쓰기 잠금 해제
		}
	}()

	go func() { // read
		for i := 0; i < 3; i++ {
			rwMutex.RLock() // 읽기 보호
			fmt.Println("read 1:", data)
			time.Sleep(1 * time.Second)
			rwMutex.RUnlock()
		}
	}()

	go func() { // read
		for i := 0; i < 3; i++ {
			rwMutex.RLock()
			fmt.Println("read 2:", data)
			time.Sleep(2 * time.Second)
			rwMutex.RUnlock()
		}
	}()

	time.Sleep(5 * time.Second)
}

/*
조건변수 : 대기하고 있는 객체 하나만 깨우거나 여러개를 동시에 깨울 때 사용
*/
func example_3_conditionalVariable() {
	var mutex = new(sync.Mutex)
	var cond = sync.NewCond(mutex) // 뮤택스를 이용하여 조건 변수 생성
	loopCount := 3
	c := make(chan bool, 3) // 비동기 채널 생성

	for i := 0; i < loopCount; i++ {
		go func(n int) {
			mutex.Lock()
			c <- true
			fmt.Println("wait begin :", n)
			cond.Wait()
			fmt.Println("wait end:", n)
			mutex.Unlock()
		}(i)
	}

	for i := 0; i < loopCount; i++ {
		<-c
	}
	/* // Signal : 고루틴을 차례대로 하나씩 깨운다.
	for i:=0; i<loopCount; i++ {
	  mutex.Lock()
	  fmt.Println("signal :", i)
	  cond.Signal()
	  mutex.Unlock()
	}
	*/
	// Broadcast : 고루틴을 모두 깨운다.
	mutex.Lock()
	fmt.Println("broadcast")
	cond.Broadcast()
	mutex.Unlock()

	fmt.Scanln()
}

/*
Once.Do()는 어떤 상황이든 함수 또는 클로저를 딱 한번만 실행시킨다.
복잡한 반복문 안에서 각종 초기화를 할 때 유용하다.
*/
func example_4_executeOnce() {
	once := new(sync.Once)

	for i := 0; i < 3; i++ {
		go func(n int) {
			fmt.Println("goroutine :", n)

			once.Do(func() { fmt.Println("hello") })
		}(i)
	}

	fmt.Scanln()
}

/*
풀은 객체(메모리)를 사용한 후 보관해두었다가 다시 사용하게 해주는 기능
객체를 반복해서 할당하면 메모리 사용량이 늘어나고 가비지컬렉터에게도 부담이 되는데,
풀은 일종의 캐시라고 할 수 있으며 메모리 할당 & 해제 횟수를 줄여 성능을 높이고자 할 때 사용한다.
*/

type Data struct {
	tag    string
	buffer []int
}

func example_5_pool() {
	pool := sync.Pool{
		New: func() interface{} {
			data := new(Data)
			data.tag = "new"
			data.buffer = make([]int, 10)

			return data
		},
	}

	for i := 0; i < 10; i++ {
		go func() {
			data := pool.Get().(*Data)
			for index := range data.buffer {
				data.buffer[index] = rand.Intn(100)
			}
			fmt.Println("random Put Data :", data)

			data.tag = "used"
			pool.Put(data)
		}()
	}

	for i := 0; i < 10; i++ {
		go func() {
			data := pool.Get().(*Data)
			n := 0
			for index := range data.buffer {
				data.buffer[index] = n
				n += 2
			}
			fmt.Println("jjacksu Put Data :", data)

			data.tag = "used"
			pool.Put(data)
		}()
	}

	fmt.Scanln()
}

/*
waitGroup : 고루틴이 모두 끝날 때까지 기다릴 때 사용한다.
*/
func example_6_waitGroup() {
	wg := new(sync.WaitGroup)

	for i := 0; i < 10; i++ {
		wg.Add(1) // 대기그룹에 고루틴 개수 추가
		go func(n int) {
			defer wg.Done() // 고루틴이 끝나기 직전에 wg.Done() 호출
			fmt.Println(n)
		}(i)
	}

	wg.Wait() // 모든 고루틴이 끝날 때까지 기다림
	fmt.Println("the end")
}

/*
원자적연산 : 더이상 쪼갤 수 없는 연산여러 스레드, CPU코어에서 같은 변수를 수정할 때 서로 영향 받지 않고 안전하게 연산할 수 있다.
보통 원자적 연산은 CPU의 명령어를 직접 사용하여 구현되어 있다.
*/
func example_7_atomicOperation() {
	// 고루틴을 사용하여 정수형 변수를 2000번은 더하고, 1000번은 뺀다.
	wg := new(sync.WaitGroup)
	var data int32 = 0

	for i := 0; i < 2000; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			// data += 1 // atomic하지 않다.
			atomic.AddInt32(&data, 1)
		}()
	}

	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			// data -= 1
			atomic.AddInt32(&data, -1)
		}()
	}

	wg.Wait()

	fmt.Println("expect result is 1000, actual result is [", data, "]")

}
