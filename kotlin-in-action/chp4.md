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

TBD..
