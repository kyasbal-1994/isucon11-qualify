package scenario

import "github.com/isucon/isucandar/score"

// スコアタグの管理

const (
	ScoreStartBenchmark        score.ScoreTag = "00.StartBenchmark        "
	ScoreGraphExcellent        score.ScoreTag = "01.GraphExcellent        "
	ScoreGraphGood             score.ScoreTag = "02.GraphGood             "
	ScoreGraphNormal           score.ScoreTag = "03.GraphNormal           "
	ScoreGraphBad              score.ScoreTag = "04.GraphBad              "
	ScoreGraphWorst            score.ScoreTag = "05.GraphWorst            "
	ScoreTodayGraphExcellent   score.ScoreTag = "06.TodayGraphExcellent   "
	ScoreTodayGraphGood        score.ScoreTag = "07.TodayGraphGood        "
	ScoreTodayGraphNormal      score.ScoreTag = "08.TodayGraphNormal      "
	ScoreTodayGraphBad         score.ScoreTag = "09.TodayGraphBad         "
	ScoreTodayGraphWorst       score.ScoreTag = "10.TodayGraphWorst      "
	ScoreReadInfoCondition     score.ScoreTag = "11.ReadInfoCondition    "
	ScoreReadWarningCondition  score.ScoreTag = "12.ReadWarningCondition "
	ScoreReadCriticalCondition score.ScoreTag = "13.ReadCriticalCondition"
	ScoreIsuInitialize         score.ScoreTag = "_1.IsuInitialize        " //scoreが0のもの
	ScoreNormalUserInitialize  score.ScoreTag = "_2.NormalUserInitialize " //全てのIsuInitializeが終わってはじめて+1
	ScoreViewerInitialize      score.ScoreTag = "_3.ViewerInitialize     "
	ScoreViewerDropout         score.ScoreTag = "_4.ViewerDropout        "
	ScoreRepairIsu             score.ScoreTag = "_5.RepairIsu            "
	ScorePostInfoCondition     score.ScoreTag = "_6.PostInfoCondition    "
	ScorePostWarningCondition  score.ScoreTag = "_7.PostWarningCondition "
	ScorePostCriticalCondition score.ScoreTag = "_8.PostCriticalCondition"
)
