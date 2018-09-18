val input = if (args.length > 0) args(0) else ""

val output = input match {
	case "test" => "case"
	case "seunghyo" => "boram"
	case "boram" => "seunghyo"
	case _ => "hul"
}

println(input + " / " + output)
