# Chapter 5. Programming with lambdas

Lambda Expression : 다른 함수에 넘길 수 있는 작은 코드 조각. 람다 사용시에 쉽게 공통 코드 구조를 라이브러리 함수로 뽑아낼 수 있다. 람다를 자주 사용하는 경우로 컬렉션 처리를 들 수 있다. 컬렉션 처리 패턴을 표준 라이브러리 함수에 람다를 넘기는 방식으로 대치하는 예제를 살펴볼거고, 자바 라이브러리와 람다를 함께 사용하는 방법도 살펴본다.

# 5.1. 람다 식과 멤버 참조

## 5.1.1. 람다 소개 : 코드 블록을 함수 인자로 넘기기

    /* annoumous class */
    button.setOnClickListener(new OnClickListener() {
      @Override
      public void onClick(View view) {}
    }); // 클릭 이벤트 발생시 annoymous class 를 통해 이벤트를 처리한다.
    
    /* lambda expression */
    button.setOnClickListener { /* 클릭 이벤트 발생시 수행할 동작  */ }

## 5.1.2. 람다와 컬렉션

이부분은 자바8과 매우 유사하기 때문에 패스한다.

## 5.1.3. 람다식의 문법

- 람다는 값처럼 여기 저기 전달할 수 있는 동작의 모음이다.

    /* 변수로 담아서 처리할 경우 */
    val sum = { x: Int, y: Int -> x + y }
    println( sum(1, 2) ) ==> 3
    
    /* 람다 식을 직접 호출하는 경우 */
    { println(42) }() ==> 42
    
    /* 위의 기능은 가독성도 좋지 않고 쓸모없다. 굳이 코드의 일부분을 블록으로 둘러싸 실행할 필요가 있으면 run을 사용 */
    run { println(42) } // run은 인자로 받은 람다를 실행해주는 라이브러리 함수이다.
    
    /**
     * Calls the specified function [block] with `this` value as its receiver and returns its result.
     */
    @kotlin.internal.InlineOnly
    public inline fun <T, R> T.run(block: T.() -> R): R {
        contract {
            callsInPlace(block, InvocationKind.EXACTLY_ONCE)
        }
        return block()
    }

- 실행 시점에 람다 호출에는 아무 부가 비용이 들지 않으며 프로그램 기본 구성요소와 비슷한 성능을 낸다.
    - 8.2절에서 그 이유를 설명한다. 궁금하군.
- 이름 붙인 인자를 이용하여 람다식 넘기기

        data class Person(val name: String, val age: Int)
        
        val people = listOf(Person("이몽룡", 29), Person("성춘향", 31))
        val names = people.joinToString(seperator = " ", transform = { it.name })

- 람다를 괄호 밖에 선언하기

        val names = people.joinToString(" ") { it.name }

- `it` 을 사용하는 관습은 코드를 아주 간단하게 만들어준다. 단 람다 안에 람다가 중첩되는 경우 파라미터를 명시하는 편이 낫다. 파라미터를 명시하지 않으면 각각의 `it`이 가리키는 파라미터가 어떤 람다에 속했는지 파악하기 어려울 수 있다. 문맥에서 람다 파라미터의 의미, 타입을 쉽게 알 수 없는 경우에도 파라미터를 명시적으로 선언하면 도움이 된다.

## 5.1.4. 현재 영역에 있는 변수에 접근

- 자바와 동일하게 람다식 안에서 람다식 밖의 변수 사용이 가능하다. 단 자바에서는 람다식 밖의 변수를 사용하는 경우 `final` 변수에 접근할 수 있었으나, 코틀린 람다식에서는 `final`이 아닌 변수에 접근 및 변경할 수 있다.

        fun printProblemCounts(responses: Collection<String>) {
          var clientErrors = 0
          var serverErrors = 0
          responses.forEach {
            if (it.startsWith("4")) {
              clientError++
            } else if (it.startsWith("5")) {
              serverError++
            }
          }
        }
        
        /* 자바 */
        void printProblemCounts(Collection<String> responses) {
          int clientErrors = 0, serverErrors = 0;
          responses.forEach(response -> {
            if (response.startsWith("4")) {
              clientErrors++; // 컴파일 에러 발생
            } else if response.startsWith("5")) {
              serverErrors++; // 컴파일 에러 발생
            }
          });
        }

- 위 예제의 람다 안에서 사용하는 외부 변수를 람다가 `capture`한 변수라고 부른다.
- 기본적으로 함수 안의 정의된 로컬변수는 함수가 반환되면 끝난다. 하지만 어떤 함수가 자신의 로컬 변수를 capture한 람다를 반환하거나 다른 변수에 저장한다면 로컬 변수의 생명주기, 함수의 생명주기가 달라질 수 있다.
- capture한 변수가 있는 람다를 저장하여 함수가 끝난 뒤 실행해도 람다의 본문 코드는 여전히 capture변수를 읽거나 쓸 수 있다.
    - `final` 변수가 아닌 경우 : 변수를 특별한 래퍼로 감싸서 나중에 변경, 읽을 수 있게 한 다음 래퍼에 대한 참조를 람다 코드와 함께 저장
    - `final` 변수일 경우 : 람다 코드를 변수 값과 함께 저장

## 5.1.5. 멤버 참조

- 자바의 `Method Reference` 와 상통되는 표현이다.
- 멤버 참조는 프로퍼티나 메소드를 단 하나만 호출하는 함수 값을 만들어 준다.

        val getAge = { p: Person -> p.age }
        val getAge = Person::age
        
        people.maxBy(Person::age)
        people.maxBy {p -> p.age}
        people.maxBy { it.age }
        
        // 최상위에 선언된 function이나 property가 있다면,
        fun salute() = println("salute!")
        run(::salute()) ==> salute!

- `Constructor Reference` 를 사용하면 클래스 생성 작업을 연기하거나 저장해둘 수 있다.

        val createPerson = ::Person
        val p = createPerson("abcde", 29)
        
        fun Person.isAdult() = age >= 21
        val adultPredicate = Person::isAdult

## 5.2. 컬렉션 함수형 API

함수형 프로그래밍에서는 람다나 다른 함수를 인자로 받거나 함수를 반환하는 함수를 고차함수 (HOF, High Order Function) 이라고 부른다. 고차 함수는 기본 함수를 조합하여 새로운 연산을 정의하거나, 다른 고차함수를 통해 조합된 함수를 또 조합해서 더 복잡한 연산을 쉽게 정의할 수 있다는 장점이 있다.

- combinator pattern : 고차함수 + 단순한 함수 조합하여 코드를 작성하는 기법
- combinator : combinator pattern에서 복잡한 연산을 만들기 위해 값 or 함수를 조합할 때 사용하는 고차함수

### 콜렉션에서 사용하는 함수들 중 기억해야할 함수들

    적재적소에 사용하라 : count & size
    count : 특정 조건의 개수를 반환
    size : 중간 컬렉션을 만들어 그 개수를 반환함
    
    people.filter(canBeInClub27).size => 조건에 해당하는 콜렉션 생성 후 개수 반환
    people.count { canBeInClub27 } => 조건에 해당하는 원소의 개수 반환
    

# 5.3. 지연 계산(lazy) 컬렉션 연산

- `map`, `filter` 등은 해당 결과 컬렉션을 즉시 생성한다.
- `sequence`를 사용하면 중간 임시 컬렉션을 사용하지 않고도 컬렉션 연산을 연쇄할 수 있다.

    people
      .map(Person::name) // Collection<String>
      .filter { it.startsWith("A") } // Collection<String>
    
    people
      .asSequence() // Creates a [Sequence] instance that wraps the original collection returning its elements when being iterated.
      .map(Person::name) // Collection<String>
      .filter { it.startsWith("A") } // Collection<String>
      .toList() // 결과 시퀀스를 다시 리스트로 변환

- 원소가 2개밖에 없다면 성능에 큰 영향이 없겠으나 원소가 수백만개라면 훨씬 효율이 떨어진다.
- `Sequence` 인터페이스의 연산자는 `iterator()` 하나 뿐이다.

        public interface Sequence<out T> {
          public operator fun iterator(): Iterator<T>
        }
        
        Sequence 인터페이스의 구현체는..
          map => TransformingSequence
          flatMap => FlatteningSequence
        ... 등등의 구현체들이 있다.

    - 이는 한번에 하나씩 열거될 수 있는 원소의 시퀀스를 표현한다.
    - 시퀀스의 원소는 필요할때 계산된다. 중간 처리 결과를 저장하지 않고 연산을 연쇄적으로 적용해서 효율적으로 계산을 수행할 수 있다.
    - 시퀀스 원소를 차례로 이터레이션 할일이 있다면 굳이 `toList()`로 변환하지 않아도 된다.
    - 시퀀스의 원소를 인덱스를 사용해 접근하는 등 다른 함수들을 이용하려면 리스트로 변환해야 한다.

## 5.3.1. Sequence 연산 실행 : 중간 연산과 최종 연산

- 중간 연산
    - 다른 시퀀스를 반환한다
    - 해당 시퀀스는 최초 시퀀스의 원소를 변환하는 방법을 안다.
    - 항상 지연 계산된다.
    - ex) `map`, `filter` 등..
- 최종 연산
    - 결과를 반환한다.
    - 결과는 최초 컬렉션에 대해 변환을 적용한 시퀀스로부터 일련의 계산을 수행해 얻을 수 있는 컬렉션이나 원소, 숫자 또는 객체다.
    - ex) `toList()`
- sequence, non sequence 연산 비교

        (1..4)
          .map { print("map($it) "); it * it }
          .filter { print("filter($it) "); it % 2 == 0 }
        
        (1..4).asSequence()
          .map { print("map($it) "); it * it }
          .filter { print("filter($it) "); it % 2 == 0 }
          .toList()
        
        map(1) map(2) map(3) map(4) filter(1) filter(4) filter(9) filter(16)
        map(1) filter(1) map(2) filter(4) map(3) filter(9) map(4) filter(16)

    - sequence의 경우 모든 연산은 각 원소에 대해 순차적으로 적용된다.
    - non sequence의 경우 `map`한 결과를 모두 담고난 후 `filter`처리를 한다.
- 자바 Stream, 코틀린 Sequence 와의 비교

    Sequence는 자바의 Stream과 개념이 같다. 코틀린에서 같은 개념을 따로 구현해 제공하는 이유는 안드로이드 등에서 예전 자바버전을 사용하는 경우 Stream이 없기 때문이다.

## 5.3.2. Sequence 만들기 : asSequence(), generateSequence()

- `asSequence()`는 위 예제에서 보았듯, 일반적인 컬렉션을 가지고 함수형 연산을 처리해야하는 경우 사용하면 된다.
- `generateSequence()`

        generateSequence(nextFunction: () -> T?)
        generateSequence(seedFunction: () -> T?, nextFunction: (T) -> T?)
        generateSequence(seed: T?, nextFunction: (T) -> T?)
        
        val naturalNums = generateSequence(0) { it + 1 }

