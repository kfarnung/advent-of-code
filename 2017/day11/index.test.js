/* global expect, test */

const Day11 = require('.');

test('Input', () => {
  const [part1, part2] = Day11.run('../private/inputs/2017/day11.txt');
  expect(part1).toBe(696);
  expect(part2).toBe(1461);
});
