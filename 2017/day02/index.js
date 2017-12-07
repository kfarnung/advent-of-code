const fs = require('fs')

if (process.argv.length < 2) {
  process.exit(-1)
}

const fileContent = fs.readFileSync(process.argv[2], 'utf8')

let min = Number.MAX_SAFE_INTEGER
let max = Number.MIN_SAFE_INTEGER
let current = ''
let sum = 0

for (const ch of fileContent) {
  if (ch === '\n' || ch === '\t') {
    let num = Number.parseInt(current)
    max = Math.max(max, num)
    min = Math.min(min, num)
    current = ''

    if (ch === '\n') {
      sum += max - min
      min = Number.MAX_SAFE_INTEGER
      max = Number.MIN_SAFE_INTEGER
    }
  } else if (ch >= '0' && ch <= '9') {
    current += ch
  }
}

console.log(sum)
