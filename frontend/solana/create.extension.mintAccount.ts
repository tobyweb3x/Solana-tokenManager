import { ExtensionType, LENGTH_SIZE, TOKEN_2022_PROGRAM_ID, TYPE_SIZE,
        createInitializeMetadataPointerInstruction, createInitializeInstruction,
        createInitializeMintCloseAuthorityInstruction, createInitializeMintInstruction, 
        getMintLen, createInitializeGroupPointerInstruction, TOKEN_GROUP_SIZE,
        createInitializeGroupInstruction,TOKEN_GROUP_MEMBER_SIZE, 
        createInitializeGroupMemberPointerInstruction,
        createInitializeMemberInstruction, createInitializePermanentDelegateInstruction, 
        createInitializeNonTransferableMintInstruction, createInitializeTransferFeeConfigInstruction, 
        createInitializeInterestBearingMintInstruction, createInitializeTransferHookInstruction, 
        createInitializeDefaultAccountStateInstruction, AccountState} 
        from "@solana/spl-token";
import { Connection, Keypair, SystemProgram, Transaction, clusterApiUrl, 
        sendAndConfirmTransaction, PublicKey , Signer, TransactionInstruction} 
        from "@solana/web3.js";
import { pack, TokenMetadata, createUpdateFieldInstruction } from "@solana/spl-token-metadata";


async function createToken22MintAccExtension(serializedData: dataObject) {

    const
    token22Extensions: ExtensionType[] = [], 
    beforeInitializeMintInstructions: TransactionInstruction[] = [], 
    afterInitializeMintInstructions: TransactionInstruction[] = [],
    rents: number[] = [], 
    data = serializedData.data;
    
    for (const ext of Object.values(serializedData.extensionTypes)) {

        let extensionType = ext as ExtensionType;
            includeToken22Extension(extensionType, token22Extensions, rents, beforeInitializeMintInstructions, afterInitializeMintInstructions, data)
    } 
    
    let mintSpace = getMintLen(token22Extensions)
    const totalRent = rents.reduce((accumulator, currentValue) => accumulator + currentValue, 0);

    const connection = new Connection(clusterApiUrl("devnet"), "confirmed")
    const rentFees = await connection.getMinimumBalanceForRentExemption(mintSpace + totalRent)
    
    // allocate the mint's account state on chain
    const createMintAccountInstruction = SystemProgram.createAccount({
        fromPubkey: data.payer.publicKey,
        lamports: rentFees,
        newAccountPubkey: data.mint.publicKey,
        space: mintSpace,
        programId: TOKEN_2022_PROGRAM_ID,
    });

    // Initialize the mint account
    const initMintInstruction = createInitializeMintInstruction(
        data.mint.publicKey,
        data.decimals,
        data.mintAuthority,
        data.freezeAuthority,
        TOKEN_2022_PROGRAM_ID,
    );

    // include transactions
 const transaction = new Transaction().add(
    createMintAccountInstruction,
    ...beforeInitializeMintInstructions,
    initMintInstruction, 
    ...afterInitializeMintInstructions
)

    //  send, sign & await transaction
  const transactionSignature = await sendAndConfirmTransaction(
    connection,
    transaction,
    [data.payer, data.mint],
  );
  
};


function includeToken22Extension(extensionType: ExtensionType, extensions: ExtensionType[], rent: number[], beforeInitializeMintInstructions: TransactionInstruction[], afterInitializeMintInstructions: TransactionInstruction[], data: token22Params) {

       switch (extensionType) {

        // Token Account Extensions
        case ExtensionType.DefaultAccountState: 
            extensions.push(ExtensionType.DefaultAccountState)
            beforeInitializeMintInstructions.push(createInitializeDefaultAccountStateInstruction(data.mint.publicKey, data.accountState, TOKEN_2022_PROGRAM_ID))

        // Mint Account Extensions
        case ExtensionType.TransferFeeConfig:
            extensions.push(ExtensionType.TransferFeeConfig)
            beforeInitializeMintInstructions.push(createInitializeTransferFeeConfigInstruction(data.mint.publicKey, data.payer.publicKey, data.payer.publicKey, data.feeBasisPoints, data.maxFee, TOKEN_2022_PROGRAM_ID))

        case ExtensionType.MintCloseAuthority:
            extensions.push(ExtensionType.MintCloseAuthority)
            beforeInitializeMintInstructions.push(createInitializeMintCloseAuthorityInstruction(data.mint.publicKey, data.closeMintAuthority, TOKEN_2022_PROGRAM_ID))
            
        case ExtensionType.NonTransferable:
            extensions.push(ExtensionType.NonTransferable)
            beforeInitializeMintInstructions.push(createInitializeNonTransferableMintInstruction(data.mint.publicKey, TOKEN_2022_PROGRAM_ID))

        case ExtensionType.InterestBearingConfig:
            extensions.push(ExtensionType.InterestBearingConfig)
            beforeInitializeMintInstructions.push(createInitializeInterestBearingMintInstruction(data.mint.publicKey,data.rateAuthority, data.rate, TOKEN_2022_PROGRAM_ID))
                    
        case ExtensionType.PermanentDelegate:
            extensions.push(ExtensionType.PermanentDelegate)
            beforeInitializeMintInstructions.push(createInitializePermanentDelegateInstruction(data.mint.publicKey, data.permanentDelegate, TOKEN_2022_PROGRAM_ID))
        
        case ExtensionType.TransferHook:
            extensions.push(ExtensionType.TransferHook)
            beforeInitializeMintInstructions.push(createInitializeTransferHookInstruction(data.mint.publicKey, data.mintAuthority, data.transferHookProgramId, TOKEN_2022_PROGRAM_ID ))
                
        case ExtensionType.MetadataPointer:
            extensions.push(ExtensionType.MetadataPointer)
            beforeInitializeMintInstructions.push(createInitializeMetadataPointerInstruction(data.mint.publicKey, data.updateAuthority, data.metadataAddress, TOKEN_2022_PROGRAM_ID))
            
        case ExtensionType.TokenMetadata:
            const tokenMetadata: TokenMetadata = {mint: data.mint.publicKey, name: data.tokenName, symbol: data.tokenSymbol, uri: data.externalURI, additionalMetadata: data.additionalFields};
            rent.push(TYPE_SIZE + LENGTH_SIZE + pack(tokenMetadata).length)
            afterInitializeMintInstructions.push(createInitializeInstruction({programId:TOKEN_2022_PROGRAM_ID, metadata: data.mint.publicKey, updateAuthority: data.payer.publicKey, 
                                            mintAuthority: data.payer.publicKey, mint: data.mint.publicKey, name: data.tokenName, symbol: data.tokenSymbol, uri: data.externalURI}))
            addAdditionalMetatdataFieldExtension(data.updateAuthority, data.metadataAddress, data.additionalFields).forEach((instruction) => afterInitializeMintInstructions.push(instruction))

        case ExtensionType.GroupPointer:
            extensions.push(ExtensionType.GroupMemberPointer)
            rent.push(TOKEN_GROUP_SIZE)
            beforeInitializeMintInstructions.push(createInitializeGroupPointerInstruction(data.mint.publicKey, data.updateAuthority, data.mint.publicKey,TOKEN_2022_PROGRAM_ID ))
            afterInitializeMintInstructions.push(createInitializeGroupInstruction({group: data.mint.publicKey, mint:data.mint.publicKey, mintAuthority:data.mintAuthority, updateAuthority:data.updateAuthority, maxSize:data.groupExtensionMaxsize, programId:TOKEN_2022_PROGRAM_ID }))

        case ExtensionType.GroupMemberPointer:
            extensions.push(ExtensionType.GroupMemberPointer)
            rent.push(TOKEN_GROUP_MEMBER_SIZE)
            beforeInitializeMintInstructions.push(createInitializeGroupMemberPointerInstruction(data.mint.publicKey, data.updateAuthority, data.mint.publicKey, TOKEN_2022_PROGRAM_ID))
            afterInitializeMintInstructions.push(createInitializeMemberInstruction({group: data.mint.publicKey, member:data.mint.publicKey,memberMint:data.mint.publicKey, memberMintAuthority:data.mintAuthority, groupUpdateAuthority:data.updateAuthority, programId:TOKEN_2022_PROGRAM_ID}))
        
        }
}


export type dataObject = {
    extensionTypes: number[],
    data: token22Params
}

export interface token22Params {
    tokenName: string,
    tokenSymbol: string,
    externalURI: string,
    decimals: number,
    rate: number,
    supply: number,
    feeBasisPoints: number,
    accountState: AccountState,
    maxFee: bigint,
    groupExtensionMaxsize: number,
    payer: Keypair,
    mint: Keypair,
    tokenAccount: Keypair,
    permanentDelegate: PublicKey,
    metadataAddress: PublicKey,
    updateAuthority: PublicKey,
    rateAuthority: PublicKey,
    mintAuthority: PublicKey,
    freezeAuthority: PublicKey | null,
    closeMintAuthority: PublicKey
    transferHookProgramId: PublicKey,
    multiSigners: (Signer | PublicKey)[],
    additionalFields:[string, string][],
}

export function addAdditionalMetatdataFieldExtension(updateAuthority: PublicKey, metadataAddress: PublicKey, additionalFields: [string, string][]) {
    let additionalFieldMetadataInstructions: TransactionInstruction[] = [];
    for (const [key, val] of Object.entries(additionalFields || [])) {
        additionalFieldMetadataInstructions.push(
            createUpdateFieldInstruction( {
                updateAuthority: updateAuthority,
                metadata: metadataAddress,
                field: val[0],
                value: val[1],
                programId: TOKEN_2022_PROGRAM_ID
            }
            ),
        );
    }
    
    return additionalFieldMetadataInstructions;
}
