'use strict';

const path = require('path');

const matchRelativePath = /^\.\.?[/\\]/;

function isAbsolutePath(str) {
  return path.posix.isAbsolute(str) || path.win32.isAbsolute(str);
}

function isRelativePath(str) {
  return matchRelativePath.test(str);
}

function stringifyRequest(loaderContext, request) {
  const splitted = request.split('!');
  const context =
    loaderContext.context ||
    (loaderContext.options && loaderContext.options.context);

  return JSON.stringify(
    splitted
      .map((part) => {
        const splittedPart = part.match(/^(.*?)(\?.*)/);
        const query = splittedPart ? splittedPart[2] : '';
        let singlePath = splittedPart ? splittedPart[1] : part;

        if (isAbsolutePath(singlePath) && context) {
          singlePath = path.relative(context, singlePath);

          if (isAbsolutePath(singlePath)) {
            return singlePath + query;
          }

          if (isRelativePath(singlePath) === false) {
            singlePath = './' + singlePath;
          }
        }

        return singlePath.replace(/\\/g, '/') + query;
      })
      .join('!')
  );
}

module.exports = stringifyRequest;
