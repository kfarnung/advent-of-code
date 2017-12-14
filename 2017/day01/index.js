const fs = require('fs')

class Day01 {
  static sumString (str, offset) {
    let sum = 0
    const length = str.length

    for (let i = 0; i < length; i++) {
      if (str[i] === str[(i + offset) % length]) {
        sum += str[i] - '0'
      }
    }

    return sum
  }

  static run (input) {
    const fileContent = fs.readFileSync(input, 'utf8').trim()

    return [
      this.sumString(fileContent, 1),
      this.sumString(fileContent, Math.floor(fileContent.length / 2))
    ]
  }
}

module.exports = Day01
