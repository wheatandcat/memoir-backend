// Code generated by moq; DO NOT EDIT.
// github.com/matryer/moq

package repository

import (
	"cloud.google.com/go/firestore"
	"context"
	"github.com/wheatandcat/memoir-backend/graph/model"
	"sync"
	"time"
)

// Ensure, that ItemRepositoryInterfaceMock does implement ItemRepositoryInterface.
// If this is not the case, regenerate this file with moq.
var _ ItemRepositoryInterface = &ItemRepositoryInterfaceMock{}

// ItemRepositoryInterfaceMock is a mock implementation of ItemRepositoryInterface.
//
// 	func TestSomethingThatUsesItemRepositoryInterface(t *testing.T) {
//
// 		// make and configure a mocked ItemRepositoryInterface
// 		mockedItemRepositoryInterface := &ItemRepositoryInterfaceMock{
// 			CreateFunc: func(ctx context.Context, f *firestore.Client, userID string, i *model.Item) error {
// 				panic("mock out the Create method")
// 			},
// 			DeleteFunc: func(ctx context.Context, f *firestore.Client, userID string, i *model.DeleteItem) error {
// 				panic("mock out the Delete method")
// 			},
// 			GetItemFunc: func(ctx context.Context, f *firestore.Client, userID string, id string) (*model.Item, error) {
// 				panic("mock out the GetItem method")
// 			},
// 			GetItemUserMultipleInPeriodFunc: func(ctx context.Context, f *firestore.Client, userID []string, stertDate time.Time, endDate time.Time, first int, cursor ItemsInPeriodCursor) ([]*model.Item, error) {
// 				panic("mock out the GetItemUserMultipleInPeriod method")
// 			},
// 			GetItemsInDateFunc: func(ctx context.Context, f *firestore.Client, userID string, date time.Time) ([]*model.Item, error) {
// 				panic("mock out the GetItemsInDate method")
// 			},
// 			GetItemsInPeriodFunc: func(ctx context.Context, f *firestore.Client, userID string, stertDate time.Time, endDate time.Time, first int, cursor ItemsInPeriodCursor) ([]*model.Item, error) {
// 				panic("mock out the GetItemsInPeriod method")
// 			},
// 			UpdateFunc: func(ctx context.Context, f *firestore.Client, userID string, i *model.UpdateItem, updatedAt time.Time) error {
// 				panic("mock out the Update method")
// 			},
// 		}
//
// 		// use mockedItemRepositoryInterface in code that requires ItemRepositoryInterface
// 		// and then make assertions.
//
// 	}
type ItemRepositoryInterfaceMock struct {
	// CreateFunc mocks the Create method.
	CreateFunc func(ctx context.Context, f *firestore.Client, userID string, i *model.Item) error

	// DeleteFunc mocks the Delete method.
	DeleteFunc func(ctx context.Context, f *firestore.Client, userID string, i *model.DeleteItem) error

	// GetItemFunc mocks the GetItem method.
	GetItemFunc func(ctx context.Context, f *firestore.Client, userID string, id string) (*model.Item, error)

	// GetItemUserMultipleInPeriodFunc mocks the GetItemUserMultipleInPeriod method.
	GetItemUserMultipleInPeriodFunc func(ctx context.Context, f *firestore.Client, userID []string, stertDate time.Time, endDate time.Time, first int, cursor ItemsInPeriodCursor) ([]*model.Item, error)

	// GetItemsInDateFunc mocks the GetItemsInDate method.
	GetItemsInDateFunc func(ctx context.Context, f *firestore.Client, userID string, date time.Time) ([]*model.Item, error)

	// GetItemsInPeriodFunc mocks the GetItemsInPeriod method.
	GetItemsInPeriodFunc func(ctx context.Context, f *firestore.Client, userID string, stertDate time.Time, endDate time.Time, first int, cursor ItemsInPeriodCursor) ([]*model.Item, error)

	// UpdateFunc mocks the Update method.
	UpdateFunc func(ctx context.Context, f *firestore.Client, userID string, i *model.UpdateItem, updatedAt time.Time) error

	// calls tracks calls to the methods.
	calls struct {
		// Create holds details about calls to the Create method.
		Create []struct {
			// Ctx is the ctx argument value.
			Ctx context.Context
			// F is the f argument value.
			F *firestore.Client
			// UserID is the userID argument value.
			UserID string
			// I is the i argument value.
			I *model.Item
		}
		// Delete holds details about calls to the Delete method.
		Delete []struct {
			// Ctx is the ctx argument value.
			Ctx context.Context
			// F is the f argument value.
			F *firestore.Client
			// UserID is the userID argument value.
			UserID string
			// I is the i argument value.
			I *model.DeleteItem
		}
		// GetItem holds details about calls to the GetItem method.
		GetItem []struct {
			// Ctx is the ctx argument value.
			Ctx context.Context
			// F is the f argument value.
			F *firestore.Client
			// UserID is the userID argument value.
			UserID string
			// ID is the id argument value.
			ID string
		}
		// GetItemUserMultipleInPeriod holds details about calls to the GetItemUserMultipleInPeriod method.
		GetItemUserMultipleInPeriod []struct {
			// Ctx is the ctx argument value.
			Ctx context.Context
			// F is the f argument value.
			F *firestore.Client
			// UserID is the userID argument value.
			UserID []string
			// StertDate is the stertDate argument value.
			StertDate time.Time
			// EndDate is the endDate argument value.
			EndDate time.Time
			// First is the first argument value.
			First int
			// Cursor is the cursor argument value.
			Cursor ItemsInPeriodCursor
		}
		// GetItemsInDate holds details about calls to the GetItemsInDate method.
		GetItemsInDate []struct {
			// Ctx is the ctx argument value.
			Ctx context.Context
			// F is the f argument value.
			F *firestore.Client
			// UserID is the userID argument value.
			UserID string
			// Date is the date argument value.
			Date time.Time
		}
		// GetItemsInPeriod holds details about calls to the GetItemsInPeriod method.
		GetItemsInPeriod []struct {
			// Ctx is the ctx argument value.
			Ctx context.Context
			// F is the f argument value.
			F *firestore.Client
			// UserID is the userID argument value.
			UserID string
			// StertDate is the stertDate argument value.
			StertDate time.Time
			// EndDate is the endDate argument value.
			EndDate time.Time
			// First is the first argument value.
			First int
			// Cursor is the cursor argument value.
			Cursor ItemsInPeriodCursor
		}
		// Update holds details about calls to the Update method.
		Update []struct {
			// Ctx is the ctx argument value.
			Ctx context.Context
			// F is the f argument value.
			F *firestore.Client
			// UserID is the userID argument value.
			UserID string
			// I is the i argument value.
			I *model.UpdateItem
			// UpdatedAt is the updatedAt argument value.
			UpdatedAt time.Time
		}
	}
	lockCreate                      sync.RWMutex
	lockDelete                      sync.RWMutex
	lockGetItem                     sync.RWMutex
	lockGetItemUserMultipleInPeriod sync.RWMutex
	lockGetItemsInDate              sync.RWMutex
	lockGetItemsInPeriod            sync.RWMutex
	lockUpdate                      sync.RWMutex
}

// Create calls CreateFunc.
func (mock *ItemRepositoryInterfaceMock) Create(ctx context.Context, f *firestore.Client, userID string, i *model.Item) error {
	if mock.CreateFunc == nil {
		panic("ItemRepositoryInterfaceMock.CreateFunc: method is nil but ItemRepositoryInterface.Create was just called")
	}
	callInfo := struct {
		Ctx    context.Context
		F      *firestore.Client
		UserID string
		I      *model.Item
	}{
		Ctx:    ctx,
		F:      f,
		UserID: userID,
		I:      i,
	}
	mock.lockCreate.Lock()
	mock.calls.Create = append(mock.calls.Create, callInfo)
	mock.lockCreate.Unlock()
	return mock.CreateFunc(ctx, f, userID, i)
}

// CreateCalls gets all the calls that were made to Create.
// Check the length with:
//     len(mockedItemRepositoryInterface.CreateCalls())
func (mock *ItemRepositoryInterfaceMock) CreateCalls() []struct {
	Ctx    context.Context
	F      *firestore.Client
	UserID string
	I      *model.Item
} {
	var calls []struct {
		Ctx    context.Context
		F      *firestore.Client
		UserID string
		I      *model.Item
	}
	mock.lockCreate.RLock()
	calls = mock.calls.Create
	mock.lockCreate.RUnlock()
	return calls
}

// Delete calls DeleteFunc.
func (mock *ItemRepositoryInterfaceMock) Delete(ctx context.Context, f *firestore.Client, userID string, i *model.DeleteItem) error {
	if mock.DeleteFunc == nil {
		panic("ItemRepositoryInterfaceMock.DeleteFunc: method is nil but ItemRepositoryInterface.Delete was just called")
	}
	callInfo := struct {
		Ctx    context.Context
		F      *firestore.Client
		UserID string
		I      *model.DeleteItem
	}{
		Ctx:    ctx,
		F:      f,
		UserID: userID,
		I:      i,
	}
	mock.lockDelete.Lock()
	mock.calls.Delete = append(mock.calls.Delete, callInfo)
	mock.lockDelete.Unlock()
	return mock.DeleteFunc(ctx, f, userID, i)
}

// DeleteCalls gets all the calls that were made to Delete.
// Check the length with:
//     len(mockedItemRepositoryInterface.DeleteCalls())
func (mock *ItemRepositoryInterfaceMock) DeleteCalls() []struct {
	Ctx    context.Context
	F      *firestore.Client
	UserID string
	I      *model.DeleteItem
} {
	var calls []struct {
		Ctx    context.Context
		F      *firestore.Client
		UserID string
		I      *model.DeleteItem
	}
	mock.lockDelete.RLock()
	calls = mock.calls.Delete
	mock.lockDelete.RUnlock()
	return calls
}

// GetItem calls GetItemFunc.
func (mock *ItemRepositoryInterfaceMock) GetItem(ctx context.Context, f *firestore.Client, userID string, id string) (*model.Item, error) {
	if mock.GetItemFunc == nil {
		panic("ItemRepositoryInterfaceMock.GetItemFunc: method is nil but ItemRepositoryInterface.GetItem was just called")
	}
	callInfo := struct {
		Ctx    context.Context
		F      *firestore.Client
		UserID string
		ID     string
	}{
		Ctx:    ctx,
		F:      f,
		UserID: userID,
		ID:     id,
	}
	mock.lockGetItem.Lock()
	mock.calls.GetItem = append(mock.calls.GetItem, callInfo)
	mock.lockGetItem.Unlock()
	return mock.GetItemFunc(ctx, f, userID, id)
}

// GetItemCalls gets all the calls that were made to GetItem.
// Check the length with:
//     len(mockedItemRepositoryInterface.GetItemCalls())
func (mock *ItemRepositoryInterfaceMock) GetItemCalls() []struct {
	Ctx    context.Context
	F      *firestore.Client
	UserID string
	ID     string
} {
	var calls []struct {
		Ctx    context.Context
		F      *firestore.Client
		UserID string
		ID     string
	}
	mock.lockGetItem.RLock()
	calls = mock.calls.GetItem
	mock.lockGetItem.RUnlock()
	return calls
}

// GetItemUserMultipleInPeriod calls GetItemUserMultipleInPeriodFunc.
func (mock *ItemRepositoryInterfaceMock) GetItemUserMultipleInPeriod(ctx context.Context, f *firestore.Client, userID []string, stertDate time.Time, endDate time.Time, first int, cursor ItemsInPeriodCursor) ([]*model.Item, error) {
	if mock.GetItemUserMultipleInPeriodFunc == nil {
		panic("ItemRepositoryInterfaceMock.GetItemUserMultipleInPeriodFunc: method is nil but ItemRepositoryInterface.GetItemUserMultipleInPeriod was just called")
	}
	callInfo := struct {
		Ctx       context.Context
		F         *firestore.Client
		UserID    []string
		StertDate time.Time
		EndDate   time.Time
		First     int
		Cursor    ItemsInPeriodCursor
	}{
		Ctx:       ctx,
		F:         f,
		UserID:    userID,
		StertDate: stertDate,
		EndDate:   endDate,
		First:     first,
		Cursor:    cursor,
	}
	mock.lockGetItemUserMultipleInPeriod.Lock()
	mock.calls.GetItemUserMultipleInPeriod = append(mock.calls.GetItemUserMultipleInPeriod, callInfo)
	mock.lockGetItemUserMultipleInPeriod.Unlock()
	return mock.GetItemUserMultipleInPeriodFunc(ctx, f, userID, stertDate, endDate, first, cursor)
}

// GetItemUserMultipleInPeriodCalls gets all the calls that were made to GetItemUserMultipleInPeriod.
// Check the length with:
//     len(mockedItemRepositoryInterface.GetItemUserMultipleInPeriodCalls())
func (mock *ItemRepositoryInterfaceMock) GetItemUserMultipleInPeriodCalls() []struct {
	Ctx       context.Context
	F         *firestore.Client
	UserID    []string
	StertDate time.Time
	EndDate   time.Time
	First     int
	Cursor    ItemsInPeriodCursor
} {
	var calls []struct {
		Ctx       context.Context
		F         *firestore.Client
		UserID    []string
		StertDate time.Time
		EndDate   time.Time
		First     int
		Cursor    ItemsInPeriodCursor
	}
	mock.lockGetItemUserMultipleInPeriod.RLock()
	calls = mock.calls.GetItemUserMultipleInPeriod
	mock.lockGetItemUserMultipleInPeriod.RUnlock()
	return calls
}

// GetItemsInDate calls GetItemsInDateFunc.
func (mock *ItemRepositoryInterfaceMock) GetItemsInDate(ctx context.Context, f *firestore.Client, userID string, date time.Time) ([]*model.Item, error) {
	if mock.GetItemsInDateFunc == nil {
		panic("ItemRepositoryInterfaceMock.GetItemsInDateFunc: method is nil but ItemRepositoryInterface.GetItemsInDate was just called")
	}
	callInfo := struct {
		Ctx    context.Context
		F      *firestore.Client
		UserID string
		Date   time.Time
	}{
		Ctx:    ctx,
		F:      f,
		UserID: userID,
		Date:   date,
	}
	mock.lockGetItemsInDate.Lock()
	mock.calls.GetItemsInDate = append(mock.calls.GetItemsInDate, callInfo)
	mock.lockGetItemsInDate.Unlock()
	return mock.GetItemsInDateFunc(ctx, f, userID, date)
}

// GetItemsInDateCalls gets all the calls that were made to GetItemsInDate.
// Check the length with:
//     len(mockedItemRepositoryInterface.GetItemsInDateCalls())
func (mock *ItemRepositoryInterfaceMock) GetItemsInDateCalls() []struct {
	Ctx    context.Context
	F      *firestore.Client
	UserID string
	Date   time.Time
} {
	var calls []struct {
		Ctx    context.Context
		F      *firestore.Client
		UserID string
		Date   time.Time
	}
	mock.lockGetItemsInDate.RLock()
	calls = mock.calls.GetItemsInDate
	mock.lockGetItemsInDate.RUnlock()
	return calls
}

// GetItemsInPeriod calls GetItemsInPeriodFunc.
func (mock *ItemRepositoryInterfaceMock) GetItemsInPeriod(ctx context.Context, f *firestore.Client, userID string, stertDate time.Time, endDate time.Time, first int, cursor ItemsInPeriodCursor) ([]*model.Item, error) {
	if mock.GetItemsInPeriodFunc == nil {
		panic("ItemRepositoryInterfaceMock.GetItemsInPeriodFunc: method is nil but ItemRepositoryInterface.GetItemsInPeriod was just called")
	}
	callInfo := struct {
		Ctx       context.Context
		F         *firestore.Client
		UserID    string
		StertDate time.Time
		EndDate   time.Time
		First     int
		Cursor    ItemsInPeriodCursor
	}{
		Ctx:       ctx,
		F:         f,
		UserID:    userID,
		StertDate: stertDate,
		EndDate:   endDate,
		First:     first,
		Cursor:    cursor,
	}
	mock.lockGetItemsInPeriod.Lock()
	mock.calls.GetItemsInPeriod = append(mock.calls.GetItemsInPeriod, callInfo)
	mock.lockGetItemsInPeriod.Unlock()
	return mock.GetItemsInPeriodFunc(ctx, f, userID, stertDate, endDate, first, cursor)
}

// GetItemsInPeriodCalls gets all the calls that were made to GetItemsInPeriod.
// Check the length with:
//     len(mockedItemRepositoryInterface.GetItemsInPeriodCalls())
func (mock *ItemRepositoryInterfaceMock) GetItemsInPeriodCalls() []struct {
	Ctx       context.Context
	F         *firestore.Client
	UserID    string
	StertDate time.Time
	EndDate   time.Time
	First     int
	Cursor    ItemsInPeriodCursor
} {
	var calls []struct {
		Ctx       context.Context
		F         *firestore.Client
		UserID    string
		StertDate time.Time
		EndDate   time.Time
		First     int
		Cursor    ItemsInPeriodCursor
	}
	mock.lockGetItemsInPeriod.RLock()
	calls = mock.calls.GetItemsInPeriod
	mock.lockGetItemsInPeriod.RUnlock()
	return calls
}

// Update calls UpdateFunc.
func (mock *ItemRepositoryInterfaceMock) Update(ctx context.Context, f *firestore.Client, userID string, i *model.UpdateItem, updatedAt time.Time) error {
	if mock.UpdateFunc == nil {
		panic("ItemRepositoryInterfaceMock.UpdateFunc: method is nil but ItemRepositoryInterface.Update was just called")
	}
	callInfo := struct {
		Ctx       context.Context
		F         *firestore.Client
		UserID    string
		I         *model.UpdateItem
		UpdatedAt time.Time
	}{
		Ctx:       ctx,
		F:         f,
		UserID:    userID,
		I:         i,
		UpdatedAt: updatedAt,
	}
	mock.lockUpdate.Lock()
	mock.calls.Update = append(mock.calls.Update, callInfo)
	mock.lockUpdate.Unlock()
	return mock.UpdateFunc(ctx, f, userID, i, updatedAt)
}

// UpdateCalls gets all the calls that were made to Update.
// Check the length with:
//     len(mockedItemRepositoryInterface.UpdateCalls())
func (mock *ItemRepositoryInterfaceMock) UpdateCalls() []struct {
	Ctx       context.Context
	F         *firestore.Client
	UserID    string
	I         *model.UpdateItem
	UpdatedAt time.Time
} {
	var calls []struct {
		Ctx       context.Context
		F         *firestore.Client
		UserID    string
		I         *model.UpdateItem
		UpdatedAt time.Time
	}
	mock.lockUpdate.RLock()
	calls = mock.calls.Update
	mock.lockUpdate.RUnlock()
	return calls
}

// Ensure, that UserRepositoryInterfaceMock does implement UserRepositoryInterface.
// If this is not the case, regenerate this file with moq.
var _ UserRepositoryInterface = &UserRepositoryInterfaceMock{}

// UserRepositoryInterfaceMock is a mock implementation of UserRepositoryInterface.
//
// 	func TestSomethingThatUsesUserRepositoryInterface(t *testing.T) {
//
// 		// make and configure a mocked UserRepositoryInterface
// 		mockedUserRepositoryInterface := &UserRepositoryInterfaceMock{
// 			CreateFunc: func(ctx context.Context, f *firestore.Client, u *model.User) error {
// 				panic("mock out the Create method")
// 			},
// 			ExistByFirebaseUIDFunc: func(ctx context.Context, f *firestore.Client, fUID string) (bool, error) {
// 				panic("mock out the ExistByFirebaseUID method")
// 			},
// 			FindByFirebaseUIDFunc: func(ctx context.Context, f *firestore.Client, fUID string) (*model.User, error) {
// 				panic("mock out the FindByFirebaseUID method")
// 			},
// 			FindByUIDFunc: func(ctx context.Context, f *firestore.Client, uid string) (*model.User, error) {
// 				panic("mock out the FindByUID method")
// 			},
// 			FindDatabaseDataByUIDFunc: func(ctx context.Context, f *firestore.Client, uid string) (*User, error) {
// 				panic("mock out the FindDatabaseDataByUID method")
// 			},
// 			UpdateFunc: func(ctx context.Context, f *firestore.Client, u *model.User) error {
// 				panic("mock out the Update method")
// 			},
// 			UpdateFirebaseUIDFunc: func(ctx context.Context, f *firestore.Client, user *User) error {
// 				panic("mock out the UpdateFirebaseUID method")
// 			},
// 		}
//
// 		// use mockedUserRepositoryInterface in code that requires UserRepositoryInterface
// 		// and then make assertions.
//
// 	}
type UserRepositoryInterfaceMock struct {
	// CreateFunc mocks the Create method.
	CreateFunc func(ctx context.Context, f *firestore.Client, u *model.User) error

	// ExistByFirebaseUIDFunc mocks the ExistByFirebaseUID method.
	ExistByFirebaseUIDFunc func(ctx context.Context, f *firestore.Client, fUID string) (bool, error)

	// FindByFirebaseUIDFunc mocks the FindByFirebaseUID method.
	FindByFirebaseUIDFunc func(ctx context.Context, f *firestore.Client, fUID string) (*model.User, error)

	// FindByUIDFunc mocks the FindByUID method.
	FindByUIDFunc func(ctx context.Context, f *firestore.Client, uid string) (*model.User, error)

	// FindDatabaseDataByUIDFunc mocks the FindDatabaseDataByUID method.
	FindDatabaseDataByUIDFunc func(ctx context.Context, f *firestore.Client, uid string) (*User, error)

	// UpdateFunc mocks the Update method.
	UpdateFunc func(ctx context.Context, f *firestore.Client, u *model.User) error

	// UpdateFirebaseUIDFunc mocks the UpdateFirebaseUID method.
	UpdateFirebaseUIDFunc func(ctx context.Context, f *firestore.Client, user *User) error

	// calls tracks calls to the methods.
	calls struct {
		// Create holds details about calls to the Create method.
		Create []struct {
			// Ctx is the ctx argument value.
			Ctx context.Context
			// F is the f argument value.
			F *firestore.Client
			// U is the u argument value.
			U *model.User
		}
		// ExistByFirebaseUID holds details about calls to the ExistByFirebaseUID method.
		ExistByFirebaseUID []struct {
			// Ctx is the ctx argument value.
			Ctx context.Context
			// F is the f argument value.
			F *firestore.Client
			// FUID is the fUID argument value.
			FUID string
		}
		// FindByFirebaseUID holds details about calls to the FindByFirebaseUID method.
		FindByFirebaseUID []struct {
			// Ctx is the ctx argument value.
			Ctx context.Context
			// F is the f argument value.
			F *firestore.Client
			// FUID is the fUID argument value.
			FUID string
		}
		// FindByUID holds details about calls to the FindByUID method.
		FindByUID []struct {
			// Ctx is the ctx argument value.
			Ctx context.Context
			// F is the f argument value.
			F *firestore.Client
			// UID is the uid argument value.
			UID string
		}
		// FindDatabaseDataByUID holds details about calls to the FindDatabaseDataByUID method.
		FindDatabaseDataByUID []struct {
			// Ctx is the ctx argument value.
			Ctx context.Context
			// F is the f argument value.
			F *firestore.Client
			// UID is the uid argument value.
			UID string
		}
		// Update holds details about calls to the Update method.
		Update []struct {
			// Ctx is the ctx argument value.
			Ctx context.Context
			// F is the f argument value.
			F *firestore.Client
			// U is the u argument value.
			U *model.User
		}
		// UpdateFirebaseUID holds details about calls to the UpdateFirebaseUID method.
		UpdateFirebaseUID []struct {
			// Ctx is the ctx argument value.
			Ctx context.Context
			// F is the f argument value.
			F *firestore.Client
			// User is the user argument value.
			User *User
		}
	}
	lockCreate                sync.RWMutex
	lockExistByFirebaseUID    sync.RWMutex
	lockFindByFirebaseUID     sync.RWMutex
	lockFindByUID             sync.RWMutex
	lockFindDatabaseDataByUID sync.RWMutex
	lockUpdate                sync.RWMutex
	lockUpdateFirebaseUID     sync.RWMutex
}

// Create calls CreateFunc.
func (mock *UserRepositoryInterfaceMock) Create(ctx context.Context, f *firestore.Client, u *model.User) error {
	if mock.CreateFunc == nil {
		panic("UserRepositoryInterfaceMock.CreateFunc: method is nil but UserRepositoryInterface.Create was just called")
	}
	callInfo := struct {
		Ctx context.Context
		F   *firestore.Client
		U   *model.User
	}{
		Ctx: ctx,
		F:   f,
		U:   u,
	}
	mock.lockCreate.Lock()
	mock.calls.Create = append(mock.calls.Create, callInfo)
	mock.lockCreate.Unlock()
	return mock.CreateFunc(ctx, f, u)
}

// CreateCalls gets all the calls that were made to Create.
// Check the length with:
//     len(mockedUserRepositoryInterface.CreateCalls())
func (mock *UserRepositoryInterfaceMock) CreateCalls() []struct {
	Ctx context.Context
	F   *firestore.Client
	U   *model.User
} {
	var calls []struct {
		Ctx context.Context
		F   *firestore.Client
		U   *model.User
	}
	mock.lockCreate.RLock()
	calls = mock.calls.Create
	mock.lockCreate.RUnlock()
	return calls
}

// ExistByFirebaseUID calls ExistByFirebaseUIDFunc.
func (mock *UserRepositoryInterfaceMock) ExistByFirebaseUID(ctx context.Context, f *firestore.Client, fUID string) (bool, error) {
	if mock.ExistByFirebaseUIDFunc == nil {
		panic("UserRepositoryInterfaceMock.ExistByFirebaseUIDFunc: method is nil but UserRepositoryInterface.ExistByFirebaseUID was just called")
	}
	callInfo := struct {
		Ctx  context.Context
		F    *firestore.Client
		FUID string
	}{
		Ctx:  ctx,
		F:    f,
		FUID: fUID,
	}
	mock.lockExistByFirebaseUID.Lock()
	mock.calls.ExistByFirebaseUID = append(mock.calls.ExistByFirebaseUID, callInfo)
	mock.lockExistByFirebaseUID.Unlock()
	return mock.ExistByFirebaseUIDFunc(ctx, f, fUID)
}

// ExistByFirebaseUIDCalls gets all the calls that were made to ExistByFirebaseUID.
// Check the length with:
//     len(mockedUserRepositoryInterface.ExistByFirebaseUIDCalls())
func (mock *UserRepositoryInterfaceMock) ExistByFirebaseUIDCalls() []struct {
	Ctx  context.Context
	F    *firestore.Client
	FUID string
} {
	var calls []struct {
		Ctx  context.Context
		F    *firestore.Client
		FUID string
	}
	mock.lockExistByFirebaseUID.RLock()
	calls = mock.calls.ExistByFirebaseUID
	mock.lockExistByFirebaseUID.RUnlock()
	return calls
}

// FindByFirebaseUID calls FindByFirebaseUIDFunc.
func (mock *UserRepositoryInterfaceMock) FindByFirebaseUID(ctx context.Context, f *firestore.Client, fUID string) (*model.User, error) {
	if mock.FindByFirebaseUIDFunc == nil {
		panic("UserRepositoryInterfaceMock.FindByFirebaseUIDFunc: method is nil but UserRepositoryInterface.FindByFirebaseUID was just called")
	}
	callInfo := struct {
		Ctx  context.Context
		F    *firestore.Client
		FUID string
	}{
		Ctx:  ctx,
		F:    f,
		FUID: fUID,
	}
	mock.lockFindByFirebaseUID.Lock()
	mock.calls.FindByFirebaseUID = append(mock.calls.FindByFirebaseUID, callInfo)
	mock.lockFindByFirebaseUID.Unlock()
	return mock.FindByFirebaseUIDFunc(ctx, f, fUID)
}

// FindByFirebaseUIDCalls gets all the calls that were made to FindByFirebaseUID.
// Check the length with:
//     len(mockedUserRepositoryInterface.FindByFirebaseUIDCalls())
func (mock *UserRepositoryInterfaceMock) FindByFirebaseUIDCalls() []struct {
	Ctx  context.Context
	F    *firestore.Client
	FUID string
} {
	var calls []struct {
		Ctx  context.Context
		F    *firestore.Client
		FUID string
	}
	mock.lockFindByFirebaseUID.RLock()
	calls = mock.calls.FindByFirebaseUID
	mock.lockFindByFirebaseUID.RUnlock()
	return calls
}

// FindByUID calls FindByUIDFunc.
func (mock *UserRepositoryInterfaceMock) FindByUID(ctx context.Context, f *firestore.Client, uid string) (*model.User, error) {
	if mock.FindByUIDFunc == nil {
		panic("UserRepositoryInterfaceMock.FindByUIDFunc: method is nil but UserRepositoryInterface.FindByUID was just called")
	}
	callInfo := struct {
		Ctx context.Context
		F   *firestore.Client
		UID string
	}{
		Ctx: ctx,
		F:   f,
		UID: uid,
	}
	mock.lockFindByUID.Lock()
	mock.calls.FindByUID = append(mock.calls.FindByUID, callInfo)
	mock.lockFindByUID.Unlock()
	return mock.FindByUIDFunc(ctx, f, uid)
}

// FindByUIDCalls gets all the calls that were made to FindByUID.
// Check the length with:
//     len(mockedUserRepositoryInterface.FindByUIDCalls())
func (mock *UserRepositoryInterfaceMock) FindByUIDCalls() []struct {
	Ctx context.Context
	F   *firestore.Client
	UID string
} {
	var calls []struct {
		Ctx context.Context
		F   *firestore.Client
		UID string
	}
	mock.lockFindByUID.RLock()
	calls = mock.calls.FindByUID
	mock.lockFindByUID.RUnlock()
	return calls
}

// FindDatabaseDataByUID calls FindDatabaseDataByUIDFunc.
func (mock *UserRepositoryInterfaceMock) FindDatabaseDataByUID(ctx context.Context, f *firestore.Client, uid string) (*User, error) {
	if mock.FindDatabaseDataByUIDFunc == nil {
		panic("UserRepositoryInterfaceMock.FindDatabaseDataByUIDFunc: method is nil but UserRepositoryInterface.FindDatabaseDataByUID was just called")
	}
	callInfo := struct {
		Ctx context.Context
		F   *firestore.Client
		UID string
	}{
		Ctx: ctx,
		F:   f,
		UID: uid,
	}
	mock.lockFindDatabaseDataByUID.Lock()
	mock.calls.FindDatabaseDataByUID = append(mock.calls.FindDatabaseDataByUID, callInfo)
	mock.lockFindDatabaseDataByUID.Unlock()
	return mock.FindDatabaseDataByUIDFunc(ctx, f, uid)
}

// FindDatabaseDataByUIDCalls gets all the calls that were made to FindDatabaseDataByUID.
// Check the length with:
//     len(mockedUserRepositoryInterface.FindDatabaseDataByUIDCalls())
func (mock *UserRepositoryInterfaceMock) FindDatabaseDataByUIDCalls() []struct {
	Ctx context.Context
	F   *firestore.Client
	UID string
} {
	var calls []struct {
		Ctx context.Context
		F   *firestore.Client
		UID string
	}
	mock.lockFindDatabaseDataByUID.RLock()
	calls = mock.calls.FindDatabaseDataByUID
	mock.lockFindDatabaseDataByUID.RUnlock()
	return calls
}

// Update calls UpdateFunc.
func (mock *UserRepositoryInterfaceMock) Update(ctx context.Context, f *firestore.Client, u *model.User) error {
	if mock.UpdateFunc == nil {
		panic("UserRepositoryInterfaceMock.UpdateFunc: method is nil but UserRepositoryInterface.Update was just called")
	}
	callInfo := struct {
		Ctx context.Context
		F   *firestore.Client
		U   *model.User
	}{
		Ctx: ctx,
		F:   f,
		U:   u,
	}
	mock.lockUpdate.Lock()
	mock.calls.Update = append(mock.calls.Update, callInfo)
	mock.lockUpdate.Unlock()
	return mock.UpdateFunc(ctx, f, u)
}

// UpdateCalls gets all the calls that were made to Update.
// Check the length with:
//     len(mockedUserRepositoryInterface.UpdateCalls())
func (mock *UserRepositoryInterfaceMock) UpdateCalls() []struct {
	Ctx context.Context
	F   *firestore.Client
	U   *model.User
} {
	var calls []struct {
		Ctx context.Context
		F   *firestore.Client
		U   *model.User
	}
	mock.lockUpdate.RLock()
	calls = mock.calls.Update
	mock.lockUpdate.RUnlock()
	return calls
}

// UpdateFirebaseUID calls UpdateFirebaseUIDFunc.
func (mock *UserRepositoryInterfaceMock) UpdateFirebaseUID(ctx context.Context, f *firestore.Client, user *User) error {
	if mock.UpdateFirebaseUIDFunc == nil {
		panic("UserRepositoryInterfaceMock.UpdateFirebaseUIDFunc: method is nil but UserRepositoryInterface.UpdateFirebaseUID was just called")
	}
	callInfo := struct {
		Ctx  context.Context
		F    *firestore.Client
		User *User
	}{
		Ctx:  ctx,
		F:    f,
		User: user,
	}
	mock.lockUpdateFirebaseUID.Lock()
	mock.calls.UpdateFirebaseUID = append(mock.calls.UpdateFirebaseUID, callInfo)
	mock.lockUpdateFirebaseUID.Unlock()
	return mock.UpdateFirebaseUIDFunc(ctx, f, user)
}

// UpdateFirebaseUIDCalls gets all the calls that were made to UpdateFirebaseUID.
// Check the length with:
//     len(mockedUserRepositoryInterface.UpdateFirebaseUIDCalls())
func (mock *UserRepositoryInterfaceMock) UpdateFirebaseUIDCalls() []struct {
	Ctx  context.Context
	F    *firestore.Client
	User *User
} {
	var calls []struct {
		Ctx  context.Context
		F    *firestore.Client
		User *User
	}
	mock.lockUpdateFirebaseUID.RLock()
	calls = mock.calls.UpdateFirebaseUID
	mock.lockUpdateFirebaseUID.RUnlock()
	return calls
}

// Ensure, that InviteRepositoryInterfaceMock does implement InviteRepositoryInterface.
// If this is not the case, regenerate this file with moq.
var _ InviteRepositoryInterface = &InviteRepositoryInterfaceMock{}

// InviteRepositoryInterfaceMock is a mock implementation of InviteRepositoryInterface.
//
// 	func TestSomethingThatUsesInviteRepositoryInterface(t *testing.T) {
//
// 		// make and configure a mocked InviteRepositoryInterface
// 		mockedInviteRepositoryInterface := &InviteRepositoryInterfaceMock{
// 			CommitFunc: func(ctx context.Context, batch *firestore.WriteBatch) error {
// 				panic("mock out the Commit method")
// 			},
// 			CreateFunc: func(ctx context.Context, f *firestore.Client, batch *firestore.WriteBatch, i *model.Invite)  {
// 				panic("mock out the Create method")
// 			},
// 			DeleteFunc: func(ctx context.Context, f *firestore.Client, batch *firestore.WriteBatch, code string)  {
// 				panic("mock out the Delete method")
// 			},
// 			FindFunc: func(ctx context.Context, f *firestore.Client, code string) (*model.Invite, error) {
// 				panic("mock out the Find method")
// 			},
// 			FindByUserIDFunc: func(ctx context.Context, f *firestore.Client, userID string) (*model.Invite, error) {
// 				panic("mock out the FindByUserID method")
// 			},
// 		}
//
// 		// use mockedInviteRepositoryInterface in code that requires InviteRepositoryInterface
// 		// and then make assertions.
//
// 	}
type InviteRepositoryInterfaceMock struct {
	// CommitFunc mocks the Commit method.
	CommitFunc func(ctx context.Context, batch *firestore.WriteBatch) error

	// CreateFunc mocks the Create method.
	CreateFunc func(ctx context.Context, f *firestore.Client, batch *firestore.WriteBatch, i *model.Invite)

	// DeleteFunc mocks the Delete method.
	DeleteFunc func(ctx context.Context, f *firestore.Client, batch *firestore.WriteBatch, code string)

	// FindFunc mocks the Find method.
	FindFunc func(ctx context.Context, f *firestore.Client, code string) (*model.Invite, error)

	// FindByUserIDFunc mocks the FindByUserID method.
	FindByUserIDFunc func(ctx context.Context, f *firestore.Client, userID string) (*model.Invite, error)

	// calls tracks calls to the methods.
	calls struct {
		// Commit holds details about calls to the Commit method.
		Commit []struct {
			// Ctx is the ctx argument value.
			Ctx context.Context
			// Batch is the batch argument value.
			Batch *firestore.WriteBatch
		}
		// Create holds details about calls to the Create method.
		Create []struct {
			// Ctx is the ctx argument value.
			Ctx context.Context
			// F is the f argument value.
			F *firestore.Client
			// Batch is the batch argument value.
			Batch *firestore.WriteBatch
			// I is the i argument value.
			I *model.Invite
		}
		// Delete holds details about calls to the Delete method.
		Delete []struct {
			// Ctx is the ctx argument value.
			Ctx context.Context
			// F is the f argument value.
			F *firestore.Client
			// Batch is the batch argument value.
			Batch *firestore.WriteBatch
			// Code is the code argument value.
			Code string
		}
		// Find holds details about calls to the Find method.
		Find []struct {
			// Ctx is the ctx argument value.
			Ctx context.Context
			// F is the f argument value.
			F *firestore.Client
			// Code is the code argument value.
			Code string
		}
		// FindByUserID holds details about calls to the FindByUserID method.
		FindByUserID []struct {
			// Ctx is the ctx argument value.
			Ctx context.Context
			// F is the f argument value.
			F *firestore.Client
			// UserID is the userID argument value.
			UserID string
		}
	}
	lockCommit       sync.RWMutex
	lockCreate       sync.RWMutex
	lockDelete       sync.RWMutex
	lockFind         sync.RWMutex
	lockFindByUserID sync.RWMutex
}

// Commit calls CommitFunc.
func (mock *InviteRepositoryInterfaceMock) Commit(ctx context.Context, batch *firestore.WriteBatch) error {
	if mock.CommitFunc == nil {
		panic("InviteRepositoryInterfaceMock.CommitFunc: method is nil but InviteRepositoryInterface.Commit was just called")
	}
	callInfo := struct {
		Ctx   context.Context
		Batch *firestore.WriteBatch
	}{
		Ctx:   ctx,
		Batch: batch,
	}
	mock.lockCommit.Lock()
	mock.calls.Commit = append(mock.calls.Commit, callInfo)
	mock.lockCommit.Unlock()
	return mock.CommitFunc(ctx, batch)
}

// CommitCalls gets all the calls that were made to Commit.
// Check the length with:
//     len(mockedInviteRepositoryInterface.CommitCalls())
func (mock *InviteRepositoryInterfaceMock) CommitCalls() []struct {
	Ctx   context.Context
	Batch *firestore.WriteBatch
} {
	var calls []struct {
		Ctx   context.Context
		Batch *firestore.WriteBatch
	}
	mock.lockCommit.RLock()
	calls = mock.calls.Commit
	mock.lockCommit.RUnlock()
	return calls
}

// Create calls CreateFunc.
func (mock *InviteRepositoryInterfaceMock) Create(ctx context.Context, f *firestore.Client, batch *firestore.WriteBatch, i *model.Invite) {
	if mock.CreateFunc == nil {
		panic("InviteRepositoryInterfaceMock.CreateFunc: method is nil but InviteRepositoryInterface.Create was just called")
	}
	callInfo := struct {
		Ctx   context.Context
		F     *firestore.Client
		Batch *firestore.WriteBatch
		I     *model.Invite
	}{
		Ctx:   ctx,
		F:     f,
		Batch: batch,
		I:     i,
	}
	mock.lockCreate.Lock()
	mock.calls.Create = append(mock.calls.Create, callInfo)
	mock.lockCreate.Unlock()
	mock.CreateFunc(ctx, f, batch, i)
}

// CreateCalls gets all the calls that were made to Create.
// Check the length with:
//     len(mockedInviteRepositoryInterface.CreateCalls())
func (mock *InviteRepositoryInterfaceMock) CreateCalls() []struct {
	Ctx   context.Context
	F     *firestore.Client
	Batch *firestore.WriteBatch
	I     *model.Invite
} {
	var calls []struct {
		Ctx   context.Context
		F     *firestore.Client
		Batch *firestore.WriteBatch
		I     *model.Invite
	}
	mock.lockCreate.RLock()
	calls = mock.calls.Create
	mock.lockCreate.RUnlock()
	return calls
}

// Delete calls DeleteFunc.
func (mock *InviteRepositoryInterfaceMock) Delete(ctx context.Context, f *firestore.Client, batch *firestore.WriteBatch, code string) {
	if mock.DeleteFunc == nil {
		panic("InviteRepositoryInterfaceMock.DeleteFunc: method is nil but InviteRepositoryInterface.Delete was just called")
	}
	callInfo := struct {
		Ctx   context.Context
		F     *firestore.Client
		Batch *firestore.WriteBatch
		Code  string
	}{
		Ctx:   ctx,
		F:     f,
		Batch: batch,
		Code:  code,
	}
	mock.lockDelete.Lock()
	mock.calls.Delete = append(mock.calls.Delete, callInfo)
	mock.lockDelete.Unlock()
	mock.DeleteFunc(ctx, f, batch, code)
}

// DeleteCalls gets all the calls that were made to Delete.
// Check the length with:
//     len(mockedInviteRepositoryInterface.DeleteCalls())
func (mock *InviteRepositoryInterfaceMock) DeleteCalls() []struct {
	Ctx   context.Context
	F     *firestore.Client
	Batch *firestore.WriteBatch
	Code  string
} {
	var calls []struct {
		Ctx   context.Context
		F     *firestore.Client
		Batch *firestore.WriteBatch
		Code  string
	}
	mock.lockDelete.RLock()
	calls = mock.calls.Delete
	mock.lockDelete.RUnlock()
	return calls
}

// Find calls FindFunc.
func (mock *InviteRepositoryInterfaceMock) Find(ctx context.Context, f *firestore.Client, code string) (*model.Invite, error) {
	if mock.FindFunc == nil {
		panic("InviteRepositoryInterfaceMock.FindFunc: method is nil but InviteRepositoryInterface.Find was just called")
	}
	callInfo := struct {
		Ctx  context.Context
		F    *firestore.Client
		Code string
	}{
		Ctx:  ctx,
		F:    f,
		Code: code,
	}
	mock.lockFind.Lock()
	mock.calls.Find = append(mock.calls.Find, callInfo)
	mock.lockFind.Unlock()
	return mock.FindFunc(ctx, f, code)
}

// FindCalls gets all the calls that were made to Find.
// Check the length with:
//     len(mockedInviteRepositoryInterface.FindCalls())
func (mock *InviteRepositoryInterfaceMock) FindCalls() []struct {
	Ctx  context.Context
	F    *firestore.Client
	Code string
} {
	var calls []struct {
		Ctx  context.Context
		F    *firestore.Client
		Code string
	}
	mock.lockFind.RLock()
	calls = mock.calls.Find
	mock.lockFind.RUnlock()
	return calls
}

// FindByUserID calls FindByUserIDFunc.
func (mock *InviteRepositoryInterfaceMock) FindByUserID(ctx context.Context, f *firestore.Client, userID string) (*model.Invite, error) {
	if mock.FindByUserIDFunc == nil {
		panic("InviteRepositoryInterfaceMock.FindByUserIDFunc: method is nil but InviteRepositoryInterface.FindByUserID was just called")
	}
	callInfo := struct {
		Ctx    context.Context
		F      *firestore.Client
		UserID string
	}{
		Ctx:    ctx,
		F:      f,
		UserID: userID,
	}
	mock.lockFindByUserID.Lock()
	mock.calls.FindByUserID = append(mock.calls.FindByUserID, callInfo)
	mock.lockFindByUserID.Unlock()
	return mock.FindByUserIDFunc(ctx, f, userID)
}

// FindByUserIDCalls gets all the calls that were made to FindByUserID.
// Check the length with:
//     len(mockedInviteRepositoryInterface.FindByUserIDCalls())
func (mock *InviteRepositoryInterfaceMock) FindByUserIDCalls() []struct {
	Ctx    context.Context
	F      *firestore.Client
	UserID string
} {
	var calls []struct {
		Ctx    context.Context
		F      *firestore.Client
		UserID string
	}
	mock.lockFindByUserID.RLock()
	calls = mock.calls.FindByUserID
	mock.lockFindByUserID.RUnlock()
	return calls
}
