# Chapter 4. Classes, objects, and interfaces

- 코틀린의 클래스와 인터페이스는 자바와 약간 다르다.
    - 코틀린의 경우, 인터페이스에 프로퍼티 선언이 들어갈 수 있다.
    - 자바와 달리 코틀린은 기본적으로 `final`, `public` 이다.
- inner class 에는 outer class 에 대한 참조가 없다.
- 클래스를 생성할 때 생성자 구문으로 대부분 다 처리할 수 있지만, 필요하면 접근자를 직접 정의할 수 있다.
- 코틀린 컴파일러는 번잡스러움을 피하기 위해 유용한 메소드를 자동으로 만들어준다.
    - `data class` : 컴파일러가 `equals()`, `hashCode()`, `toString()` 을 만들어 준다.
    - 코틀린에서 제공하는 delegation을 사용하면 위임 처리를 위한 준비 메소드를 직접 작성할 필요가 없다.
- `object` 키워드 : 클래스와 인스턴스를 동시에 선언하면서 만들 때 쓴다.
    - singleton class, companion object, object expression (자바의 annoymous class) 를 표현할 때 `object` 키워드를 쓴다.

# 4.1. 클래스 계층 정리

## 4.1.1. 코틀린 인터페이스

- 기본 사용법

        interface Clickable {
          fun click()
        }
        
        class Button : Clickable { // 코틀린에서는 : 으로 상위 인터페이스를 표현한다.
          override fun click() = println("I was clicked") // override 는 꼭 사용해야 한다.ㅖ
        }
        
        Button().click() ==> I was clicked

- default method

        interface Clickable {
          fun click()
          fun showOff() = println("I'm clickable")
        }
        
        interface Focusable {
          fun setFocus(b: Boolean) = println("I ${if (b) "got" else "lost"} focus.")
          fun showOff() = println("I'm focusable")
        }

    - 한 클래스에서 이 두 인터페이스를 함께 구현하면 컴파일 오류가 발생한다.

            class Button : Clickable, Focusable {
              override fun click() = println("I was clicked")
              override fun showOff() {
                super<Clickable>.showOff()
                super<Focusable>.showOff()
              }
            }

        - 이름과 시그니처가 같은 멤버 메소드에 대해 둘 이상의 디폴트 구현이 있는 경우 인터페이스를 구현하는 하위클래스에서 명시적으로 새로운 구현을 작성해야 한다.
        - 상위 타입의 이름을 `<>` 사이에 넣어서 `super` 로 지정하면 어떤 상위 타입의 멤버 메소드를 호출할지 지정할 수 있다.
- 자바에서 코틀린의 메소드가 있는 인터페이스 구현
    - 코틀린은 자바 6과 호환되게 설계됐다. 디폴트 메소드는 지원하지 않는다.
    - 따라서 코틀린은 디폴트 메소드가 있는 인터페이스를 일반 인터페이스와 디폴트 메소드 구현이 static 메소드로 들어있는 클래스를 조합하여 구현한다.
    - 인터페이스에는 메소드 선언만 들어가며, 인터페이스와 함께 생성되는 클래스에는 모든 디폴트 메소드 구현이 static 메소드로 들어간다.

## 4.1.2. open, final, abstract 변경자 : 기본적으로 final

- 어떤 클래스의 상속을 허용하려면 클래스 앞에 `open` 키워드를 붙여야 한다. 재구현을 허용할 프로퍼티나 메소드에도 `open` 키워드를 붙여야 한다.

        open class RichButton : Clickable {
        	fun disable() {}
          open fun animate() {}
          override fun click() {}
        }

- 오버라이드하는 메소드 구현을 하위클래스에서 오버라이드 하지 못하게 금지하려면 메소드 앞에 `final`을 명시해야 한다.

        open class RichButton : Clickable {
          final override fun click() {}
        }

- 클래스가 기본적으로 상속 가능 상태를 `final`로 함으로써 얻는 큰 이익은 다양한 경우에 스마트 캐스트가 가능하다는 점이다. 클래스 프로퍼티의 경우 `val`이면서 커스텀 접근자가 없는 경우에 스마트 캐스트를 이용할 수 있다. 프로퍼티가 `final`이 아니라면 그 프로퍼티를 다른 클래스가 상속하면서 커스텀 접근자를 정의함으로써 스마트 캐스트의 요구사항을 깰 수 있다.
- `abstract class` : 자바 처럼 코틀린에서도 사용 가능하다.

        abstract class Animated {
          abstract fun animate() // abstract는 추상 클래스에서만 붙일 수 있음.
          open fun stopAnimating() {}
          fun animateTwice() {}
        }

## 4.1.3. Visibility Modifier : 기본적으로 public

- 코틀린에서는 기본적으로 모두 public 이다.
- private, protected는 있지만 package-private는 없다. 코틀린은 패키지를 namespace를 관리하기 위한 용도로만 사용한다.
- 위의 대안으로 `internal`이라는 visibility modifier 를 도입했다.
    - `internal` : 모듈 내부에서만 볼 수 있음.
        - 모듈 : 한번에 한꺼번에 컴파일되는 코틀린 파일들. gradle, maven 등의 프로젝트가 모듈이 될 수 있음.
        - 자바에서는 패키지가 같은 클래스를 선언하기만 하면 어떤 프로젝트의 외부에 있는 코드라도 패키지 내부에 있는 패키지 전용 선언에 의해 쉽게 접근이 가능하다. 때문에 모듈의 캡슐화가 깨지게 된다.
- 최상위 선언에 대해 `private`을 하용한다.
    - 최상위에 선언된 클래스, 함수, 프로퍼티 등은 파일 내에서만 사용할 수 있다.
    - 이 또한 하위 시스템의 자세한 구현 사항을 외부에 감추고 싶을때 유용한 방법이다.

[Visibility Modifier](https://www.notion.so/a9000854498248eaacb0b578b174795d)

## 4.1.4. 내부 클래스와 중첩된 클래스 : 기본적으로 중첩 클래스

- 코틀린에서도 클래스 안에 다른 클래스를 선언할 수 있다. 다만 nested class 명시적으로 요청하지 않는 한 바깥쪽 클래스 인스턴스에 대한 접근 권한이 없다.

        interface State : Serializable
        interface View {
          fun getCurrentState(): State
          fun restoreState(state: State) {}
        }
        
        // Button 클래스의 상태를 저장하는 클래스는 Button 클래스 내부에 선언하면 편하다.
        
        /* 자바로 표현 */
        public class Button implements View {
        	@Override
          public State getCurrentState() {
            return new ButtonState(); // Button 의 상태를 담은 클래스를 반환하지만, 이를 직렬화하면 오류가 발생한다.
          }
        
          public class ButtonState implements State {}
        	// inner class 가 된다.
        }
        
        /*
        클래스 안에 정의한 클래스는 자동으로 inner class 가 된다. ButtonState는 Button에 대한 참조를 묵시적으로 포함한다.
        그 참조로 인해서 ButtonState를 직렬화할 수 없다. 이를 해결하려면 ButtonState를 static 으로 선언해야 한다.
        */
        
        /* 코틀린으로 표현 */
        class Button : View {
          override fun getCurrentState(): State = ButtonState()
          class ButtonState : State {} // 자바의 static class와 대응된다.
        }

[Untitled](https://www.notion.so/41fb45d86012454bbb1cb8bc9647d990)

    class Outer {
      inner class Inner {
        fun getOuterReference(): Outer = this@Outer // inner class에서 outer class의 참조에 접근하려면
      }
    
    }

## 4.1.5. sealed class : 클래스 계층 정의시 계층 확장 제한

- 상위 클래스를 상속한 하위 클래스 정의를 제한할 수 있다. sealed class 의 하위 클래스를 정의할 때는 반드시 상위 클래스 안에 중첩시켜야 한다.

        sealed class Expr { // sealed로 봉인!
          class Num(val value: Int) : Expr()
          class Sum(val left: Expr, val right: Expr) : Expr()
        }
        
        fun eval(e: Expr): Int {
          when (e) { // 모든 하위 클래스만을 대상으로 검사하기 때문에 else 분기가 없어도 된다.
            is Expr.Num -> e.value
            is Expr.Sum -> eval(e.left) + eval(e.right)
          }
        }

- 내부적으로 `sealed class`는 `private` 생성자를 가진다.
- `sealed interface` 는 정의할 수 없다. 봉인된 인터페이스를 만들 수 있다면 인터페이스를 자바 쪽에서 구현하지 못하게 막을 수 있는 수단이 코틀린 컴파일러에겐 없기 때문이다.

# 4.2. 뻔하지 않은 생성자와 프로퍼티를 갖는 클래스 선언

- primary 생성자 : 클래스를 초기화할 때 주로 사용되는 간략한 생성자. 클래스 본문 밖에서 정의한다.
- secondary 생성자 : 클래스 본문 안에서 정의한다.
- initializer block : 초기화 로직

## 4.2.1. 클래스 초기화 : 주 생성자와 초기화 블록

- primary constructor : 클래스 이름 뒤에 오는 괄호로 둘러싸인 코드

        class User constructor(_nickname: String) { // 파라미터가 하나만 있는 주 생성자
        	val nickname: String
          init { // 초기화 블록
            nickname = _nickname
          }
        }

- `constructor` : 주 생성자, 부 생성자 정의를 시작할 때 사용
- `init` : 객체가 만들어질 때 실행될 초기화 코드가 들어간다. 주 생성자와 함께 사용된다.
- 주생성자 앞에 별도의 어노테이션, Visibility Modifier 가 없다면 `constructor`를 생략해도 된다.

        class User(_nickname: String) { val nickname = _nickname }

- 주 생성자의 파라미터로 프로퍼티를 초기화 한다면 주 생성자 앞에 `val`을 추가하여 프로퍼티 정의 & 초기화를 간략하게 쓸 수 있다.

        class User(val nickname: String)
        
        class User(val nickname: String, 
        					 val isSubscribed: Boolean = true) // 생성자 파라미터에 대한 Default 값 제공
        
        val me = User("gogleowner")
        println(me.isSubscribed) ==> true

- 모든 생성자 파리미터에 default 값을 지정하면 파라미터가 없는 생성자를 만들어 준다.
    - 자바의 라이브러리 중에는 파라미터가 없는 생성자를 통해 객체를 생성해야만 되는 경우가 있는데 이 기능이 코틀린과의 통합을 쉽게 해준다.
- 상속받은 클래스에서 상위 클래스의 생성자를 호출해야할 경우

        open class User(val nickname: String)
        class TwitterUser(nickname: String) : User(nickname) { .. }
        
        open class Button
        class RadioButton : Button() // 아무런 인자를 받지 않는 생성자를 가진 상위 클래스를 호출해야 한다.

    - 클래스를 상속받는 경우 반드시 `()` 가 들어간다.
    - 인터페이스를 구현하는 경우 생성자가 없기 때문에 괄호가 없다.
- 어떤 클래스를 클래스 외부에서 인스턴스화하지 못하게 하고 싶다면

        class Secretive private constructor() {}

    - 유틸리티 함수를 담는 클래스는 인스턴스화할 필요가 없고, 싱글톤인 클래스는 미리 정한 팩토리 메소드 등을 통해 객체를 생성해야한다. 자바에서는 private 생성자를 정의하여 클래스를 다른 곳에서 인스턴스화하지 못하게 막는 경우가 생긴다.
    - 코틀린은 언어레벨에서 유틸리티 함수는 최상위 함수로, 싱글톤을 사용하고 싶으면 `object`로 선언하면 된다.

## 4.2.2. 부 생성자 : 상위 클래스를 다른 방식으로 초기화

- 일반적으로 코틀린에서는 디폴트 파라미터 값을 통해서 생성자가 자바에 비해 적다. 자바에서 오버로드한 생성자가 필요한 상황 중 상당 수는 코틀린의 디폴트 파라미터 값, 이름 붙은 인자 문법을 사용해 해결 가능하다.
- 디폴트 값 제공을 위해 부 생성자를 여럿 만들지 말고, 파라미터 디폴트 값을 생성자에 직접 명시하면 더 좋다.
- 프레임워크 클래스를 확장하는데 여러 가지 방법으로 인스턴스를 초기화할 수 있게 다양한 생성자를 지원해야하는 경우 부 생성자가 여럿 필요할 수 있다.

    open class View {
      constructor(ctx: Context) { .. }
      constructor(ctx: Context, attr: AttributeSet) { .. }
    }
    
    class MyButton : View {
    	constructor(ctx: Context) : super(ctx) { .. }
    	constructor(ctx: Context, attr: AttributeSet) : super(ctx, attr) { .. }
    
      // 혹은 이렇게 다른 생성자에 위임할 수 있다.
      constructor(ctx: Context) : this(ctx, MY_STYLE) { .. }
    }

## 4.2.3. 인터페이스에 선언된 프로퍼티 구현

    interface User {
      val nickname: String
    }
    
    class PrivateUser(override val nickname: String) : User
    
    class SubscribingUser(val email: String) : User {
      override val nickname: String
        get() = email.substringBefore("@") // custom getter => 호출될 때마다 substring 수행
    }
    
    class FacebookUser(val accountId: String) : User {
      override val nickname = getFacebookName(accountId) // 한번만 초기화
    }

## 4.2.4. getter, setter 에서 뒷받침하는 필드에 접근

    class User(val name: String) {
      var address: String = "unspecified"
        set(value: String) {
          println("Address was changed for $name: $field -> $value")
          field = value // 프로퍼티 변경
        }
    }

## 4.2.5. 접근자의 가시성 변경

    class LengthCounter {
      var counter: Int = 0
        private set // 클래스 밖에서 이 프로퍼티의 값을 변경할 수 없다.
    
      fun addWord(word: String) {
        counter += word.length
      }
    }

# 4.3. 컴파일러가 생성한 메소드 : data class, 클래스 위임

## 4.3.1. 모든 클래스가 정의해야 하는 메소드

자바와 마찬가지로 코틀린 클래스도 `toString(), equals(), hashCode()` 등을 오버라이드할 수 있다.

    class Client(val name: String, val postalCode: Int) {
      override fun toString() = "Client(name=$name, postalCode=$postalCode")
    }
    
    val c1 = Client("gogleowner", 12345)
    val c2 = Client("gogleowner", 12345)
    
    println(c1 == c2) ==> false

- 자바에서는 ==를 primitive 타입과 reference 타입을 비교할 때 사용한다.
    - primitive : 두 피연산자의 값이 같은지 비교한다. (equality)
    - reference : 두 피 연산자의 주소가 같은지 비교한다. (reference comparison)
    - 따라서 두 객체의 equality 를 비교하려면 `equals()` 를 호출해야 한다.
- 코틀린에서는 ==는 내부적으로 `equals()`를 호출하여 객체를 비교한다. 클래스가 `equals()`를 오버라이드한다면 `==`를 통해 안전하게 인스턴스를 비교할 수 있다.
- reference  비교를 위해서는 `===` 연산자를 사용할 수 있다.
- `equals()`를 오버라이드할 때 반드시 `hashCode()`도 오버라이드 해야 한다.
    - 위 예제에서 `equals()` 를 구현했다고 가정하면, c1과 c2는 프로퍼티가 일치하기 때문에 새 인스턴스와 집합의 기존 인스턴스는 동일하다. 두 인스턴스가 set 에 속했는지 확인하면 true가 반환되어야하지만.. false가 반환된다.

            val aset = hashSetOf(Client("gogleowner", 12345))
            println(aset.contains(Client("gogleowner", 12345)) ==> false

    - 이는 `hashCode()`를 재정의하지 않아서이다. JVM 언어에서는 `equals()` 가 true를 반환하는 두 객체는 반드시 같은 `hashCode()`를 반환해야한다는 제약이 있는데, 위에서는 이를 어기고 있다.

## 4.3.2. data class : 모든 클래스가 정의해야하는 메소드를 자동 생성

- data class는 주 생성자에 정의된 프로퍼티를 기준으로 `toString(), equals(), hashCode()` 등을 자동으로 구현해준다.
- 이 외에 몇가지 유용한 메소드들을 더 생성해준다.
- data class and immutablity : copy() method
    - data class 는 모든 프로퍼티를 읽기 전용으로 만들어 immutable 클래스로 만들라고 권장한다.
    - `copy` 메소드는 객체를 복사하면서 일부 프로퍼티를 바꿀 수 있게 해주는 메소드이다. 복사본은 원본과 다른 생명주기를 가지며 복사를 하면서 일부 프로퍼티 값 바꾸거나, 제거해도 원본을 참조하는 다른 부분에 전혀 영향을 끼치지 않는다.

## 4.3.3. 클래스 위임 : by 키워드 사용

- 상속을 허용하지 않는 클래스에 새로운 동작을 추가해야할 경우 일반적으로 데코레이터 패턴이 쓰인-다.
- 상속을 허용하지 않는 클래스 대신 사용할 수 있는 새로운 클래스(데코레이터)를 만들되 기존 클래스와 같은 인터페이스를 데코레이터가 제공하게 만들고, 기존 클래스를 데코레이터 내부에 필드로 유지 하는 것이다. 새로 정의해야하는 기능은 데코레이터의 메소드에 필드로 유지하고, 새로 정의해야하는 기능은 데코레이터에 새로 정의하면 된다.

        class DelegatingCollection<T> : Collection<T> {
          private val innerList = arrayListOf<T>()
        
          override val size: Int get() = innerList.size
          override fun isEmpty(): Boolean = innerList.isEmpty()
          ..
        }

- 이 경우 `by` 키워드를 통해 인터페이스에 대한 구현을 다른 객체에 위임 중이라는 사실을 명시할 수 있다.

        class DelegatingCollection<T>(
            innerList: Collection<T> = arrayListOf<T>()
          ) : Collection<T> by innerList {
          // Collection에 정의된 메소드를 재구현하고 싶으면 override
          // 해당 클래스 특화 기능으로 추가하고 싶으면 메소드를 구현하면 된다.
        }

# 4.4. object 키워드 : 클래스 선언과 인스턴스 생성

- `object` 키워드 사용 경우
    - object declaration : 싱글턴을 정의하는 방법 중 하나.
    - companion object : 인스턴스 메소드는 아니지만 어떤 클래스와 관련 있는 메소드, 팩토리 메소드를 담을 때 쓰인다.
    - 객체 식 : annoymous inner class 대신 사용

## 4.4.1. object declaration : 싱글턴을 쉽게 만들기

- 생성자는 객체 선언에 사용할 수 없다.

    object Payroll {
      val allEmployees = arrayListOf<Person>()
      fun calculateSalary() {
        for (person in allEmployees) { .. }
      }
    }
    
    object CaseInsensitiveFileComparator : Comparator<File> {
      override fun compare(file1: File, file2: File): Int {
        return file1.path.compareTo(file.path, ignoreCase = true)
      }
    }

## 4.4.2. compaion object : 팩토리 메소드와 static 멤버가 들어갈 장소

- 코틀린은 자바의 `static`키워드를 지원하지 않는다. 대신 패키지 수준의 최상위 함수, 객체, 멤버 선언을 활용한다.
- 하지만 최상위 함수는 `private` 으로 표시된 클래스의 멤버변수에는 접근할 수 없다.
- 클래스의 인스턴스와 관계 없이 호출해야 하지만, 클래스 내부 정보에 접근해야 하는 함수가 필요할 경우 클래스에 중첩된 객체 선언의 멤버 함수로 정의해야 한다.

    class A {
      companion object {
        fun bar() {
          println("companion object called")
        }
      }
    }
    
    A.bar() ==> companion object called
    
    class User private constructor(val nickname: String) {
      companion object {
        fun newSubscribingUser(email: String) = User(email.substringBefore("@"))
        fun newFacebookUser(accountId: String) = User(getFacebookName(accountId))
      }
    }

- 팩토리 메소드 → 목적에 따라 팩토리 매소드의 이름을 정할 수 있다.

## 4.4.3. companion object 를 일반 객체처럼 사용

- compaion object 는 클래스 안에 정의된 일반 객체다. 이름을 붙이거나, 인터페이스를 상속하거나, 여러 확장함수, 프로퍼티를 정의할 수 있다.

        class Person(val name: String) {
          companion object Loader {
            fun fromJSON(jsonText: String): Person = ..
          }
        }
        
        Person.Loader.fromJSON("{name: 'blabla'}")
        Person.fromJSON("{name: 'blabla'}")

    - 이름을 지정하지 않으면 자동으로 `Compaion`이 된다.
- 인터페이스 상속

        interface JsonFactory<T> {
          fun fromJSON(jsonText: String): T
        }
        
        class Person(val name: String) {
          companion object : JsonFactory<Person> {
            override fun fromJSON(jsonText: String): Person = ...
          }
        }

- 자바에서 코틀린 클래스의 `static` 멤버변수에 접근하기 위해서는 `@JvmStatic` 을 코틀린 멤버변수에 붙이면 된다. static 필드가 필요하다면 `@JvmField` 를 최상위 프로퍼티나 객체에서 선언된 프로퍼티 앞에 붙인다.
- 확장

        class Person(val firstName: String, val lastName: String) {
          companion object { // 비어있는 companion object 를 꼭 선언해야함
          }
        }
        
        fun Person.Companion.fromJson(jsonText: String) : Person { .. }

## 4.4.4. 객체 식 : annoymous class를 다른 방식으로 작성

    window.addMouseListener {
      object : MouseAdapter() { // annoymous class 선언
        override fun mouseClicked(e: MouseEvent) {}
        override fun mouseEntered(e: MouseEvent) {}
      }
    }

- 이는 싱글톤이 아니다. 호출될때마다 새로운 인스턴스가 생성된다
- 자바의 annoymous class 는 값이 final 이 아니면 조작할 수 없다. 코틀린에서는 final이 아닌 변수도 안에서 사용할 수 있다.

        fun countClicks(window: Window) {
          var clickCount = 0
          window.addMouseListener(object : MouseAdapter() {
            override fun mouseClicked(e: MouseEvent) {
              clickCount++
            }
          }
        }

- 객체 식은 무명 객체 안에서 여러 메소드를 오버라이드해야 하는 경우에 유용하다. 메소드가 하나 뿐인 인터페이스를 구현해야할 경우에는 lambda 를 사용하는 편이 낫다.

