// Code generated by moq; DO NOT EDIT.
// github.com/matryer/moq

package moqs

import (
	"cloud.google.com/go/firestore"
	"context"
	"github.com/wheatandcat/memoir-backend/graph/model"
	"github.com/wheatandcat/memoir-backend/repository"
	"sync"
)

// Ensure, that UserRepositoryInterfaceMock does implement repository.UserRepositoryInterface.
// If this is not the case, regenerate this file with moq.
var _ repository.UserRepositoryInterface = &UserRepositoryInterfaceMock{}

// UserRepositoryInterfaceMock is a mock implementation of repository.UserRepositoryInterface.
//
//	func TestSomethingThatUsesUserRepositoryInterface(t *testing.T) {
//
//		// make and configure a mocked repository.UserRepositoryInterface
//		mockedUserRepositoryInterface := &UserRepositoryInterfaceMock{
//			CreateFunc: func(ctx context.Context, f *firestore.Client, u *model.User) error {
//				panic("mock out the Create method")
//			},
//			DeleteFunc: func(ctx context.Context, f *firestore.Client, batch *firestore.BulkWriter, uid string) error {
//				panic("mock out the Delete method")
//			},
//			ExistByFirebaseUIDFunc: func(ctx context.Context, f *firestore.Client, fUID string) (bool, error) {
//				panic("mock out the ExistByFirebaseUID method")
//			},
//			FindByFirebaseUIDFunc: func(ctx context.Context, f *firestore.Client, fUID string) (*model.User, error) {
//				panic("mock out the FindByFirebaseUID method")
//			},
//			FindByUIDFunc: func(ctx context.Context, f *firestore.Client, uid string) (*model.User, error) {
//				panic("mock out the FindByUID method")
//			},
//			FindDatabaseDataByUIDFunc: func(ctx context.Context, f *firestore.Client, uid string) (*repository.User, error) {
//				panic("mock out the FindDatabaseDataByUID method")
//			},
//			FindInUIDFunc: func(ctx context.Context, f *firestore.Client, uid []string) ([]*model.User, error) {
//				panic("mock out the FindInUID method")
//			},
//			UpdateFunc: func(ctx context.Context, f *firestore.Client, u *model.User) error {
//				panic("mock out the Update method")
//			},
//			UpdateFirebaseUIDFunc: func(ctx context.Context, f *firestore.Client, user *repository.User) error {
//				panic("mock out the UpdateFirebaseUID method")
//			},
//		}
//
//		// use mockedUserRepositoryInterface in code that requires repository.UserRepositoryInterface
//		// and then make assertions.
//
//	}
type UserRepositoryInterfaceMock struct {
	// CreateFunc mocks the Create method.
	CreateFunc func(ctx context.Context, f *firestore.Client, u *model.User) error

	// DeleteFunc mocks the Delete method.
	DeleteFunc func(ctx context.Context, f *firestore.Client, batch *firestore.BulkWriter, uid string) error

	// ExistByFirebaseUIDFunc mocks the ExistByFirebaseUID method.
	ExistByFirebaseUIDFunc func(ctx context.Context, f *firestore.Client, fUID string) (bool, error)

	// FindByFirebaseUIDFunc mocks the FindByFirebaseUID method.
	FindByFirebaseUIDFunc func(ctx context.Context, f *firestore.Client, fUID string) (*model.User, error)

	// FindByUIDFunc mocks the FindByUID method.
	FindByUIDFunc func(ctx context.Context, f *firestore.Client, uid string) (*model.User, error)

	// FindDatabaseDataByUIDFunc mocks the FindDatabaseDataByUID method.
	FindDatabaseDataByUIDFunc func(ctx context.Context, f *firestore.Client, uid string) (*repository.User, error)

	// FindInUIDFunc mocks the FindInUID method.
	FindInUIDFunc func(ctx context.Context, f *firestore.Client, uid []string) ([]*model.User, error)

	// UpdateFunc mocks the Update method.
	UpdateFunc func(ctx context.Context, f *firestore.Client, u *model.User) error

	// UpdateFirebaseUIDFunc mocks the UpdateFirebaseUID method.
	UpdateFirebaseUIDFunc func(ctx context.Context, f *firestore.Client, user *repository.User) error

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
		// Delete holds details about calls to the Delete method.
		Delete []struct {
			// Ctx is the ctx argument value.
			Ctx context.Context
			// F is the f argument value.
			F *firestore.Client
			// Batch is the batch argument value.
			Batch *firestore.BulkWriter
			// UID is the uid argument value.
			UID string
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
		// FindInUID holds details about calls to the FindInUID method.
		FindInUID []struct {
			// Ctx is the ctx argument value.
			Ctx context.Context
			// F is the f argument value.
			F *firestore.Client
			// UID is the uid argument value.
			UID []string
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
			User *repository.User
		}
	}
	lockCreate                sync.RWMutex
	lockDelete                sync.RWMutex
	lockExistByFirebaseUID    sync.RWMutex
	lockFindByFirebaseUID     sync.RWMutex
	lockFindByUID             sync.RWMutex
	lockFindDatabaseDataByUID sync.RWMutex
	lockFindInUID             sync.RWMutex
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
//
//	len(mockedUserRepositoryInterface.CreateCalls())
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

// Delete calls DeleteFunc.
func (mock *UserRepositoryInterfaceMock) Delete(ctx context.Context, f *firestore.Client, batch *firestore.BulkWriter, uid string) error {
	if mock.DeleteFunc == nil {
		panic("UserRepositoryInterfaceMock.DeleteFunc: method is nil but UserRepositoryInterface.Delete was just called")
	}
	callInfo := struct {
		Ctx   context.Context
		F     *firestore.Client
		Batch *firestore.BulkWriter
		UID   string
	}{
		Ctx:   ctx,
		F:     f,
		Batch: batch,
		UID:   uid,
	}
	mock.lockDelete.Lock()
	mock.calls.Delete = append(mock.calls.Delete, callInfo)
	mock.lockDelete.Unlock()
	return mock.DeleteFunc(ctx, f, batch, uid)
}

// DeleteCalls gets all the calls that were made to Delete.
// Check the length with:
//
//	len(mockedUserRepositoryInterface.DeleteCalls())
func (mock *UserRepositoryInterfaceMock) DeleteCalls() []struct {
	Ctx   context.Context
	F     *firestore.Client
	Batch *firestore.BulkWriter
	UID   string
} {
	var calls []struct {
		Ctx   context.Context
		F     *firestore.Client
		Batch *firestore.BulkWriter
		UID   string
	}
	mock.lockDelete.RLock()
	calls = mock.calls.Delete
	mock.lockDelete.RUnlock()
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
//
//	len(mockedUserRepositoryInterface.ExistByFirebaseUIDCalls())
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
//
//	len(mockedUserRepositoryInterface.FindByFirebaseUIDCalls())
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
//
//	len(mockedUserRepositoryInterface.FindByUIDCalls())
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
func (mock *UserRepositoryInterfaceMock) FindDatabaseDataByUID(ctx context.Context, f *firestore.Client, uid string) (*repository.User, error) {
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
//
//	len(mockedUserRepositoryInterface.FindDatabaseDataByUIDCalls())
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

// FindInUID calls FindInUIDFunc.
func (mock *UserRepositoryInterfaceMock) FindInUID(ctx context.Context, f *firestore.Client, uid []string) ([]*model.User, error) {
	if mock.FindInUIDFunc == nil {
		panic("UserRepositoryInterfaceMock.FindInUIDFunc: method is nil but UserRepositoryInterface.FindInUID was just called")
	}
	callInfo := struct {
		Ctx context.Context
		F   *firestore.Client
		UID []string
	}{
		Ctx: ctx,
		F:   f,
		UID: uid,
	}
	mock.lockFindInUID.Lock()
	mock.calls.FindInUID = append(mock.calls.FindInUID, callInfo)
	mock.lockFindInUID.Unlock()
	return mock.FindInUIDFunc(ctx, f, uid)
}

// FindInUIDCalls gets all the calls that were made to FindInUID.
// Check the length with:
//
//	len(mockedUserRepositoryInterface.FindInUIDCalls())
func (mock *UserRepositoryInterfaceMock) FindInUIDCalls() []struct {
	Ctx context.Context
	F   *firestore.Client
	UID []string
} {
	var calls []struct {
		Ctx context.Context
		F   *firestore.Client
		UID []string
	}
	mock.lockFindInUID.RLock()
	calls = mock.calls.FindInUID
	mock.lockFindInUID.RUnlock()
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
//
//	len(mockedUserRepositoryInterface.UpdateCalls())
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
func (mock *UserRepositoryInterfaceMock) UpdateFirebaseUID(ctx context.Context, f *firestore.Client, user *repository.User) error {
	if mock.UpdateFirebaseUIDFunc == nil {
		panic("UserRepositoryInterfaceMock.UpdateFirebaseUIDFunc: method is nil but UserRepositoryInterface.UpdateFirebaseUID was just called")
	}
	callInfo := struct {
		Ctx  context.Context
		F    *firestore.Client
		User *repository.User
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
//
//	len(mockedUserRepositoryInterface.UpdateFirebaseUIDCalls())
func (mock *UserRepositoryInterfaceMock) UpdateFirebaseUIDCalls() []struct {
	Ctx  context.Context
	F    *firestore.Client
	User *repository.User
} {
	var calls []struct {
		Ctx  context.Context
		F    *firestore.Client
		User *repository.User
	}
	mock.lockUpdateFirebaseUID.RLock()
	calls = mock.calls.UpdateFirebaseUID
	mock.lockUpdateFirebaseUID.RUnlock()
	return calls
}
