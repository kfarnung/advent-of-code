const fs = require('fs');

class Day10 {
  static * parseData (str) {
    let current = '';

    for (const ch of str) {
      if (ch === '\n' || ch === ',') {
        yield Number.parseInt(current);
        current = '';
      } else if (ch >= '0' && ch <= '9') {
        current += ch;
      }
    }
  }

  static * parseDataAsAscii (str) {
    for (const ch of str) {
      if (ch !== '\n') {
        yield ch.charCodeAt(0);
      }
    }

    for (const num of [17, 31, 73, 47, 23]) {
      yield num;
    }
  }

  static generateArray (start, length) {
    const arr = [];
    for (let i = start; i < length; i++) {
      arr.push(i);
    }

    return arr;
  }

  static reverse (arr, start, length) {
    const range = Math.floor(length / 2);
    const arrLength = arr.length;

    for (let i = 0; i < range; i++) {
      const first = (start + i) % arrLength;
      const second = (start + length - 1 - i) % arrLength;

      const temp = arr[first];
      arr[first] = arr[second];
      arr[second] = temp;
    }
  }

  static runHash (lengths, numRounds) {
    const arr = this.generateArray(0, 256);
    const arrLength = arr.length;
    let currentPosition = 0;
    let skipSize = 0;

    for (let i = 0; i < numRounds; i++) {
      for (const num of lengths) {
        this.reverse(arr, currentPosition, num);
        currentPosition = (currentPosition + num + skipSize) % arrLength;
        skipSize++;
      }
    }

    return arr;
  }

  static calculateHash (str) {
    const hash = this.runHash(Array.from(this.parseDataAsAscii(str)), 64);

    const hex = [];
    let xor = 0;
    for (let i = 0; i < hash.length; i++) {
      if (i > 0 && i % 16 === 0) {
        hex.push(xor);
        xor = 0;
      }

      xor ^= hash[i];
    }

    hex.push(xor);
    return hex;
  }

  static padLeft (str, width = 2, ch = '0') {
    let retVal = '';
    for (let i = 0; i < width - str.length; i++) {
      retVal += ch;
    }

    retVal += str;
    return retVal;
  }

  static getHashString (str) {
    const hash = this.calculateHash(str);
    const hexStr = [];
    for (const byte of hash) {
      hexStr.push(this.padLeft(byte.toString(16)));
    }

    return hexStr.join('');
  }

  static run (input) {
    const fileContent = fs.readFileSync(input, 'utf8');
    const hash1 = this.runHash(Array.from(this.parseData(fileContent)), 1);

    return [
      hash1[0] * hash1[1],
      this.getHashString(fileContent)
    ];
  }
}

module.exports = Day10;
