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
