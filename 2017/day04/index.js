const fs = require('fs')

class Day04 {
  static * parseRows (str) {
    let current = ''
    let row = []

    for (const ch of str) {
      if (ch === '\n' || ch === ' ') {
        row.push(current)
        current = ''

        if (ch === '\n') {
          yield row
          row = []
        }
      } else if (ch >= 'a' && ch <= 'z') {
        current += ch
      }
    }
  }

  static isPassphraseValid (words, checkAnagrams = false) {
    let set = new Set()
    for (let word of words) {
      if (checkAnagrams) {
        word = word.split('').sort().join('')
      }

      if (set.has(word)) {
        return false
      }

      set.add(word)
    }

    return true
  }

  static run (input) {
    const fileContent = fs.readFileSync(input, 'utf8')

    let validCount = 0
    let validAnagramCount = 0
    for (const words of this.parseRows(fileContent)) {
      if (this.isPassphraseValid(words)) {
        validCount++
      }

      if (this.isPassphraseValid(words, true)) {
        validAnagramCount++
      }
    }

    return [
      validCount,
      validAnagramCount
    ]
  }
}

module.exports = Day04
