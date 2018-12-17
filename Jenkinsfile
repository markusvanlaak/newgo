pipeline {
  environment {
    registry = “markusvanlaak/gotest”
    registryCredential = ‘dockerhub’
    dockerImage = ‘’
  }
 	// Clean workspace before doing anything
    deleteDir()

    try {
      environment {
        registry = “markusvanlaak/gotest”
        registryCredential = ‘dockerhub’
        dockerImage = ‘’
      }

        stage ('Clone') {
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
<<<<<<< HEAD
        }
=======
        }
        stage(‘Building image’) {
          steps{
            script {
              dockerImage = docker.build registry + “:$BUILD_NUMBER”
            }
          }
        }
        stage(‘Deploy Image’) {
          steps{
            script {
              docker.withRegistry( ‘’, registryCredential ) {
                dockerImage.push()
              }
            }
          }
        }
>>>>>>> 39b136798a08373342b55324175e5cc75801ec4f
        stage ('Deploy') {
            sh "echo 'shell scripts to deploy to server...'"
      	}
    } catch (err) {
        currentBuild.result = 'FAILED'
        throw err
    }
}
