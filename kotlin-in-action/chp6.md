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

