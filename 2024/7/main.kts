import kotlin.io.path.Path
import kotlin.io.path.readLines
import kotlin.math.pow

fun solution(radix: Int): ULong {
    var ans = 0.toULong()

    line@ for (line in Path("input.txt").readLines()) {
        val parts = line.split(": ")
        val target = parts[0].toULong()
        val members = parts[1].split(" ").map { it.toULong() }

        permutation@ for (permutation in 0 until radix.toDouble().pow(members.size-1).toInt()) {
            var res = members[0]

            var i = 1
            for (op in permutation.toString(radix).padStart(members.size-1, '0')) {
                when (op) {
                    '0' -> res += members[i]
                    '1' -> res *= members[i]
                    '2' -> res = String.format("%s%s", res, members[i]).toULong()
                }

                if (res > target) continue@permutation

                i++
            }

            if (res == target) {ans += res; continue@line}
        }
    }

    return  ans
}

println("Part1 answer: ${solution(2)}")
println("Part2 answer: ${solution(3)}")