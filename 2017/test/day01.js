/* global describe, it */

const assert = require('assert');
const fs = require('fs');
const Day01 = require('../day01');

describe('Day01', function () {
  describe('Part 1', function () {
    it('Example tests', function () {
      assert.strictEqual(Day01.sumString('1122'), 3);
      assert.strictEqual(Day01.sumString('1111'), 4);
      assert.strictEqual(Day01.sumString('1234'), 0);
      assert.strictEqual(Day01.sumString('91212129'), 9);
    });

    it('Input tests', function () {
      const fileContent = fs.readFileSync('./day01/input', 'utf8').trim();
      assert.strictEqual(Day01.sumString(fileContent), 1102);
    });
  });

  describe('Part 2', function () {
    it('Example tests', function () {
      assert.strictEqual(Day01.sumStringHalfway('1212'), 6);
      assert.strictEqual(Day01.sumStringHalfway('1221'), 0);
      assert.strictEqual(Day01.sumStringHalfway('123425'), 4);
      assert.strictEqual(Day01.sumStringHalfway('123123'), 12);
      assert.strictEqual(Day01.sumStringHalfway('12131415'), 4);
    });

    it('Input tests', function () {
      const fileContent = fs.readFileSync('./day01/input', 'utf8').trim();
      assert.strictEqual(Day01.sumStringHalfway(fileContent), 1076);
    });
  });
});
