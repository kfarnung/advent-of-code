
const Day05 = require('.');
const { describe, test } = require('node:test');
const assert = require('node:assert/strict');

describe('Day 05', () => {
  test('Input', () => {
    const [part1, part2] = Day05.run('../private/inputs/2017/day05.txt');
    assert.strictEqual(part1, 339351);
    assert.strictEqual(part2, 24315397);
  });
});
