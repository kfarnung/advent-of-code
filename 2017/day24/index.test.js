/* global expect, test */

const Day24 = require('.');

test('Input', () => {
  const [part1, part2] = Day24.run('./day24/input');
  expect(part1).toBe(1868);
  expect(part2).toBe(1841);
});
