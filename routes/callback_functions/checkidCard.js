/*
 * Copyright IBM Corp. All Rights Reserved.
 *
 * SPDX-License-Identifier: Apache-2.0
 */

'use strict';

const { Wallets } = require('fabric-network');
const fs = require('fs');
const path = require('path');

async function main(idCard) {
    try {
        // load the network configuration
        const ccpPath = path.resolve(__dirname, '..', '..','..', '..','test-network', 'organizations', 'peerOrganizations', 'org1.example.com', 'connection-org1.json');
        const ccp = JSON.parse(fs.readFileSync(ccpPath, 'utf8'));


        // Create a new file system based wallet for managing identities.
        const walletPath = path.join(process.cwd(), 'wallet');
        const wallet = await Wallets.newFileSystemWallet(walletPath);
        console.log(`Wallet path: ${walletPath}`);

        // Check to see if we've already enrolled the user.
        const identity = await wallet.get(idCard.toString());
        if (!identity) {
            console.log('An identity for the ser '+idCard+' does not exist in the wallet');
            console.log('Run the registerUser.js application before retrying');
            return false;
        }
        return true;

    } catch (error) {
        console.error(`Failed to submit transaction${idCard}: ${error}`);
        process.exit(1);
    }
}






module.exports = main;
