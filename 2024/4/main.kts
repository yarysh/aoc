import kotlin.io.path.Path
import kotlin.io.path.readLines


fun part1(): Int {
    var ans = 0
    
    val arr =  mutableListOf<String>()
    Path("input.txt").readLines().forEach {
        arr.add(it)
    }
    
    val dirs = listOf(
        listOf(Pair(0, 0), Pair(0, -1), Pair(0, -2), Pair(0, -3)), // up
        listOf(Pair(0, 0), Pair(1, 0), Pair(2, 0), Pair(3, 0)), // right
        listOf(Pair(0, 0), Pair(0, 1), Pair(0, 2), Pair(0, 3)), // down
        listOf(Pair(0, 0), Pair(-1, 0), Pair(-2, 0), Pair(-3, 0)), // left
        listOf(Pair(0, 0), Pair(-1, -1), Pair(-2, -2), Pair(-3, -3)), // up-diagonal-right
        listOf(Pair(0, 0), Pair(1, 1), Pair(2, 2), Pair(3, 3)), // down-diagonal-right
        listOf(Pair(0, 0), Pair(-1, -1), Pair(-2, -2), Pair(-3, -3)), // up-diagonal-left
        listOf(Pair(0, 0), Pair(1, 1), Pair(2, 2), Pair(3, 3)) // down-diagonal-left
    )
    
    for (i in 0 until arr.size) {
        for (j in 0 until arr[i].length) {
            if (arr[i][j] != 'X') continue

            wordsCoords(Pair(i, j), Pair(arr.size, arr[i].length), dirs).forEach { coords ->
                buildString {
                    repeat(coords.size) {append(arr[coords[it].first][coords[it].second])}
                }.takeIf { it == "XMAS" }?.let { ans++}
            }
        }
    }
    
    return ans
}


fun part2(): Int {
    var ans = 0

    val arr =  mutableListOf<String>()
    Path("input.txt").readLines().forEach {
        arr.add(it)
    }

    val dirs = listOf(
        listOf(Pair(0, 0), Pair(-1, -1), Pair(1, -1), Pair(1, 1), Pair(-1, 1)), // X
    )

    for (i in 0 until arr.size) {
        for (j in 0 until arr[i].length) {
            if (arr[i][j] != 'A') continue

            wordsCoords(Pair(i, j), Pair(arr.size, arr[i].length), dirs).forEach { coords ->
                buildString {
                    repeat(coords.size) {append(arr[coords[it].first][coords[it].second])}
                }.takeIf { it in listOf("AMSSM", "ASMMS", "ASSMM", "AMMSS") }?.let { ans++}
            }
        }
    }

    return ans
}

fun wordsCoords(pos: Pair<Int, Int>, max: Pair<Int, Int>, dirs: List<List<Pair<Int, Int>>>): List<List<Pair<Int, Int>>> {
    return buildList {
        dirs.forEach {
            buildList { 
                it.forEach {
                    val newPos = Pair(pos.first + it.first, pos.second + it.second)
                    if (newPos.first in 0 until max.first && newPos.second in 0 until max.second) {
                        add(newPos)
                    }
                }
            }.apply {
                if (size == dirs.first().size) add(this)
            }
        }
    }
}

println("Part1 answer: ${part1()}")
println("Part2 answer: ${part2()}")
