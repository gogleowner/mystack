# Chapter 9. Generics

# 9.1. Generic Type Parameter

## 9.1.1. Generic Function and Property

    fun <T> List<T>.slice(indices: IntRange): List<T>
    
    val letters = ('a'..'z').toList()
    println(letters.slice<Char>(0..2)) // 타입 명시적으로 지정
    println(letters.slice(0..2)) // 컴파일러가 타입을 추론함
    
    val <T> List<T>.penultimate: T // 끝에서 두번째 원소
      get() = this[size - 2]
    
    println((1..2).toList().penultimate) // 3

## 9.1.2. Define Generic Class

    interface List<T> {
      operator fun get(int: Int): T
    }
    
    class StringList: List<String> {
      override fun get(index: Int): String = ...
    }
    class ArrayList<T>: List<T> {
      override fun get(index: Int): T = ...
    }

## 9.1.3. Type Parameter Constraint

- 클래스나 함수에 사용할 수 있는 타입 인자를 제한하는 기능

    fun <T : Number> List<T>.sum(): T
       타입 : 상한
    
    fun <T : Comparable<T>> max(first: T, second: T): T {
      return if (first > second) first else second
    }

- 타입 파라미터에 여러 제약을 가하기

        fun <T> ensureTrailingPeriod(seq: T)
          where T : CharSequence, T : Appendable { // 타입 파라미터 제약 목록
          if (!seq.endsWith('.')) { // CharSeqeunce 인터페이스의 함수 호출
            seq.append('.') // Appendable 인터페이스의 메소드 호출
          }
        }

