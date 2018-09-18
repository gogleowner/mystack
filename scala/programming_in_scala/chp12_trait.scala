trait Philosophical {
  def philosophize() = println("in trait")
}

class Frog extends Philosophical {
  override def toString = "frog"
}
val frog = new Frog()
frog.philosophize
println(frog)
