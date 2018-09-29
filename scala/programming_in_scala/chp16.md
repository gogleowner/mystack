# chp16. List

- list -> 변경 불가능, 구조가 재귀적임.
- array -> 배열은 평면적이다

## creation
```
List("apple", "oranges", "pears")
List(1, 2, 3, 4)

List(
  List(1, 0, 0),
  List(1, 0, 0)
)

"apple" :: ("oranges" :: ("pears" :: Nil))
```

## :: (colon)

```
List("apple", "oranges", "pears") => "apple" :: ("oranges" :: ("pears" :: Nil))
List(1, 2, 3) => 1 :: (2 :: (3 :: Nil))
```

- `::` => 오른쪽 결합법칙을 사용한다. 괄호를 아래와 같이 생략할 수 있다
    - `1 :: 2 :: 3 :: 4 :: Nil`

- 그러나 `1 :: 2 :: 3 :: 4` 은 불가하다. 맨 뒤에 Nill을 붙여줘야함.
- `1 + 2 => 1. + (2)` 이듯이 맨 뒤의 `Nil` 로부터 시작하여 체이닝이 일어난다. 때문에 마지막에 `Nil` 이 있어야함.
- 참고링크 : https://stackoverflow.com/questions/37741565/why-do-we-need-nil-while-creating-list-in-scala

## Operation function
- `.head`
- `.tail`
- `.isEmpty`

## list pattern

- 원소의 개수를 알고 있다면 `,`를 이용하여 원소를 매핑
- 원소의 개수를 모른다면 `::`를 이용하여 마지막 값은 나머지 리스트의 원소들을 매핑

```
val fruits = List("apple", "oranges", "pears")
val List(a, b, c) = fruits


```
```
val fruits = List("apple", "oranges", "pears", "grape")
val a :: b :: rest = fruits

// rest 에는 List("pears", "grape") 가 담길 것.
```

## first-order method

- `:::` : 두 리스트 연결

- usage

    ```
List(1, 2) ::: List(3, 4, 5)
xs ::: ys ::: zs => xs ::: (ys ::: zs)
    ```

- `:::` 은 `List` 안에 구현된 메소드이다. 위에서 언급한 패턴매치를 사용하여 직접 구현할 수도 있다.

```
def append[T](xs: List[T], ys: List[T]): List[T] = 
  xs match {
    case List() => ys // 온전한 리스트이면 리스트를 반환
    case firstElement :: remainElements => firstElement :: append(remainElements, ys)
    // 분할된 리스트이면 분할하여 머지하도록 함.
  }

scala> append( (1 to 3).toList , (4 to 10).toList )
res0: List[Int] = List(1, 2, 3, 4, 5, 6, 7, 8, 9, 10)
```

### 리스트의 양 끝

- `length` : List는 LinkedList 로 구현되어있기에 길이를 구하려면 전체 순회를 해야한다.
    - 때문에 성능 측면에서 `length == 0` 보다  `isEmpty` 를 사용해야함.
- `init` : 마지막 원소 제외한 요소들을 반환
- `last` : 마지막 원소를 반환
- 위 두 메소드는 `head`, `tail` 과 다르게 전체 순회를 해야함.

```
scala> val numberList = (1 to 10).toList
numberList: List[Int] = List(1, 2, 3, 4, 5, 6, 7, 8, 9, 10)

scala> numberList.init
res1: List[Int] = List(1, 2, 3, 4, 5, 6, 7, 8, 9)

scala> numberList.last
res2: Int = 10
```


### 접두사와 접미사
- `drop` : n번째 요소를 제외한 모든 요소들을 반환
- `take` : 리스트의 처음부터 n번째까지의 요소들을 반환
- `splitAt` : n번째를 전후로 리스트를 분할

```
scala> numberList drop 2
res3: List[Int] = List(3, 4, 5, 6, 7, 8, 9, 10)

scala> numberList take 5
res4: List[Int] = List(1, 2, 3, 4, 5)

scala> numberList splitAt 5
res5: (List[Int], List[Int]) = (List(1, 2, 3, 4, 5),List(6, 7, 8, 9, 10))
```

### 원소 선택하기
- `apply` : n번째 요소를 가져옴
    - `apply` 이기 때문에 `리스트변수명(n)`으로도 접근 가능.
    - 그러나 `LinkedList` 특성상 n번째까지 순회하여 요소를 가져오기 때문에 n값에 비례하여 시간이 걸림.
    - `xs apply n` == `(xs drop n).head`

- `indices` : 리스트의 모든 index 시퀀스를 반환함

```
scala> numberList.apply(1)
res6: Int = 2

scala> numberList(1)
res7: Int = 2

scala> numberList.indices
res8: scala.collection.immutable.Range = Range 0 until 10
```

### 리스트의 리스트를 한 리스트로
- `flatten`

```
scala> List( List(1, 2), List(3, 4), List(), (5 to 10).toList)
res10: List[List[Int]] = List(List(1, 2), List(3, 4), List(), List(5, 6, 7, 8, 9, 10))

scala> List( List(1, 2), List(3, 4), List(), (5 to 10).toList).flatten
res11: List[Int] = List(1, 2, 3, 4, 5, 6, 7, 8, 9, 10)

scala> List("apple", "computer").map(_.toCharArray).flatten
res14: List[Char] = List(a, p, p, l, e, c, o, m, p, u, t, e, r)
```


### 두 리스트를 순서쌍으로 묶기
- `zip` : 두 리스트를 인자로 받아서 순서쌍의 리스트를 만든다. 길이가 다르면 긴쪽의 남은 원소를 버림
- `zipWithIndex` : 리스트의 원소와 인덱스를 튜플로 묶는다.
- `unzip` :  튜플의 리스트를 리스트이 튜플로 변경

```
scala> val abcde = ('a' to 'e').toList
abcde: List[Char] = List(a, b, c, d, e)

scala> abcde.indices.zip(abcde)
res26: scala.collection.immutable.IndexedSeq[(Int, Char)] = Vector((0,a), (1,b), (2,c), (3,d), (4,e))

scala> val abcde = ('a' to 'e').toList
abcde: List[Char] = List(a, b, c, d, e)

scala> abcde.indices.zip(abcde)
res27: scala.collection.immutable.IndexedSeq[(Int, Char)] = Vector((0,a), (1,b), (2,c), (3,d), (4,e))

scala> abcde.zipWithIndex
res28: List[(Char, Int)] = List((a,0), (b,1), (c,2), (d,3), (e,4))

scala> val zipped = abcde.indices.zip(abcde)
zipped: scala.collection.immutable.IndexedSeq[(Int, Char)] = Vector((0,a), (1,b), (2,c), (3,d), (4,e))

scala> zipped.unzip
res29: (scala.collection.immutable.IndexedSeq[Int], scala.collection.immutable.IndexedSeq[Char]) = (Vector(0, 1, 2, 3, 4),Vector(a, b, c, d, e))
```


### 리스트 출력하기
- `toString` : 요소들을 `, `단위로 반환함
- `mkString` : 요소들을 구분자,  시작/끝 문자를 추가하여 반환함.

```
scala> numberList.toString
res30: String = List(1, 2, 3, 4, 5, 6, 7, 8, 9, 10)

scala> numberList.mkString
res31: String = 12345678910

scala> numberList.mkString(",")
res32: String = 1,2,3,4,5,6,7,8,9,10

scala> numberList.mkString("[ ", ",", " ]")
res33: String = [ 1,2,3,4,5,6,7,8,9,10 ]
```

### 리스트 변환
- `iterator`
- `toArray`
- `copyToArray`

## high-order operator

### 매핑
- `map` : 리스트의 모든 원소에 함수를 적용해서 나온 결과값으로 이뤄진 리스트를 반환
- `flatMap` : 리스트의 모든 원소에 함수를 적용해서 나온 모든 리스트를 연결한 단일 리스트를 반환
- `foreach` : 리스트의 모든 원소에 반환타입이 `Unit`인 함수를 적용한다.

```
scala> val words = List("the", "quick", "brown", "fox")
words: List[String] = List(the, quick, brown, fox)

scala> words map(_.toList)
res4: List[List[Char]] = List(List(t, h, e), List(q, u, i, c, k), List(b, r, o, w, n), List(f, o, x))

scala> words flatMap(_.toList)
res5: List[Char] = List(t, h, e, q, u, i, c, k, b, r, o, w, n, f, o, x)

scala> var sum = 0
sum: Int = 0

scala> List.range(1, 5).foreach(sum += _)

scala> sum
res9: Int = 10
```

### 리스트 걸러내기
- 조건에 맞는 원소들을 걸러냄
    - `filter` : 리스트의 모든 원소에 반환타입이 `Boolean`인 함수를 적용하여 반환값이 `true`인 원소의 리스트를 만든다.
    - `partition` : `filter`의 인자와 같지만 `true`인 원소 리스트, `false`인 원소 리스트의 순서쌍을 만든다.
    - `find` : 조건에 만족하는 첫번째 원소만 반환한다.
- 조건에 맞는 긴 접두사를 반환
- `takeWhile` : 함수를 만족하는 가장 긴 접두사를 반환
- `dropWhile` : 함수를 만족하는 가장 긴 접두사를 제거
- `span` : `takeWhile`, `dropWhile`을 합쳐놓은 것과 같음. 두 리스트의 순서쌍을 반환

```
scala> List.range(1, 5).filter(_ % 2 == 0)
res10: List[Int] = List(2, 4)

scala> List("the", "quick", "brown", "fox").filter(_.length == 3)
res13: List[String] = List(the, fox)

scala> List.range(1, 5).partition(_ % 2 == 0)
res14: (List[Int], List[Int]) = (List(2, 4),List(1, 3))

scala> List.range(1, 5).find(_ % 2 == 0)
res15: Option[Int] = Some(2)

scala> List(1, 2, 3, -4, 5).takeWhile(_ > 0)
res32: List[Int] = List(1, 2, 3)

scala> List(1, -1, 2, 3, -4, 5).takeWhile(_ > 0)
res33: List[Int] = List(1)

scala> List(1, 2, 3, 2, 3, -4, 5).dropWhile(_ > 0)
res37: List[Int] = List(-4, 5)

scala> List(1, 2, 3, 2, 3, -4, 5).span(_ > 0)
res38: (List[Int], List[Int]) = (List(1, 2, 3, 2, 3),List(-4, 5))
```

### 리스트 전체에 대한 술어
- `forall` : 리스트의 모든 원소가 함수를 만족할때 결과가 true
- `exists` : 리스트의 원소중의 하나라도 결과를 만족하는 원소가 존재하면 true

```
scala> List.range(1, 5).exists(_ == 0)
res39: Boolean = false

scala> List.range(1, 5).forall(_ > 0)
res41: Boolean = true
```

### fold
- `/:`, `foldLeft` : 리스트의 모든 원소를 왼쪽에서 오른쪽으로 이진연산자를 적용. 첫번째 인자는 초기값, 두번째 인자는 연산

```
scala> List.range(1, 5).foldLeft(0)(_ + _)
res42: Int = 10

scala> (0 /: List.range(1, 5)) (_ + _)
res44: Int = 10

scala> ("" /: words) (_ + " " + _)
res45: String = " the quick brown fox"

scala> (words.head /: words.tail) (_ + " " + _)
res46: String = the quick brown fox

(z /: List(a, b, c))(op) => op(op(op(z, a), b), c)
```

- `:\`, `foldRight` : 리스트의 모든 원소를 오른쪽에서 왼쪽으로 이진연산자를 적용. 첫번째 인자는 연산, 두번째 인자는 초기값

```
(List(a, b, c) :\ z)(op) => op(a, op(b, op(c, z)))
```

### 리스트 정렬
- `sortWith` : 두 원소를 비교할 수 있는 함수를 이용하여 정렬. `x before y` 에서 `x`가 `y`보다 앞에 있어야한다면, `true`를 반환해야함.

```
scala> List(1, -3, 4, 2, 6).sortWith(_ < _)
res50: List[Int] = List(-3, 1, 2, 4, 6)
```

## List 객체의 메소드

### 원소로부터 리스트 만들기
- `List.apply` : `List(1, 2, 3) == List.apply(1, 2, 3)`

### 수의 범위를 리스트로 만들기
- `List.range(from, until)`
- `List.range(from, until, incremental)`

```
scala> List.range(1, 5)
res51: List[Int] = List(1, 2, 3, 4)

scala> List.range(1, 5, 2)
res52: List[Int] = List(1, 3)
```

### 균일한 리스트 만들기
- `List.fill` : 같은 원소를 0번 이상 반복한 리스트를 만든다.

```
scala> List.fill(5)('a')
res53: List[Char] = List(a, a, a, a, a)
```

### 함수 도표화
- `List.tabulate` : 제공된 함수로 계산한 원소의 리스트를 생성한다.
    - 첫번째 인자 : 생성할 리스트의 차원
    - 두번째 인자 : 리스트의 원소를 묘사

```
scala> List.tabulate(5)(n => n * n)
res54: List[Int] = List(0, 1, 4, 9, 16)
```

### 여러 리스트 연결하기
- `concat`

```
scala> List.concat( List('a', 'b'), List('c', 'd', 'e') )
res56: List[Char] = List(a, b, c, d, e)
```

## 여러 리스트 함께 처리하기

- `zipped` : 가장 짧은 리스트의 원소개수 만큼 순서쌍을 만들어낸다. 

```
scala> ( List(10, 20), List(3, 4, 5) ).zipped.map(_ * _)
res59: List[Int] = List(30, 80)
```

