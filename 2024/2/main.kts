import kotlin.io.path.Path
import kotlin.io.path.readLines
import kotlin.math.abs

fun part1(): Int {
    var ans = 0

    Path("input.txt").readLines().forEach{
        ans += if (isSafe(it.split(" ").map { it.toInt() })) 1 else 0
    }
    
    return ans
}

fun part2(): Int {
    var ans = 0

    lineloop@ for (line in Path("input.txt").readLines()) {
        val lvl = line.split(" ").map { it.toInt() }

        if (isSafe(lvl)) {
            ans++
            continue@lineloop
        }
        
        for (idx in lvl.indices) {
            if (isSafe(lvl.filterIndexed { i, _ -> i != idx })) {
                ans++
                continue@lineloop
            }
        }
    }

    return ans
}

fun isSafe(lvl: List<Int>): Boolean {
    val decr = (lvl[0]-lvl[1]) < 0
    for (i in 0 until lvl.size-1) {
        val diff = lvl[i] - lvl[i + 1]

        if (decr != (diff < 0) || abs(diff) !in 1..3) {
            return false
        }
    }
    
    return true
}


println("Part1 answer: ${part1()}")
println("Part2 answer: ${part2()}")
