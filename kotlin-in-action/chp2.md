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
- 자바에서는 데이터를 필드에 저장하며, 멤버 필드들은 보통 private 이다. 보통은 해당 필드를 가져오기 위해 getter, setter 메소드를 제공한다.
- 필드 &  접근자를 묶어서 프로퍼티라고 부른다.
- 코틀린은 프로퍼티를 언어 기본 기능으로 제공하며, 코틀린의 프로퍼티는 자바의 필드, 접근자 메소드를 완전히 대체한다.

    class Person(
    	val name: String, // 읽기 전용. getter 메소드를 만들어낸다.
      var isMarried: Boolean // 필드는 private, public setter 메소드를 만들어낸다.
    )
    
    val p = Person("Bob", true)
    p.name
    p.isMarried = false

- getter & setter 이름을 정하는 규칙에는 예외가 있다.
    - `isXxx`
        - getter : `get` 접두사가 붙지 않고 이름 그대로 사용한다.  ( `isXxx()` )
        - setter : `is` 를 `set` 으로 바꾼 이름을 사용한다.  (  `setXxx()` )
- 자바로 선언한 클래스에 접근할때 코틀린에서는 getter 를 `val` 처럼 쓸수 있고, getter/setter 쌍은 `var` 처럼 쓸수 있다.

## 2.2.2. 커스텀 접근자

- 프로퍼티 접근자를 커스텀하게 만들수 있다.

    class Rectangle(val height: Int, val width: Int) {
    	val isSquare: Boolean
      get() = height == width // isSquare 접근시마다 값을 매번 계산한다.
    }

- 파라미터가 없는 함수로 만드는 것과 커스텀 게터를 정의하는 방식 모두 비슷하다. 구현이나 성능 차이는 없다.
- 클래스의 특성을 정의하고 싶다면 프로퍼티로 정의해야 한다.

## 2.2.3. 코틀린 소스코드 구조 : 디렉토리와 패키지

- 자바와 동일하게 모든 클래스를 패키지 단위로 관리하는 구조를 가진다. 코틀린의 경우도 파일의 맨 앞에 `package` 문을 넣을 수 있다. 해당 패키지 안에 들어간 클래스, 함수, 프로퍼티들은 직접 사용할 수 있다.
- 다른 패키지에 정의한 클래스, 함수, 프로퍼티를 사용하려면 `import` 문을 통해 불러와야 한다.
- 코틀린은 클래스, 함수 `import` 에 차이가 없고, 모든 선언을 `import` 키워드로 가져올 수 있다.

        package geometry.example
        
        import geometry.shapes.createRandomRectangle // 함수만 import
        import geometry.shapes.* // 패키지 내의 모든 선언을 import
        
        fun main(args: Array<String>) {
          println(createRandomRectangle().isSquare())
        }

- 자바에서는 파일 명과 클래스 명이 일치하여 outer 클래스는 하나만 넣을 수 있고 디렉토리가 패키지를 의미한다.
- 코틀린에서는 여러 클래스를 한 파일에 넣을 수 있고, 파일의 이름도 마음대로 정할 수 있다. 파일에 클래스가 여러개라면 파일명이 패키지명이 된다.

        ㄴ geometry
          ㄴ example.kt // package geometry.example
          ㄴ shapes.kt // package geometry.shapes

# 2.3. 선택 표현과 처리 : enum , when

## 2.3.1. enum 클래스 정의

    enum class Color {
      RED, ORANGE, YELLOW, ..
    }

- 코틀린에서 enum은 `soft keyword`라고 부르는 존재다.
    - `enum` : class 앞에 있을때는 특별한 의미를 지니지만 다른 곳에서는 사용할 수 있다.
    - `class` : 키워드이기 때문에 변수 등을 정의할 때 사용할 수 없어서 clazz, aClass 를 사용한다.

        /* java */
        class Abc {
          String enum; // enum 이 예약어이기 때문에 컴파일 오류가 발생한다.
        }
        
        /* kotlin */
        class Abc(val enum: String) // enum 이 soft keyword이기 때문에 변수명으로 사용 가능하다.

- enum 클래스 안에 프로퍼티, 메소드 정의가 가능하다.

        enum class Color(val r: Int, val g: Int, val b: Int) {
        	RED(255, 0, 0),
        	ORANGE(255, 165, 0),
        	YELLOW(255, 255, 0)
          ;
        
          fun rgb() = (r * 256 + g) * 256 + b
        }

## 2.3.2. when 으로 enum 클래스 다루기

- `when` 은 자바의 `switch`를 대치하되 훨씬 더 강력하다.
- `if` 와 마찬가지로 `when`도 값을 만들어내는 expression 이다.

    fun getMnemonic(color: Color) = 
      when (color) {
        Color.RED -> "Richard"
        Color.ORANGE -> "Of"
        Color.YELLOW -> "York" // 자바와 달리 각 분기 끝에 break 를 넣지 않아도 된다.
        Color.YELLOW, Color.ORANGE -> "York" // 여러 값을 매치 패턴으로 사용하려면 값 사이를 , 로 구분한다.
      }
    
    println(getMnemonic(Color.RED)) => Richard

## 2.3.3. when과 임의의 객체를 함께 사용

    fun mix(c1: Color, c2: Color) = 
      when (setOf(c1, c2)) { // setOf() -> Set 을 생성하는 factory method이다.
        setOf(RED, YELLOW) -> ORANGE
        setOf(YELLO, BLUE) -> GREEN
        else -> throw Exception("Dirty Color")
    	}
    
    println(mix(BLUE, YELLOW))

- when 인자와 일치하는 조건 값을 찾을 때까지 각 분기를 검사한다.
- 분기 조건에 식을 넣을 수 있기 때문에 많은 경우 코드를 더 간결하게 작성할 수 있다.

## 2.3.4. 인자 없는 when 사용

위의 예제는 함수가 호출될 때마다 Set 인스턴스를 생성하여 불필요한 가비지 객체가 많이 늘어난다. 인자가 없는 when 을 사용하면 불필요한 객체 생성을 막을 수 있다.

    fun maxOptimized(c1: Color, c2: Color) = 
      when {
        (c1 == RED && c2 == YELLOW) || (c1 == YELLOW && c2 == RED) -> ORANGE
        (c1 == YELLOW && c2 == BLUE) || (c1 == BLUE && c2 == YELLOW) -> GREEN
        ...
      }

## 2.3.5. 스마트 캐스트 : 타입 검사 & 타입 캐스트

`(1 + 2) + 4` 간단한 서술식을 계산하는 함수를 만들어보자.

    식을 트리 구조로 저장하자.
    노드는 Sum, Num 중 하나이다.
    Num은 leaf 노드지만 Sum 은 non-terminal 노드이다.
    Sum 노드의 두 자식은 덧셈을 위한 두 인자이다.
    
    interface Expr
    class Num(val value: Int): Expr
    class Sum(val left: Expr, val right: Expr): Expr

- 단순하게 `if, else` 로 구현.

        fun eval(e: Expr): Int {
            if (e is Num) {
                return e.value
            } else if (e is Sum) {
                return eval(e.left) + eval(e.right)
            } else {
                throw IllegalArgumentException()
            }
        }
        
        @Test
        fun `eval test`() {
          println(eval(Sum(Sum(Num(1), Num(2)), Num(4))))
        }

- 코틀린에서는 `is` 를 이용하여 변수 타입을 검사한다. 자바의 `instanceOf` 와 유사하다. 자바에서는 타입 검사 후 캐스팅해야 쓸수 있다. 코틀린에서는 변수를 캐스팅하지 않아도 사용 가능하다.

        if (e is Num) {
          val n = e as Num // as 는 타입캐스팅을 하는건데, 위의 is 에서 이미 검사르 했기 때문에 컴파일러가 알아서 캐스팅해주기에 해당문을 생략해도 된다.
        	return n.value
        }

## 2.3.6. 리팩토링 : if를 when으로 변경

    fun eval(e: Expr): Int =
        when (e) {
          is Num -> e.value // if 문과 마찬가지로 타입검사를 하고 나면 스마트캐스트가 이루어진다.
          is Sum -> eval(e.left) + eval(e.right)
          else -> throw IllegalArgumentException()
    		}

## 2.3.7. if와 when의 분기에서 블록 사용

when 도 블록을 넣을 수 있다.

    fun eval(e: Expr): Int =
        when (e) {
          is Num -> {
    				prinltn("num: ${e.value}")
    				e.value
    			}
          is Sum -> {
    				val left = eval(e.left) ; left = eval(e.left)
    				println("sum: $left + $right")
    				left + right
    			}
          else -> throw IllegalArgumentException()
    		}

# 2.4. 대상을 이터레이션 : while, for 루프

## 2.4.1. while 루프 → 자바와 동일하다.

    while (조건) {
    	/* ... */
    }
    do {
    	/* ... */
    } while (조건)

## 2.4.2. 수에 대한 이터레이션 : 범위와 수열

코틀린에서는 for 루프 (어떤 변수를 초기화하여 루프가 한번 실행될 때마다 갱신되고 루프조건이 거짓이 될때 반복을 마치는 형태) 가 없다. 이를 위해 코틀린에서는 `range` 를 사용한다.

    val oneToTen = 1..10

- 코틀린의 `range`는 양 끝을 포함하는 구간이다.
- 어떤 범위에 속한 값을 일정한 순서로 이터레이션하는 경우를 수열(progression)이라고 부른다.

    fun fizzBuzz(i: Int) = when {
      i % 3 == 0 -> "Fizz"
      i % 5 == 0 -> "Buzz"
      i % 15 == 0 -> "FizzBuzz"
      else i.toString()
    }
    
    for (i in 1..100) {
      print(fizzBuzz(i))
    }
    1 2 Fizz 4 Buzz Fizz 7 ...
    
    for (i in 100 downTo 1 step 2) {
    	print(fizzBuzz(i))
    }
    Buzz 98 Fizz 94 92 FizzBuzz 88 ...
    
    100 downTo 1 => 역방향 수열
    step 2 => 이터레이션을 돌때마다 2씩 증가한다. 이 예제는 역방향이니 -2 증가한다.

- 양 끝을 포함하지 않는 범위를 만들고 싶으면 `until` 함수를 사용하면 된다.

        for (x in 0 until size) 는 for (x in 0..size-1) 과 같지만 좀더 명확하다.

## 2.4.3. 맵에 대한 이터레이션

    val binaryReps = TreeMap<Char, String>()
    
    for (c in 'A'..'F') {
      binaryReps[c] = Integer.toBinaryString(c.toInt())
    }
    
    for ((letter, binary) in binaryReps) { // key, value 를 괄호로 감싼다.
      println("$letter = $binary")
    }
    
    A = 1000001
    B = 1000010
    C = 1000011
    D = 1000100
    E = 1000101
    F = 1000110
    
    val list = listOf(10, 11, 1001)
    for ((index, element) in list.withIndex()) { // list를 인덱스와 함께 이터레이션
      println("$index : $element")
    }
    
    0 : 10
    1 : 11
    2 : 1001

## 2.4.4. in 으로 컬렉션이나 범위의 원소 검사

    fun isLetter(c: Char) = 
      c in 'a'..'z' || c in 'A'..'Z' // ==> ('a' <= c && c <= 'z') || ('A' <= c && c <= 'Z')
    fun isNotLetter(c: Char) = 
      c !in '0'..'9' // ==> '0' > c && c > '9'
    
    fun recognize(c: Char) = when (c) {
      in '0'..'9' -> "It's a digit!"
      in 'a'..'z', in 'A'..'Z' -> "It's a letter!"
      else -> "I don't know"
    }

# 2.5. 코틀린의 예외처리

발생한 예외를 함수 호출단에서 처리하지 않으면 함수 호출 스택을 거슬러 올라가면서 예외를 처리하는 부분이 나올 때까지 예외를 다시 던진다.

    if (percentage !in 0..100) {
      throw IllegalArgumentException()
    }
    
    val percentage = 
      if (number in 0..100) number
      else throw IllegalArgumentException() // throw 는 expression 이다.

## 2.5.1. try, catch, finally

    fun readNumber(reader: BufferedReader): Int? {
      try {
    		return reader.readLine().toInt()
      } catch (e: NumberFormatException) {
        return null
    	} finally {
        reader.close()
      }
    }

- 자바에서는 `readLine()` 메소드 사용시에 `IOException` 이 Checked Exception 이기 때문에 명시적으로 예외를 처리하는 코드를 넣어야 하지만, 코틀린에서는 Checked / Unchecked Exception 을 구별하지 않는다.
- `NumberFormatException` 은 Checked Exception이 아니다. 자바 컴파일러는 해당 예외를 강제하지 않으나, 숫자로 파싱이 되지 않을 경우는 흔하기에 넣은 부분이라 할 수 있다.
- `close()` 호출시에도 `IOException`이 Checked Exception이긴 하지만 스트림을 닫다가 실패하는 경우 클라이언트 프로그램이 취할 수 있는 의미있는 동작은 없기에 강제하지 않는다.
- 단, 코틀린에서는 자바7부터 지원하던 `try-with-resource` 문법을 제공하지는 않는다. 라이브러리 함수로 같은 기능을 구현하긴 한다.

## 2.5.2. try를 expression 으로 사용

    val number: Int = try {
      reader.readLine().toInt()
    } catch (e: NumberFormatException) { return } // 예외 발생시 이 다음 코드는 실행되지 않는다.
    
    val number: Int? = try {
      reader.readLine().toInt()
    } catch (e: NumberFormatException) { null } // 예외 발생시에도 다음 코드를 실행하게 하려면 값을 반환하도록 해야한다. nullable값이기 때문에 타입을 Int? 로 준다.
