import Day11 from './index.js';
import { describe, test } from 'node:test';
import assert from 'node:assert/strict';

describe('Day 11', () => {
  describe('Part 1', () => {
    test('Input', () => {
      const [part1] = Day11.run('../private/inputs/2017/day11.txt');
      assert.strictEqual(part1, 696);
    });
  });

  describe('Part 2', () => {
    test('Input', () => {
      const [, part2] = Day11.run('../private/inputs/2017/day11.txt');
      assert.strictEqual(part2, 1461);
    });
  });
});
