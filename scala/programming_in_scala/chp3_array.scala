val big = new java.math.BigInteger("12345")
println(big)

val arrayLength = 3

val greetArr = new Array[String](arrayLength)
greetArr(0) = "hello"
greetArr(1) = ", "
greetArr(2) = "world!"

greetArr.foreach(print)
println()

for (i <- 0 to 2)
	print(greetArr(i))

val greetArr2: Array[String] = new Array[String](arrayLength)
greetArr2.update(0, "hello")
greetArr2.update(1, ", ")
greetArr2.update(2, "world")

for(i <- 0.to(2))
	print(greetArr2.apply(i))

for (num <- Array("one", "two", "three"))
	print(num + " ")

Array("one", "two", "three").foreach(print)
