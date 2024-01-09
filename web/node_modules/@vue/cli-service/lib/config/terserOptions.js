module.exports = options => ({
  terserOptions: {
    compress: {
      arrows: false,
      collapse_vars: false, // 0.3kb
      comparisons: false,
      computed_props: false,
      hoist_funs: false,
      hoist_props: false,
      hoist_vars: false,
      inline: false,
      loops: false,
      negate_iife: false,
      properties: false,
      reduce_funcs: false,
      reduce_vars: false,
      switches: false,
      toplevel: false,
      typeofs: false,

      booleans: true, // 0.7kb
      if_return: true, // 0.4kb
      sequences: true, // 0.7kb
      unused: true, // 2.3kb

      conditionals: true,
      dead_code: true,
      evaluate: true
    },
    mangle: {
      safari10: true
    }
  },
  sourceMap: options.productionSourceMap,
  cache: true,
  parallel: options.parallel,
  extractComments: false
})
