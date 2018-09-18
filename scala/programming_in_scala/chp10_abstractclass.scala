abstract class Element {
  def contents: Array[String]

  def height = contents.length
  def width = if (height == 0) 0 else contents(0).length

  def above(that: Element): Element = new ArrayElement(this.contents ++ that.contents)

  def beside(that: Element): Element = {
    new ArrayElement(
      for (
        (line1, line2) <- contents zip that.contents
      ) yield line1 + line2
    )
  }

  override def toString = contents mkString "\n"
}

class ArrayElement(conts: Array[String]) extends Element {
  def contents: Array[String] = conts

  override def toString: String = conts.toString
}

class LineElement(str: String) extends ArrayElement(Array(str)) {
  override val width = str.length
  override val height = 1
}
/*
 * i don't know about this class 
class UniformElement (
  ch: Char,
  override val width: Int,
  override val height: Int
) extends Element {
  private val line = ch.toString * width
  override def contents = Array.fill(height)(line)
}
*/
object Application extends App {
  val arrayElem = new ArrayElement(Array("hello", "world"))
  printElement(arrayElem)

  val lineElem = new LineElement("hello")
  printElement(lineElem)

//  val uniformElem = new UniformElement('x', 2, 3)
//  printElement(uniformElem)


  private def printElement(element: Element)  {
    println(element.height)
    println(element.width)
    println(element)
    println(element.contents)
  }
}
