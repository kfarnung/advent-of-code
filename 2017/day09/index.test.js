const Day09 = require('.');
const { describe, test } = require('node:test');
const assert = require('node:assert/strict');

describe('Day 09', () => {
  describe('Part 1', () => {
    test('Input', () => {
      const [part1] = Day09.run('../private/inputs/2017/day09.txt');
      assert.strictEqual(part1, 16021);
    });
  });

  describe('Part 2', () => {
    test('Input', () => {
      const [, part2] = Day09.run('../private/inputs/2017/day09.txt');
      assert.strictEqual(part2, 7685);
    });
  });
});
