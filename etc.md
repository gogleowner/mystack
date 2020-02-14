# 카테고리 분류하기 모호한 것들 모음.

## IntelliJ IDEA
- `Cmd + Shift + F8` : 해당 프로젝트에 걸어놓은 Break Point 목록을 보여준다. 여기에서 설정된 Break Point를 제어할 수 있음
  - Reference : https://www.jetbrains.com/help/idea/using-breakpoints.html

## Spring Framework
- CustomizableThreadFactory
  - Thread 생성시에  여러 설정 값 세팅에 도움을 줄 수 있는 클래스이다.
  - Usage

```
class MyTest {
    @Test
    internal fun `CustomizableThreadFactory practice`() {
        val executorService = Executors.newSingleThreadExecutor(object : CustomizableThreadFactory() {
            override fun getThreadNamePrefix(): String {
                return "gogleowner-"
            }
            override fun getThreadPriority(): Int {
                return Thread.MIN_PRIORITY
            }
        })
        executorService.submit { logger.info { "this is gogleowner thread!" } }
        if (executorService.awaitTermination(5, TimeUnit.SECONDS)) {
            executorService.shutdown()
        }
    }
}
```

  - Reference : https://docs.spring.io/spring/docs/current/javadoc-api/org/springframework/scheduling/concurrent/CustomizableThreadFactory.html

## Scala
- Scala shell 이 도움이 될때가 있다...
- 앱의 logger 가 json으로 출력되는 경우가 있을텐데 이 경우 stacktrace 를 확인하기 너무 어렵다.

```
{
    "errorStacktrace": "java.lang.NullPointerException\n\tat blabla.AClass.aMethod(AClass.java:10)\n\tat java.lang.Thread.run(Thread.javadoc:748)"
}
```

- 이 경우 `\n` 으로 split 하면 stacktrace가 예쁘게 나온다.

```
$ scala
scala> "java.lang.NullPointerException\n\tat blabla.AClass.aMethod(AClass.java:10)\n\tat java.lang.Thread.run(Thread.javadoc:748)".split("\n").foreach(println)

java.lang.NullPointerException
    at blabla.AClass.aMethod(AClass.java:10)
    at java.lang.Thread.run(Thread.javadoc:748)
```

- 이런게 대화형 쉘의 장점인 듯 싶다.
  - cf) python 으로도 동일하게 사용할 수 있겠다. 내일 해봐야지.

## curl
- header 넘기기 : `-H`, 데이터 넘기기 : `-d`

```
curl -XPUT localhost:8080/foo/bar -H "Content-Type: application/json" -d {}
```

## Shell Script
- 파일 내의 중복된 라인이 많은 순으로 정렬

```
aaa
bbb
ccc
ddd
aaa
aaa
aaa
bbb
bbb
bbb
ccc
ccc
```

- 위와 같은 파일이 있다고 가정..
    - `$ cat tmp.txt | sort | uniq -c | sort -nr | awk '{print $2}'`
- `sort` : 말그대로 정렬이다. `aaa aaa aaa aaa bbb bbb bbb...`
- `uniq -c` : uniq는 반복되는 라인을 필터링하는 기능을 가지고 있는데, `-c (--count)` 옵션은 개수를 기준으로 필터링 하는 옵션이다.

```
4 aaa
3 bbb
2 ccc
1 ddd
```

- `sort -nr` : `-n (--numeric-sort)`, `-r (--reverse)` 는 numlric 한 값을 역순으로 정렬한다는 의미이다.

- `awk '{print $2}'` : 두번째 필드를 출력하라는 함수이다.
    - awk : https://ko.wikipedia.org/wiki/AWK

```
aaa
bbb
ccc
ddd
```

- if 문 조건 :: 젠킨스 잡에서 파라미터에 값을 입력하지 않았을 경우 값이 있는지 판단할 때 등등...
  - `if [ -n "$파라미터명" ]; then` .. `fi`
  - `-n` : 문자열의 길이가 0이 아니면 참
  - `-z` : 문자열의 길이가 0이면 참
    - 위의 if 문은 `if [ ! -z "$파라미터명" ]; then` .. `fi` 으로도 표현 가능
  - `-eq` : 값이 같으면 참
  - `-ne` : 값이 다르면 참
  - `-gt` : 값1 > 값2
  - `-ge` : 값1 >= 값2
  - `-lt` : 값1 < 값2
  - `-le` : 값1 <= 값2
  - `-a` : &&
  - `-o` : ||
  - `-d` : 파일이 디렉토리면 참
  - `-e` : 파일이 있으면 참
  - `-L` : 파일이 심볼릭링크면 참
  - `-r` : 파일이 읽기 가능하면 참
  - `-w` : 파일이 쓰기 가능하면 참
  - `-x` : 파일이 실행 가능하면 참
  - `-s` : 파일의 크기가 0보다 크면 참
  - `파일1 -nt 파일2` : 파일1이 파일2보다 최신파일이면 참
  - `파일1 -ot 파일2` : 파일1이 파일2보다 이전파일이면 참
  - `파일1 -ef 파일2` : 파일1이 파일2와 같은 파일이면 참

