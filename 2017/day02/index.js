const fs = require('fs')

function * parseRows (str) {
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

function minMaxDiff (row) {
  let min = Number.MAX_SAFE_INTEGER
  let max = Number.MIN_SAFE_INTEGER
  for (const num of row) {
    min = Math.min(min, num)
    max = Math.max(max, num)
  }

  return max - min
}

function evenlyDivisibleQuotient (row) {
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

function run (input) {
  const fileContent = fs.readFileSync(input, 'utf8')

  let minMaxSum = 0
  let evenlyDivisibleSum = 0
  for (const row of parseRows(fileContent)) {
    minMaxSum += minMaxDiff(row)
    evenlyDivisibleSum += evenlyDivisibleQuotient(row)
  }

  console.log(`Part 1: ${minMaxSum}`)
  console.log(`Part 2: ${evenlyDivisibleSum}`)
}

if (process.argv.length < 2) {
  process.exit(-1)
}

run(process.argv[2])
