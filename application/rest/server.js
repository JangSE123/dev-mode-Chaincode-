const express = require('express');
const app = express();
let path = require('path');
let sdk = require('./sdk');

const PORT = 8001;
const HOST = '0.0.0.0';
app.use(express.json());
app.use(express.urlencoded({ extended: true }))

app.get('/init', function (req, res) {
   let a = req.query.a;
   let aval = req.query.aval;
   let b = req.query.b;
   let bval = req.query.bval;
   let c = req.query.c;
   let cval = req.query.cval;
   let args = [a, aval, b, bval, c, cval];
   sdk.send(false, 'Init', args, res);
});

app.get('/query', function (req, res) {
   let name = req.query.name;
   let args = [name];
   sdk.send(true, 'Query', args, res);
});

app.post('/registerUser', function (req, res) {
   let userID = req.body.userID;
   let name = req.body.name;
   let args = [userID, name];
   sdk.send(false, 'RegisterUser', args, res);
});

// 부동산 등록
app.post('/registerProperty', function (req, res) {
   let propertyID = req.body.propertyID;
   let ownerID = req.body.ownerID;
   let description = req.body.description;
   let args = [propertyID, ownerID, description];
   sdk.send(false, 'RegisterProperty', args, res);
});

// 부동산 정보 업데이트
app.post('/updatePropertyDescription', function (req, res) {
   let propertyID = req.body.propertyID;
   let newDescription = req.body.newDescription;
   let args = [propertyID, newDescription];
   sdk.send(false, 'UpdatePropertyDescription', args, res);
});

// 부동산 조회
app.get('/queryProperty', function (req, res) {
   let propertyID = req.query.propertyID;
   let args = [propertyID];
   sdk.send(true, 'QueryProperty', args, res);
});

// 소유권 이전
app.post('/transferPropertyOwnership', function (req, res) {
   let propertyID = req.body.propertyID;
   let newOwnerID = req.body.newOwnerID;
   let args = [propertyID, newOwnerID];
   sdk.send(false, 'TransferPropertyOwnership', args, res);
});

app.use(express.static(path.join(__dirname, '../client')));
app.listen(PORT, HOST);
console.log(`Running on http://${HOST}:${PORT}`);


