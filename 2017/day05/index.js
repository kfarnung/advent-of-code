const fs = require('fs');

class Day05 {
  static * parseInt (str) {
    let current = '';

    for (const ch of str) {
      if (ch === '\n') {
        yield Number.parseInt(current);
        current = '';
      } else if ((ch >= '0' && ch <= '9') || ch === '-') {
        current += ch;
      }
    }
  }

  static incrementModifier (arr, index) {
    arr[index]++;
  }

  static decrementThreeModifier (arr, index) {
    if (arr[index] >= 3) {
      arr[index]--;
    } else {
      this.incrementModifier(arr, index);
    }
  }

  static followSteps (arr, modifierFunc) {
    let index = 0;
    let numSteps = 0;
    while (index >= 0 && index < arr.length) {
      const nextOffset = arr[index];
      modifierFunc.call(this, arr, index);
      index += nextOffset;
      numSteps++;
    }

    return numSteps;
  }

  static run (input) {
    const fileContent = fs.readFileSync(input, 'utf8');
    const arr1 = Array.from(this.parseInt(fileContent));
    const arr2 = Array.from(arr1);

    return [
      this.followSteps(arr1, this.incrementModifier),
      this.followSteps(arr2, this.decrementThreeModifier)
    ];
  }
}

module.exports = Day05;
