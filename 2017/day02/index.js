const fs = require('fs')

class Day02 {
  static * parseRows (str) {
    let current = ''
    let row = []

    for (const ch of str) {
      if (ch === '\n' || ch === '\t') {
        row.push(Number.parseInt(current))
        current = ''

        if (ch === '\n') {
          yield row
          row = []
        }
      } else if (ch >= '0' && ch <= '9') {
        current += ch
      }
    }
  }

  static minMaxDiff (row) {
    let min = Number.MAX_SAFE_INTEGER
    let max = Number.MIN_SAFE_INTEGER
    for (const num of row) {
      min = Math.min(min, num)
      max = Math.max(max, num)
    }

    return max - min
  }

  static evenlyDivisibleQuotient (row) {
    for (const num1 of row) {
      for (const num2 of row) {
        if (num1 === num2) {
          continue
        }

        if (num1 % num2 === 0) {
          return num1 / num2
        } else if (num2 % num1 === 0) {
          return num2 / num1
        }
      }
    }
  }

  static run (input) {
    const fileContent = fs.readFileSync(input, 'utf8')

    let minMaxSum = 0
    let evenlyDivisibleSum = 0
    for (const row of this.parseRows(fileContent)) {
      minMaxSum += this.minMaxDiff(row)
      evenlyDivisibleSum += this.evenlyDivisibleQuotient(row)
    }

    return [
      minMaxSum,
      evenlyDivisibleSum
    ]
  }
}

module.exports = Day02
