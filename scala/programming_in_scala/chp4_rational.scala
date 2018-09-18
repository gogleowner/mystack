class Rational(n: Int, d: Int) {
	require(d != 0)

	private val g = gcd(n.abs, d.abs)
	val numer: Int = n / g
	val denom: Int = d / g

	def this(n: Int) = this(n, 1)

	override def toString = numer + " /  " + denom
	def add(that: Rational) : Rational = {
		new Rational(numer * that.denom + that.numer * denom, denom * that.denom)
	}

	def +(that: Rational) : Rational = add(that)
	def +(num: Int) : Rational = new Rational(numer + num * denom, denom)

	def *(that: Rational) : Rational = new Rational(numer * that.numer, denom * that.denom)
	
	def lessThan(that: Rational) = numer * that.denom < that.numer * denom

	def max(that: Rational) = if (lessThan(that)) that else this

	private def gcd(a: Int, b: Int) : Int = {
		if (b == 0) {
			a
		} else {
			gcd(b, a % b)
		}
	}
}

//implicit def intToRational(x: Int) = new Rational(x)

object Application extends App {
	
implicit def intToRational(x: Int) = new Rational(x)
	val a = new Rational(1, 2)
	val b = new Rational(2, 3)
	println(a+b)
	println(a*b)
	println(a.lessThan(b))
	println(a.max(b))
	println(new Rational(2))

	println(new Rational(22 , 42))

	println(a + 2)
	println(2 + a)
}
