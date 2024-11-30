/* global expect, test */

const Day07 = require('.');

test('Input', () => {
  const [part1, part2] = Day07.run('../private/inputs/2017/day07.txt');
  expect(part1).toBe('eqgvf');
  expect(part2).toBe(757);
});
