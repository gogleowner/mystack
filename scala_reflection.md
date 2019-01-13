# scala reflection

## Overview

리플렉션은 프로그램 자체를 검사하고 가능하면 수정하는 기능이다. 이는 객체지향, 함수형, 로직컬한 프로그래밍 패러다임 전반에서 오랜 역사를 지니고 있다. 일부 언어는 리플렉션을 기본 원칙으로 삼고 있지만 많은 언어가 점차적으로 시간이 지남에 따라 리플렉션 능력을 발전시키고 있다.

리플렉션은 프로그램의 함축적인 요소(implicit elements)를 구체화(reify, explicit)하는 기능을 포함한다. 이 요소들은 클래스, 메소드, 표현식과 같은 정적인 프로그램 기능이거나, 현재 연속(current continuation) 또는 실행 이벤트와 같은 메소드를 invoke하는 것과 필드에 엑세스하는 동적 요소일 수 있다. 리플렉션이 수행되는 시기에 따라 컴파일타임 리플렉션, 런타임 리플렉션으로 구분한다. 컴파일타임 리플렉션은 프로그램 변환기와 생성기를 개발하는 강력한 방법이며, 런타임 리플렉션은 일반적으로 언어의 의미 적용에 사용하거나 소프트웨어 요소들 간에 늦은 바인딩에 사용된다.

2.10 버전 이전 까지는, 스칼라 자체의 리플렉션 기능은 없었다. 대신 자바 리플렉션의 api 일부를 사용할 수 있었다. 그러나 많은 스칼라 특화된 요소들은 독립형 자바 리플렉션에서는 제공하기가 어려웠고, 자바와 호환되는 요소로만 사용할 수 있었다. (스칼라에서 제공되는 function, trait, type, existential, higer-kind, path-dependent, abstract types) 게다가 자바의 리플렉션은 제네릭 타입의 제약사항으로 인해 컴파일타임에 런타임 타입의 정보를 알아낼 수 없었다. 

2.10 버전에서는 새로운 리플렉션 라이브러리가 도입되어 스칼라 및 일반 형식에 대한 자바의 런타임 리플렉션의 단점을 해결할 뿐 아니라 더 강력한 툴킷을 스칼라에 적용했다. 스칼라 유형과 제네릭에 대한 완벽한 기능을 갖춘 런타임 리사이클과 함께 스칼라 2.10에는 매크로 형식의 컴파일 타임 반영 기능이 포함되어 있으며 스칼라 표현식을 추상 구문 트리로 구체화 할 수 있습니다.

### Runtime Reflection

런타임 리플렉션은 무엇인가? 런타임에 타입 또는 어떤 object를 인스턴스화 한다.

- 제네릭 타입을 포함하여, object의 타입을 추정하고,
- object를 초기화하고,
- 또는 object의 멤버들에 접근하거나, 호출할 수 있다.

예제를 통해 위의 기능을 어떻게 사용할 수 있는지 봐보자.

#### 제네릭 타입을 포함하여, object의 타입을 추정한다.
다른 JVM 기반의 언어들과 마찬가지로, 스칼라는 컴파일 타임에는 type이 지워진다.(erased) 즉, 일부 인스턴스의 런타임 유형을 검사 할 경우 스칼라 컴파일러가 컴파일 할 때 사용할 수있는 모든 유형 정보에 액세스하지 못할 수도 있습니다.

`TypeTag` 는 컴파일타임에 런타임에 사용될 수 있는 모든 type의 정보를 전달하는 객체로 생각할 수 있다. `TypeTag`는 항상 컴파일러에서 생성됨을 유의해야한다. 이 생성은 `TypeTag`를 요구하는 암시 적 매개 변수 또는 컨텍스트 바운드가 사용될 때마다 트리거됩니다. 즉, 일반적으로 암시 적 매개 변수 또는 컨텍스트 경계를 사용하는 경우에만 `TypeTag`를 얻을 수 있습니다.

```
import scala.reflect.runtime.{universe=>ru}
val l = List(1,2,3)

def getTypeTag[T: ru.TypeTag](obj: T) = ru.typeTag[T]

val theType = getTypeTag(l).tpe
theType: ru.Type = List[Int]
```

위에서 보듯, 우리는 `TypeTag`를 사용하기 위해서 반드시 `import scala.reflect.runtime.universe`를 import해야한다. `l`이라는 리스트를 만들고, 타입을 가져올 수 있는 `getTypeTag` 메소드를 정의했다. 메소드를 통해 리스트의 타입이 `Int` 임을 알아낼 수 있다.

뿐만 아니라 우리는 `Type`에 정의된 symbol들을 알아낼 수 있다.

```
scala> val decls = theType.decls.take(10)
decls: Iterable[ru.Symbol] = List(constructor List, method companion, method isEmpty, method head, method tail, method ::, method :::, method reverse_:::, method mapConserve, method ++)
```


