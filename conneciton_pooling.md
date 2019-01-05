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

# 커넥션풀 설정 값

많은 라이브러리들이 커넥션 풀을 구현하고 있다. 각각의 특성이 다르긴 하겠지만, 대부분의 라이브러리들이 가지고 있고 이 설정 값들을 통해서 커넥션 풀을 관리하고 있다.

- minimum connections
- maximum connections
- idle connections

바람직한 설정값은..!

- maxActive >= initialSize
- maxIdle >= minIdle
- maxActive = maxIdle
- 동일한 값으로 통일해도 무방하다.

## TPS와 커넥션 개수와의 관계

- Commons DBCP의 maxActive 값과 Tomcat의 maxThread 값을 잘 조절해야함.
- maxActive 가 너무 많아도 DBMS가 부담을 가질 수 있고, maxActive 값을 줄여서 maxWait를 늘리면 대기하는 쓰레드가 많아질 경우 Tomcat의 maxThread 개수가 가득 찰수도 있다. 때문에 적절한 개수로 설정해야하며, 그래도 한 장비로 버티기 어렵다면 scale out 을 해야한다.

## 커넥션 검사와 정리

- testOnBorrow, testOnReturn → 커넥션을 얻을 때 / 반환할 때 validationQuery를 수행할 것인지에 대한 여부이다. 성능이 중요한 서비스라면 두 값 모두 false로 하도록 한다.

## Evictor 쓰레드와 관련된 속성

- Evictor 쓰레드는 커넥션 자원을 정리하는 구성요소이며 별도의 쓰레드로 수행된다.
- 커넥션풀 내의 유휴 상태 커넥션 중 오랫동안 사용되지 않은 커넥션을 추출하여 제거한다.
- 커넥션에 대해 유효성 검사를 수행하여 문제가 있을 경우 해당 커넥션을 제거한다.
- 커넥션 개수가 minIdle 속성값보다 적으면 minIdle 속성값만큼 커넥션을 생성한다.
- Evictor 쓰레드는 동작시에 커넥션풀에 lock을 걸고 동작하기 때문에 너무 자주 실행하면 서비스에 부담을 줄 수 있다.
- 한번에 검사할 커넥션 개수를 크게 설정하면 한번에 Evictor 쓰레드가 동작하는 시간이 길어져 lock 상태가 길어지므로 서비스에 부담을 줄 수 있다.
- IDC 정책에 따라 소켓 연결 후 일정시간 아무 패킷도 주고받지 않으면 연결이 끊기는 현상이 있을 수 있는데 timeBetweenEvictionRunsMillis 속성으로 의도치않게 연결이 끊기는 현상을 막을 수도 있다.

## statement pooling 관련 옵션

- 풀링할 Prepared Statement 개수를 적절하게 설정.

# WAS와 DBMS 통신시의 타임아웃 계층

- TransactionTimeout
    - 프레임워크, 어플리케이션 레벨에서 발생할수 있는 타임아웃이다.
    - 전체 Statement 수행시간을 허용할 수 있는 최대시간 이내로 제한하려할 때 사용한다.
    - 스프링은 트랜젝션의 시작, 경과시간을 기록하면서 특정 시간 초과시에 예외를 발생하도록 한다.
    - `@Transactional(timeout=10)`
- StatementTimeout
    - Statement 하나가 얼마나 오래 수행되어도 괜찮은지에 대한 한계 값
    - 정상적으로 소켓 연결을 맺고 있을 때만 유효함
    - `java.sql.Statement.setQueryTimeout`
- JDBC Driver SocketTimeout
    - DBMS가 비정상 종료되었거나 네트워크 장애가 발생시 발생
        - OS의 SocketTimeout 설정에 영향을 받는다.
    - SocketTimeout > StatementTimeout 으로 잡아야 Statement가 안정적으로 수행될 수 있음
    - Socket Connect 시 타임아웃(connectTimeout): `Socket.connect(SocketAddress endpoint, int timeout)` 메서드를 위한 제한 시간
    - Socket Read/Write의 타임아웃(socketTimeout): `Socket.setSoTimeout(int timeout)` 메서드를 위한 제한 시간


# Reference

- 정의에 관한 문서
    - [https://en.wikipedia.org/wiki/Connection_pool](https://en.wikipedia.org/wiki/Connection_pool)
    - [https://stackoverflow.com/questions/4041114/what-is-database-pooling](https://stackoverflow.com/questions/4041114/what-is-database-pooling)
    - [https://en.wikipedia.org/wiki/Object_pool_pattern](https://en.wikipedia.org/wiki/Object_pool_pattern)
- 커넥션풀 라이브러리에 관한 문서
    - [https://commons.apache.org/proper/commons-dbcp/configuration.html](https://commons.apache.org/proper/commons-dbcp/configuration.html)
    - [https://d2.naver.com/helloworld/5102792](https://d2.naver.com/helloworld/5102792)
    - [https://github.com/brettwooldridge/HikariCP](https://github.com/brettwooldridge/HikariCP)
    - [http://sjh836.tistory.com/148](http://sjh836.tistory.com/148)
    - [https://d2.naver.com/helloworld/1321](https://d2.naver.com/helloworld/1321)
