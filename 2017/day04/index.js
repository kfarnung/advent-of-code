const fs = require('fs');

class Day04 {
  static *parseRows(str) {
    let current = '';

    for (const ch of str) {
      if (ch === '\n' || ch === '\r') {
        if (current) {
          yield current;
        }

        current = '';
      } else {
        current += ch;
      }
    }
  }

  static isPassphraseValid(passphrase, checkAnagrams = false) {
    const set = new Set();
    const words = passphrase.split(' ');
    for (let word of words) {
      if (checkAnagrams) {
        word = word.split('').sort().join('');
      }

      if (set.has(word)) {
        return false;
      }

      set.add(word);
    }

    return true;
  }

  static run(input) {
    const fileContent = fs.readFileSync(input, 'utf8');

    let validCount = 0;
    let validAnagramCount = 0;
    for (const passphrase of this.parseRows(fileContent)) {
      if (this.isPassphraseValid(passphrase)) {
        validCount++;
      }

      if (this.isPassphraseValid(passphrase, true)) {
        validAnagramCount++;
      }
    }

    return [validCount, validAnagramCount];
  }
}

module.exports = Day04;
