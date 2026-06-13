const Day05 = require('.');
const { describe, test } = require('node:test');
const assert = require('node:assert/strict');

describe('Day 05', () => {
  describe('Part 1', () => {
    test('Input', () => {
      const [part1] = Day05.run('../private/inputs/2017/day05.txt');
      assert.strictEqual(part1, 339351);
    });
  });

  describe('Part 2', () => {
    test('Input', () => {
      const [, part2] = Day05.run('../private/inputs/2017/day05.txt');
      assert.strictEqual(part2, 24315397);
    });
  });
});
