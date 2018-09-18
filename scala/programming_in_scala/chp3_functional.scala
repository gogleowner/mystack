println("---- imperative style ----")
def printArgs(args: Array[String]): Unit = {
	var i = 0
	while (i < args.length) {
		println(args(i))
		i += 1
	}
}

printArgs(args)

println("---- functional style ----")

def printArgsFunc(args: Array[String]): Unit = {
	args.foreach(println)
}

printArgsFunc(args)


println("---- formatArgs ----")

def formatArgs(args: Array[String]) = args.mkString("\n")
val formattedArgs = formatArgs(Array("one", "two", "three"))
assert(formattedArgs == "one\ntwo\nthree")
