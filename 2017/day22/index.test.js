/* global expect, test */

const Day22 = require('.');

test('Input', () => {
  const [part1, part2] = Day22.run('../private/inputs/2017/day22.txt');
  expect(part1).toBe(5261);
  expect(part2).toBe(2511927);
});
