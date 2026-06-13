
const Day10 = require('.');
const { describe, test } = require('node:test');
const assert = require('node:assert/strict');

describe('Day 10', () => {
  test('Input', () => {
    const [part1, part2] = Day10.run('../private/inputs/2017/day10.txt');
    assert.strictEqual(part1, 5577);
    assert.strictEqual(part2, '44f4befb0f303c0bafd085f97741d51d');
  });
});
