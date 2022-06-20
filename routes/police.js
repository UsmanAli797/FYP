var express = require('express');
var router = express.Router();

router.get('/police_login', function(req, res, next) {

     res.render('police_login.ejs');
  });



  router.get('/police_dashboard', function(req, res, next) {

     res.render('Police_Dashboard.ejs');
  });





  module.exports = router;