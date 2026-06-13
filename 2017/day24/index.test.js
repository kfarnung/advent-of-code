const Day24 = require('.');
const { describe, test } = require('node:test');
const assert = require('node:assert/strict');

describe('Day 24', () => {
  describe('Part 1', () => {
    test('Input', () => {
      const [part1] = Day24.run('../private/inputs/2017/day24.txt');
      assert.strictEqual(part1, 1868);
    });
  });

  describe('Part 2', () => {
    test('Input', () => {
      const [, part2] = Day24.run('../private/inputs/2017/day24.txt');
      assert.strictEqual(part2, 1841);
    });
  });
});
