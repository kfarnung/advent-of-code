import Day25 from './index.js';
import { describe, test } from 'node:test';
import assert from 'node:assert/strict';

describe('Day 25', () => {
  describe('Part 1', () => {
    test('Input', () => {
      const [part1] = Day25.run('../private/inputs/2017/day25.txt');
      assert.strictEqual(part1, 2832);
    });
  });
});
