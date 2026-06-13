const Day14 = require('.');
const { describe, test } = require('node:test');
const assert = require('node:assert/strict');

describe('Day 14', () => {
  describe('Part 1', () => {
    test('Input', () => {
      const [part1] = Day14.run('../private/inputs/2017/day14.txt');
      assert.strictEqual(part1, 8292);
    });
  });

  describe('Part 2', () => {
    test('Input', () => {
      const [, part2] = Day14.run('../private/inputs/2017/day14.txt');
      assert.strictEqual(part2, 1069);
    });
  });
});
