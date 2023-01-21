module.exports = {
  parser: "@typescript-eslint/parser",
  parserOptions: {
    ecmaVersion: 2020,
    sourceType: "module",
    ecmaFeatures: {
      jsx: true,
    },
  },
  extends: [
    "plugin:react/recommended",
    "plugin:@typescript-eslint/recommended",
    "plugin:prettier/recommended",
  ],
  plugins: ["simple-import-sort"],
  globals: {
    Promise: true,
  },
  rules: {
    "no-shadow": "off",
    "no-catch-shadow": "off",
    "react/prop-types": "off",
    "@typescript-eslint/no-unused-vars": ["warn", { argsIgnorePattern: "^_" }],
    "simple-import-sort/imports": [
      "error",
      {
        groups: [
          ["^\\u0000"],
          ["^@?\\w"],
          ["^src"],
          ["^"],
          ["^\\."],
          ["^.+\\u0000$"],
        ],
      },
    ],
    "simple-import-sort/exports": "error",
  },
};
