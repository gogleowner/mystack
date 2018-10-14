# chp21. Implicit Conversions and Parameters

- 라이브러리 상의 코드는 있는 그대로 사용해야한다. 이 문제를 처리하기 위해 다른 언어들은 여러 요소를 도입했다.
    - `ruby` => module
    - `smalltalk` => 패키지들이 서로 클래스를 추가할 수 있음
    - 매우 강력하지만 위험하기도 하다.
- `scala`는 같은 문제에 대해 `Implicit Conversions and Parameters`를 답으로 제공한다.

## Implicit conversions
- 이벤트를 받을수 있도록 리스너를 등록하고, 리스너에서는 이벤트를 수행하는 `java swing` 코드이다. 반복되는 코드를 아래와 같이 줄일 수 있다.(scala 2.12 미만 버전에서는 오류발생)

```
  val button = new JButton
  button.addActionListener(
    new ActionListener {
      def actionPerformed(event: ActionEvent) = {
        println("pressed!")
      }
    }
  )

  button.addActionListener( // Type mismatch!
    (_: ActionEvent) => println("pressed!")
  )
```

- 함수를 액션 리스너로 바꾸는 암시적 변환 코드

```
  implicit def function2ActionListener(f: ActionEvent => Unit) =
    new ActionListener {
      def actionPerformed(event: ActionEvent) = f(event)
    }


  button.addActionListener(
    function2ActionListener(
      (_: ActionEvent) => println("pressed!")
    )
  )

  button.addActionListener(
    (_: ActionEvent) => println("pressed!")
  )
```

## Rules for implicits

### 표시 규칙
- `implicit` 키워드는 컴파일러가 암시 처리시 사용할 선언을 표시한다.

### 스코프 규칙
- 삽입된 `implicit` 변환은 스코프 내에 단일 식별자로만 존재하거나 변환의 결과나 원래 타입과 연관있어야한다.

### 한번에 하나만 규칙
- 오직 하나의 암시적 선언만 사용한다.

### 명시적 우선 규칙
- 코드가 그 상태 그대로 컴파일된다면 암시를 통한 변환을 시도하지 않는다.

### 암시적 변환 이름 붙이기
- 아무 이름이나 붙일 수 있는데, 메소드 호출시 명시적으로 변환을 하고 싶은 경우, 프로그램의 특정 지점에서 사용 가능한 변환이 어떤 것이 있는지 파악해야할 경우.

### 암시가 쓰이는 부분
1. 값을 컴파일러가 원하는 타입으로 변환
2. 선택의 수신 객체를 변환
3. 암시적 파라미터를 지정

## Implicit conversion to an expected type

```
scala> implicit def doubleToInt(x: Double) = x.toInt
<console>:11: warning: implicit conversion method doubleToInt should be enabled
by making the implicit value scala.language.implicitConversions visible.
This can be achieved by adding the import clause 'import scala.language.implicitConversions'
or by setting the compiler option -language:implicitConversions.
See the Scaladoc for value scala.language.implicitConversions for a discussion
why the feature should be explicitly enabled.
       implicit def doubleToInt(x: Double) = x.toInt
                    ^
doubleToInt: (x: Double)Int

scala> val num: Int = 3.5
num: Int = 3
```
- `Int` 가 필요한 곳에서 `Double`을 보면 컴파일 오류가 발생해야하는데 그 전에 `implicit` 가 있는지 찾고 있으면 호출한다.


## Converting the receiver
- 메소드를 호출하는 대상이 되는 객체인 수신 객체에도 적용할 수 있다.
- 수신 객체 변환을 통해 새 클래스를 기존 클래스 계층구조에 통합
- 언어 안에서 `DSL`을 만드는 일에 지원

### 새 타입과 함께 통합

```
class Rational(n: Int, d: Int) {
  def + (that: Rational): Rational = ...
  def + (that: Int): Rational = ...
}

implicit def intToRational(x: Int) = new Rational(x, 1)
```

### 새로운 문법 흉내
- `Map`을 표현할때의 `->`는 사실 문법이 아니다.
- [`scala.Predef#ArrowAssoc`](https://www.scala-lang.org/api/current/scala/Predef$.html#ArrowAssoc[A]extendsAnyVal) 클래스의 메소드이다.
- `implicit` 를 이용하여 구현되어있다.

### 암시적클래스

```
case class Rectangle(width: Int, height: Int)

implicit class RectangleMaker(width: Int) {
  def x(height: Int) = Rectangle(width, height)
}
```

## Implicit parameters
- `someCall(a)` 가 `someCall(a)(b)`가 되게 하려면 컴파일러가 판단할 수 있도록 해야한다.
- `(b)`를 호출할 수 있도록 `implicit` 로 표시, `someCall` 정의의 마지막 파라미터도 `implicit`로 표시해야 한다.

```
class PreferredPrompt(val preference: String)

object Greeter {
  def greet(name: String)(implicit prompt: PreferredPrompt) = {
    println(s"Welcome $name. The system is ready")
    println(prompt.preference)
  }
}

scala> val prompt = new PreferredPrompt("relax> ")
prompt: PreferredPrompt = PreferredPrompt@512abf25

scala> Greeter.greet("gogleowner")(prompt)
Welcome gogleowner. The system is ready
relax>
```

- 컴파일러가 파라미터를 암시적으로 제공하게 하려면 필요한 `implicit` 변수 타입을 만들고 스코프에 속하도록 `import` 하도록 한다.

```
object JoesPrefs {
  implicit val prompt = new PreferredPrompt("Yes, master> ")
}

scala> import JoesPrefs._
import JoesPrefs._

scala> Greater.greet("gogleowner")
Welcome gogleowner. The system is ready
Yes, master>
```
- 위와 같이 암시적 파라미터를 구현한 `object`을 `import`함으로써 마지막 파라미터를 안줘도 실행할 수 있도록 한다.


## Context bounds
- 파라미터에 대해 `implicit`를 사용하는 경우 컴파일러는 그 파라미터에 암시적 값을 제공하려고 시도할 뿐 아니라 메소드 본문 안에서 사용 가능한 `implicit`를 사용한다.
- 그러나 `implicit` 파라미터가 타입이 맞지 않다면 또다른 암시적 매개변수가 있는지 살펴볼 것이다.

```
def maxList[T](elements: List[T])
      (implicit ordering: Ordering[T]): T =

  elements match {
    case List() =>
      throw new IllegalArgumentException("empty list!")
    case List(x) => x
    case x :: rest =>
      val maxRest = maxList(rest)     // (ordering) is implicit
      if (ordering.gt(x, maxRest)) x  // this ordering is
      else maxRest                    // still explicit
  }
```

- 이 경우 파라미터 `ordering`를 `implicitly`로 없앨 수 있다. => `def implicitly[T](implicit t: T) = T`
- 컴파일러는 `T`의 암시적 정의를 살펴볼 것이다.

```
if (ordering.gt(x, maxRest)) x => if (implicitly[Ordering[T]].gt(x, maxRest)) x
```
- `scala`는 파라미터의 이름을 없애고 메소드 헤더를 `Context Bound`를 사용하여 더 짧게 정의할 수 있도록 해준다.

```
(변경전) def maxList[T](elements: List[T])(implicit ordering: Ordering[T]): T =
(변경후) def maxList[T : Ordering](elements: List[T]): T =
```


## When multiple conversions apply
- 적용 가능한 암시적 변환이 스코프 안에 여러개 있을 수 있다. 이 경우 스칼라는 변환을 추가하지 않는다.
- 여러 변환을 사용하는 경우 어떤 것을 선택할지는 명확하지 않다.

```
def printLength(seq: Seq[Int]) = println(seq.length)
implicit def intToRange(i: Int) = 1 to i
implicit def intToDigits(i: Int) = i.toString.toList.map(_.toInt)

scala> printLength(12)
<console>:19: error: type mismatch;
 found   : Int(12)
 required: Seq[Int]
Note that implicit conversions are not applicable because they are ambiguous:
 both method intToRange of type (i: Int)scala.collection.immutable.Range.Inclusive
 and method intToDigits of type (i: Int)List[Int]
 are possible conversion functions from Int(12) to Seq[Int]
       printLength(12)
                   ^
```
- `v2.7`까지는 컴파일 에러가 발생하나,
- `v2.8`부터는 이 규칙을 완화하여 가능한 변환중 하나가 더 다른 하나보다 더 구체적일때 컴파일러는 더 구체적인것을 선택한다.
    - 전자의 인자 타입이 후자의 서브타입
    - 두 변환 모두 메소드인데 전자를 둘러싼 클래스가 후자를 둘러싼 클래스를 확장한다.

## Debugging implicits
- 암시적인 부분을 디버깅할 때 도움이 될 힌트를 제공한다.
- 런타임에서 컴파일러가 암시를 찾지 못하는 경우, 명시적으로 작성후 실행할 때 어느 부분에서 오류인지 잘 알려준다.
- 컴파일러가 암시적 변환을 어떻게 적용하는지 확인
    - `$ scalac -Xprint:typer`

## Conclusion
- 암시는 강력하고, 코드를 압축할 수 있도록 해주는 기능이 있다.
- 암시를 너무 자주 사용하면 코드 이해가 어려워질 수 있다. 암시적 변환을 추가하기 전 `extends, mixin, method overload` 등 다른 방법을 먼저 자문해보아야 한다.
