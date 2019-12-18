package main

import (
  "flag"
  "fmt"
)

/*
$ go build commanelineFlag.go
$ ./commandlineFlag
Usage of ./commandlineFlag:
  -rating float
    	평점
  -release
    	개봉 여부
  -runtime int
    	상영 시간
  -title string
    	영화 이름
$ ./commandlineFlag -title=기생충 -runtime=120 -rating=9.9 -release=true
영화이름 : 기생충
상영시간 : 120
평점 : 9.900000
개봉여부 : 개봉
개봉여부 : 개봉
*/
func main() {
  title := flag.String("title", "", "영화 이름")
  runtime := flag.Int("runtime", 0, "상영 시간")
  rating := flag.Float64("rating", 0.0, "평점")
  release := flag.Bool("release", false, "개봉 여부")

  flag.Parse() // 명령줄 옵션의 내용을 각 자료형별로 분석

  if flag.NFlag() == 0 {
    flag.Usage()
    return
  }

  fmt.Printf(
    "영화이름 : %s\n상영시간 : %d\n평점 : %f\n",
    *title,
    *runtime,
    *rating,
  )

  if *release {
    fmt.Println("개봉여부 : 개봉")
  } else {
    fmt.Println("개봉여부 : 미개봉")
  }
}
