const globals = require('globals');
const js = require('@eslint/js');

module.exports = [js.configs.recommended, {
  languageOptions: {
    globals: {
      ...globals.commonjs,
      ...globals.node,
    },

    ecmaVersion: 12,
    sourceType: 'commonjs',
  },

  rules: {
    indent: ['error', 2, {
      SwitchCase: 1,
    }],

    'linebreak-style': ['error', 'unix'],

    quotes: ['error', 'single', {
      avoidEscape: true,
    }],

    semi: ['error', 'always'],
  },
}];