package resource

import (
	"errors"
	"fmt"
	"sort"
	"testing"

	"github.com/hashicorp/waypoint-plugin-sdk/internal/testproto"
	pb "github.com/hashicorp/waypoint-plugin-sdk/proto/gen"
	"github.com/ryboe/q"
	"github.com/stretchr/testify/require"
)

func TestManagerCreateAll(t *testing.T) {
	t.Run("with no resources", func(t *testing.T) {
		m := NewManager()
		require.NoError(t, m.CreateAll(int(42)))
	})

	t.Run("with two non-dependent resources", func(t *testing.T) {
		require := require.New(t)

		var calledA, calledB int
		m := NewManager(
			WithResource(NewResource(
				WithName("A"),
				WithCreate(func(v int) error {
					calledA = v
					return nil
				}),
			)),

			WithResource(NewResource(
				WithName("B"),
				WithCreate(func(v int) error {
					calledB = v
					return nil
				}),
			)),
		)

		// Create
		require.NoError(m.CreateAll(int(42)))

		// Ensure we called all
		require.Equal(calledA, 42)
		require.Equal(calledB, 42)
	})

	t.Run("with two dependent resources", func(t *testing.T) {
		require := require.New(t)

		var calledB int
		m := NewManager(
			WithResource(NewResource(
				WithName("A"),
				WithState(&testState{}),
				WithCreate(func(s *testState, v int) error {
					s.Value = v
					return nil
				}),
			)),

			WithResource(NewResource(
				WithName("B"),
				WithCreate(func(s *testState) error {
					calledB = s.Value
					return nil
				}),
			)),
		)

		// Create
		require.NoError(m.CreateAll(int(42)))

		// Ensure we called all
		require.Equal(calledB, 42)

		// Ensure we have state
		require.NotNil(m.State())
	})

	t.Run("rollback on error", func(t *testing.T) {
		require := require.New(t)

		var destroyOrder []string
		m := NewManager(
			WithResource(NewResource(
				WithName("A"),
				WithState(&testState{}),
				WithCreate(func(s *testState, v int) error {
					s.Value = v
					return nil
				}),
				WithDestroy(func() error {
					destroyOrder = append(destroyOrder, "A")
					return nil
				}),
			)),

			WithResource(NewResource(
				WithName("B"),
				WithState(&testState2{}),
				WithCreate(func(s *testState) error {
					return errors.New("whelp")
				}),
				WithDestroy(func() error {
					destroyOrder = append(destroyOrder, "B")
					return nil
				}),
			)),

			WithResource(NewResource(
				WithName("C"),
				WithCreate(func(s *testState2) error {
					return nil
				}),
				WithDestroy(func() error {
					destroyOrder = append(destroyOrder, "C")
					return nil
				}),
			)),
		)

		// Create
		err := m.CreateAll(int(42))
		require.Error(err)
		require.Equal("whelp", err.Error())

		// Ensure we called destroy
		require.Equal([]string{"B", "A"}, destroyOrder)

		// Ensure we have no state
		require.NotNil(m.State())
	})
}

func TestManagerDestroyAll(t *testing.T) {
	var calledB int32
	require := require.New(t)

	// init is a function so that we can reinitialize an empty manager
	// for this test to test loading state
	var destroyOrder []string
	var destroyState int32
	init := func() *Manager {
		return NewManager(
			WithResource(NewResource(
				WithName("A"),
				WithState(&testproto.Data{}),
				WithCreate(func(s *testproto.Data, v int32) error {
					s.Number = v
					return nil
				}),
				WithDestroy(func(s *testproto.Data) error {
					destroyOrder = append(destroyOrder, "A")
					destroyState = s.Number
					return nil
				}),
			)),

			WithResource(NewResource(
				WithName("B"),
				WithCreate(func(s *testproto.Data) error {
					calledB = s.Number
					return nil
				}),
				WithDestroy(func() error {
					destroyOrder = append(destroyOrder, "B")
					return nil
				}),
			)),
		)
	}

	// Create
	m := init()
	require.NoError(m.CreateAll(int32(42)))

	// Ensure we called all
	require.Equal(calledB, int32(42))

	// Create a new manager, load the state, and verify it works
	m2 := init()
	require.NoError(m2.LoadState(m.State()))

	// Grab our resource state
	actual := m2.Resource("A").State().(*testproto.Data)
	require.NotNil(actual)
	require.Equal(actual.Number, int32(42))

	// Destroy
	require.NoError(m2.DestroyAll())

	// Ensure we destroyed
	require.Equal([]string{"B", "A"}, destroyOrder)
	require.Equal(destroyState, int32(42))
}

func TestManagerDestroyAll_noDestroyFunc(t *testing.T) {
	var calledB int32
	require := require.New(t)

	// init is a function so that we can reinitialize an empty manager
	// for this test to test loading state
	var destroyOrder []string
	init := func() *Manager {
		return NewManager(
			WithResource(NewResource(
				WithName("A"),
				WithState(&testproto.Data{}),
				WithCreate(func(s *testproto.Data, v int32) error {
					s.Number = v
					return nil
				}),
			)),

			WithResource(NewResource(
				WithName("B"),
				WithCreate(func(s *testproto.Data) error {
					calledB = s.Number
					return nil
				}),
				WithDestroy(func() error {
					destroyOrder = append(destroyOrder, "B")
					return nil
				}),
			)),
		)
	}

	// Create
	m := init()
	require.NoError(m.CreateAll(int32(42)))

	// Ensure we called all
	require.Equal(calledB, int32(42))

	// Create a new manager, load the state, and verify it works
	m2 := init()
	require.NoError(m2.LoadState(m.State()))

	// Grab our resource state
	actual := m2.Resource("A").State().(*testproto.Data)
	require.NotNil(actual)
	require.Equal(actual.Number, int32(42))

	// Destroy
	require.NoError(m2.DestroyAll())

	// Ensure we destroyed
	require.Equal([]string{"B"}, destroyOrder)
}

func TestManagerStatusAll(t *testing.T) {
	q.Q("---------")
	q.Q("starting")
	q.Q("---------")
	defer func() {
		q.Q("---------")
		q.Q("end")
		q.Q("---------")
		q.Q("")
	}()

	require := require.New(t)

	// init is a function so that we can reinitialize an empty manager
	// for this test to test loading state
	init := func() *Manager {
		return NewManager(
			WithResource(NewResource(
				WithName("A"),
				WithState(&testState3{}),
				WithCreate(func(s *testState3, v int) error {
					sAddr := fmt.Sprintf("%p", s)
					q.Q("==> sA: ", sAddr)
					s.Name = "resource A"
					s.Value = v
					return nil
				}),
				WithStatus(func(s *testState3, sr *pb.StatusReport_Resource) error {
					sAddr := fmt.Sprintf("%p", s)
					srAddr := fmt.Sprintf("%p", sr)
					q.Q("==> sA: ", sAddr)
					q.Q("==> sAr: ", srAddr)
					sr.Name = s.Name
					return nil
				}),
			)),

			WithResource(NewResource(
				WithName("B"),
				WithCreate(func(s *testState3) error {
					// no-op
					sAddr := fmt.Sprintf("%p", s)
					q.Q("==> sB (should match A): ", sAddr)
					return nil
				}),
				WithStatus(func(sr *pb.StatusReport_Resource) error {
					// sAddr := fmt.Sprintf("%p", s)
					srAddr := fmt.Sprintf("%p", sr)
					// q.Q("==> sB: ", sAddr)
					q.Q("==> sBr: ", srAddr)
					sr.Name = "no state here"
					return nil
				}),
			)),
			WithResource(NewResource(
				WithName("C"),
				WithState(&testState4{}),
				WithCreate(func(s *testState4) error {
					sAddr := fmt.Sprintf("%p", s)
					q.Q("==> sC: ", sAddr)
					s.Name = "resource C"
					return nil
				}),
				WithStatus(func(s *testState4, sr *pb.StatusReport_Resource) error {
					sAddr := fmt.Sprintf("%p", s)
					srAddr := fmt.Sprintf("%p", sr)
					q.Q("==> sC: ", sAddr)
					q.Q("==> sCr: ", srAddr)
					sr.Name = s.Name
					return nil
				}),
			)),
			WithResource(NewResource(
				WithName("D"),
				WithState(&testState{}),
				WithCreate(func(s *testState) error {
					sAddr := fmt.Sprintf("%p", s)
					q.Q("==> sD: ", sAddr)
					s.Value = 0
					return nil
				}),
				// WithStatus(func(s *testState4, sr *pb.StatusReport_Resource) error {
				// 	sAddr := fmt.Sprintf("%p", s)
				// 	srAddr := fmt.Sprintf("%p", sr)
				// 	sr.Name = s.Name
				// 	return nil
				// }),
			)),
		)
	}

	// Create
	m := init()
	require.NoError(m.CreateAll(42))

	require.NoError(m.StatusAll())
	reports := m.Status()

	require.Len(reports, 3)
	sort.Sort(byName(reports))
	// for _, r := range reports {
	// 	q.Q("report Name:", r.Name)
	// }

	require.Equal(reports[0].Name, "no state here")
	require.Equal(reports[1].Name, "resource A")
	require.Equal(reports[2].Name, "resource C")

	// Destroy
	require.NoError(m.DestroyAll())
}

// byName implements sort.Interface for sorting the results from calling
// Status(), to ensure ordering when validating the tests
type byName []pb.StatusReport_Resource

func (a byName) Len() int           { return len(a) }
func (a byName) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a byName) Less(i, j int) bool { return a[i].Name < a[j].Name }
