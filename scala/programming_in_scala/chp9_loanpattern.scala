def withPrintWriter(file: java.io.File, op: java.io.PrintWriter => Unit) = {
	val writer = new java.io.PrintWriter(file)
	try {
		op(writer)
	} finally {
		writer.close()
	}	
}

println(withPrintWriter(new java.io.File("date.txt"), writer => {
	writer.println("blabla")
	writer.println(new java.util.Date)
})
)
