/* global describe, it */

const assert = require('assert');
const fs = require('fs');
const Day04 = require('../day04');

describe('Day04', function () {
  describe('Part 1', function () {
    it('Example tests', function () {
      assert(Day04.isPassphraseValid('aa bb cc dd ee'));
      assert(!Day04.isPassphraseValid('aa bb cc dd aa'));
      assert(Day04.isPassphraseValid('aa bb cc dd aaa'));
    });

    it('Input tests', function () {
      const fileContent = fs.readFileSync('./day04/input', 'utf8');
      let count = 0;
      for (const passphrase of Day04.parseRows(fileContent)) {
        if (Day04.isPassphraseValid(passphrase)) {
          count++;
        }
      }

      assert.strictEqual(count, 455);
    });
  });

  describe('Part 2', function () {
    it('Example tests', function () {
      assert(Day04.isPassphraseValid('abcde fghij', true));
      assert(!Day04.isPassphraseValid('abcde xyz ecdab', true));
      assert(Day04.isPassphraseValid('a ab abc abd abf abj', true));
      assert(Day04.isPassphraseValid('iiii oiii ooii oooi oooo', true));
      assert(!Day04.isPassphraseValid('oiii ioii iioi iiio', true));
    });

    it('Input tests', function () {
      const fileContent = fs.readFileSync('./day04/input', 'utf8');
      let count = 0;
      for (const passphrase of Day04.parseRows(fileContent)) {
        if (Day04.isPassphraseValid(passphrase, true)) {
          count++;
        }
      }

      assert.strictEqual(count, 186);
    });
  });
});
