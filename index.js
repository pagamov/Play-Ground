import { VK } from 'vk-io';

import { DirectAuthorization, officialAppCredentials } from '@vk-io/authorization';

const direct = new DirectAuthorization({
    callbackService,

    scope: 'all',

    // Direct authorization is only available for official applications
    ...officialAppCredentials.android, // { clientId: string; clientSecret: string; }

    // Or manually provide app credentials
    // clientId: process.env.CLIENT_ID,
    // clientSecret: process.env.CLIENT_SECRET,

    login: process.env.LOGIN,
    password: process.env.PASSWORD,
    
    apiVersion: '5.199'
});

const vk = new VK({
    token: process.env.TOKEN
});

async function run() {
    const response = await vk.api.wall.get({
        owner_id: 1
    });

    console.log(response);
}

run().catch(console.log);