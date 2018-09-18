if (args.length > 0) {
	for (line <- scala.io.Source.fromFile(args(0)).getLines())
		println(line.length + " " + line)
} else {
	Console.err.println("please enter file name!")
}

val lineList = scala.io.Source.fromFile(args(0)).getLines().toList
val longestLine = lineList.reduceLeft((l, r) => if (l.length > r.length) l else r).length.toString.length

for (line <- lineList) {
	val blanks = " " * (longestLine - line.length.toString.length)
	println(blanks + line.length + " | " + line)
}
