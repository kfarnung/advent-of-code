const Day04 = require('.');
const { describe, test } = require('node:test');
const assert = require('node:assert/strict');

describe('Day 04', () => {
  describe('Part 1', () => {
    test('Examples', () => {
      assert.strictEqual(Day04.isPassphraseValid('aa bb cc dd ee'), true);
      assert.strictEqual(Day04.isPassphraseValid('aa bb cc dd aa'), false);
      assert.strictEqual(Day04.isPassphraseValid('aa bb cc dd aaa'), true);
    });

    test('Input', () => {
      const [part1] = Day04.run('../private/inputs/2017/day04.txt');
      assert.strictEqual(part1, 455);
    });
  });

  describe('Part 2', () => {
    test('Examples', () => {
      assert.strictEqual(Day04.isPassphraseValid('abcde fghij', true), true);
      assert.strictEqual(Day04.isPassphraseValid('abcde xyz ecdab', true), false);
      assert.strictEqual(Day04.isPassphraseValid('a ab abc abd abf abj', true), true);
      assert.strictEqual(Day04.isPassphraseValid('iiii oiii ooii oooi oooo', true), true);
      assert.strictEqual(Day04.isPassphraseValid('oiii ioii iioi iiio', true), false);
    });

    test('Input', () => {
      const [, part2] = Day04.run('../private/inputs/2017/day04.txt');
      assert.strictEqual(part2, 186);
    });
  });
});
