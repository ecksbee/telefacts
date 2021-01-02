  
import resolve from '@rollup/plugin-node-resolve';
import commonjs from '@rollup/plugin-commonjs';
import { terser } from 'rollup-plugin-terser';
import copy from 'rollup-plugin-copy';

const production = !process.env.ROLLUP_WATCH;
const copyConfig = {
  targets: [
    { src: 'node_modules/@webcomponents', dest: '../assets/node_modules' },
  ],
};
export default {
	input: 'main.js',
	output: {
		file: '../assets/bundle.js',
		format: 'iife',
	},
	plugins: [
		copy(copyConfig),
		resolve(),
		commonjs(),
		production && terser() // minify, but only in production
	]
};