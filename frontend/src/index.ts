import {Keypair} from '@solana/web3.js';

export function createRandomKeyPair() {
   const wallet = Keypair.generate();
   console.log('public addr', wallet.publicKey.toBase58())
   console.log('private addr', wallet.secretKey)
}


(window as any).createRandomKeyPair = createRandomKeyPair;
