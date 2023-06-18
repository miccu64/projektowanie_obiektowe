const { createProxyMiddleware } = require('http-proxy-middleware');

module.exports = function(app) {
  app.use(
    '/api',
    createProxyMiddleware({
      target: 'http://zadanie10-back.mangoriver-0ceea17e.westeurope.azurecontainerapps.io',
      changeOrigin: true,
      secure: false,
    })
  );
};