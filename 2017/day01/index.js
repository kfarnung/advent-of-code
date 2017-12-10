const fs = require('fs')

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

function run (input) {
  const fileContent = fs.readFileSync(input, 'utf8').trim()
  console.log(`Part 1: ${sumString(fileContent, 1)}`)
  console.log(`Part 2: ${sumString(fileContent, Math.floor(fileContent.length / 2))}`)
}

if (process.argv.length < 2) {
  process.exit(-1)
}

run(process.argv[2])
