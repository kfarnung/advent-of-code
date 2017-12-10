const fs = require('fs')

function * parseInt (str) {
  let current = ''

  for (const ch of str) {
    if (ch === '\n') {
      yield Number.parseInt(current)
      current = ''
    } else if ((ch >= '0' && ch <= '9') || ch === '-') {
      current += ch
    }
  }
}

function incrementModifier (arr, index) {
  arr[index]++
}

function decrementThreeModifier (arr, index) {
  if (arr[index] >= 3) {
    arr[index]--
  } else {
    incrementModifier(arr, index)
  }
}

function followSteps (arr, modifierFunc) {
  let index = 0
  let numSteps = 0
  while (index >= 0 && index < arr.length) {
    const nextOffset = arr[index]
    modifierFunc(arr, index)
    index += nextOffset
    numSteps++
  }

  return numSteps
}

function run (input) {
  const fileContent = fs.readFileSync(input, 'utf8')
  const arr1 = Array.from(parseInt(fileContent))
  const arr2 = Array.from(arr1)

  console.log(`Part 1: ${followSteps(arr1, incrementModifier)}`)
  console.log(`Part 2: ${followSteps(arr2, decrementThreeModifier)}`)
}

if (process.argv.length < 2) {
  process.exit(-1)
}

run(process.argv[2])
