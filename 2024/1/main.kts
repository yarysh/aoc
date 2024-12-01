import kotlin.io.path.Path
import kotlin.io.path.readLines
import kotlin.math.abs


fun part1(): Int {
    val lst1 = mutableListOf<Int>()
    val lst2 = mutableListOf<Int>()

    Path("input.txt").readLines().forEach {
        it.split("   ").map { s -> s.toInt() }.let { (a, b) ->
            lst1.add(a)
            lst2.add(b)
        }
    }.also {
        lst1.sort()
        lst2.sort()
    }

    return lst1.zip(lst2).sumOf { abs(it.first - it.second) }
}

fun part2(): Int {
    val lst1 = mutableListOf<Int>()
    val map2 = mutableMapOf<Int, Int>()

    Path("input.txt").readLines().forEach {
        it.split("   ").map { s -> s.toInt() }.let { (a, b) ->
            lst1.add(a)
            map2[b] = map2.getOrDefault(b, 0) + 1
        }
    }.also {
        lst1.sort()
    }

    return lst1.sumOf{it * map2.getOrDefault(it, 0)}
}


println("Part1 answer: ${part1()}")
println("Part2 answer: ${part2()}")
