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
expr1.flattMap(x => for (y <- expr2; seq>) yield expr3)
```

- 데이터베이스에서 최소한 두 권 이상의 책을 쓴 작가를 flatMap, withFilter를 이용한 경우

```
books.flatMap(b1 =>
    books.withFilter(b2 => b1 != b2)
         .flatMap(b2 => b1.authors
                          .flatMap(a1 => b2.authors.withFilter(a2 => a1 == a2).map(_ => a1))))
       .distinct
```

## 23.5 Going the other way
## 23.6 Generalizing for
## 23.7 Conclusion

