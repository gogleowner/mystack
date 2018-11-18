# chp22. Implementing List

스칼라에서 리스트를 어떻게 구현하는지 살펴보자.

## 22.1 The List class in principle

```
package scala.collection.immutable.List
sealed abstract class List[+A] 
  extends AbstractSeq[A] 
  with LinearSeq[A] 
  with Product 
  with GenericTraversableTemplate[A, List] 
  with LinearSeqOptimized[A, List[A]] 
  with Serializable

final case class ::[B](head: B, tl: List[B]) 
  extends List[B] 
  with Product 
  with Serializable

object Nil 
  extends List[Nothing] 
  with Product 
  with Serializable
```

- 리스트는 추상클래스이다.
    - 따라서 `new List`로 선언언할 수 없다.
    - `List[+A]` 리스트가 공변성임을 의미한다.
    - 이로 인해 `List[Int]` 타입의 값을 `List[Any]`같은 타입의 변수에 할당할 수 있다.

    ```
scala> val numList = (1 to 3).toList
numList: List[Int] = List(1, 2, 3)

scala> var numList2: List[Any] = numList
numList2: List[Any] = List(1, 2, 3)
    ```

- 서브클래스는 `::`, `Nil`로 이루어져있다.

### Nil
- 모든 List 타입과 서로 호환이 가능하다.
- `List`의 `head`, `isEmpty`, `tail` 메소드들을 재구현하여 빈값을 반환한다.

```
case object Nil extends List[Nothing] {
  override def isEmpty = true
  override def head: Nothing =
    throw new NoSuchElementException("head of empty list")
  override def tail: List[Nothing] =
    throw new UnsupportedOperationException("tail of empty list")
  // Removal of equals method here might lead to an infinite recursion similar to IntMap.equals.
  override def equals(that: Any) = that match {
    case that1: scala.collection.GenSeq[_] => that1.isEmpty
    case _ => false
  }
}
```

### ::

```
final case class ::[B](override val head: B, private[scala] var tl: List[B]) 
  extends List[B] {

  override def tail: List[B] = tl
  override def isEmpty: Boolean = false
}
```

- 파라미터 `head` 와 `tl`은 각각 머리와 꼬리이다. `case class`의 특성상 파라미터들은 해당 클래스의 멤버변수이다.

```
sealed abstract class List[+A] extends AbstractSeq[A]
                                  with LinearSeq[A]
                                  with Product
                                  with GenericTraversableTemplate[A, List]
                                  with LinearSeqOptimized[A, List[A]]
                                  with scala.Serializable {

  def ::[B >: A] (x: B): List[B] =
    new scala.collection.immutable.::(x, this)

  def :::[B >: A](prefix: List[B]): List[B] =
    if (isEmpty) prefix
    else if (prefix.isEmpty) this
    else (new ListBuffer[B] ++= prefix).prependToList(this)

  // 이하 생략
}
```

- `::` 메소드
  - 이 메소드는 타입 파라미터 `B`를 받는 다형성 메소드이다.
  - `B`는 `A`의 슈퍼타입이어야한다는 제약이 걸려있다.

- `:::` 메소드
  - 리스트끼리 합치는 메소드이다.

## 22.2 The ListBuffer class
- `scala.collection.mutable.ListBuffer` 
- 리스트에 원소를 축적할 수 있다. 
- 원소 추가가 다 끝나고 나면 `toList`을 통해 리스트로 변환할 수 있다.



```
List의 모든 원소를 +1한 리스트를 변환하는 경우

scala> import scala.collection.mutable.ListBuffer
import scala.collection.mutable.ListBuffer

scala> val buf = new ListBuffer[Int]
buf: scala.collection.mutable.ListBuffer[Int] = ListBuffer()

scala> val numbers = (1 to 3).toList
numbers: List[Int] = List(1, 2, 3)

scala> for (x <- numbers) buf += x + 1

scala> buf.toList
res7: List[Int] = List(2, 3, 4)
```

## 22.3 The List class in practice
- 위에서 보듯, 리스트는 `map` 을 통해 원소를 변환하는 경우 `ListBuffer`를 활용하고 있다.
- 음.. 좀더 고민해봐야할것들이 있는 부분인데, 일단 사용하는데는 문제없으니 패스!

## 22.4 Functional on the outside
- `List`가 외부에서 볼때는 완전히 함수적이지만 내부에서는 `ListBuffer`를 통해 명령형으로 작성됨을 볼 수 있었다.
- `immutable`하도록 유지하기 위해서 그렇게 작성한 것이고, `mutable`하다면 프로그램이 깨지기 쉽기에 이부분을 깨고자 내부는 명령형으로 구현했다.
- `List` - `ListBuffer`의 관계는 `java`의 `String` - `StringBuffer`와 하는 일이 비슷하다.
- 두 경우 모두 `immutable`한 데이터구조를 원했고 점진적으로 구축할 수 있는 효율적인 방법을 제공하고 싶어했다.
