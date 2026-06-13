
const Day16 = require('.');
const { describe, test } = require('node:test');
const assert = require('node:assert/strict');

describe('Day 16', () => {
  test('Input', () => {
    const [part1, part2] = Day16.run('../private/inputs/2017/day16.txt');
    assert.strictEqual(part1, 'padheomkgjfnblic');
    assert.strictEqual(part2, 'bfcdeakhijmlgopn');
  });
});
