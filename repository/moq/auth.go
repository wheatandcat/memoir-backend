// Code generated by moq; DO NOT EDIT.
// github.com/matryer/moq

package moqs

import (
	"cloud.google.com/go/firestore"
	"context"
	"github.com/wheatandcat/memoir-backend/repository"
	"sync"
)

// Ensure, that AuthRepositoryInterfaceMock does implement repository.AuthRepositoryInterface.
// If this is not the case, regenerate this file with moq.
var _ repository.AuthRepositoryInterface = &AuthRepositoryInterfaceMock{}

// AuthRepositoryInterfaceMock is a mock implementation of repository.AuthRepositoryInterface.
//
// 	func TestSomethingThatUsesAuthRepositoryInterface(t *testing.T) {
//
// 		// make and configure a mocked repository.AuthRepositoryInterface
// 		mockedAuthRepositoryInterface := &AuthRepositoryInterfaceMock{
// 			DeleteFunc: func(ctx context.Context, f *firestore.Client, batch *firestore.WriteBatch, uid string)  {
// 				panic("mock out the Delete method")
// 			},
// 		}
//
// 		// use mockedAuthRepositoryInterface in code that requires repository.AuthRepositoryInterface
// 		// and then make assertions.
//
// 	}
type AuthRepositoryInterfaceMock struct {
	// DeleteFunc mocks the Delete method.
	DeleteFunc func(ctx context.Context, f *firestore.Client, batch *firestore.WriteBatch, uid string)

	// calls tracks calls to the methods.
	calls struct {
		// Delete holds details about calls to the Delete method.
		Delete []struct {
			// Ctx is the ctx argument value.
			Ctx context.Context
			// F is the f argument value.
			F *firestore.Client
			// Batch is the batch argument value.
			Batch *firestore.WriteBatch
			// UID is the uid argument value.
			UID string
		}
	}
	lockDelete sync.RWMutex
}

// Delete calls DeleteFunc.
func (mock *AuthRepositoryInterfaceMock) Delete(ctx context.Context, f *firestore.Client, batch *firestore.WriteBatch, uid string) {
	if mock.DeleteFunc == nil {
		panic("AuthRepositoryInterfaceMock.DeleteFunc: method is nil but AuthRepositoryInterface.Delete was just called")
	}
	callInfo := struct {
		Ctx   context.Context
		F     *firestore.Client
		Batch *firestore.WriteBatch
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
	mock.DeleteFunc(ctx, f, batch, uid)
}

// DeleteCalls gets all the calls that were made to Delete.
// Check the length with:
//     len(mockedAuthRepositoryInterface.DeleteCalls())
func (mock *AuthRepositoryInterfaceMock) DeleteCalls() []struct {
	Ctx   context.Context
	F     *firestore.Client
	Batch *firestore.WriteBatch
	UID   string
} {
	var calls []struct {
		Ctx   context.Context
		F     *firestore.Client
		Batch *firestore.WriteBatch
		UID   string
	}
	mock.lockDelete.RLock()
	calls = mock.calls.Delete
	mock.lockDelete.RUnlock()
	return calls
}