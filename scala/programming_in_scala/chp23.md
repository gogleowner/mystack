# chp23. For Expressions Revisited
- `map`, `flatMap`, `filter`을 통해서 리스트를 처리할 수 있는데 추상화를 너무 많이 한 코드는 이해하기 어려울 수 있다.

## 처리하기 어려운 경우에 대한 예시

```
object Exercise1 extends App {
  val lara = Person("Lara", false)
  val bob = Person("Bob", true)
  val julie = Person("Julie", false, lara, bob)
  val persons = List(lara, bob, julie)

  val a = persons.filter(p => !p.isMale)
                 .flatMap(p => p.children.map(c => (p.name, c.name)))

  val b = persons.withFilter(p => !p.isMale)
                 .flatMap(p => p.children.map(c => (p.name, c.name)))

  val c =
    for (p <- persons; if !p.isMale; c <- p.children)
      yield (p.name, c.name)

  println(a)
  println(b)
  println(c)
}

case class Person(name: String,
                  isMale: Boolean,
                  children: Person*)
```
- `map`, `flatMap`.. 을 이용한 구문보다 `for`를 통한 구문이 더 간단하다.

## 23.1 For expressions

### `for ( seq ) yield expr`
- `seq` : generator, definition, filter를 나열한 것
    ```
for (
  p <- persons          // Generator
  n = p.name            // Definition
  if (n startwith "To") // Filter
)
    ```

### Generator
- `pat <- expr`
- `expr`은 보통 리스트
- `pat`은 리스트의 모든 원소를 하나씩 순회한다.
- Generator가 여러개라면 뒤쪽이 앞쪽보다 더 빨리 변환된다.

```
scala> for (x <- List(1, 2);
     |      y <- List("one", "two"))
     | yield (x, y)
res8: List[(Int, String)] = List((1,one), (1,two), (2,one), (2,two))
```

### Definition
- `pat = expr`
- `expr`을 `pat`에 값을 바인딩 한다.

### Filter
- `if expr`
- `expr`은 `Boolean` 타입의 값이다.


## 23.2 The n-queens problem

## 23.3 Querying with for expressions

```
object Exercise3 extends App {
  val books: List[Book] =
    List(
      Book(
        "Structure and Interpretation of Computer Programs",
        "Abelson, Harold", "Sussman, Gerald J."
      ),
      Book(
        "Principles of Compiler Design",
        "Aho, Alfred", "Ullman, Jeffrey"
      ),
      Book(
        "Programming in Modula-2",
        "Wirth, Niklaus"
      ),
      Book(
        "Elements of ML Programming",
        "Ullman, Jeffrey"
      ),
      Book(
        "The Java Language Specification", "Gosling, James",
        "Joy, Bill", "Steele, Guy", "Bracha, Gilad"
      )
    )

  // 작가의 성이 "Gosling"인 모든 책
  val goslingBooks =
    for (book <- books; author <- book.authors; if author.startsWith("Gosling"))
      yield book.title

  println(goslingBooks)

  // 제목에 Program 문자열이 들어간 모든 책의 제목
  val containsProgramBook =
    for (book <- books; title = book.title; if title.contains("Program"))
    yield book.title

  println(containsProgramBook)
  // 이 데이터베이스에서 최소한 두 권 이상의 책을 쓴 작가
  val authorsWithWrittenTwoMoreBooks =
    for (b1 <- books; b2 <- books if b1 != b2; a1 <- b1.authors; a2 <- b2.authors if a1 == a2)
    yield a1

  println(authorsWithWrittenTwoMoreBooks)

  // 중복 제거
  println(removeDuplicates(authorsWithWrittenTwoMoreBooks))

  def removeDuplicates[A](a: List[A]): List[A] =
    if (a.isEmpty) a
    else a.head :: removeDuplicates(a.tail.filter(t => a.head != t))

  def removeDuplicatesWithFor[A](elements: List[A]): List[A] =
    if (elements.isEmpty) elements
    else
      elements.head :: removeDuplicatesWithFor(
        for (t <- elements.tail if t != elements.head) yield t
      )
}

case class Book(title: String, authors: String*)
```

## 23.4 Translation of for expressions
- 모든 `for`문은 `map`, `flataMap`, `withFilter`로 표현 가능하다.
- 스칼라 컴파일러가 사용하는 변환방식에 대해 봐보자.

### Generator가 하나밖에 없는 경우

```
for (x <- expr1) yield expr2
expr1.map(x => expr2)
```

### Generator로 시작하고 Filter가 하나 있는 경우

```
for (x <- expr1 if expr2) yield expr3
for (x <- expr1.withFilter(x => expr2)) yield expr3
expr1.withFilter(x => expr2).map(x => expr3)
```

- Filter 다음에 다른 원소가 더 있다면

```
for (x <- expr1 if expr2; seq) yeild expr3
for (x <- expr1.withFilter(expr2); seq) yeild expr3
```

### Generator 2개로 시작하는 경우

```
for (x <- expr1; y <- expr2; seq) yield expr3
expr1.flatMap(x => for (y <- expr2; seq) yield expr3)
```

- 데이터베이스에서 최소한 두 권 이상의 책을 쓴 작가를 flatMap, withFilter를 이용한 경우

```
books.flatMap(b1 =>
  books.withFilter(b2 => b1 != b2)
       .flatMap(b2 => b1.authors
         .flatMap(a1 => b2.authors.withFilter(a2 => a1 == a2).map(_ => a1))))
       .distinct
```

### Generator에 있는 패턴의 변환
- Generator의 왼쪽 변수가 특정 타입의 변수가 아닌 패턴이라면 변환방법이 복잡해진다.
- 튜플이라면 아래와 같이 구현할 수 있다.

```
for ((num, ch) <- tuples) yield num + ch.toInt
tuples.map{case (num, ch) => num + ch.toInt}
```

- 그러나 임의의 패턴이라면 복잡해진다.

```
for (pat <- expr1) yield expr2

expr1 withFilter {
  case pat => true
  case _ => false
} map {
  case pat => expr2
}
```

### Definition 변환

- Definition을 꼭 for 표현식 안에 선언할 필요가 없다면 밖에다가 쓰는게 더 좋다.
- 왜냐면 iteration 할때마다 계속 Definition 부분을 수행하기 때문이다.

```
for (x <- expr1; y = expr2; seq) yield expr3

// 이것 보다는
for (x <- 1 to 1000; y = expensiveComputationNotInvolvingX) yield x * y

// 이렇게 작성하는 게 좋다.
y = expensiveComputationNotInvolvingX
for (x <- 1 to 1000) yield x * y
```

### for 루프 변환
- 지금까지는 `yield` 를 통해 for 표현식을 어떻게 변환하는지 살펴봤다.
- 일반적인 for 루프문은 `for` 로 작성해도 되고, `foreach`로도 작성이 가능하다. (일반적으로 는 `foreach`를 주로 사용)
```
for (x <- expr1) body
expr1.foreach(x => body)

for (x <- expr1; if expr2; y <- expr3) body
expr1 withFilter(x => expr2) foreach (x => expr3 foreach (y => body))
```

## 23.5 Going the other way
- for 표현식을 map, flatMap, withFilter 고차함수 호출로 변환할 수 있는데, 그 반대도 가능하다.
- map, flatMap, withFilter 도 for 표현식으로 표현할 수 있다.

## 23.6 Generalizing for

### for는 일반적인 컬렉션 객체들에서 사용 가능하다.
- 리스트나 배열이 map, flatap, withFilter 연산을 정의하기 때문에 for 표현식에서도 사용이 가능하다.
- foreach 메소드도 있어서, for 루프도 사용 가능하다.
- 그 뿐 아니라 여러 collection 객체들은 for 표현식을 지원한다.

### 어떤 종류의 컬렉션을 표시하는 파라미터화된 클래스 `C` 가 있다고 가정

```
abstract class C[A] {
  def map[B](f: A => B): C[B]            // A => B로 변환하는 함수를 받아서 C를 생성하는데 C의 타입은 B
  def flatMap[B](f: A => C[B]): C[B]     // A => B타입의 원소를 갖는 C로 매핑하는 함수를 받아서 C를 만들어냄
  def withFilter(p: A => Boolean): C[A]  // A => Boolean 을 반환하는 함수를 받아 동일한 타입의 컬렉션을 반환
  def foreach(b: A => Unit): Unit        // A => Unit 을 받아 Unit 을 반환
}
```

### `withFilter`는 호출이 있으면 매번 새로운 `C`객체를 만든다. 
- `filter`와의 차이
- https://www.scala-lang.org/api/2.12.0/scala/collection/Traversable.html#withFilter(p:A=>Boolean):scala.collection.generic.FilterMonadic[A,Repr]
- c filter p : 새로운 컬렉션을 만든다

    ```
def filter(p: A => Boolean): Repr = filterImpl(p, isFlipped = false)
def filterNot(p: A => Boolean): Repr = filterImpl(p, isFlipped = false)
private[scala] def filterImpl(p: A => Boolean, isFlipped: Boolean): Repr = {
  val b = newBuilder
  for (x <- this)
    if (p(x) != isFlipped) b += x

  b.result
}
    ```
- c withFilter p : 후속 map, flatMap, foreach 및 withFilter 작업의 도메인 만 제한한다.

    ```
def withFilter(p: A => Boolean): FilterMonadic[A, Repr] = new WithFilter(p)
class WithFilter(p: A => Boolean) extends FilterMonadic[A, Repr] {
  def map[B, That](f: A => B)(implicit bf: CanBuildFrom[Repr, B, That]): That = {
    val b = bf(repr)
    for (x <- self)
      if (p(x)) b += f(x)
    b.result
  }
  def flatMap[B, That](f: A => GenTraversableOnce[B])(implicit bf: CanBuildFrom[Repr, B, That]): That = {
    val b = bf(repr)
    for (x <- self)
      if (p(x)) b ++= f(x).seq
    b.result
  }
  def foreach[U](f: A => U): Unit =
    for (x <- self)
      if (p(x)) f(x)
}
    ```

### withFilter를 좀더 자세히 들여다 보자.
- 위에서 보듯이, `filter`는 새로운 컬렉션을 만드는 반면, `withFilter`는 `WithFilter`라는 래퍼 객체를 만든다.
- for 표현식의 변환시 `withFilter` 뒤에는 대부분 다른 메소드들을 호출한다. `map`, `flatMap`, `foreach` 등.
- 만약 컬렉션 내의 객체들이 크다면(또는 길이가 길다면) 중간 객체를 만드는 걸 피해야할 것이다.
- `withFilter` 를 통해서 나중에 처리할 수 있는 걸러진 원소를 `기억하는` 래퍼 객체를 만드는 것이다.

### Monad
- Monad 는 계산과 관계된 다양한 타입을 컬렉션을 포함해 설명할 수 있다.
    - 계산은 상태, I/O, 백트래킹, 트랜잭션 등을 예로 들수 있다.
- 위의 `C` 클래스의 메소드들을 모나드 위에서 표현할 수 있다.
- 예로 든 `C`는 실제로 스칼라에서 제공하는 [scala.collection.genericFilterMonadic](https://www.scala-lang.org/api/2.12.x/scala/collection/generic/FilterMonadic.html) trait이다.

## 23.7 Conclusion
- for 표현식, for 루프가 내부적으로 어떻게 되어있는지 살펴봤다.
- for 표현식에서는 map, flatMap, withFilter 의 고차 메소드들을 호출하여 변환한다.
- for 표현식의 개념이 단순히 컬렉션을 iteration하는 것보다 더 일반적이라는 것, 대부분의 클래스에서 for를 지원하도록 구현이 가능하다는 것을 확인했다.
