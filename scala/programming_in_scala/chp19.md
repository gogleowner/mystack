# chp19. Type Parameterization

- `Type Parameterization`을 사용하면 제네릭과 트레이트를 쓸수 있다.
- `java`의 제네릭과 다른 점
    - `java` 는 타입파라미터 생략이 가능하다. (`IDE`에서는 경고를 주기도 한다.)
    - `scala` 는 타입파라미터가 없으면 컴파일 에러가 발생한다.
- 이번 챕터에서는 함수적인 큐 데이터구조를 만들고, 내부 표현을 감추는 기법을 보여준 후, 타입 파라미터의 변성과 정보은닉의 상호작용을 설명한다.

## Functional Queue
- `head`: 큐의 첫번째 원소를 반환한다.
- `tail`: 큐의 첫번째 원소를 제외한 나머지를 반환한다.
- `enqueue`: 원소를 큐의 맨 뒤에 추가한 새로운 큐를 반환한다.

- 변경 가능한 큐와 달리 함수형 큐는 원소를 추가해도 내용을 바꾸지 않고 새로운 큐를 반환한다.
- 이상적인 것은 위 세개의 메소드들의 수행시간이 `O(1)`이 되야 하는 것이다.

```
class Queue[T](
  private val leading: List[T],
  private val trailing: List[T]
) {

  // `mirror` 연산속도는 `List`의 원소의 개수에 대비한다.
  private def mirror = 
    if (leading.isEmpty) new Queue(trailing.reverse, Nil)
    else this

  def head = mirror.leading.head

  def tail = {
    val q = mirror
    new Queue(q.leading.tail, q.trailing)
  }

  def enqueue(x: T) = new Queue(leading, x :: trailing)

  override def toString = s"leading -> ${leading.mkString(",")} / trailing -> ${trailing.mkString(",")}"
}
```

## 정보 은닉
- 위의 `Queue`클래스는 효율성 측면에서는 좋으나 불필요하게 내부 구현이 많이 노출되어있고, 초기화할때마다 `trailing`을 위해서 `Nil`을 넘겨야 한다.
- 클라이언트 입장에서 사용하기 편하도록 코드를 감출 방법을 알아보도록 하자.

### 비공개 생성자, 팩토리 메소드
- 클래스이름과 파라미터 사이의 `private`은 생성자를 `private`으로 만든다.

```
class Queue[T] private (
  private val leading: List[T],
  private val trailing: List[T]
)
```

- 외부에서 접근하여 생성할 수 있도록 생성자를 추가하자.

```
def this(elems: T*) = this(elems.toList, Nil)
```

- 다른 방법으로는 싱글톤 `object`를 이용하여 초기화를 위한 팩토리 메소드를 추가한다.

```
object Queue {
  def apply[T](elems: T*) = new Queue[T](xs.toList, Nil)
}
```

### 비공개 클래스

- 클래스 자체를 감추고 클래스에 대한 인터페이스는 `trait`으로 노출하는 방법이다.

```
trait Queue[T] {
  def head: T
  def tail: Queue[T]
  def enqueue(elem: T): Queue[T]
}

object Queue {
  def apply[T](elems: T*) = new SimpleQueue[T](xs.toList, Nil)

  prviate class SimpleQueue[T] (
    private val leading: List[T],
    private val trailing: List[T]
  ) extends Queue {
    override def head: T = (...)
    override def tail: Queue[T] = (...)
    override enqueue: Queue[T] = (...)
  }
}
```


## Variance annotations
- 위에서 정의한 `Queue`는 `trait`이다. `Queue`라는 타입의 변수를 만들수 없다.

```
scala> def doesNotCompile(q: Queue) = {}
<console>:12: error: trait Queue takes type parameters
       def doesNotCompile(q: Queue) = {}
                             ^
```

- `Queue`는 파라미터화된 타입을 지정해야한다.

```
scala> def doesCompile(q: Queue[AnyRef]) = {}
doesCompile: (q: Queue[AnyRef])Unit
```

- `Queue` => `trait`, `type constructor` 이고, `Queue[String]` => `type` 이다.
- `S`가 `T`의 서브타입이라면 `Queue[S]`를 `Queue[T]`의 서브 타입으로 간주할 수 있다. 이 경우 `Queue`는 타입 파라미터 `T`에 대해 공변적(covariant, flexable)이다고 한다.

### 표기방법
- `+` :  서브타입 관계가 파라미터에 대해 유연하다는 뜻

```
trait Queue[+T] {...}
```

- `-` : 반공변(contravariant) 서브타입 관계.
    - `T`가 `S`의 서브타입인 경우 `Queue[S]`가 `Queue[T]`의 서브타입이라는 뜻.


### 변성과 배열
- 위 경우를 `java`로 작성하여 실행하면 런타임시 오류가 발생한다.

```
String[] a1 = { "abc" };
Object[] a2 = a1;
a2[0] = new Integer(17);
String s = a1[0];

=> java.lang.ArrayStoreException: java.lang.Integer
```

- `java`는 실행시점에 원소의 타입을 저장한다. 원소를 변경할때 마다 새 원소 값을 배열에 저장된 원소의 타입과 비교한다.
- `jdk 1.5`버전이 나오기 전까지는 전역적인 타입은 `Object`를 받도록 되어있었다. 
    - `void sort(Object[] a, Comparator cmp)`
- `scala`는 컴파일 시점에 에러가 발생한다. 실수를 줄일 수 있는 듯 하다.


## 변성 표기 검사
- 컴파일러는 클래스나 트래이트 본문의 모든 위치를 `긍정적, 부정적, 중립적`으로 구분한다.
- ~~이해가 잘 안된다 ㅠ~~

## 하위 바운드
- `Queue[T]`를 정의해서 `T`를 공변적으로 만들 수 없다.
- `T`는 `enqueue` 메소드의 파라미터 타입인데 위치가 부정적임.
- 이를 해결하는 방법은 `enqueue`메소드에 타입 파라미터를 지정하여 타입파라미터에 하위바운드를 사용도록 한다.

```
class Queue[+T](
  private val leading: List[T],
  private val trailing: List[T]
) {
  def enqueue[U >: T](elem: U) = new Queue[U](leading, x :: trailing)
}
```

- `U >: T` : `T`를 `U`의 하위바운드로 지정한다. `U`는 `T`의 슈퍼타입이어야만 한다.


## 반공변성
- ~~이해가 잘 안된다 ㅠ~~

## 상위 바운드
- `T <: U ` : `T`는 `U`의 서브타입이어야한다.

```
class Fruit {
  def daldal[T <: Fruit](fruit: T): Unit = {
    println(s"blabla$fruit")
  }
}
class Apple extends Fruit
class Grape extends Fruit
```

## 객체의 비공개 데이터
- 앞서 본 `Queue`클래스의 `leading, trailing`변수를 새로운 큐로 반환하는 방식이 아닌 돌려 쓰는 것으로 변경
- `enqueue` 메소드에 하위 바운드 추가
- 비공개 `var` 변수 표현시 `[this]` 추가

```
class Queue[+T] private (
   private[this] var leading: List[T], 
   private[this] var trailing: List[T]
) {

  private def mirror() = 
    if (leading.isEmpty) {
      while (!trailing.isEmpty) {
        leading = trailing.head :: leading
        trailing = trailing.tail
      }
    }

  def head: T = { 
    mirror()
    leading.head 
  }

  def tail: Queue[T] = { 
    mirror()
    new Queue(leading.tail, trailing) 
  }

  def enqueue[U >: T](x: U) = 
    new Queue[U](leading, x :: trailing)
}
```
