import Day23 from './index.js';
import { describe, test } from 'node:test';
import assert from 'node:assert/strict';

describe('Day 23', () => {
  describe('Part 1', () => {
    test('Input', () => {
      const [part1] = Day23.run('../private/inputs/2017/day23.txt');
      assert.strictEqual(part1, 4225);
    });
  });

  describe('Part 2', () => {
    test('Input', () => {
      const [, part2] = Day23.run('../private/inputs/2017/day23.txt');
      assert.strictEqual(part2, 905);
    });
  });
});
