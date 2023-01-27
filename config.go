package gjobctl

type AppConfig struct {
	Region     string `yaml:"region"`
	JobName    string `yaml:"job_name"`
	BucketName string `yaml:"bucket_name"`
	BucketPath string `yaml:"bucket_path"`
	ScriptDIR  string `yaml:"script_dir"`
	ScriptName string `yaml:"script_name"`
}
