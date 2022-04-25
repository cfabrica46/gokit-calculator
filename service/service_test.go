package service_test

const (
// idTest       int    = 1
// usernameTest string = "username"
// passwordTest string = "password"
// emailTest    string = "email@email.com"
// secretTest   string = "secret"

// urlTest       string = "localhost:8080"
// dbHostTest    string = "db"
// tokenHostTest string = "token"
// portTest      string = "8080"
// tokenTest     string = "token"

// schemaNameTest string = "%v"
)

var (
// errWebServer        = errors.New("error from web server")
// errNotTypeIndicated = errors.New("response is not of the type indicated")
)

// func TestSignUp(t *testing.T) {
// 	for index, table := range []struct {
// 		inUsername, inPassword, inEmail string
// 		url                             string
// 		method                          string
// 		isError                         bool
// 	}{
// 		{
// 			inUsername: usernameTest,
// 			inPassword: passwordTest,
// 			inEmail:    emailTest,
// 			isError:    false,
// 			url:        "http://token:8080/generate",
// 			method:     http.MethodPost,
// 		},
// 		{
// 			inUsername: usernameTest,
// 			inPassword: passwordTest,
// 			inEmail:    emailTest,
// 			isError:    true,
// 			url:        "http://db:8080/user",
// 			method:     http.MethodPost,
// 		},
// 		{
// 			inUsername: usernameTest,
// 			inPassword: passwordTest,
// 			inEmail:    emailTest,
// 			isError:    true,
// 			url:        "http://db:8080/id/username",
// 			method:     http.MethodGet,
// 		},
// 		{
// 			inUsername: usernameTest,
// 			inPassword: passwordTest,
// 			inEmail:    emailTest,
// 			isError:    true,
// 			url:        "http://token:8080/generate",
// 			method:     http.MethodPost,
// 		},
// 		{
// 			inUsername: usernameTest,
// 			inPassword: passwordTest,
// 			inEmail:    emailTest,
// 			isError:    true,
// 			url:        "http://token:8080/token",
// 			method:     http.MethodPost,
// 		},
// 	} {
// 		t.Run(fmt.Sprintf(schemaNameTest, index), func(t *testing.T) {
// 			var resultToken string
// 			var resultErr error
// 			var tokenResponse, errorResponse string

// 			if table.isError {
// 				errorResponse = errWebServer.Error()
// 			} else {
// 				tokenResponse = tokenTest
// 			}

// 			testResp := struct {
// 				Token string `json:"token"`
// 				Err   string `json:"err"`
// 				ID    int    `json:"id"`
// 			}{
// 				ID:    idTest,
// 				Token: tokenResponse,
// 				Err:   errorResponse,
// 			}

// 			jsonData, err := json.Marshal(testResp)
// 			if err != nil {
// 				t.Error(err)
// 			}

// 			mock := service.NewMockClient(func(req *http.Request) (*http.Response, error) {
// 				response := &http.Response{Body: io.NopCloser(bytes.NewReader([]byte("{}")))}

// 				if req.URL.String() == table.url {
// 					if req.Method == table.method {
// 						response = &http.Response{
// 							StatusCode: http.StatusOK,
// 							Body:       io.NopCloser(bytes.NewReader(jsonData)),
// 						}
// 					}
// 				}

// 				return response, nil
// 			})

// 			svc := service.NewService(
// 				mock,
// 				dbHostTest,
// 				portTest,
// 				tokenHostTest,
// 				portTest,
// 				secretTest,
// 			)

// 			resultToken, resultErr = svc.SignUp(table.inUsername, table.inPassword, table.inEmail)

// 			if !table.isError {
// 				assert.Nil(t, resultErr)
// 			} else {
// 				assert.ErrorContains(t, resultErr, errorResponse)
// 			}
// 			assert.Equal(t, tokenResponse, resultToken)
// 		})
// 	}
// }
