const fs = require('fs')

function * parseData (str) {
  let current = ''

  for (const ch of str) {
    if (ch === '\n' || ch === ',') {
      yield current
      current = ''
    } else if (ch >= 'a' && ch <= 'z') {
      current += ch
    }
  }
}

function getDistance (x, y) {
  const absX = Math.abs(x)
  const absY = Math.abs(y)

  if (absX >= absY) {
    return absX
  } else {
    return absX + ((absY - absX) / 2)
  }
}

function run (input) {
  const fileContent = fs.readFileSync(input, 'utf8')
  const directions = parseData(fileContent)

  let x = 0
  let y = 0
  let max = Number.MIN_SAFE_INTEGER

  for (const dir of directions) {
    switch (dir) {
      case 'n':
        y += 2
        break

      case 's':
        y -= 2
        break

      case 'ne':
        x += 1
        y += 1
        break

      case 'se':
        x += 1
        y -= 1
        break

      case 'nw':
        x -= 1
        y += 1
        break

      case 'sw':
        x -= 1
        y -= 1
        break

      default:
        throw new Error('Unexpected direction input!')
    }

    max = Math.max(max, getDistance(x, y))
  }

  console.log(`Part 1: ${getDistance(x, y)}`)
  console.log(`Part 2: ${max}`)
}

if (process.argv.length < 2) {
  process.exit(-1)
}

run(process.argv[2])
