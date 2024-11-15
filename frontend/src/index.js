"use strict";
var __importDefault = (this && this.__importDefault) || function (mod) {
    return (mod && mod.__esModule) ? mod : { "default": mod };
};
Object.defineProperty(exports, "__esModule", { value: true });
exports.createRandomKeyPair = createRandomKeyPair;
const web3_js_1 = require("@solana/web3.js");
const bs58_1 = __importDefault(require("bs58"));
function createRandomKeyPair(tokenInfo) {
    const wallet = web3_js_1.Keypair.generate();
    tokenInfo.mintAddressPublickey = wallet.publicKey.toBase58();
    tokenInfo.mintAddressSecretkey = bs58_1.default.encode(wallet.secretKey);
    return tokenInfo;
}
window.createRandomKeyPair = createRandomKeyPair;
