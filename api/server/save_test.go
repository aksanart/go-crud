package server

import (
	"aksan-go-crud/internal/entity"
	"aksan-go-crud/internal/service/mocks"
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func Test_save(t *testing.T) {
	testCases := []struct {
		httpCode int
		nameFunc string
		request  interface{}
		expected string
		doMock   func(mock *mocks.MockServiceIntf)
	}{
		{
			httpCode: 200,
			nameFunc: "positive test",
			request:  `{"mdate": "12 2 223", "stadium": "ssss", "team1": "rma", "team2": "brc"}`,
			expected: `{"message":"success","status":"ok"}`,
			doMock: func(mock *mocks.MockServiceIntf) {
				mock.EXPECT().SaveService(gomock.Any()).Return(nil).AnyTimes()
			},
		},
		{
			httpCode: 400,
			nameFunc: "negative test 1",
			request:  `{"mdate": "", "stadium": "ssss", "team1": "rma", "team2": "brc"}`,
			expected: `{"message":"Key: 'SaveRequest.Mdate' Error:Field validation for 'Mdate' failed on the 'required' tag", "status":"FAILED"}`,
			doMock: func(mock *mocks.MockServiceIntf) {
				mock.EXPECT().SaveService(gomock.Any()).Return(nil).AnyTimes()
			},
		},
	}
	for _, v := range testCases {
		t.Run(v.nameFunc, func(t *testing.T) {
			gin.SetMode(gin.TestMode)
			mockCtrl := gomock.NewController(t)
			defer mockCtrl.Finish()
			mockInfra := mocks.NewMockServiceIntf(mockCtrl)
			v.doMock(mockInfra)
			var user entity.SaveRequest
			json.Unmarshal([]byte(v.request.(string)), &user)
			data, err := json.Marshal(user)
			if err != nil {
				t.Fatal(err)
			}
			req, err := http.NewRequest("POST", "/save", bytes.NewBuffer(data))
			if err != nil {
				t.Fatal(err)
			}
			req.Header.Set("Content-Type", "application/json")
			rr := httptest.NewRecorder()
			r := gin.Default()
			r.Use(save(mockInfra))
			r.ServeHTTP(rr, req)
			assert.Equal(t, v.httpCode, rr.Code)
			assert.JSONEq(t, v.expected, rr.Body.String())
		})
	}
}
