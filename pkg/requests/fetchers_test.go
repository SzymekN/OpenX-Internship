package requests

import (
	"bytes"
	"errors"
	"io/ioutil"
	"net/http"
	"reflect"
	"testing"

	"github.com/SzymekN/OpenX-Internship/model"
	mock_requests "github.com/SzymekN/OpenX-Internship/pkg/requests/mocks"
	"github.com/golang/mock/gomock"
)

var sampleUsersCorrectResponse = `[{"address":{"geolocation":{"lat":"-37.3159","long":"81.1496"},"city":"kilcoole","street":"new road","number":7682,"zipcode":"12926-3874"},"id":1,"email":"john@gmail.com","username":"johnd","password":"m38rmF$","name":{"firstname":"john","lastname":"doe"},"phone":"1-570-236-7033","__v":0},{"address":{"geolocation":{"lat":"-37.3159","long":"81.1496"},"city":"kilcoole","street":"Lovers Ln","number":7267,"zipcode":"12926-3874"},"id":2,"email":"morrison@gmail.com","username":"mor_2314","password":"83r5^_","name":{"firstname":"david","lastname":"morrison"},"phone":"1-570-236-7033","__v":0}]`
var sampleCartsCorrectResponse = `[{"id":1,"userId":1,"date":"2020-03-02T00:00:00.000Z","products":[{"productId":1,"quantity":4},{"productId":2,"quantity":1},{"productId":3,"quantity":6}],"__v":0},{"id":2,"userId":1,"date":"2020-01-02T00:00:00.000Z","products":[{"productId":2,"quantity":4},{"productId":1,"quantity":10},{"productId":5,"quantity":2}],"__v":0}]`
var sampleProductsCorrectResponse = `[{"id":1,"title":"Fjallraven - Foldsack No. 1 Backpack, Fits 15 Laptops","price":109.95,"description":"Your perfect pack for everyday use and walks in the forest. Stash your laptop (up to 15 inches) in the padded sleeve, your everyday","category":"men's clothing","image":"https://fakestoreapi.com/img/81fPKd-2AYL._AC_SL1500_.jpg","rating":{"rate":3.9,"count":120}},{"id":2,"title":"Mens Casual Premium Slim Fit T-Shirts ","price":22.3,"description":"Slim-fitting style, contrast raglan long sleeve, three-button henley placket, light weight & soft fabric for breathable and comfortable wearing. And Solid stitched shirts with round neck made for durability and a great fit for casual fashion wear and diehard baseball fans. The Henley style round neckline includes a three-button placket.","category":"men's clothing","image":"https://fakestoreapi.com/img/71-3HjGNDUL._AC_SY879._SX._UX._SY._UY_.jpg","rating":{"rate":4.1,"count":259}}]`

func Test_fetchData_UnmarshalBreak(t *testing.T) {
	ctrl := gomock.NewController(t)

	mockGetter := mock_requests.NewMockHttpGetter(ctrl)
	r := &http.Response{StatusCode: 200, Body: http.NoBody}

	mockGetter.EXPECT().Get(URLAllUsers).Return(r, nil)

	b, err := fetchData(mockGetter, URLAllUsers, model.Users{})

	if b != nil {
		t.Logf("want: nil, got: %t", b)
		t.Fail()
	}

	if err == nil {
		t.FailNow()
	}
}
func Test_fetchData_ErrorGot(t *testing.T) {
	ctrl := gomock.NewController(t)

	mockGetter := mock_requests.NewMockHttpGetter(ctrl)
	r := &http.Response{StatusCode: 400, Body: http.NoBody}

	mockGetter.EXPECT().Get(URLAllUsers).Return(r, errors.New("err"))

	b, err := fetchData(mockGetter, URLAllUsers, model.Users{})

	if b != nil {
		t.Logf("want: nil, got: %T", b)
		t.Fail()
	}

	if err == nil {
		t.FailNow()
	}
}
func Test_fetchData_Positive(t *testing.T) {
	ctrl := gomock.NewController(t)

	mockGetter := mock_requests.NewMockHttpGetter(ctrl)
	r := &http.Response{StatusCode: 200, Body: ioutil.NopCloser(bytes.NewReader([]byte(sampleUsersCorrectResponse)))}

	mockGetter.EXPECT().Get(URLAllUsers).Return(r, nil)

	b, err := fetchData(mockGetter, URLAllUsers, model.Users{})

	if b == nil {
		t.Logf("want: {}interface, got: nil")
		t.Fail()
	}

	if err != nil {
		t.FailNow()
	}
}
func Test_fetchUsers_ReflectError(t *testing.T) {
	ctrl := gomock.NewController(t)

	mockGetter := mock_requests.NewMockHttpGetter(ctrl)
	//calling with wrong response value, triggers type assertion error in json.Unmarshal() in fetchData() function
	r := &http.Response{StatusCode: 200, Body: ioutil.NopCloser(bytes.NewReader([]byte(`{"foo":"bar"}`)))}

	mockGetter.EXPECT().Get(URLAllUsers).Return(r, nil)

	b, err := FetchUsers(mockGetter, URLAllUsers)

	if !reflect.DeepEqual(b, model.Users{}) {
		t.Logf("want: model.Users{}, got: %T", b)
		t.Fail()
	}

	if err == nil {
		t.FailNow()
	}
}
func Test_fetchUsers_Positive(t *testing.T) {
	ctrl := gomock.NewController(t)

	mockGetter := mock_requests.NewMockHttpGetter(ctrl)
	//calling with wrong response value, triggers type assertion error in json.Unmarshal() in fetchData() function
	r := &http.Response{StatusCode: 200, Body: ioutil.NopCloser(bytes.NewReader([]byte(sampleUsersCorrectResponse)))}

	mockGetter.EXPECT().Get(URLAllUsers).Return(r, nil)

	b, err := FetchUsers(mockGetter, URLAllUsers)

	if reflect.DeepEqual(b, model.Users{}) || b == nil {
		t.Logf("want: model.Users{}, got: %T", b)
		t.Fail()
	}

	if err != nil {
		t.FailNow()
	}
}
func Test_fetchCarts_ReflectError(t *testing.T) {
	ctrl := gomock.NewController(t)

	mockGetter := mock_requests.NewMockHttpGetter(ctrl)
	//calling with wrong response value, triggers type assertion error in json.Unmarshal() in fetchData() function
	r := &http.Response{StatusCode: 200, Body: ioutil.NopCloser(bytes.NewReader([]byte(`{"foo":"bar"}`)))}

	mockGetter.EXPECT().Get(URLAllCarts).Return(r, nil)

	b, err := FetchCarts(mockGetter, URLAllCarts)

	if !reflect.DeepEqual(b, model.Carts{}) {
		t.Logf("want: model.Carts{}, got: %T", b)
		t.Fail()
	}

	if err == nil {
		t.FailNow()
	}
}
func Test_fetchCarts_Positive(t *testing.T) {
	ctrl := gomock.NewController(t)

	mockGetter := mock_requests.NewMockHttpGetter(ctrl)
	//calling with wrong response value, triggers type assertion error in json.Unmarshal() in fetchData() function
	r := &http.Response{StatusCode: 200, Body: ioutil.NopCloser(bytes.NewReader([]byte(sampleCartsCorrectResponse)))}

	mockGetter.EXPECT().Get(URLAllCarts).Return(r, nil)

	b, err := FetchCarts(mockGetter, URLAllCarts)

	if reflect.DeepEqual(b, model.Carts{}) || b == nil {
		t.Logf("want: model.Carts{}, got: %T", b)
		t.Fail()
	}

	if err != nil {
		t.FailNow()
	}
}
func Test_fetchProducts_ReflectError(t *testing.T) {
	ctrl := gomock.NewController(t)

	mockGetter := mock_requests.NewMockHttpGetter(ctrl)
	//calling with wrong response value, triggers type assertion error in json.Unmarshal() in fetchData() function
	r := &http.Response{StatusCode: 200, Body: ioutil.NopCloser(bytes.NewReader([]byte(`{"foo":"bar"}`)))}

	mockGetter.EXPECT().Get(URLAllProducts).Return(r, nil)

	b, err := FetchProducts(mockGetter, URLAllProducts)

	if !reflect.DeepEqual(b, model.Products{}) {
		t.Logf("want: model.Products{}, got: %T", b)
		t.Fail()
	}

	if err == nil {
		t.FailNow()
	}
}

func Test_fetchProducts_Positive(t *testing.T) {
	ctrl := gomock.NewController(t)

	mockGetter := mock_requests.NewMockHttpGetter(ctrl)
	//calling with wrong response value, triggers type assertion error in json.Unmarshal() in fetchData() function
	r := &http.Response{StatusCode: 200, Body: ioutil.NopCloser(bytes.NewReader([]byte(sampleProductsCorrectResponse)))}

	mockGetter.EXPECT().Get(URLAllProducts).Return(r, nil)

	b, err := FetchProducts(mockGetter, URLAllProducts)

	if reflect.DeepEqual(b, model.Products{}) || b == nil {
		t.Logf("want: model.Products{}, got: %T", b)
		t.Fail()
	}

	if err != nil {
		t.FailNow()
	}
}
