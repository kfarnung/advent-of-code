import Day15 from './index.js';
import { describe, test } from 'node:test';
import assert from 'node:assert/strict';

describe('Day 15', () => {
  describe('Part 1', () => {
    test('Input', () => {
      const [part1] = Day15.run('../private/inputs/2017/day15.txt');
      assert.strictEqual(part1, 638);
    });
  });

  describe('Part 2', () => {
    test('Input', () => {
      const [, part2] = Day15.run('../private/inputs/2017/day15.txt');
      assert.strictEqual(part2, 343);
    });
  });
});
