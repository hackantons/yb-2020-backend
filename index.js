
var Contract = require('web3-eth-contract');
var cors = require('cors')
Contract.setProvider('http://18.192.192.22:8545');

// Data set up
let abi = [ { "inputs": [], "name": "getScoreJson", "outputs": [ { "internalType": "string", "name": "", "type": "string" } ], "stateMutability": "view", "type": "function" }, { "inputs": [ { "internalType": "string", "name": "newScoreJson", "type": "string" } ], "name": "setScoreJson", "outputs": [], "stateMutability": "nonpayable", "type": "function" } ]

var contract = new Contract(abi, '0xFd93f18DE7763216830Cfe5fCD1A40304fDE4F09');


var express=require('express');
var app = express();

app.use(cors())

app.use(express.json());
app.use(express.urlencoded());

app.get('/leaderboard',async function(req,res)
{
  const scoreJson = await getScoreJson();
  
  res.send(scoreJson);
});

app.post('/leaderboard', async function (req, res) {
  const scoreJson = await getScoreJson();
  const score = JSON.parse(scoreJson)
  score.push(req.body)
  setScoreJson(JSON.stringify(score))
  res.send(score)
})

const getScoreJson = async () => {
  return await contract.methods.getScoreJson().call();
}

const setScoreJson = (scoreJson) => {
  contract.methods.setScoreJson(scoreJson).send({from: '0xadf7e0cd4e33f920a33bb4261e84589b572ccd64'})
}

var server=app.listen(5080,function() {});