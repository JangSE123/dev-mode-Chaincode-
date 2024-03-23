const express = require('express');
const app = express();
let path = require('path');
const sdk = require('./sdk');

const PORT = 8001;
const HOST = '0.0.0.0';

app.use(express.json());
app.use(express.urlencoded({ extended: true }));

// 사용자 등록
app.post('/registerUser', function (req, res) {
   let userID = req.body.userID;
   let name = req.body.name;
   let buildingName = req.body.buildingName;
   let description = req.body.description;
   let args = [userID, name, buildingName, description];
   sdk.send(false, 'RegisterUser', args, res);
});

// 부동산 조회
app.get('/viewBuilding', function (req, res) {
   let buildingName = req.query.buildingName;
   let args = [buildingName];
   sdk.send(true, 'ViewBuilding', args, res);
});

app.use(express.static(path.join(__dirname, '../client')));
app.listen(PORT, HOST);
console.log(`Running on http://${HOST}:${PORT}`);
