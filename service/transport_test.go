package service_test

// func TestDecodeSignUpRequest(t *testing.T) {
// 	url := urlTest

// 	dataJSON, err := json.Marshal(struct {
// 		Username string `json:"username"`
// 		Password string `json:"password"`
// 		Email    string `json:"email"`
// 	}{
// 		usernameTest,
// 		passwordTest,
// 		emailTest,
// 	})
// 	if err != nil {
// 		t.Error(err)
// 	}

// 	goodReq, err := http.NewRequest(http.MethodGet, url, bytes.NewBuffer(dataJSON))
// 	if err != nil {
// 		t.Error(err)
// 	}

// 	badReq, err := http.NewRequest(http.MethodGet, url, bytes.NewBuffer([]byte{}))
// 	if err != nil {
// 		t.Error(err)
// 	}

// 	for index, table := range []struct {
// 		in     *http.Request
// 		out    service.SignUpRequest
// 		outErr string
// 	}{
// 		{goodReq, service.SignUpRequest{usernameTest, passwordTest, emailTest}, ""},
// 		{badReq, service.SignUpRequest{}, "EOF"},
// 	} {
// 		t.Run(fmt.Sprintf(schemaNameTest, index), func(t *testing.T) {
// 			var result interface{}
// 			var resultErr string

// 			r, err := service.DecodeSignUpRequest(context.TODO(), table.in)
// 			if err != nil {
// 				resultErr = err.Error()
// 			}

// 			result, ok := r.(service.SignUpRequest)
// 			if !ok {
// 				if (table.out != service.SignUpRequest{}) {
// 					t.Error("result is not of the type indicated")
// 				}
// 			}

// 			assert.Equal(t, table.outErr, resultErr)
// 			assert.Equal(t, table.out, result)
// 		})
// 	}
// }
