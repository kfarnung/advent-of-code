/* global describe, expect, test */

const Day04 = require('.');

describe('Examples', () => {
  test('Part 1', () => {
    expect(Day04.isPassphraseValid('aa bb cc dd ee')).toBe(true);
    expect(!Day04.isPassphraseValid('aa bb cc dd aa')).toBe(true);
    expect(Day04.isPassphraseValid('aa bb cc dd aaa')).toBe(true);
  });

  test('Part 2', () => {
    expect(Day04.isPassphraseValid('abcde fghij', true)).toBe(true);
    expect(!Day04.isPassphraseValid('abcde xyz ecdab', true)).toBe(true);
    expect(Day04.isPassphraseValid('a ab abc abd abf abj', true)).toBe(true);
    expect(Day04.isPassphraseValid('iiii oiii ooii oooi oooo', true)).toBe(
      true
    );
    expect(!Day04.isPassphraseValid('oiii ioii iioi iiio', true)).toBe(true);
  });
});

test('Input', () => {
  const [part1, part2] = Day04.run('../private/inputs/2017/day04.txt');
  expect(part1).toBe(455);
  expect(part2).toBe(186);
});
