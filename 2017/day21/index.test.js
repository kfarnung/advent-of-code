import Day21 from './index.js';
import { describe, test } from 'node:test';
import assert from 'node:assert/strict';

describe('Day 21', () => {
  describe('Part 1', () => {
    test('Input', () => {
      const [part1] = Day21.run('../private/inputs/2017/day21.txt');
      assert.strictEqual(part1, 197);
    });
  });

  describe('Part 2', () => {
    test('Input', () => {
      const [, part2] = Day21.run('../private/inputs/2017/day21.txt');
      assert.strictEqual(part2, 3081737);
    });
  });
});
