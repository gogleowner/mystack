# Connection Pool

# What is DataBase Connection Pooling

데이터베이스 커넥션 풀링은 데이터베이스 커넥션을 유지하도록 하는 기능입니다. 그래서 커넥션은 다른이들이 사용할때 재사용될 수 있습니다.

전형적으로, 데이터베이스의 커넥션을 여는 것은 비싼 동작이다, 특히 데이터베이스가 원격인 경우. 네트워크 세션을 열고, 인증하고, 권한을 확인해야합니다. 풀링은 커넥션이 활성상태를 유지하도록 하고, 커넥션이 나중에 호출되었을때, 활성된 연결 중 하나가 다른 연결을 만드는데 우선적으로 사용되도록 합니다.

이 다이어그램을 봐봅시다. (다이어그램 출처 : https://stackoverflow.com/questions/4041114/what-is-database-pooling)

      +---------+
      |         |
      | Clients |
    +---------+ |
    |         |-+  (1)   +------+   (3)    +----------+
    | Clients | ===#===> | Open | =======> | RealOpen |
    |         |    |     +------+          +----------+
    +---------+    |         ^
                   |         | (2)
                   |     /------\
                   |     | Pool |
                   |     \------/
               (4) |         ^
                   |         | (5)
                   |     +-------+   (6)   +-----------+
                   #===> | Close | ======> | RealClose |
                         +-------+         +-----------+

1. 클라이언트가 커넥션을 여는 API를 호출합니다.  (클라이언트는 커넥션을 풀에서 가져오는지, 실제 연결을 시도하는지 모릅니다.)
2. 커넥션풀에 적합한 연결이 있는지 확인하여 사용 가능한 경우에 Client에 연결합니다.
3. 그렇지 않으면 새로 연결합니다.
4. 유사하게, 클라이언트에서 연결 종료 API를 호출시, 실제로는 DB연결을 종료하는 API를 호출하지는 않고
5. 이 연결이 재사용될 수 있도록 커넥션풀에 반납합니다.
6. 어떤 시점에서는 커넥션 풀의 연결이 실제로 종료될 수 있습니다.

이는 단순한 설명입니다. 실제 구현은 여러 서버 및 여러 사용자 계정에 대한 연결을 처리할 수 있으며 일부 기본 연결을 미리 할당하여 일부는 즉시 사용할 수 있고, 해당 연결에 대해 오래동안 사용을 안할 때는 연결을 닫을 수 있습니다.

# 커넥션풀 설정 값

많은 라이브러리들이 커넥션 풀을 구현하고 있습니다. 각각의 특성이 다르긴 하겠지만, 대부분의 라이브러리들이 가지고 있고 이 설정 값들을 통해서 커넥션 풀을 관리하고 있습니다.

- minimum connections
- maximum connections
- idle connections

# Reference

- 정의에 관한 문서
    - [https://en.wikipedia.org/wiki/Connection_pool](https://en.wikipedia.org/wiki/Connection_pool)
    - [https://stackoverflow.com/questions/4041114/what-is-database-pooling](https://stackoverflow.com/questions/4041114/what-is-database-pooling)
    - [https://en.wikipedia.org/wiki/Object_pool_pattern](https://en.wikipedia.org/wiki/Object_pool_pattern)
- 커넥션풀 라이브러리에 관한 문서
    - [https://commons.apache.org/proper/commons-dbcp/configuration.html](https://commons.apache.org/proper/commons-dbcp/configuration.html)
    - [https://d2.naver.com/helloworld/5102792](https://d2.naver.com/helloworld/5102792)
    - [https://github.com/brettwooldridge/HikariCP](https://github.com/brettwooldridge/HikariCP)
