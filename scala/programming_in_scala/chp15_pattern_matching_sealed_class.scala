// 1. case class
abstract class Expr
case class Var(name: String) extends Expr
case class Number(num: Double) extends Expr
case class UnOp(operator: String, arg: Expr) extends Expr
case class BinaryOp(operator: String, left: Expr, right: Expr) extends Expr

// 2. pattern match
def simplifyTop(expr: Expr): Expr = expr match {
  case UnOp("-", UnOp("-", e)) => e
  case BinaryOp("+", Number(0)) => e
  case BinaryOp("*", Number(1)) => e
  case _ => expr
}

def describe(x: Any) = x match {
  case 5 => "five"
  case true => "truth"
  case "hello" => "hi!"
  case Nil => "empty list!"
  case _ => "something else!"
}

// 3. sealed case class
//

sealed abstract class SealedExpr
case class Var(name: String) extends SealedExpr
case class Number(num: Double) extends SealedExpr
case class UnOp(operator: String, arg: SealedExpr) extends SealedExpr
case class BinaryOp(operator: String, left: SealedExpr, right: SealedExpr) extends SealedExpr

// sealed 로 선언하면 하위 클래스들이 match조건에 없으면 컴파일러가 경고를 발생한다.
// 이 경고를 피하려면 @unchecked 를 사용하거나, case _ 를 추가하면 되는데 이거는 추천하지 않는 방법이다.
def describe(e: SealedExpr): String = (e: @unchecked) match {
  case Number(_) => "a number"
  case Var(_) => "a variable"
}
