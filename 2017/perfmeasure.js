import fs from 'node:fs';
import path from 'node:path';
import { performance } from 'node:perf_hooks';

import minimist from 'minimist';

const argv = minimist(process.argv.slice(2));
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

const csvFile = argv.csv;

async function runPerfTests(filter = null, iterations = 1, csvFile = null) {
  const content = fs.readFileSync(
    path.resolve(import.meta.dirname, './.vscode/launch.json'),
    'utf8',
  );
  const json = JSON.parse(content.replace(/^\s*\/\/ .+/gm, ''));

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
      const input = config.args[1].replace(
        '${workspaceFolder}',
        `${import.meta.dirname}`,
      );
      const testModule = (await import(`./${testName}/index.js`)).default;

      console.log(`Running ${name}...`);
      csvContent += name;

      for (let i = 0; i < iterations; i++) {
        const start = performance.now();
        let results = testModule.run(input);

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

  if (csvFile) {
    fs.writeFileSync(path.normalize(csvFile), csvContent, 'utf8');
  }

  console.log('Done!');
}

await runPerfTests(filter, iterations, csvFile);
