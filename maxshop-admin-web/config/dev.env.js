'use strict'
const merge = require('webpack-merge')
const prodEnv = require('./prod.env')

module.exports = merge(prodEnv, {
  NODE_ENV: '"development"',
  // BASE_API: '"http://120.27.63.9:8080"'
  // BASE_API:'"http://39.107.30.137:40001/mock/11/"'
  BASE_API : '"http://192.168.0.103:8000/"'
})
