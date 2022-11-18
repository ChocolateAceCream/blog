import { resolve } from 'path'
import { defineConfig, loadEnv } from 'vite'
import vue from '@vitejs/plugin-vue'

// vite.config.js
import Components from 'unplugin-vue-components/vite';
import { AntDesignVueResolver } from 'unplugin-vue-components/resolvers';

import Icons from 'unplugin-icons/vite'
import { FileSystemIconLoader } from 'unplugin-icons/loaders'
import IconsResolver from 'unplugin-icons/resolver'

function pathResolve() {
  return resolve(__dirname, '.', ...arguments)
}
// https://vitejs.dev/config/
export default defineConfig(params => {
  const { command, mode } = params
  const ENV = loadEnv(mode, process.cwd())
  console.info(`--- running mode: ${mode}, command: ${command}, ENV: ${JSON.stringify(ENV)} ---`)
  return {
    base: './',
    resolve: {
      extensions: ['.json', '.js', '.ts', '.vue'],
      alias: {
        '@': pathResolve('src'),
        '@images': pathResolve('src/assets/images'),
      }
    },
    plugins: [
      vue(),
      Icons({
        compiler: 'vue3',
        customCollections: {
          'icon': FileSystemIconLoader('src/assets/svgs'),
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
    ],
  }
})
