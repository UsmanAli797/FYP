var express = require("express");
var fs = require("fs");
var router = express.Router();
var registerUser = require("./callback_functions/registerUser");
var createUser = require("./callback_functions/createUser");
var enrollAdmin = require("./callback_functions/enrollAdmin");
var queryUser = require("./callback_functions/queryUser");
var checkidCard = require("./callback_functions/checkidCard");
var fileComplaint = require("./callback_functions/fileComplaint");
const bcrypt = require("bcrypt");
const { create } = require("ipfs-http-client");
const fileUpload = require("express-fileupload");

const ipfs = create({ host: "localhost", port: "5001", protocol: "http" });

router.get("/enrollAdmin", function (req, res, next) {
    enrollAdmin();
    res.send("Admin created");
});

/* GET Dashboard. */
router.get("/", async function (req, res, next) {
    if (req.session.ID) {
        res.redirect("/citizen/citizen_dashboard");
        console.log("Success");
    } else {
        console.log("failure: ", req.session.ID);
        // enrollAdmin();
        var obj = {
            success: req.session.success,
            error: req.session.error,
            title: "350 Login",
        };
        req.session.success = false;
        req.session.error = false;
        res.render("Citizen_login", obj);
    }
});

/* GET Login. */
router.get("/citizen_login", function (req, res, next) {
    res.render("Citizen_login.ejs");
});

router.post("/login_c", async function (req, res, next) {
    var result = await queryUser(req.body.idCard);
    if (result == "") {
        req.session.error = true;
        res.redirect("/citizen/citizen_login");
    } else {
        var obj = JSON.parse(result);
        NAME = obj.fname;
        IDCARD = obj.idCard;
        // console.log(NAME,ID);

        if (obj.pass == req.body.pass) {
            req.session.USERNAME = NAME;
            req.session.ID = IDCARD;
            res.redirect("/citizen/citizen_dashboard");
        } else {
            req.session.error = true;
            res.redirect("/citizen");
        }
    }
});

router.get("/citizen_dashboard", async function (req, res, next) {
    if (req.session.ID == null) res.redirect("/citizen");
    else {
        var result = await queryUser(req.session.ID);
        var obj = JSON.parse(result);
        res.render("Citizen_Dashboard", {
            title: "Profile",
            success: req.session.success,
            obj: obj,
            ID: req.session.ID,
            USERNAME: req.session.USERNAME,
        });
        req.session.success = false;
    }
});

// File complaint
router.get("/citizen_complaint", async function (req, res) {
    if (req.session.ID == null) {
        res.redirect("/citizen");
    } else {
        var result = await queryUser(req.session.ID);
        var obj = JSON.parse(result);
        res.render("Citizen_filecomplaint", {
            title: "File Complaint",
            passerror: req.session.passerror,
            obj: obj,
            ID: req.session.NID,
            USERNAME: req.session.USERNAME,
        });
        req.session.passerror = false;
    }
});
router.post("/complaint_c", async function (req, res) {
    req.session.success = true;
    await fileComplaint(
        req.body.fname,
        req.body.idCard,
        req.body.email,
        req.body.phone,
        req.body.date,
        req.body.casee,
        req.body.desc
    );
    res.redirect("/citizen/citizen_dashboard");
});

router.get("/Citizen_Upload", async function (req, res) {
    if (req.session.ID == null) res.redirect("/citizen");
    else {
        var result = await queryUser(req.session.ID);
        var obj = JSON.parse(result);
        res.render("Citizen_Upload", {
            title: "Upload Evidence",
            passerror: req.session.passerror,
            obj: obj,
            ID: req.session.NID,
            USERNAME: req.session.USERNAME,
        });
        req.session.passerror = false;
    }
});

router.post("/upload", (req, res) => {
    req.session.success = true;
    const file = req.files.file;
    console.log(file + "file");
    const fileName = req.body.fileName;
    console.log(fileName + "fileName");

    const filePath = "files/" + fileName;
    console.log(filePath + "FilePath");

    file.mv(filePath, async (err) => {
        if (err) {
            console.log("Error: Failed to download the file");
            return res.status(500).send(err);
        }
        // const fileHash = await addFile(fileName,filePath);
        // fs.unlink(filePath, (err)=>{
        //     if (err) console.log(err);
        // });
        res.render("Citizen_view", { fileName, fileHash });
    });
});
// const addFile = async(fileName,filePath)=>{
//     const file = fs.readFileSync(filePath);
//     const hashes = [];
//    const fileAdded = await ipfs.add({path:fileName, content:file});
//    console.log('Your file on IPFS: ', fileAdded);
//    const fileHash = fileAdded.cid;
//    console.log('Your hash on IPFS: ', fileHash);
//    return fileHash;

// }

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

router.post(
    "/register_c",
    async function (req, res, next) {
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
        // bcrypt.hash(req.body.pass,10).async((hash)=>{
        // console.log(hash)
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

        // }).then(()=>{
        res.render("Citizen_login.ejs");
        //         res.status(200).json("User Registered Successfully");

        // }).catch((err)=>{
        // if (err){
        //     res.status(400).json({error:err});
        // }
        // })
        // console.log(password);
        //   await registerUser(req.body.idCard);
        //   await createUser(
        //       req.body.fname,
        //       req.body.email,
        //       req.body.idCard,
        //       req.body.phone,
        //       req.body.city,
        //       req.body.address,
        //       req.body.pass
        //   );
    }
    // }
);

//logout
router.get("/logout", async function (req, res, next) {
    req.session.ID = null;
    res.redirect("/citizen");
});

module.exports = router;
