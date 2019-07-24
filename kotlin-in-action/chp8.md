# Chapter 8. Higher-order functions: lambdas as parameters and return values

# 8.1. Define Higher-order function

- 람다나 함수 참조를 인자로 넘길 수 있거나 반환하는 함수이다.
- 람다나 함수 참조를 사용하여 함수를 값으로 표현할 수 있다.

## 8.1.1. 함수 타입

    val sum = { x: Int, y: Int -> x + y }
    val action = { println(42) }
    
    // 타입을 명시한다면
    val sum: (Int, Int) -> Int = { x, y -> x + y }
    val action: () -> Unit = { println(42) }

- 컴파일러는 위 두 변수가 함수 타입 임을 추론한다.
- 함수 타입을 정의하려면 함수 파라미터 타입을 괄호 안에 넣고, 그 뒤에 `->` 를 추가하여 함수타입을 지정하면 된다.

        (Int, String) -> Unit
          파라미터 타입    반환 타입

## 8.1.2. 인자로 받은 함수 호출

    fun twoAndThree(operation: (Int, Int) -> Int) {
      val result = operation(2, 3)
      println("The result is $result")
    }
    
    fun String.filter(predicate: (Char) -> Boolean): String {
      return buildString {
        for (index in 0 until length) {
          val element = get(index)
          if (predicate(element)) append(element)
        }
      }
    }

## 8.1.3. 자바에서 코틀린 함수 타입 사용

- 컴파일 코드 안에서 함수 타입은 일반 인터페이스로 바뀐다.
- `FunctionN` 인터페이스를 구현하는 객체를 저장한다. 코틀린 표준 라이브러리는 함수 인자의 개수에 따라 `Function0<R> , Function1<P1, R> ...` 등의 인터페이스를 제공한다.
- 각 인터페이스에는 `invoke()` 메소드를 구현하도록 되어있어, 호출하면 된다.

    fun processTheAnswer(f: (Int) -> Int) {
      println(f(42))
    }
    
    // 람다 식으로 표현
    processTheAnswer(num -> num + 1);
    
    // Annomous Class로 표현
    processTheAnswer(
      new Function1<Integer, Integer>() {
        @Override
        public Integer invoke(Integer num) {
          System.out.println(num);
          return num + 1;
        }
      }
    );

- 반환타입이 `Unit`인 함수나 람다를 자바로 작성할 수도 있다. 하지만 코틀린 Unit 타입에는 값이 존재하므로 자바에서는 그 값을 명시적으로 반환해줘야 한다.

        List<String> strings = new ArrayList<>();
        strings.add("42");
        CollectionKt.forEach(strings, s -> {
          System.out.println(s);
          return Unit.INSTANCE;
        });

## 8.1.4. 디폴트 값을 지정한 함수 타입 파라미터나 nullable 함수 타입 파라미터

- 디폴트 값을 지정한 함수 타입 파라미터

    fun <T> Collection<T>.joinToString(
      seperator: String = " ", prefix: String = "", postfix: String = "",
      transform: (T) -> String = { it.toString() } // 함수 타입 파라미터 선언, 디폴트값 지정
    ): String {
      val result = StringBuilder(prefix)
      for ((index, element) in this.withIndex()) {
        if (index > 0) result.append(seperator)
    		result.append(transform(element)) // 변환함수 호출
      }
      result.append(postfix)
      return result.toString()
    }

- nullable 함수 타입 파라미터

        fun <T> Collection<T>.joinToString(
          seperator: String = " ", prefix: String = "", postfix: String = "",
          transform: (T) -> String? = null // nullable 함수 타입 파라미터 선언
        ): String {
          val result = StringBuilder(prefix)
          for ((index, element) in this.withIndex()) {
            if (index > 0) result.append(seperator)
            val transformed = transform?.invoke(element) ?: element.toString(
        		result.append(transformed)
          }
          result.append(postfix)
          return result.toString()
        }

## 8.1.5. 함수를 함수에서 변환

    enum class Delivery { STANDARD, EXPEDITED }
    
    class Order(val itemCount: Int)
    
    fun getShippingCostCalculator(delivery: Delivery): (Order) -> Double {
      return when (delivery) {
        Delivery.STANDARD -> { order -> 6 + 2.1 * order.itemCount }
        Delivery.EXPEDITED -> { order -> 1.2 * order.itemCount }
      }
    }
    
    
    class ContactListFilters {
      var prefix: String = ""
      var onlyWithPhoneNumber: Boolean = false
    }
    
    data class Person(val firstName: String, val lastName: String, phoneNumber: String?)
    
    class ContactListFilters {
      var prefix: String = "",
      var onlyWithPhoneNumber: Boolean = false
      fun getPredicate(): (Person) -> Boolean {
        val startWithPrefix = {
          it.firstName.startsWith(prefix) || it.lastName.startsWith(prefix)
        }
    
        return if (!onlyWithPhoneNumber) {
          startsWithPrefix
        } else {
          return { startsWithPrefix(it) && it.phoneNumber != nul` }
        }
      }
    }
    
    val contacts = listOf(
      Person("abc", "def", "123-4567"), Person("ghi", "klm", null)
    )
    
    val contactListFilters = ContactListFilters()
    
    with (contactListFilters) {
      prefix = "Dm"
      onlyWithPhoneNumber = true
    }
    
    println( contacts.filter(contactListFilters.getPrediate()) )
    
    { Person(firstName=abc, lastName=def, phoneNumber=123-4567) }

## 8.1.6. 람다를 활용한 중복 제거

- 예제 데이터 정의

        data class SiteVisit(val path: String, val duration: Double, val os: OS)
        
        enum class OS { WINDOWS, LINUX, MAC, IOS, ANDROID }
        
        val log = listOf(
          SiteVisit("/", 34.0, OS.WINDOWS),
          SiteVisit("/", 22.0, OS.MAC),
          SiteVisit("/login", 12.0, OS.WINDOWS),
          SiteVisit("/signup", 8.0, OS.IOS),
          SiteVisit("/", 16.3, OS.ANDROID),
        )

- 각 OS별 사용자의 평균 방문 시간

        fun List<SiteVisit>.averageDurationFor(os: OS) =
          filter { it.os == os }.map(SiteVisit::duration).average()

- 위의 함수로는 모바일 디바이스 사용자의 평균 방문시간을 구하는 등의 복잡한 요구사항은 해결하지 못한다.
- 아예 predicate 조건을 파라미터로 받아서 사용자가 원하는 것만 필터링하도록 해보자.

        fun List<SiteVisit>.averageDurationFor(predicate: (SiteVisit) -> Boolean) =
          filter(predicate).map(SiteVisit::duration).average()
        
        log.averageDurationFor({ it.os in setOf(OS.ANDROID, OS.IOS })

# 8.2. 인라인 함수 : 람다의 부가 비용 없애기

- 코틀린은 보통 람다식을 annoymous class로 컴파일하지만 그렇다고 람다 식을 사용할 때마다 새로운 클래스가 만들어지는 것은 아니다.
- 단 람다가 외부 변수를 사용할 때는 실행시점마다 annoymous class를 매번 생성한다.
- 코틀린 컴파일러는 반복되는 코드를 별도의 라이브러리 함수로 빼내되 컴파일러가 자바의 일반 명령문만큼 효율적인 코드를 생성할 수 있도록 해준다.
- `inline` 키워드를 함수에 붙이면 컴파일러는 그 함수를 호출하는 모든 문장을 함수 본문에 해당하는 바이트코드로 바꿔치기 해준다.

## 8.2.1. 인라이닝이 작동하는 방식

    inline fun <T> synchronized(lock: Lock, action: () -> T): T {
      lock.lock()
      try {
        return action()
      } finally {
        lock.unlock()
      }
    }

- inline 으로 함수를 선언하면 자바에서 메소드 블록 내에 synchronized 를 건 것과 같아진다.

    fun foo(l: Lock) {
      println("before sync")
      synchronized(l) {
        println("Action")
      }
      println("after sync")
    }
    
    // 위의 코드는 아래와 같은 바이트코드로 컴파일된다.
    fun foo(l: Lock) {
      println("before sync")
      l.lock()
      try {
        println("Action")
      } finally {
        l.unlock()
      }
      println("after sync")
    }

## 8.2.2. 인라인 함수의 한계

- 모든 함수를 인라이닝 할 수는 없다.
- 함수가 인라이닝될 때 그 함수에 인자로 전달된 람다 식의 본문은 결과 코드에 직접 들어갈 수 있다.
- 함수 본문에 파라미터로 받은 람다를 호출한다면 가능하지만, 파라미터로 받은 람다를 다른 변수에 저장하고 나중에 그 변수를 사용한다면 람다를 표현하는 객체는 어딘가에 존재해야 하기 때무넹 람다를 인라이닝할 수 없다.
- 인라이닝할 수 없는 경우 컴파일러가 `Illegal usage of inline-parameter` 예외를 발생시킨다.

    class TransformingSequence<T, R>(s: Sequence<T>, transform: (T) -> R) : Sequence<R> {
        override fun iterator(): Iterator<R> {
            TODO("not implemented") //To change body of created functions use File | Settings | File Templates.
        }
    }
    
    fun <T, R> Sequence<T>.map(transform: (T) -> R): Sequence<R> =
      TransformingSequence(this, transform)
    // transform 람다식을 TransformingSequence 로 넘길 경우에는 inline 불가
    
    
    inline fun foo(inlined: () -> Unit, noinline notInlined: () -> Unit) {
      // ..
    }
    // 특정 람다식 파라미터는 인라인 되지 않아야할 경우 noinline 키워드를 넣는다.
    

## 8.2.3. 컬렉션 연산 인라이닝

- 코틀린의 컬렉션 연산에 필요한 `filter, map, flatMap` 등은 모두 `inline` 함수이다. 따라서 직접 연산을 구현하지 않더라도 성능은 보장된다.
- 단, `filter {...} .map {...} .filter {..}` 이러 형태의 연산은 계속해서 구간별로 중간 컬렉션을 만든다. element가 많아지면 이에 따른 부가비용이 커지기 때문에 `asSequence` 를 통해 중간 시퀀스를 람다를 필드에 저장하는 객체로 표현하여 최종 연산에서 람다식을 연쇄 호출하도록 하는 편이 성능상으로는 더 낫다.
- 단 시퀀스는 람다를 저장해야하므로 람다를 inline하지 않는다. 지연 연산을 통해 성능을 향상시키려는 이유로 모든 컬렉션에 `asSequence`를 붙여서는 안된다.
- 시퀀스 연산에서는 람다가 인라인 되지 않기 때문에 크기가 작은 컬렉션은 오히려 성능이 더 안좋을수도 있다.

## 8.2.4. 함수를 인라인으로 선언해야 하는 경우

- 일반 함수 호출의 경우 JVM이 이미 강력하게 인라이닝을 지원한다. JVM은  실행시점에도 최적화를 하고 있다. 따라서 각 함수 구현이 정확히 한번만 있다면 JVM이 최적화를 수행한다.
- 코틀린 인라인 함수는 바이트 코드에서 각 함수 호출 지점을 함수 본문으로 대치하기 때문에 코드 중복이 생긴다.
- 람다를 인자로 받는 함수를 인라인 할 경우 이점이 많다.
    - 람다를 표현하는 클래스와 람다 인스턴스에 해당하는 객체 생성이 필요없어 진다.
    - JVM이 함수 호출과 람다를 인라이닝해줄 정도로 똑똑하진 못하다.
    - 인라이닝 사용시 일반 람다에서 사용할 수 없는 기능 몇가지가 있다. (`non-local` 등)
- 주의해야 할 점
    - 코드 크기에 주의를 기울여야한다. 인라인 함수가 클 경우 함수 본문에 해당하는 바이트코드를 모든 호출 지점에 복사하면 바이트코드가 전체적으로 커질 수 있다.

## 8.2.5. 자원 관리를 위해 인라인된 람다 사용

- 여기서 말하는 자원은 file, database 등등이다. 보통 `try~finally` 에서 자원을 획득하고, 해제하는 작업을 한다.
- `try~with~resources` 구문을 `use` 함수에서 제공하고 있다.

    String readFirstLineFromFile(String path) throws IOException {
      try (BufferedReader br = new BufferedReader(new FileReader(path))) {
        return br.readLine();
      }
    }
    
    fun readFirstLineFromFile(path: String) {
      BufferedReader(FileReader(path)).use { it.readLine() }
    }

- use 함수 또한 inline 함수이다.

# 8.3. 고차 함수 안에서 흐름 제어

- 람다식 안에서 `return` 문은 람다로부터 값을 반환하는게 아니라 함수의 실행을 끝내는 것이다.
- `non-local return` : 자신을 둘러싸는 블록보다 더 바깥의 다른 블록을 반환하게 만드는 return문
- 이렇게 return 이 바깥쪽 함수를 반환시킬 수 있는 때는 람다를 인자로 받는 함수가 인라인 함수일 경우에만 해당한다.

## 8.3.2. 람다로부터 반환 : label을 이용한 return

    fun lookForAlice(people: List<Person>) {
      person.forEach label@ { // label을 연다.
        if (it.name == "Alice") return@label // 반환과 동시에 label을 ㅏㄷ는다.
      }
      println("alice might be somewhere")
    }
    
    fun lookForAlice(people: List<Person>) {
      person.forEach { if (it.name == "Alice") return@forEach }
      println("alice might be somewhere")
    }

- 람다식의 레이블을 명시하면 함수 이름을 레이블로 사용할 수 없다.
- 람다식에서는 레이블이 2개 이상 붙을 수 없다.

## 8.3.3. Annoymous function : 기본적으로 local return

- label을 이용한 식은 다소 장황해보인다. annoymous function은 기본적으로 local return이다.

    fun lookForAlice(people: List<Person) {
      people.forEach( fun (person) {
        if (person.name == "Alice") return
        println("${person.name} is not Alice")
      })
    }
    
    people.filter(fun (person): Boolean {
      return person.age < 30
    })
    
    people.filter(fun (person) = person.age < 30
    })

