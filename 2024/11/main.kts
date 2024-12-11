import kotlin.io.path.Path
import kotlin.io.path.readText



fun solution(n: Int): Long {
    val stones = Path("input.txt").readText().split(" ").map { it.toLong() }
    return countStones(stones, n)
}

fun countStones(stones: List<Long>, blinks: Int): Long {
    var cache = buildMap { stones.forEach { put(it, 1L) } }
    
    repeat(blinks) {
        cache = buildMap {
            cache.forEach { (stone, cnt) ->
                when {
                    stone == 0L -> 1L.let { put(it, getOrDefault(it, 0L) + cnt) }
                    stone.toString().length % 2 == 0 -> stone.toString().also {
                        it.substring(0, it.length / 2).toLong().let { put(it, getOrDefault(it, 0L) + cnt) }
                        it.substring(it.length / 2).toLong().let { put(it, getOrDefault(it, 0) + cnt) }
                    }
                    else -> (stone*2024L).let { put(it, getOrDefault(it, 0L) + cnt) }
                }
            }
        }
    }

    return cache.values.sum()
}

println("Part1 answer: ${solution(25)}")
println("Part2 answer: ${solution(75)}")
