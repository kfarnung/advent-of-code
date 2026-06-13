const Day07 = require('.');
const { describe, test } = require('node:test');
const assert = require('node:assert/strict');

describe('Day 07', () => {
  describe('Part 1', () => {
    test('Input', () => {
      const [part1] = Day07.run('../private/inputs/2017/day07.txt');
      assert.strictEqual(part1, 'eqgvf');
    });
  });

  describe('Part 2', () => {
    test('Input', () => {
      const [, part2] = Day07.run('../private/inputs/2017/day07.txt');
      assert.strictEqual(part2, 757);
    });
  });
});
