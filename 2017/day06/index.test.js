/* global expect, test */

const Day06 = require('.');

test('Input', () => {
  const [part1, part2] = Day06.run('../private/inputs/2017/day06.txt');
  expect(part1).toBe(7864);
  expect(part2).toBe(1695);
});
