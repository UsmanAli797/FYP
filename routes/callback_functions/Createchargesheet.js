/*
 * Copyright IBM Corp. All Rights Reserved.
 *
 * SPDX-License-Identifier: Apache-2.0
 */

'use strict';

const { Gateway, Wallets } = require('fabric-network');
const fs = require('fs');
const path = require('path');


async function main(namee,dateTime,sectionoflaws,investigatingoffcers,accusedperson,briefreport,chargedPerson,plea,sentence,judgement) {
    try {
        // load the network configuration
        const ccpPath = path.resolve(__dirname, '..', '..','..', '..','test-network', 'organizations', 'peerOrganizations', 'org1.example.com', 'connection-org1.json');
        let ccp = JSON.parse(fs.readFileSync(ccpPath, 'utf8'));

        // Create a new file system based wallet for managing identities.
        const walletPath = path.join(process.cwd(), 'wallet');
        const wallet = await Wallets.newFileSystemWallet(walletPath);
        console.log(`Wallet path: ${walletPath}`);

        // Create a new gateway for connecting to our peer node.
        const gateway = new Gateway();
        await gateway.connect(ccp, { wallet, identity: '1350341057523', discovery: { enabled: true, asLocalhost: true } });

        // Get the network (channel) our contract is deployed to.
        const network = await gateway.getNetwork('mychannel');

        // Get the contract from the network.
        const contract = network.getContract('fabcar');
 console.log(namee,dateTime,sectionoflaws,investigatingoffcers,accusedperson,briefreport,chargedPerson,plea,sentence,judgement)
        // Submit the specified transaction.

        await contract.submitTransaction('CreateChargesheet',dateTime,namee,dateTime,sectionoflaws,investigatingoffcers,accusedperson,briefreport,chargedPerson,plea,sentence,judgement);
        console.log('Chargesheet has been submitted');

        // Disconnect from the gateway.
        await gateway.disconnect();

    } catch (error) {
        console.error(`Failed to submit transaction: ${error}`);
        process.exit(1);
    }
}

module.exports = main;
