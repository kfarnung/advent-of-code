import { performance } from 'node:perf_hooks';

if (process.argv.length < 3) {
  console.error('usage: node index.js <day> <input>');
  process.exit(-1);
}

async function run(dayModule, input) {
  const day = (await import(`./${dayModule}/index.js`)).default;

  const start = performance.now();

  let results = day.run(input);

  if (results instanceof Promise) {
    results = await results;
  }

  let count = 1;
  for (const result of results) {
    console.log(`Part ${count++}: ${result}`);
  }

  const stop = performance.now();
  console.log(`Execution time: ${(stop - start).toFixed(3)}ms`);
}

await run(process.argv[2], process.argv[3]);
