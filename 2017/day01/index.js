const fs = require('fs')

if (process.argv.length < 2) {
  process.exit(-1)
}

function sumString (str) {
  let sum = 0
  const length = str.length

  for (let i = 0; i < length - 1; i++) {
    if (str[i] === str[i + 1]) {
      sum += str[i] - '0'
    }
  }

  if (str[length - 1] === str[0]) {
    sum += str[length - 1] - '0'
  }

  return sum
}

const fileContent = fs.readFileSync(process.argv[2], 'utf8')
console.log(sumString(fileContent.trim()))
