/* global expect, test */

const Day25 = require('.');

test('Input', () => {
  const [part1] = Day25.run('../private/inputs/2017/day25.txt');
  expect(part1).toBe(2832);
});
