pipeline {
  environment {
      dockerHubCredential = '43fb44ba-27e9-4061-8db0-d20dbe3689c6'
      imageName = 'datvc/go-webserver'
  }

  agent any

  stages {

    stage('Clone Repository') {
      steps {
        echo 'Cloning repositories'
        git([url: 'https://github.com/chpiano2000/go-webserver.git', branch: 'cicd'])
      }
    }

    stage('Build image') {
      steps{
        echo 'Starting to build docker image'
        script {
          def commitHash = sh(script: 'git log -1 --pretty=%h', returnStdout: true)
          env.COMMIT_HASH = commitHash
          docker.build("$imageName")
        }
      }
    }

    stage('Push Image') {
     steps {
       script {
        docker.withDockerRegistry(registry: 'https://registry.hub.docker.com', '43fb44ba-27e9-4061-8db0-d20dbe3689c6') {
          docker.push("${COMMIT_HASH}")
        }
       }
     }
    }

  }
}
