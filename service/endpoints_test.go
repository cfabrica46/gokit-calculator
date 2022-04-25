package service_test

// func TestSignUpEndpoint(t *testing.T) {
// 	for index, table := range []struct { in       service.SignUpRequest
// 		outToken string
// 		outErr   string
// 		isError  bool
// 	}{
// 		{
// 			in: service.SignUpRequest{
// 				Username: usernameTest,
// 				Password: passwordTest,
// 				Email:    emailTest,
// 			},
// 			outToken: tokenTest,
// 			outErr:   "",
// 			isError:  false,
// 		},
// 		{
// 			in:       service.SignUpRequest{},
// 			outToken: "",
// 			outErr:   errWebServer.Error(),
// 			isError:  true,
// 		},
// 	} {
// 		t.Run(fmt.Sprintf(schemaNameTest, index), func(t *testing.T) {
// 			testResp := struct {
// 				Token string `json:"token"`
// 				Err   string `json:"err"`
// 				ID    int    `json:"id"`
// 			}{
// 				ID:    idTest,
// 				Token: table.outToken,
// 				Err:   table.outErr,
// 			}

// 			jsonData, err := json.Marshal(testResp)
// 			if err != nil {
// 				t.Error(err)
// 			}

// 			mock := service.NewMockClient(func(req *http.Request) (*http.Response, error) {
// 				return &http.Response{
// 					StatusCode: http.StatusOK,
// 					Body:       ioutil.NopCloser(bytes.NewReader(jsonData)),
// 				}, nil
// 			})

// 			svc := service.NewService(
// 				mock,
// 				dbHostTest,
// 				portTest,
// 				tokenHostTest,
// 				portTest,
// 				secretTest,
// 			)

// 			r, err := service.MakeSignUpEndpoint(svc)(context.TODO(), table.in)
// 			if err != nil {
// 				t.Error(err)
// 			}

// 			result, ok := r.(service.SignUpResponse)
// 			if !ok {
// 				t.Error(errNotTypeIndicated)
// 			}

// 			if !table.isError {
// 				assert.Zero(t, result.Err)
// 			} else {
// 				assert.Contains(t, result.Err, table.outErr)
// 			}

// 			assert.Equal(t, table.outToken, result.Token)
// 		})
// 	}
// }
