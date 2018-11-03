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

  override def tail : List[B] = tl
  override def isEmpty: Boolean = false
}
```


## 22.2 The ListBuffer class
## 22.3 The List class in practice
## 22.4 Functional on the outside
## 22.5 Conclusion
