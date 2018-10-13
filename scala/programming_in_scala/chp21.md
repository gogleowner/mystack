# chp21. Implicit Conversions and Parameters

- 라이브러리 상의 코드는 있는 그대로 사용해야한다. 이 문제를 처리하기 위해 다른 언어들은 여러 요소를 도입했다.
    - `ruby` => module
    - `smalltalk` => 패키지들이 서로 클래스를 추가할 수 있음

- 매우 강력하지만 위험하기도 하다.
- `scala`는 같은 문제에 대해 `Implicit Conversions and Parameters`를 답으로 제공한다.

## Implicit conversions
- 이벤트를 받는 `java swing` 코드이다. 반복되는 코드를 아래와 같이 줄일 수 있다.(scala 2.12 미만 버전에서는 오류발생)

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



## Context bounds
## When multiple conversions apply
## Debugging implicits
## Conclusion

