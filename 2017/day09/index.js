const fs = require('fs')

class Day09 {
  static parseAndScore (str) {
    let score = 0
    let garbageCount = 0

    let depth = 0
    let escapeNext = false
    let processingGarbage = false

    for (let ch of str) {
      if (processingGarbage) {
        if (escapeNext) {
          escapeNext = false
        } else if (ch === '!') {
          escapeNext = true
        } else if (ch === '>') {
          processingGarbage = false
        } else {
          garbageCount++
        }
      } else if (ch === '<') {
        processingGarbage = true
      } else if (ch === '{') {
        depth++
      } else if (ch === '}') {
        score += depth
        depth--
      }
    }

    return { score, garbageCount }
  }

  static run (input) {
    const fileContent = fs.readFileSync(input, 'utf8')

    const { score, garbageCount } = this.parseAndScore(fileContent)
    return [
      score,
      garbageCount
    ]
  }
}

module.exports = Day09
