pipeline {
    agent any
    //agent { docker { image 'golang' } }
 	  // Clean workspace before doing anything
    //deleteDir()
    stages {
      stage ('Clone') {
        steps {
          sh "checkout scm"
          checkout scm
        }
      }
      stage ('Go Build') {
        steps {
          sh "echo 'shell scripts to build from go source ...'"
          sh "go build -o test main.go"
        }
      }
      stage ('Docker Image Build') {
        steps {
          sh "echo 'shell scripts to build the docker image...'"
          sh "pwd"
          sh "docker build -t gotest_build:${BUILD_ID} ."
        }
      }

      /*
      stage ('Docker Image Push') {
      sh "docker login markusvanlaak/gotest"
      sh "docker tag gotest:${BUILD_ID} markusvanlaak/gotest"
      sh "docker push markusvanlaak/gotest"
    }
    */
    stage ('Deploy') {
      steps {
        sh "echo 'shell scripts to deploy to server...'"
      }
    }
  }
}
