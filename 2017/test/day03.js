/* global describe, it */

const assert = require('assert');
const Day03 = require('../day03');

describe('Day03', function () {
  describe('Part 1', function () {
    it('Example tests', function () {
      assert.strictEqual(Day03.manhattanDistance(1), 0);
      assert.strictEqual(Day03.manhattanDistance(12), 3);
      assert.strictEqual(Day03.manhattanDistance(23), 2);
      assert.strictEqual(Day03.manhattanDistance(1024), 31);
    });

    it('Input tests', function () {
      assert.strictEqual(Day03.manhattanDistance(312051), 430);
    });
  });

  describe('Part 2', function () {
    it('Example tests', function () {
      assert.strictEqual(Day03.searchSums(0), 1);
      assert.strictEqual(Day03.searchSums(1), 2);
      assert.strictEqual(Day03.searchSums(3), 4);
      assert.strictEqual(Day03.searchSums(24), 25);
      assert.strictEqual(Day03.searchSums(350), 351);
    });

    it('Input tests', function () {
      assert.strictEqual(Day03.searchSums(312051), 312453);
    });
  });
});
