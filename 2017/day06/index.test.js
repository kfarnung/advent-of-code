/* global expect, test */

const Day06 = require('.');

test('Input', () => {
  const [part1, part2] = Day06.run('./day06/input');
  expect(part1).toBe(7864);
  expect(part2).toBe(1695);
});
