try {
	val file = new java.io.FileReader("input.txt")
} catch {
	case ex: java.io.FileNotFoundException => println("file not found !!" + ex)
} finally {
	println("finally!! ");
}
