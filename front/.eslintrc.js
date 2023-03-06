module.exports = {
  env: {
    browser: true,
    es2021: true
  },
  extends: [
    'plugin:react/recommended',
    'standard',
  ],
  parser: '@typescript-eslint/parser',
  parserOptions: {
    ecmaFeatures: {
      jsx: true
    },
    ecmaVersion: 12,
    sourceType: 'module'
  },
  plugins: [
    'react',
    '@typescript-eslint'
  ],
  rules: {
    semi: 'error',
    // let と const の区別を厳格化
    'prefer-const': 'error',
    // === の使用を推奨
    eqeqeq: 'error',
    // import React from 'react'React していないと 'React must be in scope when using JSX.'が出る
    // 設定で 出さないように https://ja.reactjs.org/blog/2020/09/22/introducing-the-new-jsx-transform.html#eslint
    'react/jsx-uses-react': 'off',
    'react/react-in-jsx-scope': 'off',
  },
  settings: {
    react: {
      // Warning: React version not specified in eslint-plugin-react settings. が出ないように
      version: 'detect'
    }
  }
}
