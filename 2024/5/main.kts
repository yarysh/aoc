import kotlin.io.path.Path
import kotlin.io.path.readLines

fun part1(): Int {
    val ordering = mutableMapOf<Int, MutableSet<Int>>()
    val pages = mutableListOf<List<Int>>()
    
    val lines = Path("input.txt").readLines()
    var isRules = true
    for (line in lines) {
        if (line.isEmpty()) {
            isRules = false
            continue
        }

        when (isRules) {
            true -> line.split("|").also {
                val a = it[0].toInt()
                val b = it[1].toInt()
                ordering[a] = ordering.getOrElse(a) { mutableSetOf() }.also { it.add(b) }
            }

            false -> line.split(",").map { it.toInt() }.also { pages.add(it) }
        }
    }

    return pages.sumOf {
        if (it == it.sortedWith { a, b ->
            when {
                ordering[a]?.contains(b) == true -> -1
                ordering[b]?.contains(a) == true -> 1
                else -> 0
            }
        }) it[it.size / 2] else 0
    }
}

fun part2(): Int {
    val ordering = mutableMapOf<Int, MutableSet<Int>>()
    val pages = mutableListOf<List<Int>>()

    val lines = Path("input.txt").readLines()
    var isRules = true
    for (line in lines) {
        if (line.isEmpty()) {
            isRules = false
            continue
        }

        when (isRules) {
            true -> line.split("|").also {
                val a = it[0].toInt()
                val b = it[1].toInt()
                ordering[a] = ordering.getOrElse(a) { mutableSetOf() }.also { it.add(b) }
            }

            false -> line.split(",").map { it.toInt() }.also { pages.add(it) }
        }
    }

    return pages.sumOf {
        val sorted = it.sortedWith { a, b ->
            when {
                ordering[a]?.contains(b) == true -> -1
                ordering[b]?.contains(a) == true -> 1
                else -> 0
            }
        }

        if (it != sorted) sorted[sorted.size / 2] else 0
    }
}

println("Part1 answer: ${part1()}")
println("Part2 answer: ${part2()}")
