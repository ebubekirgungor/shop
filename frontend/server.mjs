import fs from 'fs'

process.env.NITRO_SSL_CERT = fs.readFileSync('go-nuxt.shop/fullchain1.pem')
process.env.NITRO_SSL_KEY = fs.readFileSync('go-nuxt.shop/privkey1.pem')

await import('./.output/server/index.mjs')