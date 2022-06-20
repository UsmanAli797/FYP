var express = require('express');
var router = express.Router();


router.get('/court_login', function(req, res, next) {
    // if(req.session.NID) 
   // res.redirect('Citizen_login.ejs');
  //  else{
    // enrollAdmin();
    // var obj = {success: req.session.success, error: req.session.error, title: '350 Login'};
    // req.session.success = false;
    // req.session.error = false;
     res.render('court_login.ejs');
   // }
  });



module.exports = router;