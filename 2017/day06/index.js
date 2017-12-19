const fs = require('fs');

class Day06 {
  static * parseInt (str) {
    let current = '';

    for (const ch of str) {
      if (ch === '\n' || ch === '\t') {
        yield Number.parseInt(current);
        current = '';
      } else if (ch >= '0' && ch <= '9') {
        current += ch;
      }
    }
  }

  static findMaxIndex (arr) {
    let max = 0;
    let index = 0;

    for (let i = 0; i < arr.length; i++) {
      if (arr[i] > max) {
        max = arr[i];
        index = i;
      }
    }

    return index;
  }

  static findCycle (arr) {
    const map = new Map();
    let cycleCount = 0;

    while (true) {
      const key = arr.join();
      if (map.has(key)) {
        return { cycleCount, cycleSize: cycleCount - map.get(key) };
      }

      map.set(key, cycleCount);

      let index = this.findMaxIndex(arr);
      let memory = arr[index];
      arr[index] = 0;

      while (memory > 0) {
        arr[++index % arr.length] += 1;
        memory--;
      }

      cycleCount++;
    }
  }

  static run (input) {
    const fileContent = fs.readFileSync(input, 'utf8');

    const { cycleCount, cycleSize } = this.findCycle(Array.from(this.parseInt(fileContent)));
    return [
      cycleCount,
      cycleSize
    ];
  }
}

module.exports = Day06;
