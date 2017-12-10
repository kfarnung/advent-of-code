const fs = require('fs')

function parseAndScore (str) {
  let score = 0
  let garbageCount = 0

  let depth = 0
  let escapeNext = false
  let processingGarbage = false

  for (let ch of str) {
    if (processingGarbage) {
      if (escapeNext) {
        escapeNext = false
      } else if (ch === '!') {
        escapeNext = true
      } else if (ch === '>') {
        processingGarbage = false
      } else {
        garbageCount++
      }
    } else if (ch === '<') {
      processingGarbage = true
    } else if (ch === '{') {
      depth++
    } else if (ch === '}') {
      score += depth
      depth--
    }
  }

  return { score, garbageCount }
}

function run (input) {
  const fileContent = fs.readFileSync(process.argv[2], 'utf8')

  const { score, garbageCount } = parseAndScore(fileContent)
  console.log(`Part 1: ${score}`)
  console.log(`Part 2: ${garbageCount}`)
}

if (process.argv.length < 2) {
  process.exit(-1)
}

run(process.argv[2])
