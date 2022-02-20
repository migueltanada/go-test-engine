package questionaire

type Qna struct {
	SetID    string `json:"setId"`
	QnaID    string `json:"id,omitempty"`
	Question string `json:"question"`
	Choice   Choice `json:"choices"`
}

type Choice struct {
	CorrectChoices []string `json:"correct"`
	WrongChoices   []string `json:"wrong"`
}
