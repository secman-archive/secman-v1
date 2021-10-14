package packages

import (
	"io"
	"sync"
)

// Ensure, that PackageManagerMock does implement PackageManager.
// If this is not the case, regenerate this file with moq.
var _ PackageManager = &PackageManagerMock{}

type PackageManagerMock struct {
	// CreateFunc mocks the Create method.
	CreateFunc func(name string) error

	// DispatchFunc mocks the Dispatch method.
	DispatchFunc func(args []string, stdin io.Reader, stdout io.Writer, stderr io.Writer) (bool, error)

	// InstallFunc mocks the Install method.
	InstallFunc func(url string, stdout io.Writer, stderr io.Writer) error

	// InstallLocalFunc mocks the InstallLocal method.
	InstallLocalFunc func(dir string) error

	// ListFunc mocks the List method.
	ListFunc func(includeMetadata bool) []Package

	// RemoveFunc mocks the Remove method.
	RemoveFunc func(name string) error

	// UpgradeFunc mocks the Upgrade method.
	UpgradeFunc func(name string, force bool, stdout io.Writer, stderr io.Writer) error

	// calls tracks calls to the methods.
	calls struct {
		// Create holds details about calls to the Create method.
		Create []struct {
			// Name is the name argument value.
			Name string
		}
		// Dispatch holds details about calls to the Dispatch method.
		Dispatch []struct {
			// Args is the args argument value.
			Args []string
			// Stdin is the stdin argument value.
			Stdin io.Reader
			// Stdout is the stdout argument value.
			Stdout io.Writer
			// Stderr is the stderr argument value.
			Stderr io.Writer
		}
		// Install holds details about calls to the Install method.
		Install []struct {
			// URL is the url argument value.
			URL string
			// Stdout is the stdout argument value.
			Stdout io.Writer
			// Stderr is the stderr argument value.
			Stderr io.Writer
		}
		// InstallLocal holds details about calls to the InstallLocal method.
		InstallLocal []struct {
			// Dir is the dir argument value.
			Dir string
		}
		// List holds details about calls to the List method.
		List []struct {
			// IncludeMetadata is the includeMetadata argument value.
			IncludeMetadata bool
		}
		// Remove holds details about calls to the Remove method.
		Remove []struct {
			// Name is the name argument value.
			Name string
		}
		// Upgrade holds details about calls to the Upgrade method.
		Upgrade []struct {
			// Name is the name argument value.
			Name string
			// Force is the force argument value.
			Force bool
			// Stdout is the stdout argument value.
			Stdout io.Writer
			// Stderr is the stderr argument value.
			Stderr io.Writer
		}
	}

	lockCreate       sync.RWMutex
	lockDispatch     sync.RWMutex
	lockInstall      sync.RWMutex
	lockInstallLocal sync.RWMutex
	lockList         sync.RWMutex
	lockRemove       sync.RWMutex
	lockUpgrade      sync.RWMutex
}

// Create calls CreateFunc.
func (mock *PackageManagerMock) Create(name string) error {
	if mock.CreateFunc == nil {
		panic("PackageManagerMock.CreateFunc: method is nil but PackageManager.Create was just called")
	}

	callInfo := struct {
		Name string
	}{
		Name: name,
	}

	mock.lockCreate.Lock()
	mock.calls.Create = append(mock.calls.Create, callInfo)
	mock.lockCreate.Unlock()
	return mock.CreateFunc(name)
}

// CreateCalls gets all the calls that were made to Create.
// Check the length with:
//     len(mockedPackageManager.CreateCalls())
func (mock *PackageManagerMock) CreateCalls() []struct {
	Name string
} {
	var calls []struct {
		Name string
	}

	mock.lockCreate.RLock()
	calls = mock.calls.Create
	mock.lockCreate.RUnlock()
	return calls
}

// Dispatch calls DispatchFunc.
func (mock *PackageManagerMock) Dispatch(args []string, stdin io.Reader, stdout io.Writer, stderr io.Writer) (bool, error) {
	if mock.DispatchFunc == nil {
		panic("PackageManagerMock.DispatchFunc: method is nil but PackageManager.Dispatch was just called")
	}

	callInfo := struct {
		Args   []string
		Stdin  io.Reader
		Stdout io.Writer
		Stderr io.Writer
	}{
		Args:   args,
		Stdin:  stdin,
		Stdout: stdout,
		Stderr: stderr,
	}

	mock.lockDispatch.Lock()
	mock.calls.Dispatch = append(mock.calls.Dispatch, callInfo)
	mock.lockDispatch.Unlock()
	return mock.DispatchFunc(args, stdin, stdout, stderr)
}

// DispatchCalls gets all the calls that were made to Dispatch.
// Check the length with:
//     len(mockedPackageManager.DispatchCalls())
func (mock *PackageManagerMock) DispatchCalls() []struct {
	Args   []string
	Stdin  io.Reader
	Stdout io.Writer
	Stderr io.Writer
} {
	var calls []struct {
		Args   []string
		Stdin  io.Reader
		Stdout io.Writer
		Stderr io.Writer
	}

	mock.lockDispatch.RLock()
	calls = mock.calls.Dispatch
	mock.lockDispatch.RUnlock()
	return calls
}

// Install calls InstallFunc.
func (mock *PackageManagerMock) Install(url string, stdout io.Writer, stderr io.Writer) error {
	if mock.InstallFunc == nil {
		panic("PackageManagerMock.InstallFunc: method is nil but PackageManager.Install was just called")
	}

	callInfo := struct {
		URL    string
		Stdout io.Writer
		Stderr io.Writer
	}{
		URL:    url,
		Stdout: stdout,
		Stderr: stderr,
	}

	mock.lockInstall.Lock()
	mock.calls.Install = append(mock.calls.Install, callInfo)
	mock.lockInstall.Unlock()
	return mock.InstallFunc(url, stdout, stderr)
}

// InstallCalls gets all the calls that were made to Install.
// Check the length with:
//     len(mockedPackageManager.InstallCalls())
func (mock *PackageManagerMock) InstallCalls() []struct {
	URL    string
	Stdout io.Writer
	Stderr io.Writer
} {
	var calls []struct {
		URL    string
		Stdout io.Writer
		Stderr io.Writer
	}

	mock.lockInstall.RLock()
	calls = mock.calls.Install
	mock.lockInstall.RUnlock()
	return calls
}

// InstallLocal calls InstallLocalFunc.
func (mock *PackageManagerMock) InstallLocal(dir string) error {
	if mock.InstallLocalFunc == nil {
		panic("PackageManagerMock.InstallLocalFunc: method is nil but PackageManager.InstallLocal was just called")
	}

	callInfo := struct {
		Dir string
	}{
		Dir: dir,
	}

	mock.lockInstallLocal.Lock()
	mock.calls.InstallLocal = append(mock.calls.InstallLocal, callInfo)
	mock.lockInstallLocal.Unlock()
	return mock.InstallLocalFunc(dir)
}

// InstallLocalCalls gets all the calls that were made to InstallLocal.
// Check the length with:
//     len(mockedPackageManager.InstallLocalCalls())
func (mock *PackageManagerMock) InstallLocalCalls() []struct {
	Dir string
} {
	var calls []struct {
		Dir string
	}

	mock.lockInstallLocal.RLock()
	calls = mock.calls.InstallLocal
	mock.lockInstallLocal.RUnlock()
	return calls
}

// List calls ListFunc.
func (mock *PackageManagerMock) List(includeMetadata bool) []Package {
	if mock.ListFunc == nil {
		panic("PackageManagerMock.ListFunc: method is nil but PackageManager.List was just called")
	}

	callInfo := struct {
		IncludeMetadata bool
	}{
		IncludeMetadata: includeMetadata,
	}

	mock.lockList.Lock()
	mock.calls.List = append(mock.calls.List, callInfo)
	mock.lockList.Unlock()
	return mock.ListFunc(includeMetadata)
}

// ListCalls gets all the calls that were made to List.
// Check the length with:
//     len(mockedPackageManager.ListCalls())
func (mock *PackageManagerMock) ListCalls() []struct {
	IncludeMetadata bool
} {
	var calls []struct {
		IncludeMetadata bool
	}

	mock.lockList.RLock()
	calls = mock.calls.List
	mock.lockList.RUnlock()
	return calls
}

// Remove calls RemoveFunc.
func (mock *PackageManagerMock) Remove(name string) error {
	if mock.RemoveFunc == nil {
		panic("PackageManagerMock.RemoveFunc: method is nil but PackageManager.Remove was just called")
	}

	callInfo := struct {
		Name string
	}{
		Name: name,
	}

	mock.lockRemove.Lock()
	mock.calls.Remove = append(mock.calls.Remove, callInfo)
	mock.lockRemove.Unlock()
	return mock.RemoveFunc(name)
}

// RemoveCalls gets all the calls that were made to Remove.
// Check the length with:
//     len(mockedPackageManager.RemoveCalls())
func (mock *PackageManagerMock) RemoveCalls() []struct {
	Name string
} {
	var calls []struct {
		Name string
	}

	mock.lockRemove.RLock()
	calls = mock.calls.Remove
	mock.lockRemove.RUnlock()
	return calls
}

// Upgrade calls UpgradeFunc.
func (mock *PackageManagerMock) Upgrade(name string, force bool, stdout io.Writer, stderr io.Writer) error {
	if mock.UpgradeFunc == nil {
		panic("PackageManagerMock.UpgradeFunc: method is nil but PackageManager.Upgrade was just called")
	}

	callInfo := struct {
		Name   string
		Force  bool
		Stdout io.Writer
		Stderr io.Writer
	}{
		Name:   name,
		Force:  force,
		Stdout: stdout,
		Stderr: stderr,
	}

	mock.lockUpgrade.Lock()
	mock.calls.Upgrade = append(mock.calls.Upgrade, callInfo)
	mock.lockUpgrade.Unlock()
	return mock.UpgradeFunc(name, force, stdout, stderr)
}

// UpgradeCalls gets all the calls that were made to Upgrade.
// Check the length with:
//     len(mockedPackageManager.UpgradeCalls())
func (mock *PackageManagerMock) UpgradeCalls() []struct {
	Name   string
	Force  bool
	Stdout io.Writer
	Stderr io.Writer
} {
	var calls []struct {
		Name   string
		Force  bool
		Stdout io.Writer
		Stderr io.Writer
	}

	mock.lockUpgrade.RLock()
	calls = mock.calls.Upgrade
	mock.lockUpgrade.RUnlock()
	return calls
}
