/* global expect, test */

const Day12 = require('.');

test('Input', () => {
  const [part1, part2] = Day12.run('../private/inputs/2017/day12.txt');
  expect(part1).toBe(169);
  expect(part2).toBe(179);
});
