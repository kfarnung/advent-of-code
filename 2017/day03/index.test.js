/* global describe, expect, test */

const Day03 = require('.');

describe('Part 1', () => {
  test('Part 1', () => {
    expect(Day03.manhattanDistance(1)).toBe(0);
    expect(Day03.manhattanDistance(12)).toBe(3);
    expect(Day03.manhattanDistance(23)).toBe(2);
    expect(Day03.manhattanDistance(1024)).toBe(31);
  });

  test('Part 2', () => {
    expect(Day03.searchSums(0)).toBe(1);
    expect(Day03.searchSums(1)).toBe(2);
    expect(Day03.searchSums(3)).toBe(4);
    expect(Day03.searchSums(24)).toBe(25);
    expect(Day03.searchSums(350)).toBe(351);
  });
});

test('Input', () => {
  const [part1, part2] = Day03.run('../private/inputs/2017/day03.txt');
  expect(part1).toBe(430);
  expect(part2).toBe(312453);
});
