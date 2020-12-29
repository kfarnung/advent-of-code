/* global expect, test */

const Day11 = require('.');

test('Input', () => {
  const [part1, part2] = Day11.run('./day11/input');
  expect(part1).toBe(696);
  expect(part2).toBe(1461);
});
