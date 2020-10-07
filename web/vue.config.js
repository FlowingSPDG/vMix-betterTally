module.exports = {
    publicPath: './',
    devServer: {
      port: 8081,
      disableHostCheck: true,
      proxy: {
        '^/api': {
          target: 'http://localhost:5000',
          ws: true,
          changeOrigin: false
        }
      }
    }
  }
  