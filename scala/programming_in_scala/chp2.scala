var i = 0
while (i < args.length) {
	println("Hello " + args(i))
	i += 1
}

args.foreach(arg => println("Hello " + arg))


args.foreach(println)

for (arg <- args)
	println(arg)
