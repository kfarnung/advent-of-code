
const Day17 = require('.');
const { describe, test } = require('node:test');
const assert = require('node:assert/strict');

describe('Day 17', () => {
  test('Input', () => {
    const [part1, part2] = Day17.run('../private/inputs/2017/day17.txt');
    assert.strictEqual(part1, 419);
    assert.strictEqual(part2, 46038988);
  });
});
