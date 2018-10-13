# chp18. mutable object

## 무엇이 객체를 변경 가능하게 하는가?
- 객체 구현체를 직접 보지 않더라도 `immutable`과 `mutable` 객체 간의 차이를 관찰할 수 있다.
- `imuutable` 객체 필드에 접근하면 항상 동일한 결과가 나온다.
- `mutable` 객체의 메소드 호출 or 객체필드에 접근한 결과는 이전에 어떤 연산자를 실행했는가에 따라 결과가 다르다.
- 아래 예시를 보면 `bal`은 `mutable`객체이며, `deposit`, `withdraw` 실행 후 값이 변화한다. 

    ```
class BankAccount {
  private var bal: Int = 0

  def balance: Int = bal

  def deposit(amount: Int): Unit = {
    require(amount > 0)
    bal += amount
  }

  def withdraw(amount: Int): Boolean = 
    if (amount > bal)
      false
    else {
      bal -= amount
      true
    }
}

scala> val account = new BankAccount
account: BankAccount = BankAccount@748ac6f3

scala> account deposit 100

scala> account withdraw 80
res1: Boolean = true

scala> account withdraw 80
res2: Boolean = false
    ```

- 변수가 `var`을 포함하더라도 순수 함수일 수도 있다.
- `memoization`기법을 이용해 수행시간이 오래 걸리는 함수를 재정의한다면 반환결과가 동일하다.

```
class Keyed {
  def computeKey: Int = {
    Thread.sleep(1000)
 }
}

class MemoKeyed extends Keyed {
  private var keyCache: Option[Int] = None

  override def computeKey: Int = {
    if (!keyCache.isDefined)
      keyCache = Some(super.computeKey)

    keyCache.get
  }
}
```


## 재할당 가능한 변수와 프로퍼티
- 일반적으로 `vo` 또는 `dto`클래스는 멤버 변수를 `private`로 선언하고 `setter`, `getter` 메소드를 통해 접근한다.
- 두 메소드는 명시적으로 정의해야만 한다.
- `scala`는 객체의 멤버 중 `private`가 아닌 모든 `var`멤버에 두 메소드를 자동으로 정의해준다.
    - `var x`의 `getter`는 `x`, `setter`는 `x_=`이다.
    - 어떤 클래스 내에 `var x` 을 선언하면 이 필드에는 항상 `private[this]`가 붙는다.
    - 이 필드에 접근할 경우 내부적으로는 `var x`의 `getter`는 `x`, `setter`는 `x_=` 인 메소드를 호출하는 것이다.
    - 바이트코드를 보려면 [링크](https://www.toptal.com/scala/scala-bytecode-and-the-jvm)를 참고

    ```
class Time {
  var hour = 12
}

Time.scala로 저장.
$ scalac Time.scala
$ javap -p Time.class

public class Time {
  private int hour;
  public int hour();
  public void hour_$eq(int);
  public Time();
}
    ```

