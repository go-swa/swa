var webpackConfig = require('../../webpack.config.js')

module.exports = function (config) {
  configuration = {
    browsers: [
      'ChromeHeadless'
    ],
    frameworks: [
      'mocha',
      'sinon-chai'
    ],
    reporters: ['spec', 'coverage'],
    files: ['./index.js'],
    preprocessors: {
      './index.js': ['webpack', 'sourcemap']
    },
    webpack: webpackConfig,
    webpackMiddleware: {
      noInfo: true
    },
    coverageReporter: {
      dir: './coverage',
      reporters: [
        { type: 'lcov', subdir: '.' },
        { type: 'text-summary' }
      ]
    },
    customLaunchers: {
      ChromeHeadless: {
        base: 'Chrome',
        flags: [
          '--headless',
          '--disable-gpu',
          '--remote-debugging-port=9222',
        ]
      }
    }
  }

  config.set(configuration)
}

