import Day17 from './index.js';
import { describe, test } from 'node:test';
import assert from 'node:assert/strict';

describe('Day 17', () => {
  describe('Part 1', () => {
    test('Input', () => {
      const [part1] = Day17.run('../private/inputs/2017/day17.txt');
      assert.strictEqual(part1, 419);
    });
  });

  describe('Part 2', () => {
    test('Input', () => {
      const [, part2] = Day17.run('../private/inputs/2017/day17.txt');
      assert.strictEqual(part2, 46038988);
    });
  });
});
