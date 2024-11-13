"use strict";
var __awaiter = (this && this.__awaiter) || function (thisArg, _arguments, P, generator) {
    function adopt(value) { return value instanceof P ? value : new P(function (resolve) { resolve(value); }); }
    return new (P || (P = Promise))(function (resolve, reject) {
        function fulfilled(value) { try { step(generator.next(value)); } catch (e) { reject(e); } }
        function rejected(value) { try { step(generator["throw"](value)); } catch (e) { reject(e); } }
        function step(result) { result.done ? resolve(result.value) : adopt(result.value).then(fulfilled, rejected); }
        step((generator = generator.apply(thisArg, _arguments || [])).next());
    });
};
Object.defineProperty(exports, "__esModule", { value: true });
const spl_token_1 = require("@solana/spl-token");
const web3_js_1 = require("@solana/web3.js");
function createToken22TokenAccExtension(serializedData) {
    return __awaiter(this, void 0, void 0, function* () {
        const extensions = [];
        const instructions = [];
        for (const ext of Object.values(serializedData.extensionTypes)) {
            let extensionType = ext;
            let data = serializedData.data;
            switch (extensionType) {
                case spl_token_1.ExtensionType.ImmutableOwner:
                //  included by default on every Token22 ATA
                case spl_token_1.ExtensionType.MemoTransfer:
                    extensions.push(spl_token_1.ExtensionType.MemoTransfer);
                    instructions.push((0, spl_token_1.createEnableRequiredMemoTransfersInstruction)(data.tokenAccount.publicKey, data.updateAuthority, data.multiSigners));
                case spl_token_1.ExtensionType.CpiGuard:
                    extensions.push(spl_token_1.ExtensionType.CpiGuard);
                    instructions.push((0, spl_token_1.createEnableCpiGuardInstruction)(data.tokenAccount.publicKey, data.mintAuthority, data.multiSigners, spl_token_1.TOKEN_2022_PROGRAM_ID));
            }
        }
        const tokenAccountLen = (0, spl_token_1.getAccountLen)(extensions);
        const connection = new web3_js_1.Connection((0, web3_js_1.clusterApiUrl)("devnet"), "confirmed");
        const lamports = yield connection.getMinimumBalanceForRentExemption(tokenAccountLen);
        // construct transaction with these instructions
        const transaction = new web3_js_1.Transaction().add(web3_js_1.SystemProgram.createAccount({
            fromPubkey: serializedData.data.payer.publicKey,
            newAccountPubkey: serializedData.data.tokenAccount.publicKey,
            lamports,
            space: tokenAccountLen,
            programId: spl_token_1.TOKEN_2022_PROGRAM_ID,
        }), (0, spl_token_1.createInitializeAccountInstruction)(serializedData.data.tokenAccount.publicKey, serializedData.data.mint.publicKey, serializedData.data.payer.publicKey, spl_token_1.TOKEN_2022_PROGRAM_ID), ...instructions);
        //  send, sign & await transaction
        const transactionSignature = yield (0, web3_js_1.sendAndConfirmTransaction)(connection, transaction, [serializedData.data.payer, serializedData.data.tokenAccount]);
    });
}
