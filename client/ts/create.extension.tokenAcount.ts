import { ExtensionType, TOKEN_2022_PROGRAM_ID, getAccountLen,
        createEnableRequiredMemoTransfersInstruction,createEnableCpiGuardInstruction,
        createInitializeAccountInstruction} 
        from "@solana/spl-token";
import { Connection, SystemProgram, Transaction, clusterApiUrl, 
        sendAndConfirmTransaction, TransactionInstruction} 
        from "@solana/web3.js";
import { dataObject } from "./create.extension.mintAccount";

async function createToken22TokenAccExtension(serializedData:dataObject) {
   
    const extensions : ExtensionType[] = [];
    const instructions: TransactionInstruction[] = [];

    for (const ext of Object.values(serializedData.extensionTypes)) {
        let extensionType = ext as ExtensionType;
        let data = serializedData.data;

        switch (extensionType) {

        case ExtensionType.ImmutableOwner:
            //  included by default on every Token22 ATA

        case ExtensionType.MemoTransfer:
            extensions.push(ExtensionType.MemoTransfer)
            instructions.push(createEnableRequiredMemoTransfersInstruction(data.tokenAccount.publicKey, data.updateAuthority, data.multiSigners))
            
        case ExtensionType.CpiGuard: 
            extensions.push(ExtensionType.CpiGuard)
            instructions.push(createEnableCpiGuardInstruction(data.tokenAccount.publicKey, data.mintAuthority,data.multiSigners, TOKEN_2022_PROGRAM_ID))
        }

    }

    const tokenAccountLen = getAccountLen(extensions);
    const connection = new Connection(clusterApiUrl("devnet"), "confirmed");
    const lamports = await connection.getMinimumBalanceForRentExemption(tokenAccountLen);


    // construct transaction with these instructions
    const transaction = new Transaction().add(
    SystemProgram.createAccount({
        fromPubkey: serializedData.data.payer.publicKey,
        newAccountPubkey: serializedData.data.tokenAccount.publicKey,
        lamports,
        space: tokenAccountLen,
        programId: TOKEN_2022_PROGRAM_ID,
    }),
    createInitializeAccountInstruction(
        serializedData.data.tokenAccount.publicKey,
        serializedData.data.mint.publicKey,
        serializedData.data.payer.publicKey,
        TOKEN_2022_PROGRAM_ID
    ),
    ...instructions,
    );

    //  send, sign & await transaction
    const transactionSignature = await sendAndConfirmTransaction(
        connection,
        transaction,
        [serializedData.data.payer, serializedData.data.tokenAccount], 
        );

}