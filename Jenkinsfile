node {

 	// Clean workspace before doing anything
    deleteDir()

    try {
        stage ('Clone') {
        	checkout([$class: 'GitSCM', branches: [[name: '*/master']], doGenerateSubmoduleConfigurations: false, extensions: [], submoduleCfg: [], userRemoteConfigs: [[url: 'https://github.com/markusvanlaak/newgo']]])
        }
        stage ('Go Build') {
        	sh "echo 'shell scripts to build from go source ...'"
          sh "go build -o test main.go"
        }
        stage ('Docker Image Build') {
        	sh "echo 'shell scripts to build the docker image...'"
          sh "pwd"
          sh "docker build -t gotest_build:${BUILD_ID} ."
          app = docker.build("my-image:${env.BUILD_ID}")
          //customImage.push()
        }
        stage ('Docker Image Push') {
          //docker.withRegistry('https://hub.docker.com/markusvanlaak', 'dockerhub') {
          //sh "docker login markusvanlaak/gotest"
          //sh "docker tag gotest:${BUILD_ID} markusvanlaak/gotest"
          //sh "docker push markusvanlaak/gotest"
          //customImage.push()
          docker.withRegistry('http://localhost:5000', '') {

            def customImage = docker.build("my-image:${env.BUILD_ID}")

            /* Push the container to the custom Registry */
            customImage.push()
          }
        }
      	stage ('Deploy') {
            sh "echo 'shell scripts to deploy to server...'"
      	}
    } catch (err) {
        currentBuild.result = 'FAILED'
        throw err
    }
}
