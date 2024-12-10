import kotlin.io.path.Path
import kotlin.io.path.readLines

fun part1(): Int {
    val trailsMap = Path("input.txt").readLines()

    var ans = 0
    trailsMap.forEachIndexed { i, ln ->
        ln.forEachIndexed { j, ch -> if (ch == '0') {
            ans += findTrailScore(trailsMap, Pair(i, j))
        }}
    }

    return ans
}

fun part2(): Int {
    val trailsMap = Path("input.txt").readLines()

    var ans = 0
    trailsMap.forEachIndexed { i, ln ->
        ln.forEachIndexed { j, ch -> if (ch == '0') {
            ans += findTrailRating(trailsMap, Pair(i, j))
        }}
    }

    return ans
}

fun findTrailScore(trailsMap: List<String>, start: Pair<Int, Int>): Int {
    var steps = mutableSetOf(start)

    for (step in 1..9) {
        val nextSteps = mutableSetOf<Pair<Int, Int>>()

        steps.forEach {
            for ((i, j) in listOf(Pair(0, 1), Pair(0, -1), Pair(1, 0), Pair(-1, 0))) {
                val stepI = it.first + i
                val stepJ = it.second + j

                if (stepI >= 0 && stepI < trailsMap.size && stepJ >= 0 && stepJ < trailsMap[0].length && trailsMap[stepI][stepJ] == step.toString()[0]) {
                    nextSteps.add(Pair(stepI, stepJ))
                }
            }
        }

        steps = nextSteps
    }

    return steps.size
}

fun findTrailRating(trailsMap: List<String>, start: Pair<Int, Int>): Int {
    var rating = 0
    
    var currStep = start
    var stepPoint = trailsMap[start.first][start.second].toString().toInt() + 1

    while (stepPoint != 10) {
        val nextSteps = mutableSetOf<Pair<Int, Int>>()

        for ((i, j) in listOf(Pair(0, 1), Pair(0, -1), Pair(1, 0), Pair(-1, 0))) {
            val stepI = currStep.first + i
            val stepJ = currStep.second + j

            if (
                stepI >= 0 && stepI < trailsMap.size && stepJ >= 0 && stepJ < trailsMap[0].length &&
                trailsMap[stepI][stepJ] == stepPoint.toString()[0]
            ) {
                nextSteps.add(Pair(stepI, stepJ))
            }
        }

        if (nextSteps.isEmpty()) { break }

        nextSteps.forEachIndexed { i, step -> 
            if (i == 0) currStep = step else rating += findTrailRating(trailsMap, step)
        }

        stepPoint++
    }

    if (stepPoint == 10) { rating += 1 }

    return rating
}


fun parse(): List<String> {
    return Path("input.txt").readLines()
}


println("Part1 answer: ${part1()}")
println("Part2 answer: ${part2()}")