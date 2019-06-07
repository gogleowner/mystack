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

