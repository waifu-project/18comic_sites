const router = require('koa-router')();
const { getMirror } = require('../utils')

router.get('/', function *(next) {
  yield this.render('index', {
    title: '18comic get mirror server'
  });
});

router.get('/ping', function *(next) {
  this.body = "Ok"
});

router.get('/api/mirror', function *(next) {
  try {
    const body = yield getMirror()
    this.body = body;
  } catch (error) {
    this.body = []
  }
});

module.exports = router;
