package main

import (
	"testing"

	"github.com/golang/mock/gomock"
)

func TestFoo(t *testing.T) {
	ctrl := gomock.NewController(t)

	// Assert that Bar() is invoked.
	defer ctrl.Finish()

	m := NewMockFoo(ctrl)

	// Asserts that the first and only call to Bar() is passed 99.
	// Anything else will fail.
	m.
		EXPECT().
		Bar(gomock.Eq(99)).
		Return(101)

	// SUT is passed m, which is of type Foo, since it implements the Bar method.
	//
	// This is indeed the essense of testing (and mocking) with Go:
	//
	//   > That which you wish to test or mock should be replaced by value that
	//     is of its type.
	SUT(m)
}

func TestBaz(t *testing.T) {
	ctrl := gomock.NewController(t)

	// Assert that Bar() is invoked.
	defer ctrl.Finish()

	m := NewMockFoo(ctrl)

	// Asserts that the first and only call to Bar() is passed 99 and returns
	// 100.
	// Anything else will fail.
	m.
		EXPECT().
		Bar(gomock.Eq(99)).
		Return(100)

	// Here, m is of type MockFoo. Therefore, we replace our struct field Baz of
	// type Foo with m.
	//
	// When Bar is called, is is calling the Bar function on the mock struct, not
	// the actual struct.
	aBaz.Baz = m
	MoreSUT()
}

// func TestAPI(t *testing.T) {
// 	ctrl := gomock.NewController(t)
//
// 	// Assert that Get() is invoked.
// 	defer ctrl.Finish()
//
// 	m := NewMockAPI(ctrl)
//
// 	// Asserts that the first and only call to Bar() is passed 99 and returns
// 	// 100.
// 	// Anything else will fail.
// 	m.
// 		EXPECT().
// 		Get(gomock.Eq(99)).
// 		Return(nil)
//
// 	apiClient = m
// 	MakeAPICall()
// }

// func TestAPI(t *testing.T) {
// 	t.Parallel()
// 	cases := []struct {
// 		name             string
// 		Code             int
// 		Data             []map[string]interface{}
// 		Message          string
// 		wantReleaseCount int
// 		wantErr          error
// 	}{
// 		{
// 			name:    "bad status code",
// 			Code:    http.StatusInternalServerError,
// 			wantErr: ErrFailedAPICall,
// 		},
// 		{
// 			name: "success",
// 			Code: http.StatusOK,
// 			Data: []map[string]interface{}{
// 				"Hello": "World",
// 			},
// 			// TODO
// 			wantReleaseCount: 1,
// 		},
// 	}
// 	for _, tc := range cases {
// 		tc := tc
// 		t.Run(tc.name, func(t *testing.T) {
// 			t.Parallel()
// 			ctrl := gomock.NewController(t)
// 			mockApi := NewMockAPI(ctrl)
// 			data := io.NopCloser(strings.NewReader(tc.Data))
// 			mockApi.EXPECT().Get().Return(&http.Response{
// 				Code: tc.Code,
// 				Data: body,
// 			}, nil).Times(1)
//
// 			apiClient = mockApi
//
// 			resp, err := MakeAPICall()
//
// 			require.ErrorIs(t, err, tc.wantErr)
//
// 			if gotErr == nil {
// 				require.Len(t, gotReleases, tc.wantReleaseCount)
// 			}
// 		})
// 	}
// }
