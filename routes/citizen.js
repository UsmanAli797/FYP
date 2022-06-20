var express = require("express");
var fs = require("fs");
var router = express.Router();
var registerUser = require("./callback_functions/registerUser");
var createUser = require("./callback_functions/createUser");
var enrollAdmin = require("./callback_functions/enrollAdmin");
// var queryUser = require('./callback_functions/queryUser');
var checkidCard = require('./callback_functions/checkidCard');


router.get("/enrollAdmin", function (req, res, next) {
    enrollAdmin();
    res.send("Admin created");
});

// router.get('/', function(req, res, next) {
//     if(req.session.idCard) res.redirect('/citizen/citizen_signup');
//     else{
//     enrollAdmin();
//     var obj = {success: req.session.success, error: req.session.error, title: '350 Login'};
//     //var obj = { title: '350 Login' , Array:['First Element',2,3] , check: req.session.flag , obj1:{t:1} };
//     req.session.success = false;
//     req.session.error = false;
//      res.render('Citizen_login.ejs', obj);
//     }
//   });



/* GET Login. */
router.get("/citizen_login", async function (req, res, next) {
    res.render("Citizen_login.ejs");
});

router.post("/login_c", async function (req, res) {
    res.render("/citizen_dashboard");
});


//logout
router.get("/logout", async function (req, res, next) {
    req.session.idCard = null;
    res.redirect("/citizen_login");
});

/* GET Dashboard. */
router.get("/citizen_dashboard", async function (req, res, next) {
    res.render("Citizen_Dashboard.ejs");
});

/* GET Signup. */
router.get("/citizen_signup", async function (req, res, next) {
//   var obj = {
//       idCardexists: req.session.idCardexists,
//       passerror: req.session.passerror,
//       title: "SignUp",
//   };
//   req.session.idCardexists = false;
//   req.session.passerror = false;
//   req.session.success = false;
  res.render("Citizen_signup.ejs");
});

router.post("/register_c", async function (req, res, next) {
// var passerror = false;
// var idCardexists = false;
// if(req.body.pwd != req.body.pass)  passerror = true;
// console.log(req.body.fname);
// idCardexists = await checkidCard(req.body.idCard);
// if(idCardexists){
//   req.session.idCardexists = idCardexists;
//   req.session.passerror = passerror;
//   req.session.success = false;
// res.render("Citizen_signup.ejs");
// }
// else{
//   req.session.success = true;
  await registerUser(req.body.idCard);
  await createUser(
      req.body.fname,
      req.body.email,
      req.body.idCard,
      req.body.phone,
      req.body.city,
      req.body.address,
      req.body.pass
  );

res.render("Citizen_login.ejs");
}
// }
);


// File complaint
router.get("/citizen_complaint", function (req, res) {
    res.render("Citizen_filecomplaint.ejs", {
        errors: {},
        success: {},
    });
});

module.exports = router;
