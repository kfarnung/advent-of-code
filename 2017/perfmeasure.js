const fs = require('fs');
const path = require('path');
const { performance } = require('perf_hooks');

const argv = require('minimist')(process.argv.slice(2));
let filter = null;
if (argv.f) {
  filter = new RegExp(argv.f);
}

let iterations = 1;
if (argv.n) {
  const n = Number.parseInt(argv.n);
  if (!Number.isNaN(n)) {
    iterations = n;
  }
}

async function runPerfTests (filter = null, iterations = 1) {
  const content = fs.readFileSync(path.resolve(__dirname, './.vscode/launch.json'), 'utf8');
  const json = JSON.parse(content.replace(/^\s*\/\/ .+/mg, ''));

  let csvContent = 'Test';
  for (let i = 0; i < iterations; i++) {
    csvContent += `,Iteration ${i}`;
  }

  csvContent += '\n';

  for (const config of json.configurations) {
    if (config.args !== undefined) {
      const name = config.name;

      if (filter != null && !filter.test(name)) {
        continue;
      }

      const testName = config.args[0];
      const input = config.args[1].replace('${workspaceFolder}', `${__dirname}`); // eslint-disable-line no-template-curly-in-string

      console.log(`Running ${name}...`);
      csvContent += name;

      for (let i = 0; i < iterations; i++) {
        const day = require(`./${testName}`);

        const start = performance.now();
        let results = day.run(input);

        if (results instanceof Promise) {
          results = await results;
        }

        // Make sure we read all results so a generator doesn't just run part of
        // the test.
        if (Array.from(results).length <= 0) {
          throw new Error('invalid result');
        }

        const stop = performance.now();
        const executionTime = stop - start;

        console.log(`Iteration ${i}: ${executionTime.toFixed(3)}ms`);
        csvContent += `,${executionTime}`;
      }

      csvContent += '\n';
    }
  }

  console.info();
  console.info('CSV output follows in the error stream');
  console.info('======================================');

  console.error(csvContent);
}

runPerfTests(filter, iterations);
