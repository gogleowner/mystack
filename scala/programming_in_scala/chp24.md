# chp24. Collections in Depth

- 사용하기 쉬움 : 20~50개의 메소드만 알두면 대부분의 문제를 몇가지 연산을 조합하여 해결 가능
- 간결함 : 루프가 하나 이상 필요할 때 한 단어로 가능하다. 
- 안전함 : 실수를 대부분 컴파일 타임에 걸러낼 수 있다.
    1. 컬렉션 자체가 널리 쓰이고 잘 검증되어 있다.
    2. 입력, 출력을 함수 파라미터와 반환 값으로 명시해야 한다.
    3. 명시적 입출력이 정적 타입 검사를 통과해야 한다.
- 빠름 : 라이브러리의 컬렉션 연산은 튜닝, 최적화가 되어있다. 물론 그 안에서 더 나은 성능을 내게끔 튜닝할 수 있지만 오히려 나빠질 수도 있다.
- 보편적임 : 각 컬렉션은 같은 연산을 제공한다.

## 24.1 Mutable and immutable collections
### Mutable Collection
- `scala.collection.mutable`
- 원소의 추가,삭제,변경이 가능하다.
### Immutable Collection
- `scala.collection.immutable`
- 원소의 추가,삭제,변경이 불가능 하다. 연산이 있긴 하지만 원래 컬렉션은 변하지 않고 변경된 컬렉션을 새로 생성하여 반환한다.

### `scala.collection` 패키지에 위치한 컬렉션
- 변경 가능할 수도, 변경 불가능할 수도 있다.
- `IndexedSeq` 는 `scala.collection`에 존재하며, 이는 mutable, immutable한 trait으로도 존재한다.

### 기본 컬렉션
- 스칼라에서는 일반적으로 변경 불가능한 컬렉션을 선택한다.
- 스칼라 패키지가 기본적으로 import 하는 것은 변경 불가능 한 것이기 때문이다.
- 따라서 변경 가능한 컬렉션을 이용하고자 할때는 명시적으로 mutable 한 객체를 import 하여 사용하자.

### `scala.collection.generic`
- 컬렉션을 만드는 구성요소가 들어있다.
- 보통 컬렉션들은 자신이 제공하는 연산에 대한 구현을 `generic`에 있는 클래스에 위임한다.


## 24.2 Collections consistency

- 클래스다이어그램은 공식 홈페이지를 참고하자.
    - https://docs.scala-lang.org/overviews/collections/overview.html

## 24.3 Trait Traversable
- 유일한 메소드는 `foreach` 뿐이다.

```
def foreach[U](f: Elem => U)
```

- `foreach`는 function f 를 각 원소에 적용하는 것을 목적으로 한다.

### 메소드의 종류
- 추가 메소드 : `++` 
    - 두 컬렉션 객체를 하나로 엮거나, 어떤 순회 가능 객체의 뒤에 모든 원소를 추가한다.
- Map 연산 : `map, flatMap, collect` 
    - 어떤 함수를 컬렉션에 있는 원소에 적용해 새로운 컬렉션을 만들어낸다.
- 변환 연산 : `toIndexedSeq, toIterable, toStream, toArray, toList, toSeq, toSet, toMap`
    - Traversable 컬렉션을 더 구체적인 컬렉션으로 변환한다.
    - 만약 호출 대상 객체가 이미 그 타입이면 그 객체를 반환한다.
- 복사연산 : `copyToBuffer, copyToArray`
    - 컬렉션의 원소를 버퍼 or 배열에 복사한다.
- 크기 연산 : `isEmpty, nonEmpty, size, hasDefiniteSize`
    - 컬렉션은 유한일 수도 있고, 무한일 수도 있다.
    - 무한한 컬렉션의 예시) `Stream.from(0)`
    - `hasDefiniteSize` => 컬렉션이 무한할 가능성이 있는지 알려준다.
        - https://www.scala-lang.org/api/2.12.x/scala/collection/immutable/Stream.html#hasDefiniteSize:Boolean

- 원소를 가져오는 연산 : `head, last, headOption, lastOption, find`
    - 컬렉션이 매번 원소를 같은 순서로 반환하는 경우 => ordered collection
    - 순서가 없는 컬렉션 => HashSet 등등.
        - 이들은 순서가 없는 대신 효율성을 약간 더 얻는 쪽을 택함.
        - 순서가 있는 `HashSet`은 `LinkedHashSet`

- 하위 컬렉션을 가져오는 연산 : `takeWhile, take, init, slice, take, drop, filter, dropWhile, filterNot, withFilter`
    - 인덱스 범위 혹은 test 표현식에 따라 컬렉션의 일부를 반환한다.
- 분할연산 : `splitAt, span, partition, groupBy`
    - 컬렉션을 여러 하위 컬렉션으로 나눈다.
- 원소 테스트 메소드 : `exists, forall, count`
- 폴드 연산 : `foldLeft, foldRight, /:, :, reduceLeft, reduceRight`
    - 이항 연산을 연속된 원소에 반복 적용한다.
- 특정 폴드 메소드 : `sum, product, min, max`
    - 특정 타입(수 이거나 변경 가능한 타입)에서만 동작함.
- 문자열 연산 : `mkString, addString, stringPrefix`
    - 컬렉션을 문자열로 바꾸는 연산을 제공
- 뷰 연산 : `view`
    - 필요에 따라 나중에 계산이 이뤄지는 컬렉션
    - 24.14 에서 살펴볼 것

## 24.4 Trait Iterable

- 추상메소드 `iterator`를 기반으로 한다.

```
def iterator: Iterator[A]
```

- `java`의 `iterator` 와 동일한 기능이다.

### Traversable과 Iterable의 차이

- 간단한 이진트리 클래스를 예시로 하여 순회할때의 효율을 비교해보자.

```
// 트리의 루트블록
case class Branch(left: Tree, right: Tree) extends Tree

// 트리의 노드
case class Node(elem: Int) extends Tree
```

- `Traversable`을 이용한 방법

    ```
sealed abstract class Tree extends Traversable[Int] {
  override def foreach[U](f: Int => U): Unit = this match {
    case Node(elem) => f(elem)
    case Branch(l, r) => l foreach f; r foreach f
  }
}
    ```

    - 트리 안의 노드 개수에 비례하는 시간이 소요된다.
    - N개의 잎을 가진 균형 있는 트리의 경우 Branch 타입의 내부 노드는 `N - 1`개이다.
    - 트리를 모두 순회하기 위해서는 `N + N - 1`번 방문하면 된다.

- `Iterable`을 이용한 방법


    ```
sealed abstract class Tree extends Iterable[Int] {
  override def iterator: Iterator[Int] = this match {
    case Node(elem) => Iterator.single(elem)
    case Branch(l, r) => l.iterator ++ r.iterator
  }
}
    ```

    - `foreach` 와 얼핏 비슷해보이긴 하지만, 순회 이후에 `++` 로 둘을 결합하게 되어있다.
    - 전체적으로 N개의 잎이 있는 트리에서 잎을 가져오기 위해서는 log(N)정도의 접근을 해야한다.

- 정리하자면.. 트리의 모든 원소를 방문하는데 듣는 복잡도는
    - `foreach` : `2N`
    - `iterator` : `N log(N)`
    - 트리에 100만개의 원소가 있다면 foreach는 2백만단계, iterator는 2천만 단계를 거쳐야한다.

### Iterable 의 하위 분류
- `Seq, Set, Map` 이렇게 세가지 `Trait` 이 존재한다.
- 공통점은 `PartialFunction`을 구현하여 `apply, isDefinedAt`을 제공한다는 점이다.
- 각 트레잇이 `PartitionFunction`을 구현하는 방식은 다르다.


## 24.5 Sequence Trait : Seq, IndexedSeq, LinearSeq

- `Seq` : 시퀀스를 표현한다.
- 길이가 정해져 있고, 원소의 위치가 0부터 시작하는 고정된 인덱스로 지정할 수 있는 iterable 의 일종이다.

### 메소드들
- 인덱스와 길이 연산 : `apply, isDefinedAt, length, indices, lengthCompare`
    - `apply` => 인덱스로 원소를 찾는 것을 의미
    - `length` 는 `size` 에 대한 별칭이다.
    - `lengthCompare`는 길이를 비교하는 메소드인데, 한쪽의 길이가 무한하더라도 두 시퀀스의 길이를 비교할 수 있다.
- 인덱스 찾기 연산 : `indexOf, lastIndexOf, indexOfSlice, lastIndexOfSlice, indexWhere, lastIndexWhere, segmentLength, prefixLength`
    - 주어진 값과 같거나 어떤 조건 함수를 만족하는 원소의 인덱스를 반환한다.
- 추가 연산 : `+:, :+, padTo`
    - 시퀀스의 맨 앞, 뒤에 원소를 추가한 새 시퀀스를 반환한다.(prepend, append)
- 변경 연산 : `updated, patch`
    - 원래의 시퀀스의 일부 원소를 바꿔서 나오는 새로운 시퀀스를 반환
- 정렬 연산 : `sorted, sortWith, sortBy`
- 반전 연산 : `reverse, reverseIterator, reverseMap`
    - 시퀀스를 역순으로 반환함.
- 비교 연산 : `startsWith, endsWith, contains, corresponds, containsSlice`
    - 시퀀스 간의 관계 판단, 시퀀스에서 원소를 찾는다.
- 중복 집합 연산 : `intersect, diff, union, distinct`

### LinearSeq, IndexedSeq
- `Seq` 트레잇에는 `LinearSeq`, `IndexedSeq` 하위 트레잇이 있다.

- `LinearSeq`
    - 효율적인 `head, tail` 연산을 제공한다.
    - `List, Stream` 이 가장 많이 쓰인다.
- `IndexedSeq`
    - 효율적인 `apply, length, update` 연산을 제공한다.
    - `Array, ArrayBuffer` 이 가장 많이 쓰인다.
- `Vector` 는 `LinearSeq`와 `IndexedSeq`간의 절충을 제공한다.

### Buffer
- 기존 원소의 변경, 삽입, 제거 등의 연산을 제공한다.
- 대표적으로 `ListBuffer, ArrayBuffer` 가 있다.

## 24.6 Set
- `Set`은 중복을 허용하지 않는 Iterable이다.

### 메소드들
- 검사 : `contains, apply, subsetOf`
    - `contains`: `apply` 메소드와 동일하다.

    ```
scala> val fruit = Set("apple", "organge", "peach")
fruit: scala.collection.immutable.Set[String] = Set(apple, organge, peach)

scala> fruit("peach")
res0: Boolean = true

scala> fruit.apply("peach")
res1: Boolean = true

scala> fruit.contains("peach")
res2: Boolean = true
    ```

- 추가 연산 : `+, ++` 
- 제거 연산 : `-, --` 
- 집합 연산 : `intersect, union, diff`

### 구현방법
- `mutable set` : `HashTable`을 사용하여 원소를 저장한다.
- `imuutable set` 
    - 빈 집합은 싱글톤 객체를 사용한다. 
    - 원소가 4개 이하인 집합은 원소를 필드로 저장하는 객체 하나로 표현한다.
        - https://www.scala-lang.org/api/2.12.8/scala/collection/immutable/Set$$Set1.html
    - 4개 초과의 경우 `hash trie`를 사용해 구현한다.

## 24.7 Map
- `Map` 키와 값의 쌍을 가진 Iterable이다.
- `Predef` 클래스에는 `(key, value)` 를 `key -> value` 로 대신 사용할 수 있는 암시적 변환이 들어있다.
    - https://www.scala-lang.org/api/2.12.8/scala/Predef$$ArrowAssoc.html#->[B](y:B):(A,B)

### 메소드들
- 검색 연산 : `apply, get, getOrElse, contains, isDefinedAt`
- 추가, 변경 연산 : `+, ++, updated`
- 제거 연산 : `-, --`
- 하위 컬렉션 생성 메소드 : `keys, keySet, keysIterator, valuesIterator, values`
- 변환 연산 : `filterKeys, mapValues`

- 맵캐시로 사용하는 경우 `getOrElseUpdate` 가 유용하다.
    - 키에 대한 값이 맵에 있다면 그 값을 반환한다. 키가 존재하지 않다면 키-값 쌍을 맵에 추가하고 값을 반환한다.

## 24.8 Concrete immutable collection classes

### List
- 유한한 변경 불가능한 시퀀스
- `head, tail` 외의 다른 원소들에 접근하는 시간은 O(n)의 시간이 걸린다.

### Stream
- 리스트와 비슷하지만, 원소를 지연 계산하는 점이 다르다.
- 이로 인해 스트림은 무한할 수 있다.
- 외부에서 요청하는 원소만 계산한다.

```
scala> val str = 1 #:: 2#:: 3 #:: Stream.empty
str: scala.collection.immutable.Stream[Int] = Stream(1, ?)
```

- 초기화한 내용을 보면, 아직 연산이 이뤄지지 않은 부분은 ? 로 표시되어있는 것이다

### Vector
- `head`가 아닌 원소도 효율적으로 접근할 수 있는 컬렉션 타입이다.
- 벡터에 있는 임의의 원소에 접근하기 위해서는 `사실상 상수 시간`이 걸린다.
    - 벡터는 넓고 얕은 트리로 표현한다. 모든 트리 노드에는 32개의 벡터 원소를 넓거나 다른 트리 노드를 넣을 수 있다.
    - 원소가 32개 이하인 벡터는 노드 하나로 표현할 수 있다.
    - 원소가 1024(32\*32)개가 될때까지는 한번만 간접 노드를 거치면 접근 가능하다.
    - 2개노드를 거치면 2^15, 3개의 노드를 거치면 2^20, 4개의 노드를 거치면 2^25, 5개의 노드를 거치면 2^30.. 
    - 일반적인 경우 원소 선택은 최대 5단계의 기본 배열 선택으로 가능하다.
- 벡터 원소 값의 변경
    - 벡터는 변경 불가능하기 때문에 벡터의 원소를 그 자리에서 바꿀 수 없다.
    - `updated`를 이용해 새로운 벡터를 만든다.
    - 업데이트도 `사실상 상수시간` 소요된다.
    - 벡터의 중간 원소를 변경하려면 원소가 있는 노드를 복사하고 트리의 루트로부터 시작해서 원래의 노드를 가리키던 모든 노드를 복사해야 한다.
- 벡터가 빠른 random access, update 사이에서 균형을 잘 잡고 있어서, `IndexedSeq`의 기본 구현은 벡터이다.

### Imuutable Stack, Queue




## 24.9 Concrete mutable collection classes
## 24.10 Arrays
## 24.11 Strings
## 24.12 Performance characteristics
## 24.13 Equality
## 24.14 Views
## 24.15 Iterators
## 24.16 Creating collections from scratch
## 24.17 Conversions between Java and Scala collections
## 24.18 Conclusion

