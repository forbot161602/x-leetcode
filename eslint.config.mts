import js from '@eslint/js';
import globals from 'globals';
import tseslint from 'typescript-eslint';
import stylistic from '@stylistic/eslint-plugin';
import {defineConfig} from 'eslint/config';

export default defineConfig([
    {
        files: [
            'src/**/*.{js,mjs,cjs,ts,mts,cts}',
        ],
        extends: [
            js.configs.recommended,
            ...tseslint.configs.recommendedTypeChecked,
            ...tseslint.configs.stylisticTypeChecked,
            stylistic.configs.recommended,
        ],
        languageOptions: {
            globals: {
                ...globals.node,
            },
            parser: tseslint.parser,
            parserOptions: {
                project: './tsconfig.json',
            },
        },
        rules: {
            'no-console': [
                'off',
            ],

            '@typescript-eslint/consistent-type-definitions': [
                'off',
            ],
            '@typescript-eslint/no-empty-object-type': [
                'error', {
                    'allowInterfaces': 'with-single-extends',
                },
            ],
            '@typescript-eslint/no-unsafe-enum-comparison': [
                'off',
            ],

            '@stylistic/indent': [
                'error', 4,
            ],
            '@stylistic/indent-binary-ops': [
                'error', 4,
            ],
            '@stylistic/object-curly-spacing': [
                'error', 'never',
            ],
            '@stylistic/semi': [
                'error', 'always',
            ],
        },
    },
]);
