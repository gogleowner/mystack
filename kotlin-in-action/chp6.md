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

        fun <T> printHashCode(t: T) {
          println(t?.hashCode())
        }
        
        printHashCode(null) ==> null

