import {Keypair, PublicKey} from '@solana/web3.js';
import bs58 from 'bs58';

export function createRandomKeyPair(tokenInfo: TokenInfo) {
   const wallet = Keypair.generate();
   tokenInfo.mintAddressPublickey = wallet.publicKey.toBase58();
   tokenInfo.mintAddressSecretkey = bs58.encode(wallet.secretKey)
   return tokenInfo
}

interface TokenInfo {
   tokenStandard: string
   tokenType: string
   tokenName: string
   tokenSymbol: string
   tokenUrl: string
   mintAddressPublickey: string
   mintAddressSecretkey: string

}

(window as any).createRandomKeyPair = createRandomKeyPair;
