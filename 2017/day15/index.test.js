
const Day15 = require('.');
const { describe, test } = require('node:test');
const assert = require('node:assert/strict');

describe('Day 15', () => {
  test('Input', () => {
    const [part1, part2] = Day15.run('../private/inputs/2017/day15.txt');
    assert.strictEqual(part1, 638);
    assert.strictEqual(part2, 343);
  });
});
