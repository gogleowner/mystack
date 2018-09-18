object LongLineFileProcessor {
	def processFile(fileName: String, width: Int) = {
		def processLine(line: String) = {
			if (line.length > width) {
				println(fileName + " : " + line.trim)
			}
		}	

		val source = scala.io.Source.fromFile(fileName)
		for (line <- source.getLines())
			processLine(line)
	}
}

if (args.length == 0) throw new java.lang.IllegalArgumentException("file!!!");

LongLineFileProcessor.processFile(args(0), 45)
