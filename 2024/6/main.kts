import kotlin.io.path.Path
import kotlin.io.path.readLines


val directions = listOf("^", ">", "v", "<")


fun part1(): Int {
    val (lines, position, direction) = parse()
    
    val visited = mutableSetOf<Pair<Int, Int>>()
    findPath(lines, position, direction).first.also { 
        visited.addAll(it.map { it.first })
    }
    
    return visited.size
}

fun part2(): Int {
    val (initLines, initPos, initDir) = parse()
    val initPath = findPath(initLines, initPos, initDir).first
    
    val visited = mutableSetOf<Set<Pair<Pair<Int, Int>, Char>>>()

    for (item in initPath) {
        val (point, direction) = item
        
        if (point == initPos && direction == initDir) continue

        val lines = initLines.toMutableList()
        lines[point.first] = lines[point.first].replaceRange(point.second, point.second+1, "#")

        val (path, isCircle) = findPath(lines, initPos, initDir)
        if (isCircle) {
            visited.add(path)
        }
    }

    return visited.size
}

fun parse(): Triple<List<String>, Pair<Int, Int>, Char> {
    val lines = Path("input.txt").readLines()

    var direction = ' '
    var position = Pair(0, 0)
    for ((i, line) in lines.withIndex()) {
        val p = line.findAnyOf(directions)
        if (p != null) {
            direction = p.second[0]
            position = Pair(i, p.first)
            break
        }
    }
    
    return Triple(lines, position, direction)
}

fun findPath(lines: List<String>, initPos: Pair<Int, Int>, initDir: Char): Pair<Set<Pair<Pair<Int, Int>, Char>>, Boolean> {
    var isCircle = false
    val path = mutableSetOf(Pair(initPos, initDir))

    var position = initPos
    var direction = initDir

    while (
        position.first != 0 &&
        position.second != 0 &&
        position.first != lines[0].length - 1 &&
        position.second != lines.size - 1
    ) {
        val nextPosition = when (direction) {
            '^' -> Pair(position.first-1, position.second)
            '>' -> Pair(position.first, position.second+1)
            'v' -> Pair(position.first+1, position.second)
            '<' -> Pair(position.first, position.second-1)
            else -> throw Exception("Invalid direction")
        }

        if (lines[nextPosition.first][nextPosition.second] == '#') {
            directions.indexOf(direction.toString()).let {
                direction = directions[(it + 1) % 4][0]
            }
            continue
        }

        position = nextPosition
        
        if(!path.add(Pair(position, direction))) {
            isCircle = true
            break
        }
    }

    return Pair(path, isCircle)
}


println("Part1 answer: ${part1()}")
println("Part2 answer: ${part2()}")
