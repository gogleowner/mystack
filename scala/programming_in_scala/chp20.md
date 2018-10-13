# chp20. Abstract Members

- `Abstract Member` : 클래스나 트레이트의 멤버가 그 클래스 안에 완전한 정의를 갖고 있지 않음 
- `java`는 메소드만 추상으로 선언할 수 있었지만, `scala`는 클래스내의 멤버변수도 추상필드로 선언할 수 있다.
- `val`, `var`, `method`, `type`, `pre-initialized field`, `lazy val`, `path-dependent type`, `enumeration`

## Type Members
- 클래스나 트레이트의 멤버. 정의 없이 선언만 된 타입
- `type` 사용 이유 : 실제 이름이 너무 길거나 의미가 불명확할 때 더 간단하고 의도를 잘 전달할 수 있는 별명을 선언하는 것.

```
trait Abstract {
  type T
  def transform(x: T): T
  val initial: T
  var current: T
}

class Concrete extends Abstract {
  type T = String
  def transform(x: String) = x + x
  val initial = "hi"
  var current = initial
}
```

## Abstract `val`s
- `val`에 대해 이름과 타입은 주지만 값은 지정하지 않는다. 값은 서브클래스에서 정의해야한다.
- `val initial: String`
- `val` vs `def`
    - 공통점 : 외부에서 `initial`을 호출할때는 메소드 호출 방식과 동일하게 작성할 것이다.
    - 차이 : `val`이라면 항상 동일한 값을 얻을 수 있을 것인데, `def`이라면 같은 값임을 보장할 수 없다.
    - 의도 : 서브클래스 작성자에게 `initial`을 `val`로 구현해야하는 제약을 가한다.

```
  abstract class Fruit {
    val v: String // `v' for value
    def m: String // `m' for method
  }

  abstract class Apple extends Fruit {
    val v: String
    val m: String // OK to override a `def' with a `val'
  }

  abstract class BadApple extends Fruit {
    def v: String // ERROR: cannot override a `val' with a `def'
    def m: String
  }
```

## Abstract `var`s
- `var`는 컴파일시 `getter`, `setter`메소드를 생성한다.

## Initializing abstract `val`s

```
trait RationalTrait {
  val numerArg: Int
  val denomArg: Int
  require(denomArg != 0)
  private val g = gcd(numerArg, denomArg)
  val numer = numerArg / g
  val denom = denomArg / g
  private def gcd(a: Int, b: Int): Int =
  if (b == 0) a else gcd(b, a % b)
  override def toString = numer + "/" + denom
}
```

- 위 `trait`은 `numerArg`, `denomArg`가 추상필드로 선언되어있어 초기화시에 값을 할당해야 한다.
- 이때 `g`의 경우 `numerArg`, `denomArg`를 활용하게 되어있는데 초기 값은 `0`이다. 때문에 초기화시 `require`에서 컴파일 오류가 발생한다.

```
scala> new RationalTrait {
     |   val numerArg = 1
     |   val denomArg = 2
     | }
java.lang.IllegalArgumentException: requirement failed
  at scala.Predef$.require(Predef.scala:264)
  at RationalTrait.$init$(<console>:14)
  ... 32 elided
```

### 필드를 미리 초기화하기
- 슈퍼클래스를 호출하기 전 서브 클래스의 필드를 미리 초기화하는 방법이다.
- 초기화와 슈퍼 트레이트 사이는 `with`로 구분한다.

```
// 익명 클래스 표현식에서 필드를 미리 초기화
scala> new {
     |   val numerArg = 1
     |   val denomArg = 2
     | } with RationalTrait
res8: RationalTrait = 1/2


// 객체 정의에서 필드를 미리 초기화
scala> object Sample extends {
     |   val numerArg = 2
     |   val denomArg = 3
     | } with RationalTrait
defined object Sample

// 클래스 정의에서 필드를 미리 초기화
class RationalClass(n: Int, d: Int) extends {
  val numerArg = n
  val denomArg = d
} with RationalTrait {
  
}
```

### lazy val
- `lazy val` : `val`변수 정의시 오른쪽의 초기화 표현식을 계산하지 않는다. 프로그램에서 `val`의 값을 처음 사용할때 계산한다.
- 이점 : `lazy val` 초기화에 부수효과가 없거나 다른 부수 효과에 의존하지 않는 경우에만 존재함.

```
trait LazyRationalTrait { 
  val numerArg: Int 
  val denomArg: Int 
  lazy val numer = numerArg / g
  lazy val denom = denomArg / g

  override def toString = numer + "/" + denom

  private lazy val g = {
    require(denomArg != 0)
    gcd(numerArg, denomArg)
  }

  private def gcd(a: Int, b: Int): Int = 
    if (b == 0) a else gcd(b, a % b)
}
```

### Abstract `type`s

- 클래스 관계가 계층 구조를 가진다면 상위바운드를 `type`으로 쓰도록 한다.

```
class Food
abstract class Animal {
  type SuitableFood <: Food

  def eat(food: SuitableFood)
}
```

### Path-dependent types

- `java`의 `inner class`와 유사함..
- 내부 클래스 인스턴스는 외부 클래스의 인스턴스를 가리키는 참조가 들어있어서 내부에서는 외부의 멤버에 접근할 수 있다.

```
class Outer {
  class Inner
}

val o1 = new Outer
val o2 = new Outer
```

### Refinement types
- `class A extends B` => `A`가 `B`의 `이름에 의한 서브타입(nominal subtype)`
    - 타입에 이름이 있고, 서브타입 관계를 선언하면서 각 클래스의 이름을 명시하기 때문
- `scala`는 `구조적인 서브 타이핑(structual subtyping)`을 지원한다. 이는 `세분화한 타입(refinement type)`을 이용하면 된다.

```
풀을 먹는 동물을 포함할 수 있는 목초지(Pasture)

trait AnimalThatEatsGrass => 장황함.

풀을 먹는 동물

Animal { type SuitableFood = Grass }

목초지

class Pasture {
  var animals: List[ Animal { type SuitableFood = Grass } ] = Nil
}
```

### Enumerations

- `java`, `C#`과는 다르게 `scala`에서는 `Enumration`을 위한 특별한 문법은 존재하지 않다.
- 열거형을 구현하려면 [`scala.Enumration`](https://www.scala-lang.org/api/current/scala/Enumeration.html)을 활용한다.

```
object Color extends Enumeration {
  val Red, Green, Blue = Value
}
```

- [`Enumeration#Value`](https://www.scala-lang.org/api/current/scala/Enumeration$Value.html)를 이용하여 여러가지 연산을 할 수 있다.
    - `val Red = Value("Red")` 등등..
    - `iteration`은 `values`를 이용한다.

