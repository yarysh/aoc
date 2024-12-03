import kotlin.io.path.Path
import kotlin.io.path.readText


fun part1(): Int {
    var ans = 0

    val re = Regex("mul[(](\\d+),(\\d+)[)]")
    re.findAll(Path("input.txt").readText()).forEach {
        val (a, b) = it.destructured
        ans += a.toInt() * b.toInt()
    }

    return ans
}


fun part2(): Int {
    var ans = 0

    val text = Path("input.txt").readText()

    var instruction = "do"
    
    var i = 0
    loop@ while (i < text.length) {
        when (instruction) {
            "do" -> {
                if (i+3 < text.length && text.substring(i..i+3) == "mul(") {
                    var j = i+4
                    var commaFound = false
                    while (j < text.length && text[j] != ')') {
                        if (!text[j].isDigit()) {
                            if (text[j] == ',' && !commaFound) {
                                commaFound = true
                            } else {
                                i = j
                                continue@loop
                            }
                        }

                        j++
                    }
                    
                    try {
                        text.substring(i+4..<j).split(",").map { it.toInt() }.let {
                            ans += it[0] * it[1]
                        }
                    } catch (e: Exception) {
                        //
                    }
                    
                    i = j
                    continue
                } else if (i+6 < text.length && text.substring(i..i+6) == "don't()") {
                    instruction = "don't"
                    i += 7
                    continue
                }
            }
            "don't" -> {
                if (i+3 < text.length && text.substring(i..i+3) == "do()") {
                    instruction = "do"
                    i += 4
                    continue
                }
            }
        }
        
        i++
    }

    return ans
}

println("Part1 answer: ${part1()}")
println("Part2 answer: ${part2()}")
