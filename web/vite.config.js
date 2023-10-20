import { resolve } from 'path'
import { defineConfig, loadEnv } from 'vite'
import vue from '@vitejs/plugin-vue'
import Banner from 'vite-plugin-banner'

// vite.config.js
import AutoImport from 'unplugin-auto-import/vite'
import Components from 'unplugin-vue-components/vite'
import { ElementPlusResolver } from 'unplugin-vue-components/resolvers'

import Icons from 'unplugin-icons/vite'
import { FileSystemIconLoader } from 'unplugin-icons/loaders'
import IconsResolver from 'unplugin-icons/resolver'


import { visualizer } from 'rollup-plugin-visualizer'
import importToCDN from 'vite-plugin-cdn-import'

import externalGlobals from 'rollup-plugin-external-globals'

function pathResolve() {
  return resolve(__dirname, './', ...arguments)
  // return resolve(__dirname, '.', ...arguments)
}

const externalGlobalsObj = {
  vue: 'Vue',
  'vue-demi': 'VueDemi',
  'vue-router': 'VueRouter',
  'element-plus': 'ElementPlus',
}


// https://vitejs.dev/config/
export default defineConfig((params) => {
  const { command, mode } = params
  const ENV = loadEnv(mode, process.cwd())
  const timestamp = Date.parse(new Date())

  console.info(`--- running mode: ${mode}, command: ${command}, ENV: ${JSON.stringify(ENV)} ---`)
  return {
    base: './',
    root: './', // js导入的资源路径，src
    resolve: {
      extensions: ['.json', '.js', '.ts', '.vue'],
      alias: {
        '@': pathResolve('src'),
        '/img': pathResolve('src/assets/images'),
        'vue-i18n': 'vue-i18n/dist/vue-i18n.cjs.js'
      },
    },
    server: {
      port: ENV.VITE_APP_PORT,
      host: ENV.VITE_APP_HOST,
      proxy: {
        '/backend': {
          target: ENV.VITE_APP_DEV_PROXY,
          changeOrigin: true,
          rewrite: (path) => path.replace(/^\/backend/, '')
        },
        '/websocket': {
          target: ENV.VITE_WEBSOCKET_LOCAL,
          changeOrigin: true,
          rewrite: (path) => path.replace(/^\/websocket/, ''),
          ws: true,
          // configure: (proxy, _options) => {
          //   proxy.on('error', (err, _req, _res) => {
          //     console.log('proxy error', err)
          //   })
          //   proxy.on('proxyReq', (proxyReq, req, _res) => {
          //     console.log('Sending Request to the Target:', req.method, req.url)
          //   })
          //   proxy.on('proxyRes', (proxyRes, req, _res) => {
          //     console.log('Received Response from the Target:', proxyRes.statusCode, req.url)
          //   })
          // },
        }
      },
    },
    build: {
      minify: 'terser', // 必须启用：terserOptions配置才会有效
      terserOptions: {
        compress: {
          // 生产环境时移除console.log调试代码
          drop_console: true,
          drop_debugger: true,
        }
      },
      target: 'es2015',
      manifest: false,
      sourcemap: false,
      outDir: 'dist',
      build: {
        rollupOptions: {
          output: {
            manualChunks: {
              // moment: ['moment'],
              'lodash-es': ['lodash-es'],
              'md-editor-v3': ['md-editor-v3'],
              'dayjs': 'dayjs',
            },
          },
        }
      }
    },
    css: {
      preprocessorOptions: {
        scss: {
          // additionalData: `$injectedColor: orange;`
          additionalData: `
            @import "@/assets/styles/globalInjectedData.scss";
          `,
        }
      }
    },
    plugins: [
      // analyze pkg size
      visualizer({
        open: true,
        gzipSize: true,
        brotliSize: true,
      }),


      importToCDN.Plugin({
        modules: [
          {
            name: 'vue',
            var: 'Vue',
            path: `https://unpkg.com/vue@3.3.0/dist/vue.global.prod.js`,
          },
          {
            name: 'vue-demi',
            var: 'VueDemi',
            path: `https://cdn.jsdelivr.net/npm/vue-demi@0.12.5`,
          },
          {
            name: 'vue-router',
            var: 'VueRouter',
            path: `https://unpkg.com/vue-router@4.1.6`,
          },
          {
            name: 'md-editor-v3',
            var: 'MdEditorV3',
            path: 'https://unpkg.com/md-editor-v3@4.0.4/lib/umd/index.js'
          },
          {
            name: 'element-plus/lib/locale/lang/en',
            var: 'ElementPlusLocaleEn',
            path: 'https://unpkg.com/element-plus@2.3.14/dist/locale/en.js'
          },
          {
            name: 'element-plus/lib/locale/lang/zh-cn',
            var: 'ElementPlusLocaleZhCn',
            path: 'https://unpkg.com/element-plus@2.3.14/dist/locale/zh-cn'
          },
          {
            name: 'element-plus',
            var: 'ElementPlus',
            path: 'https://unpkg.com/element-plus@2.3.14/dist/index.full.min.js'
          },
          {
            name: 'emoji-mart-vue-fast',
            var: 'EmojiMart',
            path: 'https://cdn.jsdelivr.net/npm/emoji-mart-vue-fast@15.0.0/dist/emoji-mart.min.js'
          },
          {
            name: 'dayjs',
            var: 'dayjs',
            path: 'https://unpkg.com/dayjs@1.8.21/dayjs.min.js'
          },
          {
            name: 'dayjs/locale/zh-cn',
            var: 'dayjs_locale_zh_cn',
            path: 'https://unpkg.com/dayjs@1.8.21/locale/zh-cn.js'
          },
          {
            name: 'dayjs/plugin/relativeTime',
            var: 'dayjs_plugin_relativeTime',
            path: 'https://unpkg.com/dayjs@1.8.21/plugin/relativeTime.js'
          },
        ],
      }),



      AutoImport({
        resolvers: [ElementPlusResolver({
          importStyle: false,
        })],
        imports: ['vue', 'vue-router']
      }),
      {
        ...externalGlobals(externalGlobalsObj),
        enforce: 'post',
        apply: 'build',
      },
      [Banner(`
  #####                                                           #                   #####
#     # #    #  ####   ####   ####  #        ##   ##### ######   # #    ####  ###### #     # #####  ######   ##   #    #
#       #    # #    # #    # #    # #       #  #    #   #       #   #  #    # #      #       #    # #       #  #  #    #
#       ###### #    # #      #    # #      #    #   #   #####  #     # #      #####  #       #    # #####  #    # ##  ##
#       #    # #    # #      #    # #      ######   #   #      ####### #      #      #       #####  #      ###### # ## #
#     # #    # #    # #    # #    # #      #    #   #   #      #     # #    # #      #     # #   #  #      #    # #    #
 #####  #    #  ####   ####   ####  ###### #    #   #   ###### #     #  ####  ######  #####  #    # ###### #    # #    #
        \n Build on Time : ${timestamp}`)],
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
          ElementPlusResolver({
            importStyle: false,
          }),
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
