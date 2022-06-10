const path = require("path")
module.exports = {
  publicPath: "./",
  outputDir: "dist",
  assetsDir: "static",
  lintOnSave: process.env.NODE_ENV === "development",
  productionSourceMap: false,
  devServer: {
    hot: true,
    open: true,
    overlay: {
      warnings: false,
      errors: false
    },
    proxy: {
      "/api": {
        target: "http://localhost:8282",   //请求地址
        changeOrigin: true, //是否跨域
        pathRewrite: {
          "^/api": "" //使用'/api'代替target里面的地址
        }
      }
    }
  },
  configureWebpack: {
    resolve: {
      alias: {
        '@': path.resolve(__dirname, 'src'),
        'images': path.resolve(__dirname, 'src/assets/images')
      }
    }
  }
}