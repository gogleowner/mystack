val fileNames = new java.io.File(".").listFiles

for (fileName <- fileNames 
	if fileName.isFile
	if fileName.getName.endsWith(".scala"))
	println(fileName)

def scalaFiles = 
	for {
		file <- fileNames
		if file.getName.endsWith(".scala")
		line <- fileLines(file)
		trimmedLine = line.trim
		if trimmedLine.matches(".*for.*")
	} yield trimmedLine.length

def fileLines(file: java.io.File) =
	scala.io.Source.fromFile(file).getLines().toList

scalaFiles.foreach(println)
