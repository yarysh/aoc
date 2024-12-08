import kotlin.io.path.Path
import kotlin.io.path.readLines


fun part1(): Int {
    val (lines, map) = parse()

    val antinodes = mutableSetOf<Pair<Int, Int>>()

    for (points in map.values) {
        for (i in points.indices) {
            var j = i + 1
            while (j < points.size) {
                val diff = Pair(points[i].first - points[j].first, points[i].second - points[j].second)

                listOf(
                    Pair(points[i].first + diff.first, points[i].second + diff.second),
                    Pair(points[j].first - diff.first, points[j].second - diff.second)
                ).filter {
                    inBounds(it, lines.size-1, lines[0].length-1)
                }.forEach {
                    antinodes.add(Pair(it.first, it.second))
                }

                j++
            }
        }
    }

    return antinodes.size
}

fun part2(): Int {
    val (lines, map) = parse()

    val allAntinodes = mutableSetOf<Pair<Int, Int>>()

    for (points in map.values) {
        for (i in points.indices) {
            var j = i + 1
            while (j < points.size) {
                val diff = Pair(points[i].first - points[j].first, points[i].second - points[j].second)
                
                var antenna: Pair<Int, Int>
                val antinodes = mutableListOf<Pair<Int, Int>>()

                // Emitting to the top
                antenna = Pair(points[i].first, points[i].second)
                while (inBounds(antenna, lines.size-1, lines[0].length-1)) {
                    antinodes.add(antenna)
                    antenna = Pair(antenna.first + diff.first, antenna.second + diff.second)
                }

                // Emitting to the bottom
                antenna = Pair(points[j].first, points[j].second)
                while (inBounds(antenna, lines.size-1, lines[0].length-1)) {
                    antinodes.add(antenna)
                    antenna = Pair(antenna.first - diff.first, antenna.second - diff.second)
                }

                antinodes.forEach {
                    allAntinodes.add(Pair(it.first, it.second))
                }

                j++
            }
        }
    }
    
    return allAntinodes.size
}

fun parse(): Pair<List<String>, Map<Char, List<Pair<Int, Int>>> > {
    val lines = Path("input.txt").readLines().toMutableList()

    val map = mutableMapOf<Char, MutableList<Pair<Int, Int>>>()
    for ((i, line) in lines.withIndex()) {
        for ((j, point) in line.withIndex()) {
            if (point == '.') continue
            map.getOrPut(point) { mutableListOf() }.add(Pair(i, j))
        }
    }

    return Pair(lines, map)
} 

fun inBounds(point: Pair<Int, Int>, maxI: Int, maxJ: Int): Boolean {
    return point.first in 0..maxI && point.second in 0 .. maxJ
}


println("Part1 answer: ${part1()}")
println("Part2 answer: ${part2()}")
