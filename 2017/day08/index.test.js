/* global expect, test */

const Day08 = require('.');

test('Input', () => {
  const [part1, part2] = Day08.run('./day08/input');
  expect(part1).toBe(5966);
  expect(part2).toBe(6347);
});
