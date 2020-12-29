/* global expect, test */

const Day14 = require('.');

test('Input', () => {
  const [part1, part2] = Day14.run('./day14/input');
  expect(part1).toBe(8292);
  expect(part2).toBe(1069);
});
