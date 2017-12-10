const fs = require('fs')
const toMarkdown = require('to-markdown')

const htmlContent = fs.readFileSync(process.argv[2], 'utf8')
const markdownContent = toMarkdown(htmlContent, { gfm: true })
fs.writeFileSync(process.argv[3], markdownContent)
