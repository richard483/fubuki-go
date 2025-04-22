@Library('global-pipeline') _

GlobalPipeline() {
	dockerImage = "fbk-go:latest"
	projectName = "fbk-go"
	appPort = "234"
    networkName = "fubuki"
    buildArgs = [
        PORT: appPort,
        GEMINI_API_KEY: "\"${this.env.GEMINI_API_KEY}\"",
        POSTGRES_URI: "\"${this.env.FBK_POSTGRES_URI}\"",
        HOST: "\"fbk.nephren.xyz\"",
        RETRIEVE_HISTORY: "true",
        GEMINI_MODEL: "gemini-2.5-pro-exp-03-25",
        RELEASE_MODE: "true"
    ]
}
