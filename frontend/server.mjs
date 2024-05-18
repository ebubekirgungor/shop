import fs from 'fs'

process.env.NITRO_SSL_CERT = fs.readFileSync('/etc/letsencrypt/live/go-nuxt.shop/fullchain.pem')
process.env.NITRO_SSL_KEY = fs.readFileSync('/etc/letsencrypt/live/go-nuxt.shop/privkey.pem')

await import('./.output/server/index.mjs')