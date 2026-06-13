const Day25 = require('.');
const { describe, test } = require('node:test');
const assert = require('node:assert/strict');

describe('Day 25', () => {
  describe('Part 1', () => {
    test('Input', () => {
      const [part1] = Day25.run('../private/inputs/2017/day25.txt');
      assert.strictEqual(part1, 2832);
    });
  });
});
