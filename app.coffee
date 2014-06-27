request = require 'request'
_ =  require 'lodash'

['CLOUDFLARE_EMAIL', 'CLOUDFLARE_TOKEN', 'CLOUDFLARE_SUBDOMAIN', 'CLOUDFLARE_ROOT_DOMAIN'].map (param) ->
  throw new Error "Missing required param: #{param}" unless process.env[param]

fulldomain = process.env.CLOUDFLARE_SUBDOMAIN + '.' + process.env.CLOUDFLARE_ROOT_DOMAIN

request.get 'http://jsonip.com', (err, res, body) ->
  throw err if err?
  ip = JSON.parse(body).ip
  request.post
    url: 'https://www.cloudflare.com/api_json.html'
    form:
      a: 'rec_load_all'
      tkn: process.env.CLOUDFLARE_TOKEN
      z: process.env.CLOUDFLARE_ROOT_DOMAIN
      email: process.env.CLOUDFLARE_EMAIL
  , (err, res, body) ->
    throw err if err?
    record = _.find JSON.parse(body).response.recs.objs, (item) ->
      return true if item.name is fulldomain
    request.post
      url: 'https://www.cloudflare.com/api_json.html'
      form:
        a: 'rec_edit'
        tkn: process.env.CLOUDFLARE_TOKEN
        id: record.rec_id
        email: process.env.CLOUDFLARE_EMAIL
        z: process.env.CLOUDFLARE_ROOT_DOMAIN
        type: 'A'
        content: ip
        service_mode: 0
        ttl: 120
        name: process.env.CLOUDFLARE_SUBDOMAIN
    , (err, res, body) ->
      throw err if err?
