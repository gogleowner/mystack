# Chapter 7. Operator overloading and other conventions

- 어떤 언어 기능과 미리 정해진 이름의 함수를 연결해주는 기법을 코틀린에서는 convention이라고 부른다.
- 언어 기능을 타입에 의존하는 자바와 달리, 코틀린은 convention에 의존한다. 기존 자바 클래스를 코틀린 언어에 적용하기 위해, 이 관레를 택하였다. 기존 자바 클래스가 구현하는 인터페이스는 이미 고정되있고 코틀린 쪽에서 자바 클래스가 새로운 인터페이스를 구현하게 만들 수는 없다.
- 반면 확장함수를 사용하면 기존 클래스에 새로운 메소드를 추가할 수 있다. 기존 자바 클래스에 대해 확장 함수를 구현하면서 관례에 따라 이름을 붙이면 기존 자바 코드를 바꾸지 않고도 새로운 기능을 부여할 수 있다.

# 7.1. 산술 연산자 오버로딩

## 7.1.1. 이항 산술 연산 오버로딩

    data class Point(val x: Int, val y: Int) {
      operator fun plus(other: Point): Point {
        return Point(x + other.x, y + other.y)
      }
    }
    
    val p1 = Point(10, 20)
    val p2 = Point(30, 40)
    println(p1 + p2) => Point(x=40, y=60)
    println(p1.plus(p2))

- 연산자 오버로딩 함수 앞에는 꼭 `operator` 가 있어야 한다. 관례를 따르는 함수임을 명확히 할 수 있다.

    a * b => times
    a / b => div
    a % b => rem
    a + b => plus
    a - b => minus

- 연산자 우선순위는 표준 숫자 타입에 대한 연산자 우선순위와 같다. ( *, / , %   >   + , - )

    operator fun Point.times(scale: Double): Point =
      Point( (x * scale).toInt(), (y * scale).toInt() )
    
    val p = Point(10, 20)
    println(p * 1.5) ==> Point(x=15, y=30)

- 단 위의 경우 연산자 교환법칙은 성립하지 않기에, `1.5 * p` 는 적용되지 않는다.
- 결과타입이 피연산자 타입과 다른 연산자 정의하기

        operator fun Char.times(count: Int): String = toString().repeat(count)
        
        println('a' * 3) ==> aaa

- 비트연산자에 대해 특별한 연산자 함수를 사용하지 않는다.

        shl => <<
        shr => >>
        ushr => >>>
        and => &
        or => |
        xor => ^
        inv => ~
        
        println(0x0F and 0xF0) => 0
        println(0x0F or 0xF0) => 255
        println(0x1 shl 4) => 16

## 7.1.2. 복합 대입 연산자 오버로딩

- plus 연산자를 오버로딩하면 `+` 뿐 아니라 `+=` 도 함께 지원한다. 이는 복합 대입 (compound assignment)라고 불린다.
- 반환 타입이 `Unit` 인 `plusAssign` 함수를 정의하면 코틀린은 `+=` 연산자에 대해 그 함수를 사용한다.

        operator fun <T> MutableCollection<T>.plusAssign(element: T) {
          this.add(element)
        }
        
        val list = arrayListOf(1, 2)
        list += 3

- 코틀린 표준 라이브러리에서의 +, - 연산은 항상 새로운 컬렉션을 반환하며, += , -= 연산자는 항상 mutable  collectio에 적용해 메모리에 있는 객체 상태를 변화시킨다. 읽기 전용 컬렉션에서는 복사본을 반환한다.

## 7.1.3. 단항 연산자 오버로딩

    operator fun Point.unaryMinus(): Point = Point(-x, -y)
    
    val p = Point(10, 20)
    pritnln(-p) => Point(x=-10, -20)

- 오버로딩할 수 있는 단항 산술 연산자

        +a => unaryPlus
        -a => unaryMinus
        !a => not
        ++a, a++ => inc
        --a, a-- => dec

- BigDecimal 클래스의 증가 연산자 정의하기

        operator fun BigDecimal.inc() = this + BigDecimal.ONE
        
        varbd = BigDecimal.ZERO
        println(bd++) => 0
        println(++bd) => 2

# 7.2. 비교 연산자 오버로딩

## 7.2.1. 동등성 연산자 : equals

- `==, !=`  연산자는 `equals()` 호출로 컴파일된다.

        a == b => a?.equals(b) ?: (b == null)
        
        class Point(val x: Int, val y: Int) {
          override fun equals(obj: Any?): Boolean {
            if (obj === this) return true // 파라미터가 this 와 같은 객체인지 검사
            if (obj !is Point) return false // 파라미터 타입을 검사
            return obj.x == x && obj.y == y
          }
        }

- `===` : 식별자 비교 (identity equals) 로 equals 의 파라미터가 수신 객체와 같은지 확인한다.
    - 이는 자바의 `==` 와 같다. 따라서 `===`는 서로 같은 객체인지를 비교한다.
    - `===` 를 이용해 자기 자신과의 비교를 최적화한다.
- `Any` 클래스의 `equals()`에는 `operator` 이 붙어있지만 이를 구현하는 하위 클래스의 메소드에서는 `operator`를 붙이지 않아도 자동으로 상위 클래스의 `operator` 지정이 적용된다.

## 7.2.2. 순서 연산자 : compareTo

- 자바에서 정렬, 최대, 최소값 등의 값을 비교해야하는 연산에 `Comparable` 인터페이스에 있는 `compareTo()` 메소드를 통해서 구현할 수 있다.
- 코틀린에서는 이를 짧게 줄여서 호출할 수 있도록 관례를 제공한다.

    a >= b ==> a.compareTo(b) >= 0
    
    compareValuesBy => compareTo 메소드 사용을 쉽고 간결하게 정의할 수 있는 함수이다.

# 7.3. 컬렉션과 범위에 대해 쓸 수 있는 관례

## 7.3.1. 인덱스로 원소에 접근 : get , set

    val value = map[key]
    mutableMap[key] = newValue

- 이 코드가 어떻게 동작하는지 살펴보자.

- `get()` convention 구현하기

        operator fun Point.get(index: Int): Int {
          return when (index) {
            0 -> x
            1 -> y
            else -> IndexOutOfBoundException()
          }
        }
        
        x[a] => x.get(a)
        
        // 두개 이상의 파라미터도 표현 가능하다.
        operator fun get(rowIndex: Int, collIndex: Int)
        
        martix[row, col]

- `set()` convention 구현하기

        data class MutablePoint(var x: Int, var y: Int)
        
        operator fun MutablePoint.set(index: Int, value: Int) {
          when (index) {
            0 -> x = value
            1 -> y = value
            else -> IndexOutOfBoundException()
          }
        }
        
        val p = MutablePoint(10, 20)
        p[1] = 42
        println(p) => MutablePoint(x=10, y=42)
        
        x[a, b] = c ==> x.set(a,b,c)

## 7.3.2. in convention

- `in` 은 객체가 컬렉션에 들어있는지 검사한다. (contains() 와 대응)

    data class Rectangle(val upperLeft: Point, val lowerRight: Point)
    
    operator fun Rectangle.contains(p: Point): Boolean {
      return p.x in upperLeft.x until lowerRight.x &&
             p.y in upperLeft.y until lowerRight.y
    }
    
    val rect = Rectangle(Point(10, 20), Point(30, 40))
    println(Point(20, 30) in rect) => true
    println(Point(5, 5) in rect) => false
    
    a in collection => collection.contains(a)

## 7.3.2. rangeTo convention

    start..end => start.rangeTo(end)
    
    operator fun <T: Comparable<T>> T.rangeTo(that: T): CloseRange<T>
    
    val now = LocalDate.now()
    val vacation = now..now.plusDays(10)
    println(now.plusWeeks(1) in vacation) => true
    
    (0..n).forEach { println(it) } => 012345..n

## 7.3.4. for 루프를 위한 iterator convention

    for (x in list) { ... }
    =>
    for (Iterator<T> i = list.iterator() ; i.hasNext() ; i.next()) { ..}
    
    // 이 또한 관례이므로 iterator 메소드를 확장함수로 정으할 수 있다.
    
    // CharSequence Iterator
    operator fun CharSequence.iterator(): CharIterator
    for (c in "abc") { .. }
    
    // 날짜 범위에 대한 iterator 구현
    operator fun ClosedRange<LocalDate>.iterator(): Iterator<LocalDate> =
      object : Iterator<LocalDate> {
        var current = start
        override fun hasNext(): Boolean = current <= endInclusive
        override fun next(): LocalDate = current.apply { // 현재 날짜를 저장한 다음 날짜를 변경. 그 후 저장한 날짜를 반환한다.
          current = current.plusDays(1) 
        }
    }
    
    val newYear = LocalDate.ofYearDay(2019,1)
    val daysOff = newYear.minusDays(1)..newYear
    
    for (dayOff in daysOff) {
      println(dayOff)
    }
    
    2018-12-31
    2019-01-01

## 7.4.1. 구조 분해 선언과 루프

- `componentN`

        data class NameComponents(val name: String, val extension: String)
        
        fun splitFileName(fullName: String): NameComponents {
          val (name, extension) = fullName.split(".", limit = 2)
          return NameComponents(name, extension)
        }

- 아래 예제는 객체를 이터레이션, 구조 분해 선언 이 두가지 관례를 활용한다.

        fun printEntries(map: Map<String, String>) {
          for ((key, value) in map) { // 구조 분해 선언!
            println("$key -> $value")
          }
        }
        
        for (entry in map.entries) {
          val key = entry.component1()
          val value = entry.component2()
        }

