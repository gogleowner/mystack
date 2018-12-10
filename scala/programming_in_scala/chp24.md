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


## 24.4 Trait Iterable
## 24.5 Sets
## 24.6 Maps
## 24.7 Concrete immutable collection classes
## 24.8 Concrete mutable collection classes
## 24.9 Arrays
## 24.10 Strings
## 24.11 Performance characteristics
## 24.12 Equality
## 24.13 Views
## 24.14 Iterators
## 24.15 Creating collections from scratch
## 24.16 Conversions between Java and Scala collections
## 24.17 Conclusion

