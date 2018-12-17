pipeline {
  environment {
    registry = “markusvanlaak/gotest”
    registryCredential = ‘dockerhub’
    dockerImage = ‘’
  }

 	  // Clean workspace before doing anything
    //deleteDir()

    try {
        stage ('Clone') {
          sh "checkout scm"
          checkout scm
        }
        stage ('Go Build') {
        	sh "echo 'shell scripts to build from go source ...'"
          sh "go build -o test main.go"
        }
        stage ('Docker Image Build') {
        	sh "echo 'shell scripts to build the docker image...'"
          sh "pwd"
          sh "docker build -t gotest_build:${BUILD_ID} ."
        }
        stage ('Docker Image Push') {
          sh "docker login markusvanlaak/gotest"
          sh "docker tag gotest:${BUILD_ID} markusvanlaak/gotest"
          sh "docker push markusvanlaak/gotest"
        }
      	stage ('Deploy') {
            sh "echo 'shell scripts to deploy to server...'"
      	}
    } catch (err) {
        currentBuild.result = 'FAILED'
        throw err
    }
}
