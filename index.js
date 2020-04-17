const Web3 = require("web3")

const web3 = new Web3()

const account = web3.eth.accounts.privateKeyToAccount("29a1ad7fe945e254474891bb530114b8706cdf89844bce9445570743182a69b1")

console.log(account.sign("abc"))