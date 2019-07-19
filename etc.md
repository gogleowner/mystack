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

