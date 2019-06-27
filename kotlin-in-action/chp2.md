# Chapter 2. Kotlin basics

# 2.1. 기본 요소 : 함수와 변수

## 2.1.1. Hello World!

    fun main(args: Array<String>) {
      println("Hello World!")
    }

- `fun` : 함수 선언시 사용하는 키워드
- `args: Array<String>` : 파라미터 이름 뒤에 파라미터 타입을 쓴다.
- 함수를 클래스 밖에 정의할 수 있다.
- `Array` 도 일반적인 클래스와 마찬가지이다.
- `System.out.println() → println()` 으로 쓴다.
    - 코틀린 표준 라이브러리는 여러 자바 라이브러리 함수를 간결하게 사용할수 있게 wrapper 들을 제공한다.

        ConsoleKt.class 참고
        package kotlin.io // 코틀린은 기본적으로 kotlin.io 가 import 되어있다. (java의 java.lang 패키지가 import 되어있는 것처럼)
        
        (...)
        
        /** Prints the given [message] and the line separator to the standard output stream. */
        @kotlin.internal.InlineOnly
        public actual inline fun println(message: Any?) {
            System.out.println(message)
        }
        
        (...)

## 2.1.2. 함수

    fun max(a: Int, b: Int): Int {
      return if (a > b) a else b
    }

- 위의 `Hello World` 프로그램의 main 함수는 반환 값이 없는 함수이다.
- 반환 타입은 위와 같이 파라미터 뒤에 붙여주면 된다.
- 코틀린의 if 문은 값을 만들어내지 못하는 statement이 아니고 결과를 만드는 expression 이다.
    - expression : 값을 만들어내며 다른 식의 하위 요소로 계산에 참여할 수 있음.
    - statement : 자신을 둘러싸고 있는 가장 안쪽 블록의 최상위 요소로 존재. 아무런 값을 만들어내지 않는다.
    - 자바에서는 모든 제어 구조가 statement 인 반면, 코틀린에서는 loop 를 제외한 모든 제어 구조가 statement이다.

### expression 이 본문인 함수

    fun max(a: Int, b: Int): Int = if (a > b) a else b

- block 이 본문인 함수 : 본문이 중괄호로 둘러쌓임
- expression 이 본문인 함수 : 등호와 식으로 이루어짐

    fun max(a: Int, b: Int) = if (a > b) a else b

- type inference 로 인해서 반환타입을 생략할수도 있다.
- expression 이 본문인 경우 컴파일러가 함수 본문인 expression을 분석하여 결과 타입을 함수 반환 타입으로 정해준다.
- expression 이 본문인 경우에만 반환타입을 생략할 수 있다. block이 본문인 함수가 값을 반환한다면 반드시 반환타입을 지정하고 return 문을 통해 반환 값을 명시해야 한다.

## 2.1.3. 변수

    val question = "삶, 우주, 그리고 ..." or val question: String = "삶, 우주, 그리고 ..."
    val answer = 42 or val answer: Int = 42
    val yearsToCompute: Double = 7.5e6 // 7.5 * 10^6 = 7500000.0

- 컴파일러가 초기화 expression 을 분석하여 변수 타입으로 지정한다.
- 부동소수점 상수 사용시 변수 타입은 Double이 된다.
- 초기화 값이 없으면 변수에 저장될 값의 정보가 없기 때문에 반드시 타입을 명시해야한다.

### Immutable 변수, mutable 변수

- `val` (value) : immutable 참조 (자바의 final 변수와 같음.)
    - 블록을 실행할 때 정확히 한번만 초기화되야 한다.
    - 한번만 초기화되는 컴파일러가 확인할 수 있다면 조건에 따라 val 값을 다른 여러 값으로 초기화할 수 있다.

            val message: String
            
            if (canPerformOperation()) {
            	message = "Success"
            } else {
              message = "Failed"
            }

- `var` (variable) : mutable 참조

## 2.1.4. 더 쉽게 문자열 형식 지정 : 문자열 템플릿

    fun main(args: Array<String>) {
      val name = if (args.isNotEmpty()) args.first() else "Kotlin"
      println("Hello $name!")
    }

- 변수를 문자열 안에 사용할 경우에 변수 앞에 `$` 를 붙이면 된다
- `$` 문자를 넣고 싶으면 `\` 를 사용하여 `$` 를 escape시켜야 한다.
- 아래와 같이 더 간단하게 출력할수도 있다. 변수에 표현식이 들어간다면 `{}` 로 둘러싸면 된다.

        fun main(args: Array<String>) {
        	println("Hello ${if (args.isNonEmpty()) args.first() else "Kotlin"}")
        }

# 2.2. 클래스와 프로퍼티

    /* java */
    public class Person {
    	private final String name;
    	public Person(String name) { this.name = name; }
      public String getName() { return name; }
    }
    // 필드가 늘어날때마다 생성자, getter 가 반복적으로 추가된다.
    
    /* kotlin */
    class Person(val name: String)
    
    // 코틀린은 기본 access modifier 가 public 이다.

## 2.2.1. 프로퍼티

- 클래스라는 개념의 목적은 데이터를 캡슐화하여 데이터를 다루는 코드를 클래스 내에 가두는 것이다.
- 자바에서는 데이터를 필드에 저장하며, 멤버 필드들은 보통 private 이다.
-
