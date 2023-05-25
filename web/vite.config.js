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
function pathResolve() {
  return resolve(__dirname, './', ...arguments)
  // return resolve(__dirname, '.', ...arguments)
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
              moment: ['moment'],
              'lodash-es': ['lodash-es'],
            }
          }
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
      AutoImport({
        resolvers: [ElementPlusResolver()],
      }),
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
          ElementPlusResolver(),
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
