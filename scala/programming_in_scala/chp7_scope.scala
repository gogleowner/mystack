def printMultiTable() = {
	var i = 1
	while (i <= 10) {
		var j = 1
		while (j <= 10) {
			var prod = (i * j).toString
			var k = prod.length
			while (k < 4) {
				print(" ")
				k += 1
			}

			print(prod)
			j += 1
		}
		println()
		i += 1
	}
}

printMultiTable

def printMultiTableFunctionalStyle() = {
	for (i <- 1 to 10) {
		for (j <- 1 to 10) {
			val prod = (i * j).toString
			for (blank <- prod.length until 4) {
				print(" ")
			}
			print(prod)
		}
		println()
	}
}

println
printMultiTableFunctionalStyle

def makeRowSeq(row: Int) =
	for (col <- 1 to 10) yield {
		val prod = (row * col).toString
		val padding = " " * (4 - prod.length)

		prod + padding
	}

def makeRow(row: Int) = makeRowSeq(row).mkString

def multiTable() = {
	val tableSeq = 
		for (row <- 1 to 10) yield makeRow(row)

	tableSeq.mkString("\n")
}

println
println(multiTable)
