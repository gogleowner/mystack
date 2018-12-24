# ScalaTest Matchers

Custom Matchers 를 구현해보자.

## Json Diff를 활용해보자.

```
import org.json4s._
import org.json4s.jackson.JsonMethods._

val a = parse(
  s"""
     |{
     |  "field": "value"
     |}
       """.stripMargin)

val b: JValue = parse(
  s"""
     |{
     |  "field": "value",
     |  "field2": "value2"
     |}
       """.stripMargin)


val diffResult = Diff.diff(a, b)

if (diffResult.changed != JNothing) {
  println(s"changed : ${diffResult.changed}")
}
if (diffResult.added != JNothing) {
  println(s"added : ${diffResult.added}")
}
if (diffResult.deleted != JNothing) {
  println(s"deleted : ${diffResult.deleted}")
}
```
