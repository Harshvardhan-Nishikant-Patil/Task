import svelte from 'rollup-plugin-svelte';
import commonjs from '@rollup/plugin-commonjs';
import resolve from '@rollup/plugin-node-resolve';
import livereload from 'rollup-plugin-livereload';
import { terser } from 'rollup-plugin-terser';

const production = !process.env.ROLLUP_WATCH;

export default {
    input: 'src/main.js',
    output: {
        sourcemap: true,
        format: 'iife',
        name: 'app',
        file: 'public/build/bundle.js',
    },
    plugins: [
        svelte({
            preprocess: require('svelte-preprocess')({
                sourceMap: !production,
                postcss: true,
            }),
            compilerOptions: {
                dev: !production,
            },
        }),
        resolve({
            browser: true,
            dedupe: ['svelte'],
        }),
        commonjs(),
        !production && serve(),
        !production && livereload('public'),
        production && terser(),
    ],
    watch: {
        clearScreen: false,
    },
};

function serve() {
    let started = false;

    return {
        writeBundle() {
            if (!started) {
                started = true;

                require('child_process').spawn('npm', ['run', 'start'], {
                    stdio: ['ignore', 'inherit', 'inherit'],
                    shell: true,
                });
            }
        },
    };
}
