package coach

type SportsmanStatus int

const (
	SportsmanStatus_Training SportsmanStatus = 2 // 训练中
	SportsmanStatus_Online   SportsmanStatus = 1 // 在线
	MatchType_Offline        SportsmanStatus = 0 // 离线
)

type AthleteTrainingList []*AthleteTraining
type AthleteTraining struct {
	StartTime                   int64           `json:"start_time"`                    // 运动员开始运动的时间戳
	SportImg                    string          `json:"sport_img"`                     // 运动员头像
	AthleteID                   int             `json:"athlete_id"`                    // 运动员ID
	AthleteName                 string          `json:"athlete_name"`                  // 运动员姓名
	Status                      SportsmanStatus `json:"status"`                        // 当前状态
	Distance                    float64         `json:"distance"`                      // 学员训练距离，单位：m
	InstantaneousSpeed          float64         `json:"instantaneous_speed"`           // 加速度，单位：m/s2（米每二次方秒）
	AverageSpeed                float64         `json:"average_speed"`                 // 平均时速,平均速度，单位：m/s
	TotalOars                   int32           `json:"total_oars"`                    // 总桨数
	InstantaneousPropellerSpeed float64         `json:"instantaneous_propeller_speed"` // 瞬时桨速
	Stroke                      float64         `json:"stroke"`                        // 划行距离
	Acceleration                float64         `json:"acceleration"`                  // 功率（加速度）
	TrainingStatus              bool            `json:"athlete_training_status"`       //
}
