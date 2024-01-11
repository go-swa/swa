import legacyPlugin from '@vitejs/plugin-legacy'
import AutoImport from 'unplugin-auto-import/vite'
import Components from 'unplugin-vue-components/vite'
import {ElementPlusResolver} from 'unplugin-vue-components/resolvers'
import {viteLogo} from './src/core/config'
import Banner from 'vite-plugin-banner'
import * as path from 'path'
import * as dotenv from 'dotenv'
import * as fs from 'fs'
import vuePlugin from '@vitejs/plugin-vue'
import Position from './vitePlugin/Position'
import PositionServer from './vitePlugin/codeServer'
import fullImportPlugin from './vitePlugin/fullImport/fullImport.js'
export default ({
                    command,
                    mode
                }) => {
    const NODE_ENV = process.env.VITE_APP_ENV
    const envFiles = [
        `.env.${NODE_ENV}`
    ]
    for (const file of envFiles) {
        const envConfig = dotenv.parse(fs.readFileSync(file))
        for (const k in envConfig) {
            process.env[k] = envConfig[k]
        }
    }

    viteLogo(process.env)

    const timestamp = Date.parse(new Date())

    const optimizeDeps = {}

    const alias = {
        '@': path.resolve(__dirname, './src'),
        'vue$': 'vue/dist/vue.runtime.esm-bundler.js',
    }

    const esbuild = {}
    const config = {
        base: '/',
        root: './',
        resolve: {
            alias,
        },
        define: {
            'process.env': {}
        },
        server: {
            open: false,
            port: process.env.VITE_CLI_PORT,
            proxy: {
                [process.env.VITE_BASE_API]: {
                    target: `${process.env.VITE_BASE_PATH}:${process.env.VITE_SERVER_PORT}/`,
                    changeOrigin: true,
                    rewrite: path => path.replace(new RegExp('^' + process.env.VITE_BASE_API), ''),
                }
            },
        },
        build: {
            minify: 'terser',
            manifest: false,
            sourcemap: false,
            outDir: 'dist',
            chunkSizeWarningLimit: 1500,
        },
        preview: {
            port: process.env.VITE_CLI_PORT,
        },
        esbuild,
        optimizeDeps,
        plugins: [
            PositionServer(),
            Position(),
            legacyPlugin({
                targets: ['Android > 39', 'Chrome >= 60', 'Safari >= 10.1', 'iOS >= 10.3', 'Firefox >= 54', 'Edge >= 15'],
            }),
            vuePlugin(),
            [Banner(`\n Build based on swa \n Time : ${timestamp}`)]
        ],
        css: {
            preprocessorOptions: {
                scss: {
                    additionalData: `@use "@/styles/element/index.scss" as *;`,
                }
            }
        },
    }

    if (NODE_ENV === 'development') {
        config.plugins.push(
            fullImportPlugin()
        )
    } else {
        config.plugins.push(AutoImport({
                resolvers: [ElementPlusResolver()]
            }),
            Components({
                resolvers: [ElementPlusResolver({
                    importStyle: 'sass'
                })]
            }))
    }
    return config
}
