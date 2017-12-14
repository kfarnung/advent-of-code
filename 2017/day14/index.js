function * parseDataAsAscii (str) {
  for (const ch of str) {
    if (ch !== '\n') {
      yield ch.charCodeAt(0)
    }
  }

  for (const num of [17, 31, 73, 47, 23]) {
    yield num
  }
}

function generateArray (start, length) {
  const arr = []
  for (let i = start; i < length; i++) {
    arr.push(i)
  }

  return arr
}

function reverse (arr, start, length) {
  const range = Math.floor(length / 2)
  const arrLength = arr.length

  for (let i = 0; i < range; i++) {
    const first = (start + i) % arrLength
    const second = (start + length - 1 - i) % arrLength

    const temp = arr[first]
    arr[first] = arr[second]
    arr[second] = temp
  }
}

function calculateHash (lengths, numRounds = 64) {
  const arr = generateArray(0, 256)
  const arrLength = arr.length
  let currentPosition = 0
  let skipSize = 0

  for (let i = 0; i < numRounds; i++) {
    for (const num of lengths) {
      reverse(arr, currentPosition, num)
      currentPosition = (currentPosition + num + skipSize) % arrLength
      skipSize++
    }
  }

  let hex = []
  let xor = 0
  for (let i = 0; i < arr.length; i++) {
    if (i > 0 && i % 16 === 0) {
      hex.push(xor)
      xor = 0
    }

    xor ^= arr[i]
  }

  hex.push(xor)
  return hex
}

class UsageMap {
  constructor (key) {
    this._rows = []

    for (let i = 0; i < 128; i++) {
      this._rows.push(calculateHash(Array.from(parseDataAsAscii(key + '-' + i))))
    }
  }

  getBitCount () {
    let count = 0
    for (let i = 0; i < 128; i++) {
      for (let j = 0; j < 128; j++) {
        count += this._getBitAt(i, j)
      }
    }

    return count
  }

  getGroupCount () {
    let groupCount = 0
    const visited = new Set()
    for (let i = 0; i < 128; i++) {
      for (let j = 0; j < 128; j++) {
        if (!visited.has([i, j].join(','))) {
          if (this._findGroup(visited, i, j)) {
            groupCount++
          }
        }
      }
    }

    return groupCount
  }

  _getBitAt (row, column) {
    const index = Math.floor(column / 8)
    const bit = 7 - column % 8

    let val = this._rows[row][index]

    return ((val >> bit) & 1)
  }

  _findGroup (visited, row, column) {
    const queue = []

    if (this._getBitAt(row, column) === 0) {
      return false
    }

    queue.push([row, column])

    while (queue.length > 0) {
      let [ currentRow, currentColumn ] = queue.shift()

      if (!visited.has([currentRow, currentColumn].join(',')) &&
          this._getBitAt(currentRow, currentColumn) === 1) {
        visited.add([currentRow, currentColumn].join(','))

        if (currentRow > 0) {
          queue.push([currentRow - 1, currentColumn])
        }

        if (currentColumn > 0) {
          queue.push([currentRow, currentColumn - 1])
        }

        if (currentRow < 127) {
          queue.push([currentRow + 1, currentColumn])
        }

        if (currentColumn < 127) {
          queue.push([currentRow, currentColumn + 1])
        }
      }
    }

    return true
  }
}

function run (input) {
  const usageMap = new UsageMap(input)

  console.log(`Part 1: ${usageMap.getBitCount()}`)
  console.log(`Part 2: ${usageMap.getGroupCount()}`)
}

if (process.argv.length < 2) {
  process.exit(-1)
}

run(process.argv[2])
