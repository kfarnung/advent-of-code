
const Day07 = require('.');
const { describe, test } = require('node:test');
const assert = require('node:assert/strict');

describe('Day 07', () => {
  test('Input', () => {
    const [part1, part2] = Day07.run('../private/inputs/2017/day07.txt');
    assert.strictEqual(part1, 'eqgvf');
    assert.strictEqual(part2, 757);
  });
});
