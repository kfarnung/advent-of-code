const fs = require('fs');

class Day15 {
  static parseInput(str) {
    const search = /[0-9]+/g;
    return str.match(search).map((match) => Number.parseInt(match));
  }

  static *generateValues(startingValue, factor, requiredMultiple = 1) {
    let previousValue = startingValue;
    for (;;) {
      previousValue = (previousValue * factor) % 2147483647;

      if (previousValue % requiredMultiple === 0) {
        yield previousValue;
      }
    }
  }

  static areBitsEqual(val1, val2, numBits) {
    const mask = Math.pow(2, numBits) - 1;
    return (val1 & mask) === (val2 & mask);
  }

  static countMatches(
    startingValueA,
    startingValueB,
    numIterations,
    requiredMultipleA = 1,
    requiredMultipleB = 1
  ) {
    const generatorA = this.generateValues(
      startingValueA,
      16807,
      requiredMultipleA
    );
    const generatorB = this.generateValues(
      startingValueB,
      48271,
      requiredMultipleB
    );

    let count = 0;

    for (let i = 0; i < numIterations; i++) {
      const valueA = generatorA.next().value;
      const valueB = generatorB.next().value;

      if (this.areBitsEqual(valueA, valueB, 16)) {
        count++;
      }
    }

    return count;
  }

  static run(input) {
    const fileContent = fs.readFileSync(input, 'utf8');
    const [startingValueA, startingValueB] = this.parseInput(fileContent);

    return [
      this.countMatches(startingValueA, startingValueB, 40000000),
      this.countMatches(startingValueA, startingValueB, 5000000, 4, 8),
    ];
  }
}

module.exports = Day15;
