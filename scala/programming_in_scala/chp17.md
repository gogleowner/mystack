# chp17. collection

## Sequence
- 순서가 정해진 콜렉션

### List
- `List`는 앞절 챕터 16에서 다뤘듯, `LinkedList`의 특성을 가진다.
- 왜 `LinkedList`로 구현되어있는지에 대해서는 아래 스택오버플로우 페이지를 참고
- https://stackoverflow.com/questions/5130097/why-are-scalas-lists-implemented-as-linked-lists

### Array
- `Array`는 원소의 시퀀스를 저장하며, 임의의 위치에 있는 원소에 효율적으로 접근하게 해준다.
- `List` => 재귀적, `Array` => 평면적.

### ListBuffer
- `List`는 `head`에 대해서는 빠른 접근을 제공하나, `tail`의 앞쪽은 그렇지 않다. 끝부분에 원소를 추가하려면 앞에 원소를 추가해서 `reverse` 를 통해 원하는 순서의 `List`를 얻어야 한다.
- `scala.collection.mutable.ListBuffer`는 `mutable`하다.
- 원소를 추가시의 시간복잡도는 `O(1)`이다.
- 재귀형으로 구현하는 경우가 아니면 `for`,  `while` 루프를 사용하여 구현이 가능하다.

```
scala> buf += 1
res60: buf.type = ListBuffer(1)

scala> buf += 2
res61: buf.type = ListBuffer(1, 2)

scala> 0 +=: buf
res63: buf.type = ListBuffer(0, 1, 2)
```

### ArrayBuffer
- `ArrayBuffer`는 시작, 끝 부분에 원소 추가/삭제가 가능한 배열이다.
- `Array`와 `ArrayBuffer`의 비교는 아래 스택오버플로우에 잘 명시되어있음..
- https://stackoverflow.com/questions/31213733/what-is-the-difference-between-arraybuffer-and-array-in-scala

### 문자열(through `StringOps`)
- `StringOps`에는 많은 시퀀스 메소드들이 구현되어있다.
- `scala.Predef.String`에 제공하는 래퍼역할을 한다.
- 예를들어, `String`에는 정의되어있지 않은 `exist()` 메소드는 스칼라 컴파일러가 `StringOps`로 변환하는 것이다.

```
def hasUpperCase(s: String) = s.exists(_.isUpper)
```

## set, map
- `Set`의 핵심 특징은 특정 객체는 최대 하나만 들어가도록 보장하는 것
    - 이때 객체가 같은지의 비교는 `==` 연산자로 결정한다.
- `Map`은 어떤 값과 집합의 각 원소사이에 연관 관계를 만든다.

- 기본적으로 `mutable`한 `Set`, `Map`을 생성하면 내부적으로 `Hash*`를 생성된다.
    - `scala.collection.mutable.Set()` => `scala.collection.mutable.HashSet()`
    - `scala.collection.mutable.Map()` => `scala.collection.mutable.HashMap()`
- `immutable`의 경우 원소가 5개 이상일 경우에만 `Hash*`을 생성한다.
    - 0개의 경우 => `scala.collection.immutable.EmptySet`
    - 1개의 경우 => `scala.collection.immutable.Set1`
    - 2개의 경우 => `scala.collection.immutable.Set2`
    - 5개의 경우 => `scala.collection.immutable.HashSet`

- `SortedSet`, `SortedMap` => `TreeSet`, `TreeMap`
    - 순서를 나타내기 위해 `Red-Black Tree`를 사용한다.

## mutable collection, immutable collection
- 케바케지만.. `immutable collection`이 `mutable collection` 보다 프로그램 추론하기가 더 쉽다고 함

## initialize collection
- 초기 원소를 콜렉션 동반 객체의 팩토리 메소드에 넘기어 초기화함.

### 배열, 리스트 변환
- `ToArray`, `ToList` 를 이용.
- 이 연산들은 콜렉션의 모든 원소들을 복사하는 것이기에 대신 콜렉션의 크기가 큰 경우  느릴 수 있음.

### immutable set(map), mutable set(map) 사이의 변환
- `++, ++=` 을 활용..

## tuple
- `tuple`은 정해진 개수의 원소를 한데 묵는다.
- 데이터만 저장하는 단순한 클래스를 정의해야하는 번거로움을 덜 수 있다.
- 튜플 원소에 접근할때는 첫번째 원소는 `_1`, 두번째 원소는 `_2`.. 이렇게 접근한다.

