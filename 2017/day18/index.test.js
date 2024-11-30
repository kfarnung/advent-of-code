/* global expect, test */

const Day18 = require('.');

test('Input', () => {
  const [part1, part2] = Day18.run('../private/inputs/2017/day18.txt');
  expect(part1).toBe(2951);
  expect(part2).toBe(7366);
});
