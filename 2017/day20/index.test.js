const Day20 = require('.');
const { describe, test } = require('node:test');
const assert = require('node:assert/strict');

describe('Day 20', () => {
  describe('Part 1', () => {
    test('Input', () => {
      const [part1] = Day20.run('../private/inputs/2017/day20.txt');
      assert.strictEqual(part1, 300);
    });
  });

  describe('Part 2', () => {
    test('Input', () => {
      const [, part2] = Day20.run('../private/inputs/2017/day20.txt');
      assert.strictEqual(part2, 502);
    });
  });
});
