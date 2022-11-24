import { resolve } from 'path'
import { defineConfig, loadEnv } from 'vite'
import vue from '@vitejs/plugin-vue'

// vite.config.js
import Components from 'unplugin-vue-components/vite'
import { AntDesignVueResolver } from 'unplugin-vue-components/resolvers'

import Icons from 'unplugin-icons/vite'
import { FileSystemIconLoader } from 'unplugin-icons/loaders'
import IconsResolver from 'unplugin-icons/resolver'

import AutoImport from 'unplugin-auto-import/vite'
import { visualizer } from "rollup-plugin-visualizer"
function pathResolve() {
  return resolve(__dirname, '.', ...arguments)
}

const plugin = [
  AutoImport({
    // targets to transform
    include: [
      /\.[tj]sx?$/,
      /\.vue$/,
      /\.vue\?vue/,
      /\.md$/,
      /\.js$/,
      './auto-imports.d.ts',
    ],

    // global imports to register
    imports: [
      // 插件预设支持导入的api
      'vue',
      'vue-router',
      'pinia',
      // 自定义导入的api
    ],

    // Generate corresponding .eslintrc-auto-import.json file.
    // eslint globals Docs - https://eslint.org/docs/user-guide/configuring/language-options#specifying-globals
    eslintrc: {
      enabled: true, // Default `false`
      filepath: './.eslintrc-auto-import.json', // Default `./.eslintrc-auto-import.json`
      globalsPropValue: true, // Default `true`, (true | false | 'readonly' | 'readable' | 'writable' | 'writeable')
    },

    // Filepath to generate corresponding .d.ts file.
    // Defaults to './auto-imports.d.ts' when `typescript` is installed locally.
    // Set `false` to disable.
    // dts: 'src/auto-imports.d.ts',
    dts: true,
  }),
  vue(),
  Icons({
    compiler: 'vue3',
    customCollections: {
      icon: FileSystemIconLoader('src/assets/svgs'),
      // usage:
      //  <i-svg-vue style="font-size: 50px; fill: red;" />
      //  <i-icon-vue style="font-size: 50px; fill: red;" />
    },
  }),
  Components({
    resolvers: [
      AntDesignVueResolver(),
      IconsResolver({
        alias: {
          svg: 'icon',
        },
        customCollections: ['icon'],
      }),
    ],
  }),
]
// https://vitejs.dev/config/
export default defineConfig((params) => {
  const { command, mode } = params
  const ENV = loadEnv(mode, process.cwd())
  // if (mode == 'development') {
  //   plugin.push(visualizer({
  //     open: true,
  //     gzipSize: true,
  //     brotliSize: true,
  //   }))
  // }
  console.info(`--- running mode: ${mode}, command: ${command}, ENV: ${JSON.stringify(ENV)} ---`)
  return {
    base: './',
    resolve: {
      extensions: ['.json', '.js', '.ts', '.vue'],
      alias: {
        '@': pathResolve('src'),
        '@images': pathResolve('src/assets/images'),
      },
    },
    plugins: [
      visualizer({
        open: true,
        gzipSize: true,
        brotliSize: true,
      }),
      AutoImport({
        // targets to transform
        include: [
          /\.[tj]sx?$/,
          /\.vue$/,
          /\.vue\?vue/,
          /\.md$/,
          /\.js$/,
          './auto-imports.d.ts',
        ],

        // global imports to register
        imports: [
          // 插件预设支持导入的api
          'vue',
          'vue-router',
          'pinia',
          // 自定义导入的api
        ],

        // Generate corresponding .eslintrc-auto-import.json file.
        // eslint globals Docs - https://eslint.org/docs/user-guide/configuring/language-options#specifying-globals
        eslintrc: {
          enabled: true, // Default `false`
          filepath: './.eslintrc-auto-import.json', // Default `./.eslintrc-auto-import.json`
          globalsPropValue: true, // Default `true`, (true | false | 'readonly' | 'readable' | 'writable' | 'writeable')
        },

        // Filepath to generate corresponding .d.ts file.
        // Defaults to './auto-imports.d.ts' when `typescript` is installed locally.
        // Set `false` to disable.
        // dts: 'src/auto-imports.d.ts',
        dts: true,
      }),
      vue(),
      Icons({
        compiler: 'vue3',
        customCollections: {
          icon: FileSystemIconLoader('src/assets/svgs'),
          // usage:
          //  <i-svg-vue style="font-size: 50px; fill: red;" />
          //  <i-icon-vue style="font-size: 50px; fill: red;" />
        },
      }),
      Components({
        resolvers: [
          AntDesignVueResolver(),
          IconsResolver({
            alias: {
              svg: 'icon',
            },
            customCollections: ['icon'],
          }),
        ],
      }),
    ]
  }
})
