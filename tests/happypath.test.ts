import axios from 'axios'

describe('Happy path', () => {
  it('should create token and view protected routes', async () => {
    const {
      data: { client_id, client_secret },
    } = await axios.get('http://localhost:8080/credentials')

    expect(client_id).toBeDefined()
    expect(client_secret).toBeDefined()

    const {
      data: { access_token, expires_in, scope, token_type },
    } = await axios({
      method: 'post',
      url: 'http://localhost:8080/oauth2/token',
      headers: {},
      data: `grant_type=client_credentials&client_id=${client_id}&client_secret=${client_secret}&scope=read`,
    })

    expect(access_token).toBeDefined()

    const {
      data: { message },
    } = await axios.get(`http://localhost:8080/api/protected?access_token=${access_token}`)

    expect(message).toBe("Hello, I'm protected")
  })
})
