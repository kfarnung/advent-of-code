/* global expect, test */

const Day12 = require('.');

test('Input', () => {
  const [part1, part2] = Day12.run('./day12/input');
  expect(part1).toBe(169);
  expect(part2).toBe(179);
});
