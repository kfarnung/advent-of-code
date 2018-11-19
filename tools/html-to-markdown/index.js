const cheerio = require('cheerio');
const fs = require('fs');
const gfm = require('turndown-plugin-gfm').gfm;
const os = require('os');
const TurndownService = require('turndown');
const yargs = require('yargs');

const year = yargs.argv.year;
const day = yargs.argv.day;

const htmlContent = fs.readFileSync(yargs.argv._[0], 'utf8');
const $ = cheerio.load(htmlContent);

const turndownService = new TurndownService({
  headingStyle: 'atx'
});
turndownService.use(gfm);

function stripHyphens (str) {
  const regex = /^--- (.+) ---$/;
  return str.replace(regex, (match, p1) => {
    return p1;
  });
}

const newDoc = $('<div>');
newDoc.append('<h1>');
newDoc.append($('<p>').text(`http://adventofcode.com/${year}/day/${day}`));
newDoc.append($('<h2>').text('Description'));

$('article').each((index, element) => {
  const headingElement = $('h2', element);
  let heading = stripHyphens(headingElement.text());

  if (index === 0) {
    $('h1', newDoc).text(heading);
    heading = 'Part One';
  }

  headingElement.replaceWith($('<h3>').text(heading));

  newDoc.append(element);
});

const markdownContent = turndownService.turndown(newDoc.html()).concat(os.EOL);

fs.writeFileSync(yargs.argv._[1], markdownContent);
