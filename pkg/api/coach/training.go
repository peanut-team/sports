package coach

type SportsmanStatus string

const (
	SportsmanStatus_Training SportsmanStatus = "training" // 训练中
	SportsmanStatus_Online   SportsmanStatus = "online"   // 在线
	MatchType_Offline        SportsmanStatus = "offline"  // 离线
)

type AthleteTrainingList []*AthleteTraining
type AthleteTraining struct {
	SportImg                    string          `json:"sport_img"`                     // 运动员头像
	AthleteID                   int           `json:"athlete_id"`                    // 运动员ID
	AthleteName                 string          `json:"athlete_name"`                  // 运动员姓名
	Status                      SportsmanStatus `json:"status"`                        // 当前状态
	Distance                    float64         `json:"distance"`                      // 距离（路程）
	InstantaneousSpeed          float64         `json:"instantaneous_speed"`           // 瞬时时速
	AverageSpeed                float64         `json:"average_speed"`                 // 平均时速
	TotalOars                   int32           `json:"total_oars"`                    // 总桨数
	InstantaneousPropellerSpeed float64         `json:"instantaneous_propeller_speed"` // 瞬时桨速
	Stroke                      float64         `json:"stroke"`                        // 划行距离
	Acceleration                float64         `json:"acceleration"`                  // 功率（加速度）
	TrainingStatus              bool            `json:"athlete_training_status"`
}
