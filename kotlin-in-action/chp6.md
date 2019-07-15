# Chapter 6. The Kotlin type system

# 6.1. Nullable

- 코틀린 타입 시스템은 null이 될 수 있는 타입을 명시적으로 지원한다.

    fun strLenSafe(s: String?): Int = if (s != null) s.length else 0

- 실행시점에 Null이 될 수 있는 타입이나 Null이 될수 없는 타입의 객체는 같다. Null이 될 수 있는 타입은 Null이 될수 없는 타입을 감싼 래퍼가 아니라, 컴파일 시점에 수행된다. 코틀린에서는 Null이 될 수 있는 타입을 처리하는데 별도의 실행시점 부가 비용이 들지 않는다.
- 안전한 호출 연산자 : `?.`
    - null 검사, 메소드 호출을 한번의 연산으로 수행
        - `s?.toUpperCase()` ⇒ `if (s != null) s.toUpperCase() else null`
- 엘비스(elvis) 연산자 : `?:`
    - null 대신 사용할 디폴트 값을 지정할 때 사용할 수 있는 연산자

            fun foo(s: String?) {
              val t: String = s ?: ""
            }

- 안전한 캐스트 : `as?`
    - 자바의 타입캐스트와 마찬가지로 대상 값을 지정한 타입으로 바꿀 수 없으면 `ClassCastException` 발생.

        foo as? Type
          foo is Type  =>  foo as Type
          foo !is Type =>  null
        
        class Person(val firstName: String, val lastName: String) {
          override fun equals(o: Any?): Boolean {
            val other = o as? Person ?: return false
            return other.firstName == firstName && other.lastName == lastName
          }
          override fun hashCode(): Int = firstName.hashCode() * 37 + lastName.hashCode()
        }

- Not-null assertion : `!!`

        foo!!
          foo != null => foo
          foo == null => NullPointerException

- `let` 함수

        foo?.let { it .. }
          foo != null ==> 함수 블록으로 들어옴.
          foo == null ==> 함수 블록에 들어오지 않음.

- 나중에 초기화할 프로퍼티 : `lateinit var`
    - 클래스를 객체로 생성한 다음 나중에 프로퍼티들을 초기화하는 프레임워크가 많다.
        - 안드로이드 : `onCreate` 에서 Activity 를 초기화
        - jUnit : `@Before` 안에서 초기화 로직을 수행해야함.
    - 코틀린에선 클래스 안에 non-nullable 프로퍼티를 생성자 안에서 초기화하지 않고 특별한 메소드 안에서 초기화할 수는 없다. 일반적으로 생성자에서 모든 프로퍼티를 초기화해야한다.
    - `val` 프로퍼티는 `final` 로 컴파일되고, 생성자 안에서 반드시 초기화해야한다.
    - 나중에 초기화하는 프로퍼티는 항상 `var`이여야 한다. 단 그 프로퍼티를 초기화하기 전에 접근하면 `lateinit property fooService has not initialized` 예외가 발생한다.
- null이 될 수 있는 타입 확장

        public inline fun CharSequence?.isNullOrBlank(): Boolean =
          this == null || this.isBlank() // 두번째 this에는 스마트 캐스트가 적용된다.
        
        isNullOrBlank() , isNullEmpty() 는 null을 검사하면서 공백인지 확인한다.

- 타입 파라미터의 null 가능성
    - 함수나 클래스의 모든 타입 파라미터는 기본적으로 null이 될 수 있다.
    - 타입파라미터 T를 클래스나 함수 안에서 타입으로 사용하면 변수명 끝에 `?`가 없더라도 nullable 타입이다.

        fun <T> printHashCode(t: T) {
          println(t?.hashCode())
        }
        
        printHashCode(null) ==> null
        
        // kotlin 1.3.x 부터는 Any?.hashCode()로 선언된 확장함수가 호출된다.
        fun <T> printHashCode(t: T) {
          println(t.hashCode())
        }
        
        printHashCode(null) ==> null ==> 0
        
        @SinceKotlin("1.3")
        @InlineOnly
        public inline fun Any?.hashCode(): Int = this?.hashCode() ?: 0

- nullable과 자바
    - 코틀린은 자바와의 상호운용성을 아주 강조하는 언어이다.
    - 자바에서는 nullable 을 지원하지 않지만, `@Nullable` 어노테이션을 통해서 알수 있긴 하다.
    - 코틀린에서 자바코드를 활용할 경우, 이 정보를 활용하여 nullable인지 판단한다.

        @Nullable String => String?
        @NotNull String => String

    - 코틀린은 아래 어노테이션들을 통해 nullable 여부를 판단한다.
        - JSR-305 표준 `javax.annotation`
        - 안드로이드 `android.support.annotation`
        - JetBrain `org.brains.annotation`
- 플랫폼 타입
    - nullable 어노테이션이 소스코드에 없는 경우, 자바의 타입은 코틀린의 platform type이 된다.
    - 플랫폼 타입은 코틀린이 null 관련 정보를 알 수 없는 타입을 말한다.
    - 이 경우 nullable 프로퍼티로 보던지, final 프로퍼티로 보던지.. 이건 개발자의 몫이다.

            Type == Type?  or  Type

    - 코틀린은 보통 null이 될 수 없는 타입의 값에 대해 null 안전성 검사를 수행하면 경고를 표시하지만, 플랫폼 타입 값에 대해 null 안전성 검사 연산 수행시 경고를 표시하지 않는다.
    - 이 때문에 코틀린에서 자바 API를 다룰 때 조심해야한다.
    - 코틀린이 왜 플랫폼 타입을 도입했는가?
        - 모든 자바 타입을 nullable타입으로 다루면 더 안전하긴 하겠지만, 결코 null이 될 수 없는 값에 대해서도 처리를 해야하기 때문에 불필요한 검사가 들어간다.
        - 제네릭을 다룰 때 상황이 더 나빠진다. `ArrayList<String>` 을 `ArrayList<String?>?` 처럼 다루면 이 리스트의 원소에 접근할 때마다 null검사를 수행해야 한다. 모든 타입의 값에 대해 항상 null 검사를 작성하는 것은 너무 성가신 것이다.
        - 따라서 코틀린 설계자들은 자바의 타입을 가져온 경우 ㅍ개발자에게 그 타입을 제대로 처리할 책임을 부여하는 접근 방법을 택했다.
- 상속
    - 코틀린에서 자바 메소드를 오버라이드할 때 그 메소드의 파라미터,반환타입을 nullable로 선언할지에 대해 결정해야 한다.
    - 플랫폼 타입이기 때문에 두 구현 모두 가능하다.

        /* Java */
        interface StringProcessor {
          void process(String value);
        }
        
        /* Kotlin */
        class StringPrinter : StringProcessor {
          override fun process(value: String) { .. } // value가 null인 경우 예외가 발생함
        }
        class NullableStringPrinter : StringProcessor {
          override fun process(value: String?) { .. }
        }

# 6.2. 코틀린의 원시 타입

코틀린은 원시타입과 래퍼 타입을 구분하지 않는다. 코틀린 내부에서 어떻게 원시 타입에 대한 래핑이 작동하는지도 알아본다.

## 6.2.1. 원시 타입 : Int, Boolean, Char, Byte, Short, Long, Double, Float

- 자바는 primitive, reference 타입을 구분한다.
    - primitive 변수에는 값이 직접 어싸인 되고,
    - reference 변수에는 메모리 상의 객체의 주소가 들어간다.
    - primitive 변수을 콜렉션에 담는 등 참조 타입이 필요할 경우에는 `java.lang.Integer` 등의 래퍼객체를 사용한다.
- 코틀린은 primitive, reference 타입 구분을 하지 않는다.
- 단 컴파일 시점에 숫자 타입은 가장 효율적인 방식으로 표현한다. `Int` 타입은 대부분 `int`로 컴파일된다. 이런 컴파일이 불가능한 경우는 컬렉션과 같은 제네릭 클래스를 사용하는 경우 뿐이다.
    - `List<Int>` ⇒ `List<Integer>`

## 6.2.2. Nullable primitive type : Int?, Boolean? 등

- null 참조는 reference 변수에만 대입할 수 있기 때문에 null 이 될 수 있는 코틀린 타입은 primitive 로 표현할 수 없다. 따라서 코틀린에서 nullable 타입을 사용하면 그 타입은 자바의 reference 타입으로 컴파일된다.

## 6.2.3. 숫자 변환

- 코틀린과 자바의 가장 큰 차이점 중 하나는 숫자를 변환하는 방식이다.
- 코틀린은 한 타입의 숫자를 다른 타입의 숫자로 자동 변환하지 않는다. 결과 타입이 허용하는 숫자 범위가 원래 타입의 범위보다 넓은 경우 조차도 자동 변환은 불가능하다.
- 직접 변환 메소드를 호출해야 한다.

        val intNum = 1
        val longNum = i // Error: type mismatch 컴파일 오류 발생
        
        val longNum = intNum.toLong()

- 코틀린은 Boolean 을 제외한 모든 primitive 타입에 대한 변환 함수를 제공한다.
- 개발자의 혼란을 피하기 위해 타입 변환을 명시하기로 결정했다.
- Boxed Type 을 비교하는 경우 문제가 많다.

        new Integer(42).equals(new Long(42)) => false
        
        val x = 1
        val list = listOf(1L, 2L, 3L)
        
        x in list => 묵시적 타입 변환으로 인해 false
        x.toLong() in list => true

## 6.2.4. Any, Any? : 최상위 타입

- 자바의 Object 클래스가 최상위 클래스이듯, 코틀린에서는 Any 타입이 최상위 타입이다.
- 자바 메소드에서 Object 를 인자로 받거나 반환하면 코틀린에서는 Any 로 타입을 취급한다.

        class Any {
            public open operator fun equals(other: Any?): Boolean
            public open fun hashCode(): Int
            public open fun toString(): String
        }

## 6.2.5. Unit 타입

- 코틀린의 Unit 타입은 자바의 void 와 같은 기능을 한다.
- 코틀린의 Unit과 자바의 void 와 다른점?
    - Unit은 모든 기능을 갖는 일반적인 타입
    - Unit을 타입 인자로 쓸 수 있다.
    - Unit 타입의 함수는 Unit값을 묵시적으로 반환한다.

        interface Processor<T> {
          fun process(): T
        }
        class NoResultProcessor : Processor<Unit> {
          override fun process(): Unit {
            
        	}
        }

- 자바에서 타입인자로 값 없음을 표현할 경우 `java.lang.Void` 타입을 사용하는 방법도 있지만, null 을 반환하기 위해서 `return null;` 을 명시해야 한다. 코틀린에서는 위의 예제처럼 리턴문을 명시하지 않아도 된다.
- Unit 이라는 이름을 사용한 것은, 함수형 프로그래밍에서 전통적으로 **Unit은 단 하나의 인스턴스만 갖는 타입**을 의미해 왔고 그 유일한 인스턴스의 유무가 자바의 void, 코틀린의 Unit을 구분하는 가장 큰 차이이다.

## 6.2.6. Nothing 타입 : 함수가 정상적으로 끝나지 않음.

- 코틀린에는 성공적으로 값을 반환하는 값이 없어, 반환값이라는 개념 자체가 없는 함수가 일부 존재한다.

        fun fail(message: String): Nothing { // 테스트에 예외를 던져서 테스트를 실패시킨다.
          throw IllegalStateException(message)
        }
        
        val address = company.address ?: fail("No address")
        println(address.city)

- Nothing 타입은 아무 값도 갖지 않는다. 따라서 함수의 반환타입 & 반환타입으로 쓰일 타입 파라미터로만 쓸 수 있다.

# 6.3. 컬렉션과 배열

## 6.3.1. Nullable과 컬렉션

    val elementNullableList: List<Int?> // element들이 null이 될수 있다.
    val nullableList: List<Int>? // element들은 null이 될수 없고, 컬렉션 자체가 nullable이다.
    
    elementNullableList.filterNotNull() // null이 아닌 값들만 필터링해서 쓸 수 있다.

## 6.3.2. 읽기 전용과 변경 가능한 컬렉션

- 코틀린에서는 컬렉션 안의 데이터에 **접근**하는 인터페이스 & 컬렉션 안의 데이터 **변경** 인터페이스를 분리했다.

        package kotlin.collections.Collection // 컬렉션 안의 데이터 **접근** 인터페이스
        package kotlin.collections.MutableCollection // 컬렉션 안의 데이터 **변경** 인터페이스

- 컬렉션을 변경할 필요가 있을 때만 MutableCollection을 사용한다.

## 6.3.3. 코틀린 컬렉션과 자바

- 자바의 ArrayList와 HashSet은 코틀린의 MutableCollection을 확장한다.

## 6.3.4. 컬렉션을 플랫폼 타입으로 다루기

- 자바와 코틀린 코드를 오갈 경우, 자바 코드에서 정의한 타입을 코틀린에서는 플랫폼 타입으로 본다.
- 플랫폼 타입의 경우 코틀린 쪽에는 null 관련  정보가 없어서 Nullable , non-nullable 타입 둘다 허용한다.
- 대부분 문제가 되지 않지만, 컬렉션 타입이 시그니처에 들어간 자바 메소드 구현을 오버라이드 할 경우 immutable collection, mutable collection간의 차이가 문제가 된다. 이 경우 여러가지를 고려하여 선택해야 한다.
    - 컬렉션이 null이 될 수 있는가?
    - 컬렉션의 원소가 null이 될 수 있는가?
    - 오버라이드하는 메소드가 컬렉션을 변경할 수 있는가?

    interface FileContentProcessor {
      void processContents(File path, byte[] binaryContents,
                           List<String> textContents)
    }
    
    class FileIndexer : FileContentProcessor {
      override fun processContents(path: File, BinaryContents: ByteArray?, textContents: List<String>?) {
      }
    }
    
    interface DataParser<T> {
      void parseData(String input, List<T> output, List<String> errors)
    }
    
    class PersonParser : DataParser<Person> {
      override fun parseData(input: String, output: MutableList<Person>, errors: MutableList<String?>) {}
    }

## 6.3.5. 객체의 배열과 원시 타입의 배열

- 코틀린 배열은 타입 파라미터를 받는 클래스이다. 배열의 원소 타입은 타입 파라미터에 의해 정해진다.

    val letters = Array<String>(26) { i -> ('a'+i).toString() }
    val strings = listOf("a", "b", "c")

- 원시타입의 배열을 표현하는 별도 클래스들을 제공한다.
    - 각 배열 타입의 생성자는 `size` 인자를 받아서 해당 원시 타입의 디폴트 값으로 초기화된 배열을 반환한다.
    - 가변인자로 받아 해당 값이 들어간 배열을 반환하는 factory method 를 제공한다.
    - 크기와 람다를 인자로 받는 생성자를 사용한다.

    val fiveZeros = IntArray(5)
    val fiveZeros = intArrayOf(0, 0, 0, 0, 0)
    val squares = IntArray(5) { i -> (i+1) * (i+1) }
    
    squares.forEachIndexed { index, element -> 
      println("element $index is $element")
    }

