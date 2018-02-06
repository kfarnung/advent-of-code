const fs = require('fs');
const wasmDay15 = require('../wasm/day15');

const readyPromise = new Promise((resolve) => {
  wasmDay15.onRuntimeInitialized = function () {
    resolve();
  };
});

class Generator {
  constructor (startingValue, factor, requiredMultiple = 1) {
    this._previousValue = startingValue;
    this._factor = factor;
    this._requiredMultiple = requiredMultiple;
  }

  next () {
    this._previousValue = wasmDay15._calculateNext(this._previousValue, this._factor, this._requiredMultiple);
    return this._previousValue;
  }
}

class Day15 {
  static parseInput (str) {
    const search = /[0-9]+/g;
    return str.match(search).map((match) => Number.parseInt(match));
  }

  static areBitsEqual (val1, val2, numBits) {
    const mask = (1 << numBits) - 1;
    return (val1 & mask) === (val2 & mask);
  }

  static countMatches (startingValueA, startingValueB, numIterations, requiredMultipleA = 1, requiredMultipleB = 1) {
    const generatorA = new Generator(startingValueA, 16807, requiredMultipleA);
    const generatorB = new Generator(startingValueB, 48271, requiredMultipleB);

    let count = 0;

    for (let i = 0; i < numIterations; i++) {
      if (this.areBitsEqual(generatorA.next(), generatorB.next(), 16)) {
        count++;
      }
    }

    return count;
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
