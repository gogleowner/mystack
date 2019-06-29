# Chapter 3. Defining and calling functions

# 3.1. 코틀린에서 컬렉션 만들기

    val set = hashSetOf(1, 7, 53)
    val list = arrayListOf(1, 7, 53)
    val map = hashMapOf(1 to "one", 7 to "seven", 53 to "fifty-three") // to 는 특별한 키워드가 아니라 일반 함수이다.
    
    // javaClass 는 자바의 getClass()에 해당하는 코틀린 코드이다.
    println(set.javaClass) ==> class java.util.HashSet
    println(list.javaClass) ==> class java.util.ArrayList
    println(map.javaClass) ==> class java.util.HashMap

- 코틀린은 자신만의 컬렉션 기능을 제공하지 않고, 기존 자바 컬렉션을 활용할 수 있게 해준다.
- 따라서 자바 코드 & 코틀린 코드와의 상호작용이 쉽게 되어있다.
- 코틀린 컬렉션은 자바 컬렉션과 똑같은 메소드이지만, 자바보다 더 많은 기능을 제공한다.

        val strings = listOf("first", "second", "third")
        string.first() ==> first
        string.last() ==> third
        
        val numbers = setOf(1, 14, 2)
        numbers.max() ==> 14

# 3.2. 함수를 호출하기 쉽게 만들기

`toString()` 메소드에 prefix, postfix, seperator 를 인자로 넘길 수 있는 메소드를 만들어보자.

    fun <T> joinToString(collection: Collection<T>, 
    										 seperator: String, prefix: String, postfix: String): String {
      val result = StringBuilder(prefix)
    
      for ((index, element) in collection.withIndex()) {
        if (index > 0) result.append(seperator)
    		result.append(element)
      }
      result.append(postfix)
    
    	return result.toString()
    }
    
    val list = listOf(1, 2, 3)
    println(joinToString(list, "; ", "(", ")")) ==> (1; 2; 3)

## 3.2.1. 이름 붙인 인자

- 함수 호출부분의 가독성

    위 메소드를 `joinToString(list, " ", " ", ".")` 로 호출한다면 인자로 전달한 각 문자열이 어떤 역할을 하는지 구분하기 어렵다.

    코틀린에서는 함수의 인자에 이름을 명시할 수 있다.

    `joinToString(list, seperator = " ", prefix = " ", postfix = ".")`

## 3.2.2. 디폴트 파라미터 값

- 자바에서는 일부 클래스에서 overloading 한 메소드가 너무 많아져서 문제가 있다. 오버로딩 메소드들은 하위 호환성 유지 or 사용자에게 편의를 더하는 이유로 만들어지지만, 중복이 많아진다.
    - 자바의 Thread 클래스의 생성자는 무려 8개나 있다. ([https://docs.oracle.com/javase/8/docs/api/java/lang/Thread.html#constructor.summary](https://docs.oracle.com/javase/8/docs/api/java/lang/Thread.html#constructor.summary))
- 코틀린에서는 함수 선언에서 파라미터의 디폴트 값을 지정할 수 있으므로 오버로딩 중 상당수를 피할 수 있다.

        fun <T> joinToString(collection: Collection<T>,
        										 seperator: String = ", ",
        										 prefix: String = "",
        										 postfix: String = ""): String { ... }

- 자바에는 디폴트 파라미터 값이라는 개념이 없어서 코틀린 함수를 자바에서 호출하는 경우에 모든 값을 명시해야 한다. 자바에서 좀더 편하게 코틀린 함수를 호출하고 싶을 경우 `@JvmOverloads` 어노테이션을 함수에 적용해놓으면 된다. 코틀린 컴파일러가 자동으로 맨 마지막 파라미터로부터 파라미터를 하나씩 생략한 오버로딩된 자바 메소드를 추가해준다.

        // 위의 joinToString() 위에 @JvmOverloads 를 붙이면 아래 4개의 메소드가 만들어진다.
        <T> String joinToString(Collection<T> collection, String seperator, String prefix, String postfix);
        <T> String joinToString(Collection<T> collection, String seperator, String prefix);
        <T> String joinToString(Collection<T> collection, String seperator);
        <T> String joinToString(Collection<T> collection);

## 3.2.3. 정적인 유틸리티 클래스 없애기 : 최상위 함수와 프로퍼티

- 자바는 모든 메소드는 어떤 클래스의 하위에 있어야 한다. 다만 비슷한 역할을 하는 클래스가 둘 이상 있을 경우가 생기게 되기도 한다.
- 그 결과 유틸성 역할을 하는 static 메소드를 모아두는 클래스가 만들어지고, 해당 클래스에 상태나 인스턴스 메소드가 없는 클래스가 만들어진다.
    - `java.util.Colllections' ([https://docs.oracle.com/javase/8/docs/api/?java/util/Collections.html](https://docs.oracle.com/javase/8/docs/api/?java/util/Collections.html))
- 코틀린에서는 함수를 최상위 수준에 정의할 수 있어서 이런 클래스를 생성하지 않아도 되고, import 만 하면 된다.
- 코틀린 최상위 함수가 포함되는 클래스의 이름을 바꾸고 싶다면 파일에 `@JvmName` 어노테이션을 추가하여 클래스명을 지정할 수 있따.

        @file:JvmName("StringFunctions")
        
        package strings
        
        fun joinToString( .. ): String { .. }
        
        /* 자바 */
        import strings.StringFunctions;
        StringFunctions.joinToString(list, ", ", "", "");

- 최상위 수준에 프로퍼티도 지정이 가능하다.

        const val UNIX_LINE_SEPERATOR = "\n" // const는 primitive, String타입만 지정할 수 있다.
        ==> public static final String UNIX_LINE_SEPERATOR = "\n";

# 3.3. 메소드를 다른 클래스에 추가 : 확장 함수와 확장 프로퍼티

코틀린의 핵심 목표 중 하나는, 기존 자바 코드 & 코틀린 코드를 자연스럽게 통합하는 것이다. 코틀린을 기존 자바 프로젝트에 통합하는 경우에는 코틀린으로 직접 변환할 수 없거나 미처 변환하지 않은 기존 자바 코드를 처리할 수 있어야 한다. 확장 함수가 이런 역할을 해줄 수 있다.

확장함수는 어떤 클래스의 멤버 메소드인 것처럼 호출할 수 있지만 그 클래스 밖에 선언되어있다.

    package strings
    fun String.lastChar(): Char = this.get(this.length - 1)
       수신객체타입                 수신객체
    
    // this는 생략 가능하다.
    
    println("kotlin".lastChar()) ==> n

확장 함수 내부에서는 수신 객체의 메소드, 프로퍼티를 바로 사용할 수 있다. 다만 클래스 내부에서만 사용할 수 있는 private, protected 멤버 변수는 사용할 수 없다.

## 3.3.1. 임포트와 확장 함수

- 확장함수를 사용하기 위해서는 그 함수를 다른 클래스나 함수와 마찬가지로 임포트해야만 한다.

    import strings.lastChar
    
    val c = "kotlin".lastChar()

## 3.3.2. 자바에서 확장 함수 호출

- 내부적으로 확장 함수는 수신 객체를 첫번째 인자로 받는 static 메소드이다. 그래서 확장 함수를 호출해도 다른 adapter 객체나 실행시점 부가 비용이 들지 않는다.
- 위 확장 함수를 `StringUtil.kt` 파일에 정의했다면 자바에서는 `StringUtilKt.lastChar("java");` 로 호출하면 된다.

## 3.3.3. 확장 함수로 유틸리티 함수 정의

    fun <T> Collection<T>.joinToString(seperator: String = ", ", prefix: String = "", postfix: String = ""): String {
      val result = StringBuilder(prefix)
    	for ((index, element) in this.withIndex()) { // this는 수신객체를 가리킨다.
        if (index > 0) result.append(seperator)
        result.append(element)
      }
      result.append(postfix)
    	return result.toString()
    }

## 3.3.4. 확장 함수는 override 할 수 없다.

    open class View {
      open fun click() = println("View clicked")
    }
    
    class Button: View() {
      override fun click() = println("Button clicked")
    }
    
    // Button 클래스는 View클래스의 click() 을 위임하여 구현하였으므로 Button 의 click() 메소드가 동작한다.

- 확장 함수는 클래스의 일부가 아니라, 클래스 밖에 선언된다.
- 확장함수를 기반 클래스, 하위 클래스에 정의하더라도 컴파일되면 클래스 밖에 선언되어 정적 메소드로 호출되기 때문에 오버라이딩 되지 않는다.
- 어떤 클래스를 확장한 함수와 그 클래스의 멤버 함수의 이름과 시그니처가 같다면 멤버 함수가 호출된다. (멤버 함수의 우선순위가 더 높다.)

## 3.3.5. 확장 프로퍼티

- 기존 클래스 객체에 대한 프로퍼티 형식의 구문으로 사용할 수 있는 API를 추가할 수 있다. 프로퍼티라고 부르긴 하지만 상태를 저장할 적절한 방법이 없기 때문에 실제로 확장 프로퍼티는 아무 상태도 가질 수 없다.

        val String.lastChar: Char
        	get() = get(length - 1) // 기본 getter 구현을 제공할 수 없으므로 getter는 꼭 정의해야 한다.
        
        
        /* 변경 가능한 확장 프로퍼티 선언하기 */
        var StringBuilder.lastChar: Char
          get() = get(length - 1)
          set(value: Char) {
            this.setCharAt(length - 1, value)
          }
        
        println("kotlin".lastChar) ==> n
        val sb = StringBuilder("Kotlin")
        sb.lastChar = "!"
        println(sb) ==> Kotlin!
        
        StringUtilKt.getLastChar("Java");

# 3.4. 컬렉션 처리 : 가변 길이 인자, 중위 함수 호출, 라이브러리 지원

- `vararg` : variable argument
- `infix` 함수 호출 구문 사용하면, 인자가 하나 뿐인 메소드를 간편하게 호출할 수 있다.
- 구조 분해 선언 (destructing declaration) : 복합적인 값을 분해하여 여러 변수에 나눠 담을 수 있다.

## 3.4.1. 자바 컬렉션 API 확장

- 컬렉션의 `first()`, `last()` 메소드는 코틀린으로 작성한 확장함수이다.

    fun <T> List<T>.last(): T {
    	if (isEmpty())
        throw NoSuchElementException("List is empty.")
      return this[lastIndex]
    }
    
    public fun <T : Comparable<T>> Iterable<T>.max(): T? {
        val iterator = iterator()
        if (!iterator.hasNext()) return null
        var max = iterator.next()
        while (iterator.hasNext()) {
            val e = iterator.next()
            if (max < e) max = e
        }
        return max
    }

## 3.4.2. 가변 인자 함수 : 인자의 개수가 달라질 수 있는 함수 정의

    val list = listOf(2, 1, 3, 4, 5)
    public fun <T> listOf(vararg elements: T): List<T> = 
      if (elements.size > 0) elements.asList() else emptyList()

- variable argument 문법
    - 자바 : `aaa(T... values)`
    - 코틀린 : `aaa(vararg values: T)`
- 이미 배열에 들어있는 원소를 가변 길이 인자로 넘길 때는 변수 앞에 `*`만 붙여주면 된다.

        fun main(args: Array<String>) {
          val list = listOf("args: ", *args)
        	println(list)
        }

## 3.4.3. 값의 쌍 다루기 : 중위 호출과 구조 분해 선언

    val map = mapOf(1 to "one", 7 to "seven", 53 to "fifty-three")

- 여기서 `to` 키워드는 코틀린 키워드가 아니라 `infix call` 로, `to` 라는 일반 메소드를 호출한 것이다.
- infix call 시에는 수신 객체와 유일한 메소드 인자 사이에 메소드 이름을 넣는다. 이때, 객체, 메소드 이름, 유일한 인자 사이에는 공백이 들어가야 한다.

        1.to("one") // to 메소드를 일반적인 형식으로 호출
        1 to "one"  // to 메소드를 infix call 형식으로 호출

- 인자가 하나 뿐인 일반 메소드, 확장 메소드에 infix call을 사용할 수 있다.

        infix fun <A, B> A.to(that: B): Pair<A, B> = Pair(this, that)
        
        val (number, name) = 1 to "one"

- Pair 를 통해 두 변수를 즉시 초기화 할 수 있다. 이를 구조 분해 선언 (destructuring declaration) 이라고 부른다.

        for ((index, element) in collection.withIndex()) { // (index, element)도 구조 분해 선언이라 볼 수 있다.
          println("$index : $element")
        }

# 3.5. 문자열과 정규식 다루기

- 코틀린 코드가 만들어낸 문자열을 아무 자바 메소드에 넘겨도 되고, 자바 코드에서 받은 문자열을 아무 코틀린 라이브러리 함수에 넘겨도 문제 없다.
- 코틀린은 다양한 확장 함수를 제공함으로써 표준 자바 문자열을 더 다루기 쉽게 해준다. 혼동이 야기될 수 있는 일부 메소드에 대해 더 명확한 코틀린 확장 함수를 제공함으로써 개발자의 실수를 줄여준다.

## 3.5.1. 문자열 나누기

- `split()` 메소드의 인자는 정규식 문자열이다.

        "12.345-6.A".split(".") ==> [12, 345-6, A] 를 기대하지만 빈 배열이 반환된다.

- 이런 혼동을 없애기 위해 코틀린에서는 여러 가지 다른 조합의 파라미터를 받는 `split()` 확장함수를 제공한다.
- 정규식을 파라미터로 받는 함수는 `String` 이 아닌 `Regex` 타입의 값이다.

        "12.345-6.A".split("\\.|-".toRegex()) ==> [12, 345, 6, A]
        "12.345-6.A".split(".", "-") ==> [12, 345, 6, A]

## 3.5.2. 정규식과 3중 따옴표로 묶은 문자열

- 전체 파일 경로명을 디렉토리 / 파일이름 / 확장자로 구분하는 함수를 만들어보자.
    - String 확장함수를 이용

            /Users/yole/kotlin-book/chapter.adoc
            
            fun parsePath(path: String) {
            	val dir = path.substringBeforeLast("/")
              val file = path.substringAfterLast("/")
              val fileName = file.substringBeforeLast(".")
              val fileExtension = file.substringAfterLast(".")
            
              println("Dir: $dir, name: $fileName, ext: $fileExtension")
            }

    - 경로 파싱에 정규식 사용하기

            """(.+)/(.+)\.(.+)""".toRegex().matchEntire(path)?.let {
              val (dir, fileName, fileExtension) = it.destructured
              println("Dir: $dir, name: $fileName, ext: $fileExtension")
            }

## 3.5.3. 여러 줄 3중 따옴표 문자열

- 여러줄 문자열에는 들여쓰기, 줄 바꿈을 포함한 모든 문자가 들어간다.
- 이부분은.. 겪어보면 아는 내용일거라 과감히 생략하겠다.

# 3.6. 코드 다듬기 : 로컬 함수와 확장

    class User(val id: Int, val name: String, val address: String)
    
    fun saveUser(user: User) {
    	if (user.name.isEmpty()) {
        throw IllegalArgumentException("Can't save user ${user.id}: empty Name")  // 코드 중복 발생
      }
    	if (user.address.isEmpty()) {
        throw IllegalArgumentException("Can't save user ${user.id}: empty Address")  // 코드 중복 발생
      }
    
      // save user
    }
    
    // 1차 개선
    fun saveUser(user: User) {
      fun validate(value: String, fieldName: String) {
        if (value.isEmpty()) {
    			throw IllegalArgumentException("Can't save user ${user.id}: empty $fieldName")
        }
      }
    
      validate(user.name, "Name")
      validate(user.address, "Address")
    
      // save user
    }
    
    // 2차 개선
    fun User.validateBeforeSave() {
      fun validate(value: String, fieldName: String) {
        if (value.isEmpty()) {
    			throw IllegalArgumentException("Can't save user $id: empty $fieldName")
        }
      }
    
      validate(name, "Name")
      validate(address, "Address")
    }
    
    fun saveUser(user: User) {
      user.validateBeforeSave()
      // save user
    }

