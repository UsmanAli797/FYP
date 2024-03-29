const express = require('express');
const path = require('path');
var createError = require('http-errors');
const fileUpload = require('express-fileupload');


var cookieParser = require('cookie-parser');
var sessions = require('express-session');
var bodyParser     =   require("body-parser");

var indexRouter = require('./routes/citizen');
var forensicRouter = require('./routes/forensic');
var policeRouter = require('./routes/police');
var courtRouter = require('./routes/court');





const app = express();
app.use(fileUpload());
app.use(bodyParser.urlencoded({ extended: true }));
app.use(bodyParser.json());
app.use(express.urlencoded({ extended: false }));
app.use(express.static(path.join(__dirname, 'public')));
app.use(cookieParser());
const oneDay = 1000 * 60 * 60 * 24;
app.use(sessions({
    secret: "thisismysecrctekeyfhrgfgrfrty84fwir767",
    saveUninitialized:true,
    cookie: { maxAge: oneDay },
    resave: false 
}));


app.use('/citizen', indexRouter);
app.use('/police', policeRouter);
app.use('/forensic', forensicRouter);
app.use('/court', courtRouter);
app.use('/citizen', express.static('citizen/citizen'));
app.use('/police', express.static('police/police'));
app.use('/forensic', express.static('forensic/forensic'));
app.use('/court', express.static('court/court'));




app.set('view engine', 'ejs');
app.set('views',path.join(__dirname,'views'));

app.use('/public',express.static('public'));
app.get("/",(req,res) =>{
    res.render('index',{});
})

app.listen(3000, ()=>{
    console.log('App is running on port 3000');
})






module.exports = app;

