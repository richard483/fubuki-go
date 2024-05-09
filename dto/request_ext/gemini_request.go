package request

type GeminiPart struct {
	Text string `json:"text,omitempty" binding:"required"`
}

type GeminiContent struct {
	Role  string        `json:"role,omitempty" binding:"required"`
	Parts *[]GeminiPart `json:"parts,omitempty" binding:"required"`
}

type GeminiContents struct {
	Contents *[]GeminiContent `json:"contents,omitempty" binding:"required"`
}

type GeminiHyperparameter struct {
	BatchSize    int     `json:"batch_size,omitempty" binding:"required"`
	LearningRate float32 `json:"learning_rate,omitempty" binding:"required"`
	EpochCount   int     `json:"epoch_count,omitempty" binding:"required"`
}

type GeminiExampleTrainingData struct {
	TextInput string `json:"text_input,omitempty" binding:"required"`
	Output    string `json:"output,omitempty" binding:"required"`
}

type GeminiExamplesTrainingData struct {
	Examples *[]GeminiExampleTrainingData `json:"examples,omitempty" binding:"required"`
}

type GeminiTrainingData struct {
	Examples *GeminiExamplesTrainingData `json:"examples,omitempty" binding:"required"`
}

type GeminiTuningTask struct {
	Hyperparameter *GeminiHyperparameter `json:"hyperparameters,omitempty" binding:"required"`
	TrainingData   *GeminiTrainingData   `json:"training_data,omitempty" binding:"required"`
}

type GeminiTuneModel struct {
	DisplayName string            `json:"display_name,omitempty" binding:"required"`
	BaseModel   string            `json:"base_model,omitempty" binding:"required"`
	TuningTask  *GeminiTuningTask `json:"tuning_task,omitempty" binding:"required"`
}
