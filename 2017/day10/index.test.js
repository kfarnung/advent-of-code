/* global expect, test */

const Day10 = require('.');

test('Input', () => {
  const [part1, part2] = Day10.run('../private/inputs/2017/day10.txt');
  expect(part1).toBe(5577);
  expect(part2).toBe('44f4befb0f303c0bafd085f97741d51d');
});
