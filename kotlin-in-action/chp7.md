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

# 7.5. 프로퍼티 접근자 로직 재활용 : 위임 프로퍼티 (Delegated property)

- 값을 뒷받침하는 필드에 단순히 저장하는 것보다 더 복잡한 방식으로 작동하는 프로퍼티를 쉽게 구현할 수 있따.
- 프로퍼티는 위임을 사용해 자신의 값을 필드가 아니라 데이터베이스 테이블이나 브라우저 세션, 맵 등에 저장할 수 있다.
- Delegate Design Pattern : 객체가 직접 작업을 수행하지 않고 다른 도우미 객체가 그 작업을 처리하게 맡기는 패턴. 작업을 처리하는 도우미 객체를 delegate라 부른다.

## 7.5.1. 위임 프로퍼티 소개

    class Foo {
      var p: Type by Delegate() // 접근자 로직을 다른 객체에게 위임한다.
    }
    
    // 위 코드는 아래와 같이 컴파일 된다.
    class Foo {
      private val delegate = Delegate()
      var p: Type
        set(value: Type) = delegate.setValue(.., value)
        get() = delegate.getValue(..)
        // p 프로퍼티를 위해 컴파일러가 생성한 접근자는 getValue, setValue 를 호출한다.
    }
    
    class Delegate {
      operator fun getValue(...) { .. }
      operator fun setValue(..., value: Type) { .. }
    }
    
    val foo = Foo()
    val oldValue = foo.p // delegate.getValue()
    foo.p = newValue // delegate.setValue(.., newValue)

## 7.5.2. 위임 프로퍼티 사용 : 초기화 지연 `by lazy()`

- lazy initialization 은 객체의 일부분을 초기화하지 않고 남겨뒀다가 실제로 그 값이 필요할 때 초기화하는 패턴이다.

    class Person(val name: String) {
      private var _emails: List<Email>? = null
    
      val emails: List<Email>
        get() {
          if (_emails == null) {
            _emails = loadEmails(this)
          }
          return _emails!!
        }
    }

- backing property : `_emails` 프로퍼티에 값을 저장하고 `emails` 는 읽기 함수를 제공함.
- 하지만 이런 코드를 만드는 것은 성가신 일이다. lazy initialization 프로퍼티가 많아지면? thread-safe 하지 않은 점이 있는데, 코틀린에서는 더 나은 해법을 제공한다.

    class Person(val name: String) {
      val emails by lazy { loadEmails(this) }
    }

- `lazy` 함수는 코틀린 관례에 맞는 시그니처의 `getValue` 메소드가 들어있는 객체를 반환한다.
- 기본적으로 thread safe 하다. 필요에 따라 사용할 lock을 함수에 전달할 수도 있고 다중 쓰레드 환경에서 사용하지 않을 프로퍼티를 위해 `lazy`함수가 동기화를 못하게 막을 수도 있다.
- `lazy` 함수를 통해서 프로퍼티가 변경될 때마다 특정 이벤트를 발생시키는 코드를 넣기도 수월하다.

- 프로퍼티 변경 통지를 직접 구현

        open class PropertyChangeAware {
          protected val changeSupport = PropertyChangeSupport(this)
        
          fun addPropertyChangeListener(listener: PropertyChangeListener) {
            changeSupport.addPropertyChangeListener(listener)
          }
        
          fun remotePropertyChangeListener(listener: PropertyChangeListener) {
            changeSupport.removePropert9yChangeListener(listener)
          }
        }
        
        class Person(val name: String, val age: Int, val salary: Int) : PropertyChangeAware() {
          var age: Int = age
            set(newValue) {
              val oldValue = field
              field = newValue
              changeSupport.firePropertyChange("age", oldValue, newValue)
            }
        
          var salary: Int = salary // .. 위와 같음. setter에 중복이 매우 많다.
        }

- 도우미 클래스를 이용하여 프로퍼티 변경 통지 구현

        class ObservableProperty(propName: String, var propValue: Int, val chaneSupport: PropertyChangeSupport) {
          fun getValue(): Int = propValue
          fun setValue(newValue: Int) {
            val oldValue = field
            propValue = newValue
            changeSupport.firePropertyChange(propName, oldValue, newValue)
          }
        }
        
        class Person(val name: String, val age: Int, val salary: Int) : PropertyChangeAware() {
        
          val _age = ObservableProperty("age", age, changeSupport)
            get() = _age.getValue
            set(value) { _age.setValue(value) }
          ..
        }
        
        // 중복을 상당부분 제거했지만 ObservableProperty 에 넣는 준비 코드의 양이 많다.

- 프로퍼티 위임 기능을 이용

        class ObservableProperty(var propValue: Int, val changeSupport: PropertyChangeSupport) {
          operator fun getValue(p: Person, prop: KProperty<*>): Int = propValue
          operator fun setValue(p: Person, prop: KProperty<*>), newValue: Int) {
            val oldValue = propValue
            propValue = newValue
            changeSupport.firePropertyChange(prop.name, oldValue, newValue)
          }
        }
        
        /*
        - operator 
        - KProperty<*>
        */
        
        class Person(val name: String, val age: Int, val salary: Int) : PropertyChangeAware() {
          var age: Int by ObservableProperty(age, changeSupport)
          var salary: Int by ObservableProperty(salary, changeSupport)
        }
        
        // by 키워들르 이용해 위임 객체를 지정하면 위의 다소 수고스러운 로직을 컴파일러가 자동으로 처리해줌.
        // 코틀린 표준 라이브러리에는 위의 ObservableProperty 와 비슷한 클래스인 Delegates.observable 클래스가 있다.

- `Delegates.observable`을 이용한 프로퍼티 변경 통지 구현

        class Person(val name: String, val age: Int, val salary: Int) : PropertyChangeAware() {
          private val observer = { prop: KProperty<*>, oldValue: Int, newValue: Int -> 
            changeSupport.firePropertyChange(prop.naem, oldValue, newValue)
          }
          var age: Int by Delegates.observable(age, observer)
          var salary: Int by Delegates.observable(salary, observer)
        }

## 7.5.4. 위임 프로퍼티 컴파일 규칙

    class C {
      var prop: Type by MyDelegate()
    }
    
    val c = C()
    val x = c.prop 
    c.prop = x
    
    // 위의 코드가 컴파일되면 아래와 같은 코드를 생성한다.
    class C {
      private val <delegate> = MyDelegate()
      var prop: Type
        get() = <delegate>.getValue(this, <property>)
        set(value: Type) = <delegate>.setValue(this, <property>, value)
    }
    
    val x = <delegate>.getValue(c, <property>)
    <delegate>.setValue(c, <property>, x)

- 프로퍼티 값이 저장되는 장소가 변경될 경우에 간결하게 처리가 가능하다.

## 7.5.5. 프로퍼티 값을 맵에 저장

- `expando object` : 자신의 프로퍼티를 동적으로 정의할 수 있는 객체를 만들때 사용하는 위임 프로퍼티

    class Person {
      private val _attributes = hashMapOf<String, String>()
    
      fun setAttribute(attrName: String, value: String) {
        _attributes[attrName] = value
      }
    
      val name: String
        get() = _attributes["name"]!! // 수동으로 맵에서 정보를 꺼낸다.
    
      val name: String: String by _attributes // 위임 프로퍼티로 맵을 사용한다.
    }

- 위의 코드가 동작하는 이유는 Map, MutableMap의 인터페이스에 `getValue(), setValue()` 확장 함수를 제공하기 때문에 가능하다.

