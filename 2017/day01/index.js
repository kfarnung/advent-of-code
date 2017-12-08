const fs = require('fs')

if (process.argv.length < 2) {
  process.exit(-1)
}

function sumString (str, offset) {
  let sum = 0
  const length = str.length

  for (let i = 0; i < length; i++) {
    if (str[i] === str[(i + offset) % length]) {
      sum += str[i] - '0'
    }
  }

  return sum
}

const fileContent = fs.readFileSync(process.argv[2], 'utf8').trim()
console.log(`Part 1: ${sumString(fileContent, 1)}`)
console.log(`Part 2: ${sumString(fileContent, Math.floor(fileContent.length / 2))}`)
