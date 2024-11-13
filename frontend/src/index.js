"use strict";
Object.defineProperty(exports, "__esModule", { value: true });
exports.createRandomKeyPair = createRandomKeyPair;
const web3_js_1 = require("@solana/web3.js");
function createRandomKeyPair() {
    const wallet = web3_js_1.Keypair.generate();
    console.log('public addr', wallet.publicKey.toBase58());
    console.log('private addr', wallet.secretKey);
}
window.createRandomKeyPair = createRandomKeyPair;
