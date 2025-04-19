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
        GEMINI_API: "false",
        GOOGLE_PROJECT_ID: "tech-395517",
        HOST: "\"fbk-go.nephren.xyz\"",
        GOOGLE_ACCESS_TOKEN: "\"${this.env.GOOGLE_ACCESS_TOKEN}\"",
        RETRIEVE_HISTORY: "true",
        RELEASE_MODE: "true"
    ]
}
