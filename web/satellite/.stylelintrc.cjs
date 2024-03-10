// Copyright (C) 2021 Storj Labs, Inc.
// See LICENSE for copying information.

module.exports = {
    'env': {
        'es2020': true,
        'node': true,
        'jest': true,
    },
    'plugins': [
        'stylelint-scss',
        '@stylistic/stylelint-plugin',
    ],
    'extends': 'stylelint-config-standard-vue/scss',
    'customSyntax': 'postcss-html',
    'rules': {
        'at-rule-no-unknown': null,
        'comment-empty-line-before': 'always',
        'comment-whitespace-inside': 'always',
        'custom-property-pattern': '.*',
        'function-url-quotes': 'always',
        'font-family-name-quotes': 'always-unless-keyword',
        'font-family-no-missing-generic-family-keyword': true,
        'no-duplicate-selectors': true,
        'no-descending-specificity': null,
        'rule-empty-line-before': 'always-multi-line',
        'scss/at-rule-no-unknown': true,
        'selector-class-pattern': '.*',
        'selector-max-attribute': 1,
        'selector-max-type': 3,
        'selector-pseudo-element-colon-notation': 'single',
        'selector-pseudo-element-no-unknown': [
            true,
            {
                'ignorePseudoElements': ['v-deep'],
            },
        ],
        '@stylistic/indentation': 4,
        '@stylistic/string-quotes': 'single',
        '@stylistic/selector-combinator-space-after': 'always',
        '@stylistic/selector-attribute-operator-space-before': 'never',
        '@stylistic/selector-attribute-operator-space-after': 'never',
        '@stylistic/selector-attribute-brackets-space-inside': 'never',
        '@stylistic/declaration-block-trailing-semicolon': 'always',
        '@stylistic/declaration-colon-space-before': 'never',
        '@stylistic/declaration-colon-space-after': 'always-single-line',
        '@stylistic/number-leading-zero': 'always',
        '@stylistic/selector-pseudo-class-parentheses-space-inside': 'never',
        '@stylistic/media-feature-range-operator-space-before': 'always',
        '@stylistic/media-feature-range-operator-space-after': 'always',
        '@stylistic/media-feature-parentheses-space-inside': 'never',
        '@stylistic/media-feature-colon-space-before': 'never',
        '@stylistic/media-feature-colon-space-after': 'always',
    },
};