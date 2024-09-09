package main



type ExtensionType int

const (
    Uninitialized ExtensionType = iota
    TransferFeeConfig
    TransferFeeAmount
    MintCloseAuthority
    ConfidentialTransferMint
    ConfidentialTransferAccount
    DefaultAccountState
    ImmutableOwner
    MemoTransfer
    NonTransferable
    InterestBearingConfig
    CpiGuard
    PermanentDelegate
    NonTransferableAccount
    TransferHook
    TransferHookAccount
    ConfidentialTransferFee // Not implemented yet
    ConfidentialTransferFeeAmount // Not implemented yet
    MetadataPointer 
    TokenMetadata  
    GroupPointer 
    TokenGroup 
    GroupMemberPointer 
    TokenGroupMember 
)
