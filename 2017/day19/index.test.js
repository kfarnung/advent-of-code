/* global expect, test */

const Day19 = require('.');

test('Input', () => {
  const [part1, part2] = Day19.run('../private/inputs/2017/day19.txt');
  expect(part1).toBe('RUEDAHWKSM');
  expect(part2).toBe(17264);
});
