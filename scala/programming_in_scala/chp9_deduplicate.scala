object FileMatcher {
	def fileMatching(matcher: String => Boolean) =
		for (file <- new java.io.File(".").listFiles
			if (matcher(file.getName)))
		yield file

	def filesEnding(query: String) = fileMatching(_.endsWith(query))
	def filesContaining(query: String) = fileMatching(_.contains(query))
}

FileMatcher.filesEnding(".scala").foreach(println)
FileMatcher.filesContaining("deduplicate").foreach(println)

def containsOdd(nums: IndexedSeq[Int]) = nums.exists(_ % 2 == 1)

val numList = for (num <- 1 to 100) yield num

println(containsOdd(numList)

