
const Day21 = require('.');
const { describe, test } = require('node:test');
const assert = require('node:assert/strict');

describe('Day 21', () => {
  test('Input', () => {
    const [part1, part2] = Day21.run('../private/inputs/2017/day21.txt');
    assert.strictEqual(part1, 197);
    assert.strictEqual(part2, 3081737);
  });
});
