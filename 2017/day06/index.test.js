const Day06 = require('.');
const { describe, test } = require('node:test');
const assert = require('node:assert/strict');

describe('Day 06', () => {
  describe('Part 1', () => {
    test('Input', () => {
      const [part1] = Day06.run('../private/inputs/2017/day06.txt');
      assert.strictEqual(part1, 7864);
    });
  });

  describe('Part 2', () => {
    test('Input', () => {
      const [, part2] = Day06.run('../private/inputs/2017/day06.txt');
      assert.strictEqual(part2, 1695);
    });
  });
});
