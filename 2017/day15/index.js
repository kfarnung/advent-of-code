const fs = require('fs');
const wasmDay15 = require('../wasm/day15');

const readyPromise = new Promise((resolve) => {
  wasmDay15.onRuntimeInitialized = function () {
    resolve();
  };
});

class Day15 {
  static parseInput (str) {
    const search = /[0-9]+/g;
    return str.match(search).map((match) => Number.parseInt(match));
  }

  static countMatches (startingValueA, startingValueB, numIterations, requiredMultipleA = 1, requiredMultipleB = 1) {
    return wasmDay15._countMatches(startingValueA, startingValueB, numIterations, requiredMultipleA, requiredMultipleB);
  }

  static * runScenarios (input) {
    const fileContent = fs.readFileSync(input, 'utf8');
    const [ startingValueA, startingValueB ] = this.parseInput(fileContent);

    yield this.countMatches(startingValueA, startingValueB, 40000000);
    yield this.countMatches(startingValueA, startingValueB, 5000000, 4, 8);
  }

  static async run (input) {
    await readyPromise;
    return this.runScenarios(input);
  }
}

module.exports = Day15;
