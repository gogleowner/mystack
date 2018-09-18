var jetSet = Set("a", "b")
jetSet += ("c")
println(jetSet.contains("d"))
println(jetSet.contains("c"))

val mutableSet = scala.collection.mutable.Set("aaa", "bbb", "ccc")
mutableSet += "ddd"
println(mutableSet)

val mutableMap = scala.collection.mutable.Map[Int, String]()
mutableMap += 1 -> "a"
mutableMap += 2 -> "b"

println(mutableMap(1) + " "  +  mutableMap(2))

val immutableMap = Map(1 -> "a", 2 -> "b", 3 -> "c")
immutableMap += 5 -> "e"
println(immutableMap)
