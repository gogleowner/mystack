val oneToThree = List(1, 2, 3)
println(oneToThree)

val oneTwo = List(1, 2)
val threeFour = List(3, 4)
val oneTwoThreeFour = oneTwo ::: threeFour

println(oneTwo + " and " + threeFour + " were not mutated.")
println("Thus, " + oneTwoThreeFour + "is a new list.")

val twoThree = List(2, 3)
val oneTwoThree = 1 :: twoThree
println(oneTwoThree)

println(List('a', 'b') ::: List('c', 'd'))

val aList = "a" :: "ab" :: "abc" :: "abcd" :: Nil
print(aList.count(value => value.length >=3 ))

print(aList.drop(2))
