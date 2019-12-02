# 패키지 만들기, 문서화하기

## 패키지 만들기
```
# GOPATH 설정
$ export GOPATH={YOUR\_WORKSPACE}/mystack/golang/go-playground/examplepackage

# 패키지 컴파일
$ cd $GOPATH/src/calc
$ go install

# 디렉토리 위치와 상관없이 패키지를 컴파일 할 수 있다.
$ go install calc

# $GOPATH/pkg/darwin_amd64 (macos 기준) calc.a 라이브러리가 생성되어 있을 것이다.
pkg
└── darwin_amd64
    └── calc.a
```

## 문서화하기

- 콘솔에서 확인

```
$ go doc calc
package calc // import "calc"

계산 패키지

func Sum(a, b int) int
$ go doc calc Sum
package calc // import "calc"

func Sum(a, b int) int
    두 정수를 더함

# 브라우저로 확인
```

- 브라우저로 확인

  ```
$ godoc -http=localhost:6060
$ godoc -http=:6060
  ```

- `http://localhost:6060/pkg/calc` 에서 확인

