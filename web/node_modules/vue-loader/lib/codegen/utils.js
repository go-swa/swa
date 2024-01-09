const qs = require('querystring')

const ignoreList = ['id', 'index', 'src', 'type']

exports.attrsToQuery = (attrs, langFallback) => {
  let query = ``
  for (const name in attrs) {
    const value = attrs[name]
    if (!ignoreList.includes(name)) {
      query += `&${qs.escape(name)}=${value ? qs.escape(value) : ``}`
    }
  }
  if (langFallback && !(`lang` in attrs)) {
    query += `&lang=${langFallback}`
  }
  return query
}
