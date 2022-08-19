import axios from 'axios'

describe('Scopes', () => {
  let clientId: string | undefined;
  let clientSecret: string | undefined;

  beforeEach(async() => {
    const {
      data: { client_id, client_secret },
    } = await axios.get('http://localhost:8080/credentials')

    clientId = client_id;
    clientSecret = client_secret;
  });

  it("should create credentials", () => {
    expect(clientId).toBeDefined()
    expect(clientSecret).toBeDefined()
  })

  it('should create scopes A,B,C', async () => {
    const {
      data: { access_token, expires_in, scope, token_type },
    } = await axios({
      method: 'post',
      url: 'http://localhost:8080/oauth2/token',
      headers: {
        "Content-Type": "application/x-www-form-urlencoded"
      },
      data: `grant_type=client_credentials&client_id=${clientId}&client_secret=${clientSecret}&scope=A B C`,
    })

    expect(access_token).toBeDefined()
    expect(scope).toBe("A B C")

    const {
      data: { message },
    } = await axios.get(`http://localhost:8080/api/protected?access_token=${access_token}`)

    expect(message).toBe("Hello, I'm protected")

    console.log(JSON.stringify((await axios.get(`http://localhost:8080/api/protected?access_token=${access_token}`)).data, null, 4))
  })

})
