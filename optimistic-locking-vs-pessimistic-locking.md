# optimistic-locking-vs-pessimistic-locking

[https://www.quora.com/What-is-optimistic-locking-vs-pessimistic-locking](https://www.quora.com/What-is-optimistic-locking-vs-pessimistic-locking) 번역

이러한 방법론은 다중 사용자 문제를 처리하는 데 사용됩니다. 2 명이 동일한 레코드를 동시에 업데이트하려고 한다면 이걸 어떻게 처리할 것인가?

## 1. Do Nothing

    유저1 : 레코드를 읽음
    유저2 : 동일한 레코드를 읽음
    유저1 : 레코드를 업데이트함
    유저2 : 동일한 레코드를 업데이트함

유저2는 유저1이 업데이트한 레코드에 over-write 할 것이다. 유저1이 업데이트 한 것은 없던 것처럼 완전히 사라지게 된다. 이를 lost-update 라고 한다.

## 2. Pessimistic locking

레코드가 읽힐때 레코드에 Lock을 건다.

    유저1 : 레코드를 읽고, 레코드에 exclusive lock을 건다. (FOR UPDATRE 절)
    유저2 : 동일한 레코드를 읽으려고 lock을 시도하지만 유저1이 먼저 locking 했으니 뒤에 대기한다.
    유저1 : 레코드를 업데이트한다. (물론 commit도 한다.)
    유저2 : 레코드를(유저1이 변경한) 읽고 lock을 걸수 있게 된다.
    유저2 : 유저1이 변경한 레코드를 업데이트한다.

lost-update 문제가 해결되었다. 이 접근법의 문제점은 동시성이다. 유저1이 가져온 레코드는 유저1이 업데이트하지 않을수도 있는데, 이 레코드에 lock을 걸고 있다. 유저1과 유저2는 상호 베타적인 lock을 걸길 원하지만, 유저1이 해당 레코드에 대한 lock을 놓을때까지 유저2는 레코드를 읽을 수 없다. 이 방법은 너무 많은 베타적 lock을 필요로 하며 레코드를 locking하고 있는 시간이 너무 길기도 하다. 이 때문에 이 방법은 거의 구현되지 않는다.

## 3. Use Optimistic Locking

Optimistic Locking은 상호 베타적인 lock을 사용하지 않는다. 대신 레코드를 읽은 후 레코드가 변경되지 않았는지 확인하기 위해 업데이트 과정 중 검사가 수행된다. 이 작업은 테이블의 모든 필드를 검사하여 수행할 수 있다.

(예시 쿼리)

    UPDATE Table1 
    SET Col2 = x 
    WHERE COL1=:OldCol1 AND COl2=:OldCol AND Col3=:OldCol3 AND...

물론 이 방법도 몇가지 단점이 있다. 첫째, 테이블의 모든 열을 `SELECT` 해야한다. 둘째, 이 막대한 쿼리를 작성하여 실행해야한다. **대부분의** 사람들은 이를 구현할때 timestamp라는 하나의 열을 통해 구현한다. 이 컬럼은 Optimistic Concurrency를 구현하는 것 **이외의 용도로는 사용되지 않는다.** timestamp 컬럼은 숫자 또는 날짜가 될 수 있다. 이 아이디어는 row가 추가될 때마다 해당 row에 timestamp 값도 같이 추가되는 것이다. 레코드를 읽을 때마다 timestamp 컬럼도 읽는다. 갱신이 수행되면 timestamp컬럼을 확인한다. 만약 해당 row의 timestamp값이 `UPDATE` 시간에 같은 값을 가지고 있다면 동일한 값의 `UPDATE`가 수행되고 timestamp가 변경된다. 만약 `UPDATE` 시간에 timestamp값이 다르다면 사용자에게 오류가 반환된다.즉 레코드를 다시 읽고나서 변경 내용을 작성하여 다시 업데이트를 해야한다.

    유저1 : timestamp값이 21인 레코드를 읽는다.
    유저2 : timestamp값이 21인 레코드를 읽는다.
    유저1 : 레코드 갱신를 시도한다. timestamp 값 21이 데이터베이스에 있는 값과 동일하므로 timestamp값을 22로 업데이트한다.
    유저2 : 레코드 갱신를 시도한다. timestamp값 21이 데이터베이스에 있는 값(22)과 다르므로, 에러를 반환한다.
    유저2 : 레코드를 다시 읽어서 유저1이 변경한 내용에 유저2의 변경 내용을 덮고 갱신한다.

더 많은 의미와 고려사항이 있지만 이정도 연구로 충분하다.

## 추가적으로 읽어봐야할 문서
- [wikipedia-optimistic_concurrency_control](https://en.wikipedia.org/wiki/Optimistic_concurrency_control)
- [wikipedia-record-locking](https://en.wikipedia.org/wiki/Record_locking)
- [stackoverflow-optimistic-vs-pessimistic-locking](https://stackoverflow.com/questions/129329/optimistic-vs-pessimistic-locking)

