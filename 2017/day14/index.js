const Day10 = require('../day10')

class UsageMap {
  constructor (key) {
    this._rows = []

    for (let i = 0; i < 128; i++) {
      this._rows.push(Day10.calculateHash(key + '-' + i))
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

class Day14 {
  static run (input) {
    const usageMap = new UsageMap(input)

    return [
      usageMap.getBitCount(),
      usageMap.getGroupCount()
    ]
  }
}

module.exports = Day14
