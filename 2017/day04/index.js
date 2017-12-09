const fs = require('fs')

if (process.argv.length < 2) {
  process.exit(-1)
}

function * parseRows (str) {
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

function isPassphraseValid (words, checkAnagrams = false) {
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

const fileContent = fs.readFileSync(process.argv[2], 'utf8')

let validCount = 0
let validAnagramCount = 0
for (const words of parseRows(fileContent)) {
  if (isPassphraseValid(words)) {
    validCount++
  }

  if (isPassphraseValid(words, true)) {
    validAnagramCount++
  }
}

console.log(`Part 1: ${validCount}`)
console.log(`Part 2: ${validAnagramCount}`)
