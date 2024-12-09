import kotlin.io.path.Path
import kotlin.io.path.readText


fun part1(): ULong {
    val diskMap = mutableListOf<Int>()
    with(diskMap, {
        Path("input.txt").readText().forEach {
            add(it.toString().toInt())
        }
    })
    
    var ans = (0).toULong()
    
    var forwardCursor = 0
    var backwardCursor = diskMap.size - 1
    var multiplier = 0
    
    while (forwardCursor <= backwardCursor) {
        when {
            forwardCursor % 2 == 0 -> {
                repeat(diskMap[forwardCursor]) {
                    ans += (multiplier * forwardCursor / 2).toULong()

                    multiplier++
                }
            }
            else -> {
                while (diskMap[forwardCursor] > 0 && forwardCursor < backwardCursor) {
                    if (diskMap[backwardCursor] == 0) {
                        backwardCursor -= 2
                        continue
                    }

                    ans += (multiplier * backwardCursor / 2).toULong()
                    diskMap[forwardCursor]--
                    diskMap[backwardCursor]--

                    multiplier++
                }
            }
        }

        forwardCursor++
    }
    
    return ans
}

fun part2(): ULong {
    data class Block(val id: Int, val size: Int)

    val diskMap = mutableListOf<Block>()
    with(diskMap, {
        Path("input.txt").readText().forEachIndexed { i, it ->
            val size = it.toString().toInt()
            if (size != 0) {
                add(Block(
                    id = if(i % 2 == 0) i/2 else -1,
                    size = size
                ))
            }
        }
    })
    
    var backwardCursor = diskMap.size - 1
    while (backwardCursor >= 0) {
        val currBlock = diskMap[backwardCursor]
        if (currBlock.id == -1) {
            backwardCursor--
            continue
        }
        
        for ((i, block) in diskMap.withIndex()) {
            if (i >= backwardCursor) break

            if (block.id == -1 && block.size >= currBlock.size) {
                if (block.size > currBlock.size) {
                    diskMap[i] = Block(id = -1, size = block.size - currBlock.size)
                    diskMap.add(i, Block(id = -1, size = currBlock.size))

                    backwardCursor++
                }

                diskMap[backwardCursor] = diskMap[i]
                diskMap[i] = currBlock

                break
            }
        }

        backwardCursor--
    }
    
    var ans = (0).toULong()
    var multiplier = 0
    for (block in diskMap) {
        if (block.id == -1) {
            multiplier += block.size
            continue
        }

        repeat(block.size) {
            ans += (multiplier * block.id).toULong()
            multiplier++
        }
    }
    
    return ans
}


println("Part1 answer: ${part1()}")
println("Part2 answer: ${part2()}")
