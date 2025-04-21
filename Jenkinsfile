@Library('global-pipeline') _

GlobalPipeline() {
	dockerImage = "fbk-go-canary:latest"
	projectName = "fbk-go-canary"
	appPort = "234"
    networkName = "fubuki"
    buildArgs = [
        PORT: appPort,
        GEMINI_API_KEY: "\"${this.env.GEMINI_API_KEY}\"",
        POSTGRES_URI: "\"${this.env.FBK_POSTGRES_URI}\"",
        HOST: "\"fbk-canary.nephren.xyz\"",
        RETRIEVE_HISTORY: "true",
        RELEASE_MODE: "true"
    ]
}
