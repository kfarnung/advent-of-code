
const Day09 = require('.');
const { describe, test } = require('node:test');
const assert = require('node:assert/strict');

describe('Day 09', () => {
  test('Input', () => {
    const [part1, part2] = Day09.run('../private/inputs/2017/day09.txt');
    assert.strictEqual(part1, 16021);
    assert.strictEqual(part2, 7685);
  });
});
