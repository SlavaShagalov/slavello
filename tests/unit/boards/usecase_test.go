package boards

import (
	"context"
	"github.com/SlavaShagalov/slavello/tests/utils/builder"
	"reflect"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/ozontech/allure-go/pkg/framework/provider"
	"github.com/ozontech/allure-go/pkg/framework/suite"
	"github.com/pkg/errors"

	"github.com/SlavaShagalov/slavello/internal/boards/mocks"
	boardsUsecase "github.com/SlavaShagalov/slavello/internal/boards/usecase"
	"github.com/SlavaShagalov/slavello/internal/models"
	pkgErrors "github.com/SlavaShagalov/slavello/internal/pkg/errors"
)

type BoardsUsecaseSuite struct {
	suite.Suite

	boardBuilder *builder.BoardBuilder
}

func (s *BoardsUsecaseSuite) BeforeEach(t provider.T) {
	t.WithNewStep("SetupTest step", func(ctx provider.StepCtx) {})

	s.boardBuilder = builder.NewBoardBuilder()
}

func (s *BoardsUsecaseSuite) TestList(t provider.T) {
	type fields struct {
		repo        *mocks.MockRepository
		workspaceID int
		boards      []models.Board
	}

	type testCase struct {
		prepare     func(f *fields)
		workspaceID int
		boards      []models.Board
		err         error
	}

	ctx := context.Background()

	tests := map[string]testCase{
		"normal": {
			prepare: func(f *fields) {
				f.repo.EXPECT().List(ctx, f.workspaceID).Return(f.boards, nil)
			},
			workspaceID: 27,
			boards: []models.Board{
				s.boardBuilder.
					WithID(21).
					WithWorkspaceID(27).
					WithTitle("University").
					WithDescription("University Board").
					Build(),
				s.boardBuilder.
					WithID(22).
					WithWorkspaceID(27).
					WithTitle("Life").
					WithDescription("Life Board").
					Build(),
				s.boardBuilder.
					WithID(23).
					WithWorkspaceID(27).
					WithTitle("Work").
					WithDescription("Work Board").
					Build(),
			},
			err: nil,
		},
		"empty result": {
			prepare: func(f *fields) {
				f.repo.EXPECT().List(ctx, f.workspaceID).Return(f.boards, nil)
			},
			workspaceID: 27,
			boards:      []models.Board{},
			err:         nil,
		},
		"board not found": {
			prepare: func(f *fields) {
				f.repo.EXPECT().List(ctx, f.workspaceID).Return(f.boards, pkgErrors.ErrWorkspaceNotFound)
			},
			workspaceID: 27,
			boards:      nil,
			err:         pkgErrors.ErrWorkspaceNotFound,
		},
		"db error": {
			prepare: func(f *fields) {
				f.repo.EXPECT().List(ctx, f.workspaceID).Return(f.boards, pkgErrors.ErrDb)
			},
			workspaceID: 27,
			boards:      nil,
			err:         pkgErrors.ErrDb,
		},
	}

	for name, test := range tests {
		test := test
		t.Run(name, func(t provider.T) {
			t.Parallel()

			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			f := fields{
				repo:        mocks.NewMockRepository(ctrl),
				workspaceID: test.workspaceID, boards: test.boards}
			if test.prepare != nil {
				test.prepare(&f)
			}

			serv := boardsUsecase.New(f.repo)
			boards, err := serv.ListByWorkspace(ctx, test.workspaceID)
			if !errors.Is(err, test.err) {
				t.Errorf("\nExpected: %s\nGot: %s", test.err, err)
			}
			if !reflect.DeepEqual(boards, test.boards) {
				t.Errorf("\nExpected: %v\nGot: %v", test.boards, boards)
			}
		})
	}
}

func TestSuiteRunner(t *testing.T) {
	suite.RunSuite(t, new(BoardsUsecaseSuite))
}
