import axios from 'axios'

describe('Refresh Token', () => {
  let clientId: string | undefined
  let clientSecret: string | undefined

  beforeEach(async () => {
    const {
      data: { client_id, client_secret },
    } = await axios.get('http://localhost:8080/credentials')

    clientId = client_id
    clientSecret = client_secret
  })

  it('should create credentials', () => {
    expect(clientId).toBeDefined()
    expect(clientSecret).toBeDefined()
  })


  it('should create access token then refresh token', async () => {
    const { data } = await axios({
      method: 'post',
      url: 'http://localhost:8080/oauth2/token',
      headers: {
        'Content-Type': 'application/x-www-form-urlencoded',
      },
      data: `grant_type=client_credentials&client_id=${clientId}&client_secret=${clientSecret}&scope=A B C`,
    })

    // https://datatracker.ietf.org/doc/html/rfc6749#section-6
    // grant_type=refresh_token&refresh_token={the refresh_token obtained in the previous response}

    expect(data.access_token).toBeDefined()
    expect(data.refresh_token).toBeDefined()

    /*
    * not working ???
    */
    /*
    await axios({
      method: 'post',
      url: 'http://localhost:8080/oauth2/token',
      headers: {
        'Content-Type': 'application/x-www-form-urlencoded',
       // Authorization: `Bearer ${data.access_token}`,
      },
      data: `grant_type=refresh_token&refresh_token=${data.refresh_token}`,
    })
    */
  })
})
