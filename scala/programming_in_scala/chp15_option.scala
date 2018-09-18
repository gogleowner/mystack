// option
// 스칼라로 구현된 대부분의 구현체들은 값을 반환할때 Option으로 감싸서 반환함.
// 아래 Map의 get() 메소드도 자바와는 다르게 Option으로 반환함.
val capitals:Map[String, String] = Map("france" -> "paris", "japan" -> "tokyo")

println(capitals get "france") // Some(paris)

// partial function
// case 의 나열도 함수 리터럴.?

val withDefault: Option[Int] => Int = {
  case Some(x) => x
  case None => 0
}

println(withDefault(Some(10))) // 10
println(withDefault(None)) // 0

// partial function with for statement
for ((country, city) <- capitals) // 무조건 정상
  println(s"The capital of ${country} is ${city}")

// 패턴에 일치하는 값만 순회한다.
for (Some(f) <- List(Some("apple"), None, Some("orange"), None)) {
  println(f)
}

