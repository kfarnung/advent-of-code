const fs = require('fs');
const path = require('path');
const { performance } = require('perf_hooks');

async function runPerfTests () {
  const content = fs.readFileSync(path.resolve(__dirname, './.vscode/launch.json'), 'utf8');
  const json = JSON.parse(content.replace(/^\s*\/\/ .+/mg, ''));

  let executionTimes = new Map();
  const testIterations = 5;

  function addExecutionTime (name, time) {
    let bucket = executionTimes.get(name);
    if (bucket === undefined) {
      bucket = [];
      executionTimes.set(name, bucket);
    }

    bucket.push(time);
  }

  for (const config of json.configurations) {
    if (config.args !== undefined) {
      const name = config.name;
      const testName = config.args[0];
      const input = config.args[1].replace('${workspaceFolder}', `${__dirname}`); // eslint-disable-line no-template-curly-in-string

      console.log(`Running ${name}...`);

      for (let i = 0; i < testIterations; i++) {
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
        addExecutionTime(name, executionTime);
      }
    }
  }

  console.info();
  console.info('CSV output follows in the error stream');
  console.info('======================================');

  let line = 'Test';
  for (let i = 0; i < testIterations; i++) {
    line += `,Iteration ${i}`;
  }

  console.error(line);

  for (const test of executionTimes) {
    line = test[0];
    for (const time of test[1]) {
      line += `,${time}`;
    }

    console.error(line);
  }
}

runPerfTests();
