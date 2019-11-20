package main

import (
	"fmt"
	"runtime"
	"time"
)

func sum(a int, b int, c chan int) {
	c <- a + b // 채널에 값을 보냄
}

func main() {
	//startChannel()

	// example_1_synchronousChannel()
	// example_2_channelBuffering()
	// example_3_rangeClose()
	// example_3_isChannelClose()
	// example_4_sendReceiveChannel()
	// example_4_sendReceiveChannel()
	//   example_4_returnValueToChannel()
	// example_4_sumUseOnlyChannel()
	example_5_select()
}

func startChannel() {
	c := make(chan int) // chan int 형 변수 초기화

	go sum(1, 2, c)

	n := <-c // 채널에서 값을 가져옴
	fmt.Println(n)
}

func example_1_synchronousChannel() {
	done := make(chan bool)
	count := 3

	go func() {
		for i := 0; i < count; i++ {
			done <- true
			fmt.Println("고루틴 :", i)
			time.Sleep(time.Second)
		}
	}()

	for i := 0; i < count; i++ {
		<-done // 채널에 값이 들어올 때까지 대기했다가 값을 꺼냄.
		fmt.Println("메인 함수 :", i)
	}
}

// 버퍼가 2개인 비동기 채널 생성. 비동기 채널은 보내는 쪽에서 버퍼가 가득 차면 실행을 멈추고 대기, 받는쪽에서 값이 없으면 대기
func example_2_channelBuffering() {
	runtime.GOMAXPROCS(1)

	done := make(chan bool, 2)
	count := 4

	go func() {
		for i := 0; i < count; i++ {
			done <- true // 채널에 true 보냄.
			fmt.Println("고루틴 :", i)
			// time.Sleep(time.Second)
		}
	}()

	for i := 0; i < count; i++ {
		<-done // 값을 꺼냄.
		fmt.Println("메인 함수 :", i)
		time.Sleep(time.Second)
	}

	fmt.Scanln()
}

func example_3_rangeClose() {
	c := make(chan int)

	go func() {
		for i := 0; i < 5; i++ {
			c <- i // 채널에 값을 보냄.
		}

		close(c) // 채널을 닫음. 이미 닫힌 채널에 값을 보내면 패닉 발생.
	}()

	for i := range c { // range를 사용하여 값을 꺼냄.
		fmt.Println(i)
	}
}

func example_3_isChannelClose() {
	c := make(chan int, 1)

	go func() {
		c <- 1
	}()

	receivedValue, ok := <-c
	fmt.Println("channel open?", ok, "/ responseValue :", receivedValue)
}

/*
./channel.go:111:17: invalid operation: <-c (receive from send-only type chan<- int)
./channel.go:120:7: invalid operation: c <- 1 (send to receive-only type <-chan int)
*/
func example_4_sendReceiveChannel() {
	c := make(chan int)

	go func(c chan<- int) { // send: 채널에 값을 보낸다.
		for i := 0; i < 5; i++ {
			c <- i
		}

		c <- 100

		// fmt.Println(<-c) // 채널에서 값을 꺼내면 컴파일 에러
	}(c)

	go func(c <-chan int) { // receive: 채널에서 값을 꺼내온다.
		for i := range c {
			fmt.Println(i)
		}
		fmt.Println(<-c)

		// c <- 1 // 채널에서 값을 보내면 컴파일 에러
	}(c)

	fmt.Scanln()
}

func example_4_returnValueToChannel() {
	c := func(a, b int) <-chan int {
		out := make(chan int)

		go func() {
			out <- a + b
		}()

		return out
	}(1, 2)

	// c <- 4 // 받기 전용 채널 타입이라 컴파일 에러 발생.

	fmt.Println(<-c)
}

func example_4_sumUseOnlyChannel() {
	c := func(a, b int) <-chan int {
		numChannel := make(chan int)

		go func() {
			numChannel <- a
			numChannel <- b
			close(numChannel)
		}()

		return numChannel
	}(1, 2)

	sum := func(numChannel <-chan int) <-chan int {
		sumChannel := make(chan int)

		go func() {
			sum := 0
			for n := range numChannel {
				sum += n
			}

			sumChannel <- sum
			close(sumChannel)
		}()

		return sumChannel
	}(c)

	fmt.Println(<-sum)
}

func example_5_select() {
	c1 := make(chan int)
	c2 := make(chan string)

	go func() {
		for {
			c1 <- 10
			time.Sleep(100 * time.Millisecond)
		}
	}()

	go func() {
		for {
			c2 <- "Hello world"
			time.Sleep(300 * time.Millisecond)
		}
	}()

	go func() {
		for { // 무한루프..
			select {
			case i := <-c1:
				fmt.Println("c1 :", i)
			case s := <-c2:
				fmt.Println("c2 :", s)
			case <-time.After(50 * time.Millisecond):
				fmt.Println("timeout")
				/*
				   default:
				     fmt.Println("아직 c1, c2가 초기화되지 않음.")
				     time.Sleep(100 * time.Millisecond)
				*/
			}
		}
	}()

	time.Sleep(2 * time.Second)
}
