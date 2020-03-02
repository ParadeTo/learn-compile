const path = require('path')

class AddExeBanner {
  apply(compiler) {
    compiler.hooks.emit.tap('AddExeBanner', compilation => {
      const src = compilation.assets['you.js'].source()
      compilation.assets['you.js'].source = () => {
        return '#!/usr/bin/env node\n' + src
      }
    })
  }
}

module.exports = {
  mode: 'production',
  target: 'node',
  entry: './src/index.ts',
  output: {
    filename: 'you.js',
    path: path.resolve(__dirname, 'dist')
  },

  module: {
    rules: [
      {
        test: /\.ts$/,
        use: 'ts-loader'
      }
    ]
  },
  resolve: {
    extensions: ['.ts', '.js']
  },
  node: {
    fs: 'empty',
    readline: 'empty'
  },
  plugins: [new AddExeBanner()]
}
