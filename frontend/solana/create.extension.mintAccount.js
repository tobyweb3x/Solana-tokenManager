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
exports.addAdditionalMetatdataFieldExtension = addAdditionalMetatdataFieldExtension;
const spl_token_1 = require("@solana/spl-token");
const web3_js_1 = require("@solana/web3.js");
const spl_token_metadata_1 = require("@solana/spl-token-metadata");
function createToken22MintAccExtension(serializedData) {
    return __awaiter(this, void 0, void 0, function* () {
        const token22Extensions = [], beforeInitializeMintInstructions = [], afterInitializeMintInstructions = [], rents = [], data = serializedData.data;
        for (const ext of Object.values(serializedData.extensionTypes)) {
            let extensionType = ext;
            includeToken22Extension(extensionType, token22Extensions, rents, beforeInitializeMintInstructions, afterInitializeMintInstructions, data);
        }
        let mintSpace = (0, spl_token_1.getMintLen)(token22Extensions);
        const totalRent = rents.reduce((accumulator, currentValue) => accumulator + currentValue, 0);
        const connection = new web3_js_1.Connection((0, web3_js_1.clusterApiUrl)("devnet"), "confirmed");
        const rentFees = yield connection.getMinimumBalanceForRentExemption(mintSpace + totalRent);
        // allocate the mint's account state on chain
        const createMintAccountInstruction = web3_js_1.SystemProgram.createAccount({
            fromPubkey: data.payer.publicKey,
            lamports: rentFees,
            newAccountPubkey: data.mint.publicKey,
            space: mintSpace,
            programId: spl_token_1.TOKEN_2022_PROGRAM_ID,
        });
        // Initialize the mint account
        const initMintInstruction = (0, spl_token_1.createInitializeMintInstruction)(data.mint.publicKey, data.decimals, data.mintAuthority, data.freezeAuthority, spl_token_1.TOKEN_2022_PROGRAM_ID);
        // include transactions
        const transaction = new web3_js_1.Transaction().add(createMintAccountInstruction, ...beforeInitializeMintInstructions, initMintInstruction, ...afterInitializeMintInstructions);
        //  send, sign & await transaction
        const transactionSignature = yield (0, web3_js_1.sendAndConfirmTransaction)(connection, transaction, [data.payer, data.mint]);
    });
}
;
function includeToken22Extension(extensionType, extensions, rent, beforeInitializeMintInstructions, afterInitializeMintInstructions, data) {
    switch (extensionType) {
        // Token Account Extensions
        case spl_token_1.ExtensionType.DefaultAccountState:
            extensions.push(spl_token_1.ExtensionType.DefaultAccountState);
            beforeInitializeMintInstructions.push((0, spl_token_1.createInitializeDefaultAccountStateInstruction)(data.mint.publicKey, data.accountState, spl_token_1.TOKEN_2022_PROGRAM_ID));
        // Mint Account Extensions
        case spl_token_1.ExtensionType.TransferFeeConfig:
            extensions.push(spl_token_1.ExtensionType.TransferFeeConfig);
            beforeInitializeMintInstructions.push((0, spl_token_1.createInitializeTransferFeeConfigInstruction)(data.mint.publicKey, data.payer.publicKey, data.payer.publicKey, data.feeBasisPoints, data.maxFee, spl_token_1.TOKEN_2022_PROGRAM_ID));
        case spl_token_1.ExtensionType.MintCloseAuthority:
            extensions.push(spl_token_1.ExtensionType.MintCloseAuthority);
            beforeInitializeMintInstructions.push((0, spl_token_1.createInitializeMintCloseAuthorityInstruction)(data.mint.publicKey, data.closeMintAuthority, spl_token_1.TOKEN_2022_PROGRAM_ID));
        case spl_token_1.ExtensionType.NonTransferable:
            extensions.push(spl_token_1.ExtensionType.NonTransferable);
            beforeInitializeMintInstructions.push((0, spl_token_1.createInitializeNonTransferableMintInstruction)(data.mint.publicKey, spl_token_1.TOKEN_2022_PROGRAM_ID));
        case spl_token_1.ExtensionType.InterestBearingConfig:
            extensions.push(spl_token_1.ExtensionType.InterestBearingConfig);
            beforeInitializeMintInstructions.push((0, spl_token_1.createInitializeInterestBearingMintInstruction)(data.mint.publicKey, data.rateAuthority, data.rate, spl_token_1.TOKEN_2022_PROGRAM_ID));
        case spl_token_1.ExtensionType.PermanentDelegate:
            extensions.push(spl_token_1.ExtensionType.PermanentDelegate);
            beforeInitializeMintInstructions.push((0, spl_token_1.createInitializePermanentDelegateInstruction)(data.mint.publicKey, data.permanentDelegate, spl_token_1.TOKEN_2022_PROGRAM_ID));
        case spl_token_1.ExtensionType.TransferHook:
            extensions.push(spl_token_1.ExtensionType.TransferHook);
            beforeInitializeMintInstructions.push((0, spl_token_1.createInitializeTransferHookInstruction)(data.mint.publicKey, data.mintAuthority, data.transferHookProgramId, spl_token_1.TOKEN_2022_PROGRAM_ID));
        case spl_token_1.ExtensionType.MetadataPointer:
            extensions.push(spl_token_1.ExtensionType.MetadataPointer);
            beforeInitializeMintInstructions.push((0, spl_token_1.createInitializeMetadataPointerInstruction)(data.mint.publicKey, data.updateAuthority, data.metadataAddress, spl_token_1.TOKEN_2022_PROGRAM_ID));
        case spl_token_1.ExtensionType.TokenMetadata:
            const tokenMetadata = { mint: data.mint.publicKey, name: data.tokenName, symbol: data.tokenSymbol, uri: data.externalURI, additionalMetadata: data.additionalFields };
            rent.push(spl_token_1.TYPE_SIZE + spl_token_1.LENGTH_SIZE + (0, spl_token_metadata_1.pack)(tokenMetadata).length);
            afterInitializeMintInstructions.push((0, spl_token_1.createInitializeInstruction)({ programId: spl_token_1.TOKEN_2022_PROGRAM_ID, metadata: data.mint.publicKey, updateAuthority: data.payer.publicKey,
                mintAuthority: data.payer.publicKey, mint: data.mint.publicKey, name: data.tokenName, symbol: data.tokenSymbol, uri: data.externalURI }));
            addAdditionalMetatdataFieldExtension(data.updateAuthority, data.metadataAddress, data.additionalFields).forEach((instruction) => afterInitializeMintInstructions.push(instruction));
        case spl_token_1.ExtensionType.GroupPointer:
            extensions.push(spl_token_1.ExtensionType.GroupMemberPointer);
            rent.push(spl_token_1.TOKEN_GROUP_SIZE);
            beforeInitializeMintInstructions.push((0, spl_token_1.createInitializeGroupPointerInstruction)(data.mint.publicKey, data.updateAuthority, data.mint.publicKey, spl_token_1.TOKEN_2022_PROGRAM_ID));
            afterInitializeMintInstructions.push((0, spl_token_1.createInitializeGroupInstruction)({ group: data.mint.publicKey, mint: data.mint.publicKey, mintAuthority: data.mintAuthority, updateAuthority: data.updateAuthority, maxSize: data.groupExtensionMaxsize, programId: spl_token_1.TOKEN_2022_PROGRAM_ID }));
        case spl_token_1.ExtensionType.GroupMemberPointer:
            extensions.push(spl_token_1.ExtensionType.GroupMemberPointer);
            rent.push(spl_token_1.TOKEN_GROUP_MEMBER_SIZE);
            beforeInitializeMintInstructions.push((0, spl_token_1.createInitializeGroupMemberPointerInstruction)(data.mint.publicKey, data.updateAuthority, data.mint.publicKey, spl_token_1.TOKEN_2022_PROGRAM_ID));
            afterInitializeMintInstructions.push((0, spl_token_1.createInitializeMemberInstruction)({ group: data.mint.publicKey, member: data.mint.publicKey, memberMint: data.mint.publicKey, memberMintAuthority: data.mintAuthority, groupUpdateAuthority: data.updateAuthority, programId: spl_token_1.TOKEN_2022_PROGRAM_ID }));
    }
}
function addAdditionalMetatdataFieldExtension(updateAuthority, metadataAddress, additionalFields) {
    let additionalFieldMetadataInstructions = [];
    for (const [key, val] of Object.entries(additionalFields || [])) {
        additionalFieldMetadataInstructions.push((0, spl_token_metadata_1.createUpdateFieldInstruction)({
            updateAuthority: updateAuthority,
            metadata: metadataAddress,
            field: val[0],
            value: val[1],
            programId: spl_token_1.TOKEN_2022_PROGRAM_ID
        }));
    }
    return additionalFieldMetadataInstructions;
}
