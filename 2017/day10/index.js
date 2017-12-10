const fs = require('fs')

function * parseData (str) {
  let current = ''

  for (const ch of str) {
    if (ch === '\n' || ch === ',') {
      yield Number.parseInt(current)
      current = ''
    } else if (ch >= '0' && ch <= '9') {
      current += ch
    }
  }
}

function * parseDataAsAscii (str) {
  for (const ch of str) {
    if (ch !== '\n') {
      yield ch.charCodeAt(0)
    }
  }

  for (const num of [17, 31, 73, 47, 23]) {
    yield num
  }
}

function generateArray (start, length) {
  const arr = []
  for (let i = start; i < length; i++) {
    arr.push(i)
  }

  return arr
}

function reverse (arr, start, length) {
  const range = Math.floor(length / 2)
  const arrLength = arr.length

  for (let i = 0; i < range; i++) {
    const first = (start + i) % arrLength
    const second = (start + length - 1 - i) % arrLength

    const temp = arr[first]
    arr[first] = arr[second]
    arr[second] = temp
  }
}

function runHash (lengths, numRounds) {
  const arr = generateArray(0, 256)
  const arrLength = arr.length
  let currentPosition = 0
  let skipSize = 0

  for (let i = 0; i < numRounds; i++) {
    for (const num of lengths) {
      reverse(arr, currentPosition, num)
      currentPosition = (currentPosition + num + skipSize) % arrLength
      skipSize++
    }
  }

  return arr
}

function padLeft (str, width = 2, ch = '0') {
  let retVal = ''
  for (let i = 0; i < width - str.length; i++) {
    retVal += ch
  }

  retVal += str
  return retVal
}

function getHashString (hash) {
  let hex = []
  let xor = 0
  for (let i = 0; i < hash.length; i++) {
    if (i > 0 && i % 16 === 0) {
      hex.push(padLeft(xor.toString(16)))
      xor = 0
    }

    xor ^= hash[i]
  }

  hex.push(padLeft(xor.toString(16)))
  return hex.join('')
}

function run (input) {
  const fileContent = fs.readFileSync(input, 'utf8')
  const hash1 = runHash(Array.from(parseData(fileContent)), 1)
  const hash2 = runHash(Array.from(parseDataAsAscii(fileContent)), 64)
  const hashString = getHashString(hash2)

  console.log(`Part 1: ${hash1[0] * hash1[1]}`)
  console.log(`Part 2: ${hashString}`)
}

if (process.argv.length < 2) {
  process.exit(-1)
}

run(process.argv[2])
