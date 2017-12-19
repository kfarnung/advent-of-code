const fs = require('fs');

class Day01 {
  static sumString (str, offset = 1) {
    let sum = 0;
    const length = str.length;

    for (let i = 0; i < length; i++) {
      if (str[i] === str[(i + offset) % length]) {
        sum += str[i] - '0';
      }
    }

    return sum;
  }

  static sumStringHalfway (str) {
    return this.sumString(str, Math.floor(str.length / 2));
  }

  static run (input) {
    const fileContent = fs.readFileSync(input, 'utf8').trim();

    return [
      this.sumString(fileContent),
      this.sumStringHalfway(fileContent)
    ];
  }
}

module.exports = Day01;
