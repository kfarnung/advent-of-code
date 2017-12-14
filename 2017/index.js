if (process.argv.length < 3) {
  console.error('usage: node index.js <day> <input>')
  process.exit(-1)
}

const day = require(`./${process.argv[2]}`)

const results = day.run(process.argv[3])
results.forEach((value, index) => {
  console.log(`Part ${index + 1}: ${value}`)
})
