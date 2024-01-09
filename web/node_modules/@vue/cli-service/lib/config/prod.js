module.exports = (api, options) => {
  api.chainWebpack(webpackConfig => {
    if (process.env.NODE_ENV === 'production') {
      webpackConfig
        .mode('production')
        .devtool(options.productionSourceMap ? 'source-map' : false)

      webpackConfig
        .plugin('hash-module-ids')
          .use(require('webpack/lib/HashedModuleIdsPlugin'), [{
            hashDigest: 'hex'
          }])

      if (process.env.VUE_CLI_TEST) {
        webpackConfig.optimization.minimize(false)
      }
    }
  })
}
