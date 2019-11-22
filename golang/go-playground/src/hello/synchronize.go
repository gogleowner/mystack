package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

func main() {
	//  example_1_raceCondition()
	//  example_1_raceCondition_mutex()
	//  example_2_readWriteLock()
	// example_3_conditionalVariable()
	example_4_executeOnce()
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
