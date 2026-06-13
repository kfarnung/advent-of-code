
const Day14 = require('.');
const { describe, test } = require('node:test');
const assert = require('node:assert/strict');

describe('Day 14', () => {
  test('Input', () => {
    const [part1, part2] = Day14.run('../private/inputs/2017/day14.txt');
    assert.strictEqual(part1, 8292);
    assert.strictEqual(part2, 1069);
  });
});
