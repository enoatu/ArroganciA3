const isProd = process.env.NODE_ENV == 'production'
module.exports = {
  basePath: isProd ? '/ArroganciA3' : '', // for github-pages subdir
}
