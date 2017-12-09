const fs = require('fs')

if (process.argv.length < 2) {
  process.exit(-1)
}

function * parseInt (str) {
  let current = ''

  for (const ch of str) {
    if (ch === '\n' || ch === '\t') {
      yield Number.parseInt(current)
      current = ''
    } else if (ch >= '0' && ch <= '9') {
      current += ch
    }
  }
}

function findMaxIndex (arr) {
  let max = 0
  let index = 0

  for (let i = 0; i < arr.length; i++) {
    if (arr[i] > max) {
      max = arr[i]
      index = i
    }
  }

  return index
}

function findCycle (arr) {
  const map = new Map()
  let cycleCount = 0

  while (true) {
    const key = arr.join()
    if (map.has(key)) {
      return { cycleCount, cycleSize: cycleCount - map.get(key) }
    }

    map.set(key, cycleCount)

    let index = findMaxIndex(arr)
    let memory = arr[index]
    arr[index] = 0

    while (memory > 0) {
      arr[++index % arr.length] += 1
      memory--
    }

    cycleCount++
  }
}

const fileContent = fs.readFileSync(process.argv[2], 'utf8')

const { cycleCount, cycleSize } = findCycle(Array.from(parseInt(fileContent)))
console.log(`Part 1: ${cycleCount}`)
console.log(`Part 2: ${cycleSize}`)
